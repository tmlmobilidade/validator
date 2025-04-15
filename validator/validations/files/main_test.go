package file_validation

import (
	"fmt"
	"main/validator/types"
	"testing"
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
				Files: types.GtfsFiles{
					"agency":     []map[string]string{{"id": "1"}},
					"routes":     []map[string]string{{"id": "1"}},
					"trips":      []map[string]string{{"id": "1"}},
					"stop_times": []map[string]string{{"id": "1"}},
					"stops":      []map[string]string{{"id": "1"}},
					"calendar":   []map[string]string{{"id": "1"}},
				},
			},
			wantMessages: 0,
			wantSeverity: types.SEVERITY_ERROR,
		},
		{
			name: "missing required files",
			gtfs: types.Gtfs{
				Files: types.GtfsFiles{
					"routes":   []map[string]string{{"id": "1"}},
					"calendar": []map[string]string{{"id": "1"}},
					"stops":    []map[string]string{{"id": "1"}},
				},
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
				Files: types.GtfsFiles{
					"agency":     []map[string]string{{"id": "1"}},
					"routes":     []map[string]string{{"id": "1"}},
					"trips":      []map[string]string{{"id": "1"}},
					"stop_times": []map[string]string{{"id": "1"}},
					"calendar":   []map[string]string{{"id": "1"}},
				},
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
				Files: types.GtfsFiles{
					"agency":     []map[string]string{{"id": "1"}},
					"routes":     []map[string]string{{"id": "1"}},
					"trips":      []map[string]string{{"id": "1"}},
					"stop_times": []map[string]string{{"id": "1"}},
					"calendar":   []map[string]string{{"id": "1"}},
					"locations":  []map[string]string{{"id": "1"}},
				},
			},
			wantMessages: 0,
			wantSeverity: types.SEVERITY_ERROR,
		},
		{
			name: "missing both calendar.txt and calendar_dates.txt",
			gtfs: types.Gtfs{
				Files: types.GtfsFiles{
					"agency":     []map[string]string{{"id": "1"}},
					"routes":     []map[string]string{{"id": "1"}},
					"trips":      []map[string]string{{"id": "1"}},
					"stop_times": []map[string]string{{"id": "1"}},
					"stops":      []map[string]string{{"id": "1"}},
				},
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
				Files: types.GtfsFiles{
					"agency":     []map[string]string{{"id": "1"}},
					"routes":     []map[string]string{{"id": "1"}},
					"trips":      []map[string]string{{"id": "1"}},
					"stop_times": []map[string]string{{"id": "1"}},
					"stops":      []map[string]string{{"id": "1"}},
					"calendar":   []map[string]string{{"id": "1"}},
					"pathways":   []map[string]string{{"pathway_mode": "5"}}, // elevator
				},
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
				Files: types.GtfsFiles{
					"agency":       []map[string]string{{"id": "1"}},
					"routes":       []map[string]string{{"id": "1"}},
					"trips":        []map[string]string{{"id": "1"}},
					"stop_times":   []map[string]string{{"id": "1"}},
					"stops":        []map[string]string{{"id": "1"}},
					"calendar":     []map[string]string{{"id": "1"}},
					"translations": []map[string]string{{"id": "1"}},
				},
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
				Files: types.GtfsFiles{
					"agency":     []map[string]string{{"id": "1"}},
					"routes":     []map[string]string{{"id": "1", "network_id": "net1"}},
					"trips":      []map[string]string{{"id": "1"}},
					"stop_times": []map[string]string{{"id": "1"}},
					"stops":      []map[string]string{{"id": "1"}},
					"calendar":   []map[string]string{{"id": "1"}},
					"networks":   []map[string]string{{"id": "1"}},
				},
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
				Files: types.GtfsFiles{
					"agency":         []map[string]string{{"id": "1"}},
					"routes":         []map[string]string{{"id": "1", "network_id": "net1"}},
					"trips":          []map[string]string{{"id": "1"}},
					"stop_times":     []map[string]string{{"id": "1"}},
					"stops":          []map[string]string{{"id": "1"}},
					"calendar":       []map[string]string{{"id": "1"}},
					"route_networks": []map[string]string{{"id": "1"}},
				},
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
			messages := v.Validate(tt.gtfs)

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
