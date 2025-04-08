package main

import (
	"fmt"
	"log"
	"main/src/services"
	"main/src/validations"
)

func main() {

	// Clear terminal
	fmt.Print("\033c")

	// Message Service
	messageService := services.NewMessageService()

	// gtfsData, err := services.ReadGTFSZip("data/BOM.zip")
	gtfsData, err := services.ReadGTFSZip("data/bad-format.zip")
	if err != nil {
		log.Fatalf("Error reading GTFS: %v", err)
	}

	fmt.Println("\n\n ==== ---- ==== \n\n")

	// Run Validations for each file
	for fileName := range gtfsData {

		// If fileName is not in the GTFS_FILE_RULES_MAP, skip
		if _, ok := validations.GTFS_FILE_RULES_MAP[fileName]; !ok {
			// TODO: Add to warning messages
			continue
		}

		validations.GTFS_FILE_RULES_MAP[fileName](gtfsData, *messageService)
	}

	// Print Summary
	messageService.PrintSummary()
}
