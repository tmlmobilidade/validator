package main

import (
	"log"
	"main/src/lib"
	"main/src/services"
	"main/src/validations"
)

func main() {

	// Clear the terminal
	lib.AppLogger.Clear()
	lib.AppLogger.Divider("GTFS Validator")

	// gtfsData, err := services.ReadGTFSZip("data/BOM.zip")
	gtfsData, err := services.ReadGTFSZip("data/bad-format.zip")
	if err != nil {
		log.Fatalf("Error reading GTFS: %v", err)
	}

	lib.AppLogger.Divider("Running Validations")

	// Run Validations for each file
	for fileName := range gtfsData {

		// If fileName is not in the GTFS_FILE_RULES_MAP, skip
		if _, ok := validations.GTFS_FILE_RULES_MAP[fileName]; !ok {
			// TODO: Add to warning messages
			continue
		}

		validations.GTFS_FILE_RULES_MAP[fileName](gtfsData)
	}

	// Print Summary
	// services.AppMessageService.PrintSummary()
	services.AppMessageService.PrintTable()
}
