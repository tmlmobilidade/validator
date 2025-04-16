package file_validation

import (
	"fmt"
	"main/types"
)

type FileValidation struct {
	*types.Validation
}

func NewFileValidation(severity *types.Severity) *FileValidation {
	s := types.SEVERITY_ERROR
	if severity != nil {
		s = *severity
	}

	return &FileValidation{
		Validation: &types.Validation{
			ID:          "file_validation",
			Description: "Validate file data",
			Severity:    s,
		},
	}
}

func (v *FileValidation) Validate(gtfs types.Gtfs) (messages []types.Message) {
	// Required files
	requiredFiles := []string{
		"agency.txt",
		"routes.txt",
		"trips.txt",
		"stop_times.txt",
	}

	// Check required files
	for _, file := range requiredFiles {
		if _, exists := gtfs.Files[file[:len(file)-4]]; !exists {
			messages = append(messages, types.Message{
				Field:        "N/A",
				FileName:     file,
				Message:      fmt.Sprintf("Required file \"%s\" is missing", file),
				ValidationID: v.ID,
				Severity:     v.Severity,
			})
		}
	}

	// Check conditionally required files

	// Check stops.txt - Required unless demand-responsive zones are defined in locations.geojson
	if _, hasLocations := gtfs.Files["locations"]; !hasLocations {
		if _, hasStops := gtfs.Files["stops"]; !hasStops {
			messages = append(messages, types.Message{
				Field:        "N/A",
				FileName:     "stops.txt",
				Message:      "stops.txt is required when locations.geojson is not present",
				ValidationID: v.ID,
				Severity:     v.Severity,
			})
		}
	}

	// Check calendar.txt and calendar_dates.txt
	_, hasCalendar := gtfs.Files["calendar"]
	_, hasCalendarDates := gtfs.Files["calendar_dates"]

	if !hasCalendar && !hasCalendarDates {
		messages = append(messages, types.Message{
			Field:        "N/A",
			FileName:     "calendar.txt",
			Message:      "Either calendar.txt or calendar_dates.txt must be present",
			ValidationID: v.ID,
			Severity:     v.Severity,
		})
	}

	// Check levels.txt - Required when describing pathways with elevators
	if pathways, hasPathways := gtfs.Files["pathways"]; hasPathways {
		hasElevator := false
		for _, pathway := range pathways {
			if pathwayMode, ok := pathway["pathway_mode"]; ok && pathwayMode == "5" { // 5 = elevator
				hasElevator = true
				break
			}
		}

		if hasElevator {
			if _, hasLevels := gtfs.Files["levels"]; !hasLevels {
				messages = append(messages, types.Message{
					Field:        "N/A",
					FileName:     "levels.txt",
					Message:      "levels.txt is required when pathways.txt contains elevators (pathway_mode=5)",
					ValidationID: v.ID,
					Severity:     v.Severity,
				})
			}
		}
	}

	// Check feed_info.txt - Required if translations.txt is provided
	if _, hasTranslations := gtfs.Files["translations"]; hasTranslations {
		if _, hasFeedInfo := gtfs.Files["feed_info"]; !hasFeedInfo {
			messages = append(messages, types.Message{
				Field:        "N/A",
				FileName:     "feed_info.txt",
				Message:      "feed_info.txt is required when translations.txt is present",
				ValidationID: v.ID,
				Severity:     v.Severity,
			})
		}
	}

	// Check networks.txt and route_networks.txt - Forbidden if network_id exists in routes.txt
	if routes, hasRoutes := gtfs.Files["routes"]; hasRoutes {
		hasNetworkId := false
		for _, route := range routes {
			if _, ok := route["network_id"]; ok {
				hasNetworkId = true
				break
			}
		}

		if hasNetworkId {
			if _, hasNetworks := gtfs.Files["networks"]; hasNetworks {
				messages = append(messages, types.Message{
					Field:        "N/A",
					FileName:     "networks.txt",
					Message:      "networks.txt is forbidden when network_id exists in routes.txt",
					ValidationID: v.ID,
					Severity:     v.Severity,
				})
			}
			if _, hasRouteNetworks := gtfs.Files["route_networks"]; hasRouteNetworks {
				messages = append(messages, types.Message{
					Field:        "N/A",
					FileName:     "route_networks.txt",
					Message:      "route_networks.txt is forbidden when network_id exists in routes.txt",
					ValidationID: v.ID,
					Severity:     v.Severity,
				})
			}
		}
	}

	return messages
}
