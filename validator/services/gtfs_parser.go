package services

import (
	"archive/zip"
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"main/lib"
	"main/types"
	"os"
	"runtime"
	"strings"
	"sync"
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
		return types.Gtfs{}, fmt.Errorf("file %s does not exist", zipPath)
	}

	zipReader, err := zip.OpenReader(zipPath)
	if err != nil {
		return types.Gtfs{}, err
	}
	defer zipReader.Close()

	gtfsFiles := make(types.GtfsFiles)
	gtfsFieldCount := make(types.GtfsFieldCount)
	gtfsIdsMap := make(types.GtfsIdMap)

	// Add mutexes to protect concurrent access to shared maps
	// What is a mutex? https://stackoverflow.com/questions/34524/what-is-a-mutex
	var fieldCountMutex sync.Mutex
	var idsMapMutex sync.Mutex

	type result struct {
		fileNameWithoutExt string
		data               []map[string]string
		err                error
	}

	// Channels
	jobs := make(chan *zip.File, len(zipReader.File))
	results := make(chan result, len(zipReader.File))

	// Worker pool
	numWorkers := runtime.NumCPU()
	var wg sync.WaitGroup
	for range numWorkers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for file := range jobs {
				fileName := file.Name
				fileNameWithoutExt := strings.TrimSuffix(fileName, ".txt")

				// Validate
				if _, valid := GTFS_FILES[fileName]; !valid {
					lib.AppLogger.Debug("Skipping invalid GTFS file: " + fileName)
					continue
				}

				f, err := file.Open()
				if err != nil {
					lib.AppLogger.Error("Error opening file: " + fileName + " " + err.Error())
					continue
				}
				content, err := io.ReadAll(f)
				f.Close() // Close immediately after reading
				if err != nil {
					lib.AppLogger.Error("Error reading file: " + fileName + " " + err.Error())
					continue
				}

				parsedData, err := parseCSV(content, fileNameWithoutExt, &gtfsFieldCount, &gtfsIdsMap, &fieldCountMutex, &idsMapMutex)
				if err != nil {
					lib.AppLogger.Error("Error parsing file: " + fileName + " " + err.Error())
					continue
				}

				results <- result{fileNameWithoutExt, parsedData, nil}
			}
		}()
	}

	// Feed jobs
	go func() {
		for _, file := range zipReader.File {
			jobs <- file
		}
		close(jobs)
	}()

	// Close results after workers are done
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results
	for res := range results {
		if res.err != nil {
			lib.AppLogger.Error("Error processing file: " + res.fileNameWithoutExt + " " + res.err.Error())
			continue
		}
		gtfsFiles[res.fileNameWithoutExt] = res.data
	}

	return types.Gtfs{
		Files:        gtfsFiles,
		FieldCounter: gtfsFieldCount,
		IdMap:        gtfsIdsMap,
	}, nil
}

// handlePrimaryKeyMapping processes the primary key mapping for a given file, header, and value.
// It handles both single and composite primary keys, updating the idsMap accordingly.
func handlePrimaryKeyMapping(primaryKey any, header string, value string, fileNameWithoutExt string, rowIndex int, idsMap *types.GtfsIdMap, idsMapMutex *sync.Mutex) {
	if primaryKey == nil {
		return
	}

	switch pk := primaryKey.(type) {
	case string:
		// Single primary key case
		if pk == header && value != "" {
			idsMapMutex.Lock()
			(*idsMap)[fileNameWithoutExt][value] = append((*idsMap)[fileNameWithoutExt][value], rowIndex)
			idsMapMutex.Unlock()
		}
	case []string:
		// Composite primary key case
		for _, key := range pk {
			if key == header && value != "" {
				idsMapMutex.Lock()
				if _, exists := (*idsMap)[fileNameWithoutExt]; !exists {
					(*idsMap)[fileNameWithoutExt] = make(map[string][]int)
				}
				// Store each component separately with a prefix to avoid collisions
				(*idsMap)[fileNameWithoutExt][value] = append((*idsMap)[fileNameWithoutExt][value], rowIndex)
				idsMapMutex.Unlock()
			}
		}
	}
}

// parseCSV parses CSV content into a slice of maps where each map represents a row
// with column headers as keys and cell values as values.
// Returns an error if the CSV is empty or cannot be parsed.
func parseCSV(content []byte, fileNameWithoutExt string, fieldCount *types.GtfsFieldCount, idsMap *types.GtfsIdMap, fieldCountMutex *sync.Mutex, idsMapMutex *sync.Mutex) ([]map[string]string, error) {
	reader := csv.NewReader(bytes.NewReader(content))
	reader.TrimLeadingSpace = true

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	if len(records) < 1 {
		return nil, errors.New("CSV file is empty")
	}

	headers := records[0]
	result := make([]map[string]string, 0, len(records)-1)

	localCounts := make(map[string]int)

	primaryKey, ok := types.GTFS_PRIMARY_KEYS[fileNameWithoutExt]

	if !ok {
		panic("primary key not found for file: " + fileNameWithoutExt)
	}

	// Initialize the inner map for this file if it doesn't exist
	idsMapMutex.Lock()
	if (*idsMap)[fileNameWithoutExt] == nil {
		(*idsMap)[fileNameWithoutExt] = make(map[string][]int)
	}
	idsMapMutex.Unlock()

	for rowIndex, row := range records[1:] {
		entry := make(map[string]string, len(headers))
		for i, value := range row {
			if i >= len(headers) {
				continue
			}
			header := headers[i]
			entry[header] = value
			if value != "" {
				localCounts[header]++
			}

			handlePrimaryKeyMapping(primaryKey, header, value, fileNameWithoutExt, rowIndex + 1, idsMap, idsMapMutex)
		}
		result = append(result, entry)
	}

	// Now update fieldCount once
	fieldCountMutex.Lock()
	if (*fieldCount)[fileNameWithoutExt] == nil {
		(*fieldCount)[fileNameWithoutExt] = make(map[string]int)
	}
	for header, count := range localCounts {
		(*fieldCount)[fileNameWithoutExt][header] += count
	}
	fieldCountMutex.Unlock()

	return result, nil
}
