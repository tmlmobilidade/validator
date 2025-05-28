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
	messages = append(messages, v.checkRequiredFiles(gtfs)...)
	messages = append(messages, v.checkStopsConditional(gtfs)...)
	messages = append(messages, v.checkCalendarFiles(gtfs)...)
	messages = append(messages, v.checkLevelsIfElevator(gtfs)...)
	messages = append(messages, v.checkFeedInfoWithTranslations(gtfs)...)
	messages = append(messages, v.checkForbiddenNetworks(gtfs)...)
	return
}

func (v *FileValidation) newMessage(file, msg string) types.Message {
	return types.Message{
		Field:        "N/A",
		Rows:         []int{},
		FileName:     file,
		Message:      msg,
		ValidationID: v.ID,
		Severity:     v.Severity,
	}
}

func (v *FileValidation) checkRequiredFiles(gtfs types.Gtfs) []types.Message {
	required := []string{"agency.txt", "routes.txt", "trips.txt", "stop_times.txt"}
	var messages []types.Message

	for _, file := range required {
		if _, exists := gtfs.Files[file[:len(file)-4]]; !exists {
			messages = append(messages, v.newMessage(file, fmt.Sprintf("Required file \"%s\" is missing", file)))
		}
	}
	return messages
}

func (v *FileValidation) checkStopsConditional(gtfs types.Gtfs) []types.Message {
	if _, hasLocations := gtfs.Files["locations"]; !hasLocations {
		if _, hasStops := gtfs.Files["stops"]; !hasStops {
			return []types.Message{
				v.newMessage("stops.txt", "stops.txt is required when locations.geojson is not present"),
			}
		}
	}
	return nil
}

func (v *FileValidation) checkCalendarFiles(gtfs types.Gtfs) []types.Message {
	_, hasCalendar := gtfs.Files["calendar"]
	_, hasDates := gtfs.Files["calendar_dates"]

	if !hasCalendar && !hasDates {
		return []types.Message{
			v.newMessage("calendar.txt", "Either calendar.txt or calendar_dates.txt must be present"),
		}
	}
	return nil
}

func (v *FileValidation) checkLevelsIfElevator(gtfs types.Gtfs) []types.Message {
	pathways, hasPathways := gtfs.Files["pathways"]
	if !hasPathways {
		return nil
	}

	for _, pathway := range pathways {
		if mode, ok := pathway["pathway_mode"]; ok && mode == "5" {
			if _, hasLevels := gtfs.Files["levels"]; !hasLevels {
				return []types.Message{
					v.newMessage("levels.txt", "levels.txt is required when pathways.txt contains elevators (pathway_mode=5)"),
				}
			}
			break
		}
	}
	return nil
}

func (v *FileValidation) checkFeedInfoWithTranslations(gtfs types.Gtfs) []types.Message {
	if _, hasTranslations := gtfs.Files["translations"]; hasTranslations {
		if _, hasFeedInfo := gtfs.Files["feed_info"]; !hasFeedInfo {
			return []types.Message{
				v.newMessage("feed_info.txt", "feed_info.txt is required when translations.txt is present"),
			}
		}
	}
	return nil
}

func (v *FileValidation) checkForbiddenNetworks(gtfs types.Gtfs) []types.Message {
	routes, hasRoutes := gtfs.Files["routes"]
	if !hasRoutes {
		return nil
	}

	for _, route := range routes {
		if _, ok := route["network_id"]; ok {
			var messages []types.Message
			if _, hasNetworks := gtfs.Files["networks"]; hasNetworks {
				messages = append(messages, v.newMessage("networks.txt", "networks.txt is forbidden when network_id exists in routes.txt"))
			}
			if _, hasRouteNetworks := gtfs.Files["route_networks"]; hasRouteNetworks {
				messages = append(messages, v.newMessage("route_networks.txt", "route_networks.txt is forbidden when network_id exists in routes.txt"))
			}
			return messages
		}
	}
	return nil
}