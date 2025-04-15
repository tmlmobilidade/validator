package main

import (
	"log"
	"main/src/lib"
	"main/src/services"
	"main/src/types"
	"main/src/validations"
	file_validation "main/src/validations/files"
	"strings"
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

	// Clear the terminal
	lib.AppLogger.Clear()
	lib.AppLogger.Divider("GTFS Validator")

	// Start Performance Tracker
	tracker := lib.AppLogger.StartPerformanceTracker("Reading GTFS")

	// Read GTFS from zip file
	// gtfs, err := services.ReadGTFSZip("data/CMET.zip")
	gtfs, err := services.ReadGTFSZip("data/bad-files.zip")
	if err != nil {
		log.Fatalf("Error reading GTFS: %v", err)
	}
	// Check File Requirements
	fileNames := make([]string, 0, len(gtfs.Files))
	for name := range gtfs.Files {
		fileNames = append(fileNames, name)
	}
	lib.AppLogger.Accent("GTFS Files: ", strings.Join(fileNames, ", "))
	if errs := file_validation.NewFileValidation(nil).Validate(gtfs); len(errs) > 0 {
		lib.AppLogger.Error("Errors found in file requirements")
		for _, err := range errs {
			lib.AppLogger.Error(err.Message)
		}
		panic("Errors found in file requirements, aborting...")
	}

	// Run Validations for each file
	runValidations(gtfs, tracker)

	// Print Table
	services.AppMessageService.PrintTable()
}
