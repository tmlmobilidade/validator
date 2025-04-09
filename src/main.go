package main

import (
	"log"
	"main/src/lib"
	"main/src/services"
	"main/src/validations"
	"sync"
)

func main() {

	// Clear the terminal
	lib.AppLogger.Clear()
	lib.AppLogger.Divider("GTFS Validator")
	tracker := lib.AppLogger.StartPerformanceTracker("Reading GTFS")

	// gtfs, err := services.ReadGTFSZip("data/BOM.zip")
	gtfs, err := services.ReadGTFSZip("data/bad-format.zip")
	if err != nil {
		log.Fatalf("Error reading GTFS: %v", err)
	}

	tracker.End()
	tracker = lib.AppLogger.StartPerformanceTracker("Running Validations")

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

	// Print Summary
	// services.AppMessageService.PrintSummary()
	services.AppMessageService.PrintTable()
}
