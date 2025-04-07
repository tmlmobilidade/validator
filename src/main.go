package main

import (
	"fmt"
	"log"
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

	var totalRecords int
	for fileName, data := range gtfsData {
		switch fileName {
		case "agency.txt":
			records := len(data)
			totalRecords += records
			fmt.Printf("File: %s, Records: %d\n", fileName, records)
		}
	}
}
