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
	messages = append(messages, v.checkWarningFiles(gtfs, rules)...)
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
		tableName := file[:len(file)-4]
		if gtfs.HasTable(tableName) {
			messages = append(messages, v.newMessage(file, fmt.Sprintf(i18n.AppTranslator.Get("file_validations.forbidden"), file)))
		}
	}
	return messages
}

func (v *FileValidation) checkWarningFiles(gtfs types.Gtfs, rules *types.GtfsRules) []types.Message {
	warningFromRules := services.NewRulesParser(services.AppCLI.Options.RulesPath).GetWarningFiles(rules)

	var messages []types.Message

	for _, file := range warningFromRules {
		tableName := file[:len(file)-4]
		if !gtfs.HasTable(tableName) {
			messages = append(messages, v.newMessage(file, fmt.Sprintf(i18n.AppTranslator.Get("file_validations.warning"), file)))
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
		tableName := file[:len(file)-4]
		if !gtfs.HasTable(tableName) {
			messages = append(messages, v.newMessage(file, fmt.Sprintf(i18n.AppTranslator.Get("file_validations.required"), file)))
		}
	}
	return messages
}

func (v *FileValidation) checkStopsConditional(gtfs types.Gtfs) []types.Message {
	if !gtfs.HasTable("locations") {
		if !gtfs.HasTable("stops") {
			return []types.Message{
				v.newMessage("stops.txt", i18n.AppTranslator.Get("file_validations.stops_required_when_locations_missing")),
			}
		}
	}
	return nil
}

func (v *FileValidation) checkCalendarFiles(gtfs types.Gtfs) []types.Message {
	hasCalendar := gtfs.HasTable("calendar")
	hasDates := gtfs.HasTable("calendar_dates")

	if !hasCalendar && !hasDates {
		return []types.Message{
			v.newMessage("calendar.txt", i18n.AppTranslator.Get("file_validations.calendar_files_required")),
		}
	}
	return nil
}

func (v *FileValidation) checkLevelsIfElevator(gtfs types.Gtfs) []types.Message {
	pathwayCount, err := gtfs.GetTableCount("pathways")
	if err != nil || pathwayCount == 0 {
		return nil
	}

	err = gtfs.IteratePathways(func(_ int, pathway types.PathwaysRaw) error {
		if pathway.PathwayMode == "5" {
			if !gtfs.HasTable("levels") {
				return fmt.Errorf("levels required")
			}
		}
		return nil
	})
	if err != nil && err.Error() == "levels required" {
		return []types.Message{
			v.newMessage("levels.txt", i18n.AppTranslator.Get("file_validations.levels_required_when_elevator")),
		}
	}
	return nil
}

func (v *FileValidation) checkFeedInfoWithTranslations(gtfs types.Gtfs) []types.Message {
	translationCount, err := gtfs.GetTableCount("translations")
	if err != nil || translationCount == 0 {
		return nil
	}
	feedInfoCount, err := gtfs.GetTableCount("feed_info")
	if err != nil || feedInfoCount == 0 {
		return []types.Message{
			v.newMessage("feed_info.txt", i18n.AppTranslator.Get("file_validations.feed_info_required_when_translations")),
		}
	}
	return nil
}

func (v *FileValidation) checkForbiddenNetworks(gtfs types.Gtfs) []types.Message {
	routeCount, err := gtfs.GetTableCount("routes")
	if err != nil || routeCount == 0 {
		return nil
	}

	var messages []types.Message
	err = gtfs.IterateRoutes(func(_ int, route types.RouteRaw) error {
		if route.NetworkId != "" {
			networkCount, _ := gtfs.GetTableCount("networks")
			if networkCount > 0 {
				messages = append(messages, v.newMessage("networks.txt", i18n.AppTranslator.Get("file_validations.networks_forbidden_when_network_id")))
			}
			routeNetworkCount, _ := gtfs.GetTableCount("route_networks")
			if routeNetworkCount > 0 {
				messages = append(messages, v.newMessage("route_networks.txt", i18n.AppTranslator.Get("file_validations.route_networks_forbidden_when_network_id")))
			}
			return fmt.Errorf("found network_id") // Signal to stop iteration
		}
		return nil
	})
	if err != nil && err.Error() == "found network_id" {
		return messages
	}
	return nil
}
