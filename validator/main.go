package main

import (
	"fmt"
	"log"
	"main/i18n"
	"main/lib"
	"main/services"
	"main/types"
	"main/validations"
	file_validation "main/validations/files"
	"sync"
)

func runValidations(gtfs types.Gtfs, tracker *lib.PerformanceTracker, rules *types.GtfsRules) {
	// Create a wait group to wait for all validations to complete
	var wg sync.WaitGroup

	// Run Validations for each file concurrently
	for fileName := range gtfs.IdMap {
		// If fileName is not in the GTFS_FILE_RULES_MAP, skip
		if _, ok := validations.GTFS_FILE_RULES_MAP[fileName]; !ok {
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
		go func(name string) {
			defer wg.Done()
			validations.GTFS_FILE_RULES_MAP[name](gtfs, rules)
		}(fileName)
	}

	// Wait for all validations to complete
	wg.Wait()
	tracker.End()
}

func main() {
	services.AppCLI.Run()
	lib.AppLogger.SetLogLevel(services.AppCLI.Options.LogLevel)

	// Set Translator Language
	if services.AppCLI.Options.RulesLang != "" {
		i18n.AppTranslator.SetLanguage(services.AppCLI.Options.RulesLang)
	}

	// Parse Rules
	rules, err := services.NewRulesParser(services.AppCLI.Options.RulesPath).ParseRules()
	if err != nil {
		log.Fatalf("Error parsing rules: %v", err)
	}

	// Clear the terminal
	lib.AppLogger.Clear()
	lib.AppLogger.Divider("GTFS Validator")

	// Start Performance Tracker
	tracker := lib.AppLogger.StartPerformanceTracker("Reading GTFS")

	// Read GTFS from zip file
	gtfs, err := services.ReadGTFSZip(services.AppCLI.Options.InputPath)
	if err != nil {
		log.Fatalf("Error reading GTFS: %v", err)
	}

	// Check File Requirements
	if errs := file_validation.NewFileValidation(nil).Validate(gtfs, rules); len(errs) > 0 {
		for _, err := range errs {
			services.AppMessageService.AddMessage(err)
		}

		// Print JSON
		// services.AppMessageService.PrintJSON()

		services.AppMessageService.PrintTable()
		return
	}

	// Run Validations for each file
	runValidations(gtfs, tracker, rules)

	// Output Summary
	if services.AppCLI.Options.OutputPath != "" {
		services.AppMessageService.WriteToFile(services.AppCLI.Options.OutputPath)
		lib.AppLogger.Info("Summary written to: " + services.AppCLI.Options.OutputPath)
	} else {
		services.AppMessageService.PrintTable()
	}
}
