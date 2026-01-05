package file_validation

import (
	"database/sql"
	"fmt"
	"main/types"
	"os"
	"testing"

	_ "modernc.org/sqlite"
)

func TestFileValidation(t *testing.T) {
	tests := []struct {
		name          string
		gtfs          types.Gtfs
		wantMessages  int
		wantSeverity  types.Severity
		checkMessages func([]types.Message) bool
	}{
		{
			name: "all required files present",
			gtfs: types.Gtfs{
				Agency:    []types.AgencyRaw{{AgencyId: "1"}},
				Route:    []types.RouteRaw{{RouteId: "1", RouteType: "1"}},
				Trip:     []types.TripRaw{{TripId: "1", RouteId: "1"}},
				StopTime: []types.StopTimeRaw{{TripId: "1", StopSequence: "1"}},
				Stop:     []types.StopRaw{{StopId: "1", StopName: "1"}},
				Calendar:  []types.CalendarRaw{{ServiceId: "1"}},
			},
			wantMessages: 0,
			wantSeverity: types.SEVERITY_ERROR,
		},
		{
			name: "missing required files",
			gtfs: types.Gtfs{
				Route:   []types.RouteRaw{{RouteId: "1", RouteType: "1"}},
				Trip:    []types.TripRaw{{TripId: "1", RouteId: "1"}},
				StopTime: []types.StopTimeRaw{{TripId: "1", StopSequence: "1"}},
				Stop:     []types.StopRaw{{StopId: "1", StopName: "1"}},
				Calendar:  []types.CalendarRaw{{ServiceId: "1"}},
			},
			wantMessages: 3, // missing agency.txt, trips.txt, stop_times.txt
			wantSeverity: types.SEVERITY_ERROR,
			checkMessages: func(messages []types.Message) bool {
				requiredFiles := map[string]bool{
					"agency.txt":     false,
					"trips.txt":      false,
					"stop_times.txt": false,
				}

				for _, msg := range messages {
					if _, ok := requiredFiles[msg.FileName]; ok {
						requiredFiles[msg.FileName] = true
					}
				}

				for _, found := range requiredFiles {
					if !found {
						return false
					}
				}
				return true
			},
		},
		{
			name: "missing stops.txt without locations.geojson",
			gtfs: types.Gtfs{
				Agency:    []types.AgencyRaw{{AgencyId: "1"}},
				Route:     []types.RouteRaw{{RouteId: "1", RouteType: "1"}},
				Trip:      []types.TripRaw{{TripId: "1", RouteId: "1"}},
				StopTime:  []types.StopTimeRaw{{TripId: "1", StopSequence: "1"}},
				Stop:      []types.StopRaw{{StopId: "1", StopName: "1"}},
				Calendar:  []types.CalendarRaw{{ServiceId: "1"}},
			},
			wantMessages: 1,
			wantSeverity: types.SEVERITY_ERROR,
			checkMessages: func(messages []types.Message) bool {
				for _, msg := range messages {
					if msg.FileName == "stops.txt" {
						return true
					}
				}
				return false
			},
		},
		{
			name: "missing stops.txt with locations.geojson is valid",
			gtfs: types.Gtfs{
				Agency:    []types.AgencyRaw{{AgencyId: "1"}},
				Route:     []types.RouteRaw{{RouteId: "1", RouteType: "1"}},
				Trip:      []types.TripRaw{{TripId: "1", RouteId: "1"}},
				StopTime:  []types.StopTimeRaw{{TripId: "1", StopSequence: "1"}},
				Stop:      []types.StopRaw{{StopId: "1", StopName: "1"}},
				Calendar:  []types.CalendarRaw{{ServiceId: "1"}},
				// Location: []types.LocationRaw{{LocationId: "1", LocationName: "1"}},
			},
			wantMessages: 0,
			wantSeverity: types.SEVERITY_ERROR,
		},
		{
			name: "missing both calendar.txt and calendar_dates.txt",
			gtfs: types.Gtfs{
				Agency:    []types.AgencyRaw{{AgencyId: "1"}},
				Route:     []types.RouteRaw{{RouteId: "1", RouteType: "1"}},
				Trip:      []types.TripRaw{{TripId: "1", RouteId: "1"}},
				StopTime:  []types.StopTimeRaw{{TripId: "1", StopSequence: "1"}},
				Stop:      []types.StopRaw{{StopId: "1", StopName: "1"}},
				Calendar:  []types.CalendarRaw{{ServiceId: "1"}},
			},
			wantMessages: 1, // One for either required, one for calendar_dates required when no calendar
			wantSeverity: types.SEVERITY_ERROR,
			checkMessages: func(messages []types.Message) bool {
				for _, msg := range messages {
					if msg.FileName == "calendar.txt" && msg.Message == "Either calendar.txt or calendar_dates.txt must be present" {
						return true
					}
				}

				return false
			},
		},
		{
			name: "levels.txt required with elevator pathways",
			gtfs: types.Gtfs{
				Agency:    []types.AgencyRaw{{AgencyId: "1"}},
				Route:     []types.RouteRaw{{RouteId: "1", RouteType: "1"}},
				Trip:      []types.TripRaw{{TripId: "1", RouteId: "1"}},
				StopTime:  []types.StopTimeRaw{{TripId: "1", StopSequence: "1"}},
				Stop:      []types.StopRaw{{StopId: "1", StopName: "1"}},
				Calendar:  []types.CalendarRaw{{ServiceId: "1"}},
				Pathways:  []types.PathwaysRaw{{PathwayMode: "5"}}, // elevator
			},
			wantMessages: 1,
			wantSeverity: types.SEVERITY_ERROR,
			checkMessages: func(messages []types.Message) bool {
				for _, msg := range messages {
					if msg.FileName == "levels.txt" {
						return true
					}
				}
				return false
			},
		},
		{
			name: "feed_info.txt required with translations.txt",
			gtfs: types.Gtfs{
				Agency:    []types.AgencyRaw{{AgencyId: "1"}},
				Route:     []types.RouteRaw{{RouteId: "1", RouteType: "1"}},
				Trip:      []types.TripRaw{{TripId: "1", RouteId: "1"}},
				StopTime:  []types.StopTimeRaw{{TripId: "1", StopSequence: "1"}},
				Stop:      []types.StopRaw{{StopId: "1", StopName: "1"}},
				Calendar:  []types.CalendarRaw{{ServiceId: "1"}},
				Translations: []types.TranslationsRaw{{Translation: "1"}},
			},
			wantMessages: 1,
			wantSeverity: types.SEVERITY_ERROR,
			checkMessages: func(messages []types.Message) bool {
				for _, msg := range messages {
					if msg.FileName == "feed_info.txt" {
						return true
					}
				}
				return false
			},
		},
		{
			name: "networks.txt forbidden with network_id in routes",
			gtfs: types.Gtfs{
				Agency:    []types.AgencyRaw{{AgencyId: "1"}},
				Route:     []types.RouteRaw{{RouteId: "1", RouteType: "1", NetworkId: "net1"}},
				Trip:      []types.TripRaw{{TripId: "1", RouteId: "1"}},
				StopTime:  []types.StopTimeRaw{{TripId: "1", StopSequence: "1"}},
				Stop:      []types.StopRaw{{StopId: "1", StopName: "1"}},
				Calendar:  []types.CalendarRaw{{ServiceId: "1"}},
				RouteNetwork:  []types.RouteNetworkRaw{{NetworkId: "1"}},
			},
			wantMessages: 1,
			wantSeverity: types.SEVERITY_ERROR,
			checkMessages: func(messages []types.Message) bool {
				for _, msg := range messages {
					if msg.FileName == "networks.txt" {
						return true
					}
				}
				return false
			},
		},
		{
			name: "route_networks.txt forbidden with network_id in routes",
			gtfs: types.Gtfs{
				Agency:    []types.AgencyRaw{{AgencyId: "1"}},
				Route:     []types.RouteRaw{{RouteId: "1", RouteType: "1", NetworkId: "net1"}},
				Trip:      []types.TripRaw{{TripId: "1", RouteId: "1"}},
				StopTime:  []types.StopTimeRaw{{TripId: "1", StopSequence: "1"}},
				Stop:      []types.StopRaw{{StopId: "1", StopName: "1"}},
				Calendar:  []types.CalendarRaw{{ServiceId: "1"}},
				RouteNetwork:  []types.RouteNetworkRaw{{NetworkId: "1"}},
			},
			wantMessages: 1,
			wantSeverity: types.SEVERITY_ERROR,
			checkMessages: func(messages []types.Message) bool {
				for _, msg := range messages {
					if msg.FileName == "route_networks.txt" {
						return true
					}
				}
				return false
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NewFileValidation(&tt.wantSeverity)
			messages := v.Validate(tt.gtfs, nil)

			if len(messages) != tt.wantMessages {
				t.Errorf("[%v] FileValidation.Validate() got %v messages, want %v", tt.name, len(messages), tt.wantMessages)
				for _, msg := range messages {
					fmt.Println(msg)
				}
			}

			for _, msg := range messages {
				if msg.Severity != tt.wantSeverity {
					t.Errorf("[%v] FileValidation.Validate() got message with severity %v, want %v", tt.name, msg.Severity, tt.wantSeverity)
				}
				if msg.ValidationID != v.ID {
					t.Errorf("[%v] FileValidation.Validate() got message with validation ID %v, want %v", tt.name, msg.ValidationID, v.ID)
				}
			}

			if tt.checkMessages != nil && !tt.checkMessages(messages) {
				t.Errorf("[%v] FileValidation.Validate() messages did not match expected conditions", tt.name)
			}
		})
	}
}

// createMockGtfsDB creates a temporary SQLite database with specified tables
func createMockGtfsDB(t *testing.T, tables []string) (*types.Gtfs, func()) {
	tmpDB, err := os.CreateTemp("", "test_gtfs_*.db")
	if err != nil {
		t.Fatalf("Failed to create temp database: %v", err)
	}
	tmpDBPath := tmpDB.Name()
	tmpDB.Close()

	db, err := sql.Open("sqlite", tmpDBPath)
	if err != nil {
		os.Remove(tmpDBPath)
		t.Fatalf("Failed to open database: %v", err)
	}

	// Create the specified tables
	for _, table := range tables {
		_, err := db.Exec(fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (id TEXT)", table))
		if err != nil {
			db.Close()
			os.Remove(tmpDBPath)
			t.Fatalf("Failed to create table %s: %v", table, err)
		}
	}

	gtfs := types.NewGtfsFromSQLite(db, tmpDBPath)

	cleanup := func() {
		gtfs.Close()
		os.Remove(tmpDBPath)
	}

	return gtfs, cleanup
}

func TestVehiclesRequiredWithRules(t *testing.T) {
	tests := []struct {
		name          string
		tables        []string // Tables to create in the mock database
		rules         *types.GtfsRules
		wantMessages  int
		wantSeverity  types.Severity
		checkMessages func([]types.Message) bool
	}{
		{
			name:   "vehicles.txt required when rules specify error - missing vehicles",
			tables: []string{"agency", "routes", "trips", "stop_times", "stops", "calendar"},
			rules: &types.GtfsRules{
				Vehicles: types.VehiclesRules{
					File: types.SEVERITY_ERROR,
				},
			},
			wantMessages: 1,
			wantSeverity: types.SEVERITY_ERROR,
			checkMessages: func(messages []types.Message) bool {
				for _, msg := range messages {
					if msg.FileName == "vehicles.txt" {
						return true
					}
				}
				return false
			},
		},
		{
			name:   "vehicles.txt required when rules specify error - vehicles present",
			tables: []string{"agency", "routes", "trips", "stop_times", "stops", "calendar", "vehicles"},
			rules: &types.GtfsRules{
				Vehicles: types.VehiclesRules{
					File: types.SEVERITY_ERROR,
				},
			},
			wantMessages: 0,
			wantSeverity: types.SEVERITY_ERROR,
		},
		{
			name:   "vehicles.txt not required when rules specify ignore",
			tables: []string{"agency", "routes", "trips", "stop_times", "stops", "calendar"},
			rules: &types.GtfsRules{
				Vehicles: types.VehiclesRules{
					File: types.SEVERITY_IGNORE,
				},
			},
			wantMessages: 0,
			wantSeverity: types.SEVERITY_ERROR,
		},
		{
			name:   "vehicles.txt forbidden when rules specify forbidden - vehicles present",
			tables: []string{"agency", "routes", "trips", "stop_times", "stops", "calendar", "vehicles"},
			rules: &types.GtfsRules{
				Vehicles: types.VehiclesRules{
					File: types.SEVERITY_FORBIDDEN,
				},
			},
			wantMessages: 1,
			wantSeverity: types.SEVERITY_ERROR,
			checkMessages: func(messages []types.Message) bool {
				for _, msg := range messages {
					if msg.FileName == "vehicles.txt" {
						return true
					}
				}
				return false
			},
		},
		{
			name:         "no error when vehicles.txt not required and not present",
			tables:       []string{"agency", "routes", "trips", "stop_times", "stops", "calendar"},
			rules:        nil,
			wantMessages: 0,
			wantSeverity: types.SEVERITY_ERROR,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gtfs, cleanup := createMockGtfsDB(t, tt.tables)
			defer cleanup()

			v := NewFileValidation(&tt.wantSeverity)
			messages := v.Validate(*gtfs, tt.rules)

			if len(messages) != tt.wantMessages {
				t.Errorf("[%v] FileValidation.Validate() got %v messages, want %v", tt.name, len(messages), tt.wantMessages)
				for _, msg := range messages {
					fmt.Println(msg)
				}
			}

			for _, msg := range messages {
				if msg.Severity != tt.wantSeverity {
					t.Errorf("[%v] FileValidation.Validate() got message with severity %v, want %v", tt.name, msg.Severity, tt.wantSeverity)
				}
				if msg.ValidationID != v.ID {
					t.Errorf("[%v] FileValidation.Validate() got message with validation ID %v, want %v", tt.name, msg.ValidationID, v.ID)
				}
			}

			if tt.checkMessages != nil && !tt.checkMessages(messages) {
				t.Errorf("[%v] FileValidation.Validate() messages did not match expected conditions", tt.name)
			}
		})
	}
}
