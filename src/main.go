package main

import (
	"fmt"
	"log"
	"main/src/models"
	"main/src/services"
)

func main() {

	// Clear terminal
	fmt.Print("\033c")

	gtfsData, err := services.ReadGTFSZip("data/BOM.zip")
	// gtfsData, err := services.ReadGTFSZip("data/CMET.zip")
	if err != nil {
		log.Fatalf("Error reading GTFS: %v", err)
	}

	for fileName, data := range gtfsData {
		fmt.Printf("File: %s, Records: %d\n", fileName, len(data))

		// Run Validations for current file
		models.GTFS_FILE_RULES_MAP[fileName]()
	}
}
