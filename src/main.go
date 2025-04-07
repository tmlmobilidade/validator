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

	// gtfsData, err := services.ReadGTFSZip("data/BOM.zip")
	gtfsData, err := services.ReadGTFSZip("data/CMET.zip")
	if err != nil {
		log.Fatalf("Error reading GTFS: %v", err)
	}

	// Run Validations for each file
	for fileName := range gtfsData {

		// If fileName is not in the GTFS_FILE_RULES_MAP, skip
		if _, ok := validations.GTFS_FILE_RULES_MAP[fileName]; !ok {
			// TODO: Add to warning messages
			continue
		}

		fmt.Println("Running Validations for", fileName)
		validations.GTFS_FILE_RULES_MAP[fileName](gtfsData)
	}
}
