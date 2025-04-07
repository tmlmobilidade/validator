package main

import (
	"fmt"
	"log"
	"main/src/lib"
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

	// Run Validations for each file
	for fileName := range gtfsData {
		lib.GTFS_FILE_RULES_MAP[fileName](gtfsData)
	}
}
