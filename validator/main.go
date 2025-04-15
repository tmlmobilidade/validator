package main

import (
	"log"
	"main/validator/lib"
	"main/validator/services"
	"main/validator/types"
	"main/validator/validations"
	file_validation "main/validator/validations/files"
	"sync"
)

func runValidations(gtfs types.Gtfs, tracker *lib.PerformanceTracker) {
	// Create a wait group to wait for all validations to complete
	var wg sync.WaitGroup

	// Run Validations for each file concurrently
	for fileName := range gtfs.Files {
		// If fileName is not in the GTFS_FILE_RULES_MAP, skip
		if _, ok := validations.GTFS_FILE_RULES_MAP[fileName]; !ok {
			// TODO: Add to warning messages
			continue
		}

		wg.Add(1)
		go func(name string) {
			defer wg.Done()
			validations.GTFS_FILE_RULES_MAP[name](gtfs)
		}(fileName)
	}

	// Wait for all validations to complete
	wg.Wait()
	tracker.End()
}

func main() {
	services.AppCLI.Run()

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
	if errs := file_validation.NewFileValidation(nil).Validate(gtfs); len(errs) > 0 {
		for _, err := range errs {
			services.AppMessageService.AddMessage(err)
		}

		services.AppMessageService.PrintJSON(services.AppCLI.Options.OutputPath)
		return
	}

	// Run Validations for each file
	runValidations(gtfs, tracker)

	// Print Table
	// services.AppMessageService.PrintTable()

	// Print JSON
	services.AppMessageService.PrintJSON(services.AppCLI.Options.OutputPath)
}
