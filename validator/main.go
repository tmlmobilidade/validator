package main

import (
	"fmt"
	"log"
	"main/config"
	"main/i18n"
	"main/lib"
	"main/services"
	"main/types"
	"main/validations"
	file_validation "main/validations/files"
	"os"
	"path/filepath"
	"strings"
	"sync"

	// Import all validation packages to trigger their init() functions
	_ "main/validations/afetacao"
	_ "main/validations/agency"
	_ "main/validations/archives"
	_ "main/validations/areas"
	_ "main/validations/attributions"
	_ "main/validations/booking_rules"
	_ "main/validations/calendar"
	_ "main/validations/calendar_dates"
	_ "main/validations/fare_attributes"
	_ "main/validations/fare_leg_join_rules"
	_ "main/validations/fare_leg_rules"
	_ "main/validations/fare_media"
	_ "main/validations/fare_products"
	_ "main/validations/fare_rules"
	_ "main/validations/fare_transfer_rules"
	_ "main/validations/feed_info"
	_ "main/validations/frequencies"
	_ "main/validations/levels"
	_ "main/validations/location_group_stops"
	_ "main/validations/location_groups"
	_ "main/validations/municipalities"
	_ "main/validations/networks"
	_ "main/validations/pathways"
	_ "main/validations/periods"
	_ "main/validations/rider_categories"
	_ "main/validations/route_networks"
	_ "main/validations/routes"
	_ "main/validations/shapes"
	_ "main/validations/stop_areas"
	_ "main/validations/stop_times"
	_ "main/validations/stops"
	_ "main/validations/timeframes"
	_ "main/validations/transfers"
	_ "main/validations/translations"
	_ "main/validations/trips"
	_ "main/validations/vehicles"
)

func resolveMunicipalityCoordinatesPath(cliPath string) string {
	if strings.TrimSpace(cliPath) != "" {
		lib.AppLogger.Info("Using municipalities.json.")
	}

	candidates := []string{
		"municipalities.json",
		filepath.Join("..", "municipalities.json"),
	}

	for _, candidate := range candidates {
		if _, err := os.Stat(candidate); err == nil {
			lib.AppLogger.Info(fmt.Sprintf("Using municipality coordinates file: %s", candidate))
			return candidate
		}
	}

	lib.AppLogger.Info("Municipality coordinates file not found (checked: municipalities.json, ../municipalities.json). Coordinate-to-municipality validation will be skipped.")
	return ""
}

func runValidations(gtfs types.Gtfs, tracker *lib.PerformanceTracker, rules *types.GtfsRules) {
	// Create a wait group to wait for all validations to complete
	var wg sync.WaitGroup

	// List of all possible GTFS tables (from config)
	gtfsTables := config.GTFSTables

	// Run Validations for each file concurrently
	for _, fileName := range gtfsTables {
		// Check if table exists in database
		if !gtfs.HasTable(fileName) {
			continue
		}

		// Check if validation is registered for this table
		validationFn, ok := validations.Get(fileName)
		if !ok {
			services.AppMessageService.AddMessage(types.Message{
				Rows:         []int{},
				Field:        "N/A",
				FileName:     fileName,
				Message:      fmt.Sprintf(i18n.AppTranslator.Get("file_validations.not_supported"), fileName),
				ValidationID: "file_not_found_in_rules",
				Severity:     types.SEVERITY_WARNING,
			})
			continue
		}

		wg.Add(1)
		go func(name string, fn validations.ValidationFunction) {
			defer wg.Done()
			fn(gtfs, rules)
		}(fileName, validationFn)
	}

	// Wait for all validations to complete
	wg.Wait()
	tracker.End()
}

func preloadMunicipalityCoordinates(cliPath string) {
	coordinatesPath := resolveMunicipalityCoordinatesPath(cliPath)
	if coordinatesPath == "" {
		// Ensure the service remains explicitly disabled when no file is available.
		_ = services.LoadMunicipalityCoordinatesFromFile("")
		return
	}

	if err := services.LoadMunicipalityCoordinatesFromFile(coordinatesPath); err != nil {
		lib.AppLogger.Info(fmt.Sprintf("WARNING: failed to load municipality coordinates from %q: %v. Coordinate-to-municipality validation will be skipped.", coordinatesPath, err))
		_ = services.LoadMunicipalityCoordinatesFromFile("")
	}
}

func runStartupOrchestrator() (*types.GtfsRules, error) {
	// 0.1 Initialize CLI
	services.AppCLI.Run()
	lib.AppLogger.SetLogLevel(services.AppCLI.Options.LogLevel)

	// 0.2 Initialize Translator
	if services.AppCLI.Options.RulesLang != "" {
		i18n.AppTranslator.SetLanguage(services.AppCLI.Options.RulesLang)
	}

	// 0.3 Pre-load optional municipality coordinates map.
	preloadMunicipalityCoordinates(services.AppCLI.Options.MunicipalityCoordinatesPath)

	// 0.4 Parse Rules
	rules, err := services.NewRulesParser(services.AppCLI.Options.RulesPath).ParseRules()
	if err != nil {
		return nil, fmt.Errorf("error parsing rules: %w", err)
	}

	return rules, nil
}

func main() {
	rules, err := runStartupOrchestrator()
	if err != nil {
		log.Fatalf("Startup orchestrator failed: %v", err)
	}

	//
	// lib.AppLogger.Clear()
	lib.AppLogger.Divider("GTFS Validator")

	//
	// 0.5 Start Performance Tracker
	tracker := lib.AppLogger.StartPerformanceTracker("Reading GTFS")

	//
	// 1. Read GTFS from zip file
	gtfs, err := services.ReadGTFSZip(services.AppCLI.Options.InputPath)
	if err != nil {
		log.Fatalf("Error reading GTFS: %v", err)
	}

	// If there are errors in the GTFS, print the errors and exit
	if services.AppMessageService.GetSummary().TotalErrors > 0 {
		services.AppMessageService.PrintJSON()
		return
	}

	//
	// 2. Check File Requirements
	// File validations add messages directly to AppMessageService
	// Only exit early if there are errors (warnings are ok to continue)
	if hasErrors := file_validation.NewFileValidation().Validate(gtfs, rules); hasErrors {
		services.AppMessageService.WriteToFile(services.AppCLI.Options.OutputPath)
		lib.AppLogger.Error("File validations found errors. Exiting.")
		return
	}

	//
	// 3. Run Validations for each file
	runValidations(gtfs, tracker, rules)

	//
	// 4. Clean up database file
	defer func() {
		if err := gtfs.Close(); err != nil {
			lib.AppLogger.Error(fmt.Sprintf("Error closing database: %v", err))
		}
		if gtfs.DBPath() != "" {
			if err := os.Remove(gtfs.DBPath()); err != nil {
				lib.AppLogger.Debug(fmt.Sprintf("Error removing temp database file: %v", err))
			}
		}
	}()

	//
	// 5. Output Summary
	if services.AppCLI.Options.OutputPath != "" {
		services.AppMessageService.WriteToFile(services.AppCLI.Options.OutputPath)
		lib.AppLogger.Info("Summary written to: " + services.AppCLI.Options.OutputPath)
	} else {
		services.AppMessageService.PrintJSON()
	}
}
