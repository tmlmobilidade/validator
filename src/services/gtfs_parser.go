package services

import (
	"archive/zip"
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"main/src/types"
	"os"
	"strings"
)

// GTFS_FILES defines the set of valid GTFS filenames that will be processed.
var GTFS_FILES = map[string]struct{}{
	"agency.txt":          {},
	"stops.txt":           {},
	"routes.txt":          {},
	"trips.txt":           {},
	"stop_times.txt":      {},
	"calendar.txt":        {},
	"archives.txt":        {},
	"calendar_dates.txt":  {},
	"dates.txt":           {},
	"fare_attributes.txt": {},
	"fare_rules.txt":      {},
	"feed_info.txt":       {},
	"municipalities.txt":  {},
	"periods.txt":         {},
	"shapes.txt":          {},
}

// ReadGTFSZip reads and parses a GTFS zip file at the specified path.
// It returns a Gtfs map containing the parsed data from all valid GTFS files,
// or an error if the file cannot be read or processed.
func ReadGTFSZip(zipPath string) (types.Gtfs, error) {
	if _, err := os.Stat(zipPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("file %s does not exist", zipPath)
	}

	zipReader, err := zip.OpenReader(zipPath)
	if err != nil {
		return nil, err
	}
	defer zipReader.Close()

	gtfsData := make(types.Gtfs)

	// Print all files in the zip
	for _, file := range zipReader.File {
		fmt.Println(file.Name)
	}

	for _, file := range zipReader.File {
		fileName := file.Name

		// Validate against known GTFS file types
		if _, valid := GTFS_FILES[fileName]; !valid {
			fmt.Printf("Skipping invalid GTFS file: %s\n", fileName)
			continue
		}

		// Open the file
		f, err := file.Open()
		if err != nil {
			fmt.Printf("Error opening file %s: %v\n", fileName, err)
			continue
		}
		defer f.Close()

		// Read the file
		content, err := io.ReadAll(f)
		if err != nil {
			fmt.Printf("Error reading file %s: %v\n", fileName, err)
			continue
		}

		// Parse the file
		parsedData, err := parseCSV(content)
		if err != nil {
			fmt.Printf("Error parsing file %s: %v\n", fileName, err)
			continue
		}

		// Add the parsed data to the gtfsData map
		gtfsData[strings.TrimSuffix(fileName, ".txt")] = parsedData
	}

	return gtfsData, nil
}

// parseCSV parses CSV content into a slice of maps where each map represents a row
// with column headers as keys and cell values as values.
// Returns an error if the CSV is empty or cannot be parsed.
func parseCSV(content []byte) ([]map[string]string, error) {
	reader := csv.NewReader(bytes.NewReader(content))
	reader.TrimLeadingSpace = true

	// Read all records from the CSV
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	// Check if the CSV is empty
	if len(records) < 1 {
		return nil, errors.New("CSV file is empty")
	}

	// Extract headers from the first row
	headers := records[0]

	// Initialize the result slice
	var result []map[string]string

	// Process each row (skipping the header row)
	for _, row := range records[1:] {
		entry := make(map[string]string)
		for i, value := range row {
			entry[headers[i]] = value
		}
		result = append(result, entry)
	}

	return result, nil
}
