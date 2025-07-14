package file_validation

import (
	"fmt"
	"main/i18n"
	"main/lib"
	"main/services"
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
			Severity:    &s,
		},
	}
}

func (v *FileValidation) Validate(gtfs types.Gtfs, rules *types.GtfsRules) (messages []types.Message) {
	messages = append(messages, v.checkForbiddenFiles(gtfs, rules)...)
	messages = append(messages, v.checkRequiredFiles(gtfs, rules)...)
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
		Severity:     v.GetSeverity(),
	}
}

func (v *FileValidation) checkForbiddenFiles(gtfs types.Gtfs, rules *types.GtfsRules) []types.Message {
	var messages []types.Message

	forbiddenFiles := services.NewRulesParser(services.AppCLI.Options.RulesPath).GetForbiddenFiles(rules)

	for _, file := range forbiddenFiles {
		if _, exists := gtfs.IdMap[file[:len(file)-4]]; exists {
			messages = append(messages, v.newMessage(file, fmt.Sprintf(i18n.AppTranslator.Get("file_validations.forbidden"), file)))
		}
	}
	return messages
}

func (v *FileValidation) checkRequiredFiles(gtfs types.Gtfs, rules *types.GtfsRules) []types.Message {
	minRequired := []string{"agency.txt", "routes.txt", "trips.txt", "stop_times.txt"}
	requiredFromRules := services.NewRulesParser(services.AppCLI.Options.RulesPath).GetRequiredFiles(rules)

	mergedRequired := lib.RemoveDuplicates(append(minRequired, requiredFromRules...))

	var messages []types.Message

	for _, file := range mergedRequired {
		if _, exists := gtfs.IdMap[file[:len(file)-4]]; !exists {
			messages = append(messages, v.newMessage(file, fmt.Sprintf(i18n.AppTranslator.Get("file_validations.required"), file)))
		}
	}
	return messages
}

func (v *FileValidation) checkStopsConditional(gtfs types.Gtfs) []types.Message {
	if _, hasLocations := gtfs.IdMap["locations"]; !hasLocations {
		if _, hasStops := gtfs.IdMap["stops"]; !hasStops {
			return []types.Message{
				v.newMessage("stops.txt", i18n.AppTranslator.Get("file_validations.stops_required_when_locations_missing")),
			}
		}
	}
	return nil
}

func (v *FileValidation) checkCalendarFiles(gtfs types.Gtfs) []types.Message {
	_, hasCalendar := gtfs.IdMap["calendar"]
	_, hasDates := gtfs.IdMap["calendar_dates"]

	if !hasCalendar && !hasDates {
		return []types.Message{
			v.newMessage("calendar.txt", i18n.AppTranslator.Get("file_validations.calendar_files_required")),
		}
	}
	return nil
}

func (v *FileValidation) checkLevelsIfElevator(gtfs types.Gtfs) []types.Message {
	if len(gtfs.Pathways) == 0 {
		return nil
	}

	for _, pathway := range gtfs.Pathways {
		if pathway.PathwayMode == "5" {
			if _, hasLevels := gtfs.IdMap["levels"]; !hasLevels {
				return []types.Message{
					v.newMessage("levels.txt", i18n.AppTranslator.Get("file_validations.levels_required_when_elevator")),
				}
			}
			break
		}
	}
	return nil
}

func (v *FileValidation) checkFeedInfoWithTranslations(gtfs types.Gtfs) []types.Message {
	if len(gtfs.Translations) > 0 {
		if len(gtfs.FeedInfo) == 0 {
			return []types.Message{
				v.newMessage("feed_info.txt", i18n.AppTranslator.Get("file_validations.feed_info_required_when_translations")),
			}
		}
	}
	return nil
}

func (v *FileValidation) checkForbiddenNetworks(gtfs types.Gtfs) []types.Message {
	if len(gtfs.Route) == 0 {
		return nil
	}

	for _, route := range gtfs.Route {
		if route.NetworkId != "" {
			var messages []types.Message
			if len(gtfs.Network) > 0 {
				messages = append(messages, v.newMessage("networks.txt", i18n.AppTranslator.Get("file_validations.networks_forbidden_when_network_id")))
			}
			if len(gtfs.RouteNetwork) > 0 {
				messages = append(messages, v.newMessage("route_networks.txt", i18n.AppTranslator.Get("file_validations.route_networks_forbidden_when_network_id")))
			}
			return messages
		}
	}
	return nil
}
