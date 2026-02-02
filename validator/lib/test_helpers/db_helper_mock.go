package test_helpers

import (
	"fmt"
	"main/types"
)

// MockGtfs is a generic mock implementation of types.Gtfs for testing.
// It can be used to create fake GTFS data without requiring a database connection.
type MockGtfs struct {
	// TableCounts maps table names to their row counts.
	// If not set, counts will be calculated from IdMapData.
	TableCounts map[string]int

	// IdMapData contains the ID mapping structure: table -> id -> []row_indices
	IdMapData types.GtfsIdMap
}

// ToGtfs converts the MockGtfs to a types.Gtfs struct.
// It properly initializes IdMap and creates slice fields with appropriate lengths
// based on TableCounts or calculated from IdMapData.
func (m MockGtfs) ToGtfs() types.Gtfs {
	// Initialize IdMap by copying all tables from IdMapData
	idMap := make(types.GtfsIdMap)
	if m.IdMapData != nil {
		for table, ids := range m.IdMapData {
			// Deep copy the map
			idMap[table] = make(map[string][]int)
			for id, rows := range ids {
				// Copy the slice
				idMap[table][id] = make([]int, len(rows))
				copy(idMap[table][id], rows)
			}
		}
	}

	// Calculate table counts if not provided
	tableCounts := make(map[string]int)
	if m.TableCounts != nil {
		for table, count := range m.TableCounts {
			tableCounts[table] = count
		}
	}

	// Calculate missing counts from IdMap
	if m.IdMapData != nil {
		for table, ids := range m.IdMapData {
			if _, exists := tableCounts[table]; !exists {
				// Count unique row indices across all IDs in this table
				uniqueRows := make(map[int]bool)
				for _, rows := range ids {
					for _, row := range rows {
						uniqueRows[row] = true
					}
				}
				tableCounts[table] = len(uniqueRows)
			}
		}
	}

	// Create Gtfs struct with appropriate slice lengths
	gtfs := types.Gtfs{
		IdMap: idMap,
	}

	// Set slice fields based on table counts
	// This ensures len(gtfs.Agency), len(gtfs.Route), etc. work correctly
	if count, ok := tableCounts["agency"]; ok {
		gtfs.Agency = make([]types.AgencyRaw, count)
	}
	if count, ok := tableCounts["routes"]; ok {
		gtfs.Route = make([]types.RouteRaw, count)
	}
	if count, ok := tableCounts["stops"]; ok {
		gtfs.Stop = make([]types.StopRaw, count)
	}
	if count, ok := tableCounts["trips"]; ok {
		gtfs.Trip = make([]types.TripRaw, count)
	}
	if count, ok := tableCounts["stop_times"]; ok {
		gtfs.StopTime = make([]types.StopTimeRaw, count)
	}
	if count, ok := tableCounts["calendar"]; ok {
		gtfs.Calendar = make([]types.CalendarRaw, count)
	}
	if count, ok := tableCounts["calendar_dates"]; ok {
		gtfs.CalendarDates = make([]types.CalendarDatesRaw, count)
	}
	if count, ok := tableCounts["route_networks"]; ok {
		gtfs.RouteNetwork = make([]types.RouteNetworkRaw, count)
	}

	return gtfs
}

// GetTableCount returns the number of rows in a table.
// It first checks TableCounts, then calculates from IdMapData if needed.
func (m MockGtfs) GetTableCount(table string) (int, error) {
	// Check if explicitly set
	if m.TableCounts != nil {
		if count, ok := m.TableCounts[table]; ok {
			return count, nil
		}
	}

	// Calculate from IdMapData
	if m.IdMapData != nil {
		if tableMap, ok := m.IdMapData[table]; ok {
			uniqueRows := make(map[int]bool)
			for _, rows := range tableMap {
				for _, row := range rows {
					uniqueRows[row] = true
				}
			}
			return len(uniqueRows), nil
		}
	}

	return 0, fmt.Errorf("table %s not found in mock", table)
}

// GetRowsById returns the row indices for a given table and ID.
// This matches the signature of types.Gtfs.GetRowsById which returns []int.
func (m MockGtfs) GetRowsById(table, id string) ([]int, error) {
	if m.IdMapData != nil {
		if tableMap, ok := m.IdMapData[table]; ok {
			if indices, found := tableMap[id]; found {
				// Return a copy of the slice
				result := make([]int, len(indices))
				copy(result, indices)
				return result, nil
			}
		}
	}
	return []int{}, nil
}

// IdMap returns the IdMapData directly.
// This is provided for convenience when you need direct access to the IdMap.
func (m MockGtfs) IdMap() types.GtfsIdMap {
	return m.IdMapData
}
