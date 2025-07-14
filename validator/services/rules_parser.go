package services

import (
	"encoding/json"
	"fmt"
	"io"
	"main/lib"
	"main/types"
	"os"
)

// RulesParser handles parsing of GTFS validation rules from JSON files
type RulesParser struct {
	rulesPath string
}

// NewRulesParser creates a new rules parser with the specified rules file path
func NewRulesParser(rulesPath string) *RulesParser {
	return &RulesParser{
		rulesPath: rulesPath,
	}
}

// ParseRules reads and parses the rules JSON file into a GtfsRules structure
func (rp *RulesParser) ParseRules() (*types.GtfsRules, error) {
	if rp.rulesPath == "" {
		return nil, nil
	}

	// Check if file exists
	if _, err := os.Stat(rp.rulesPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("rules file does not exist: %s", rp.rulesPath)
	}

	// Open and read the file
	file, err := os.Open(rp.rulesPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open rules file: %w", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read rules file: %w", err)
	}

	// Parse JSON into GtfsRules structure
	var rules types.GtfsRules
	if err := json.Unmarshal(data, &rules); err != nil {
		return nil, fmt.Errorf("failed to parse rules JSON: %w", err)
	}

	// Validate the parsed rules
	if err := rp.validateRules(&rules); err != nil {
		return nil, fmt.Errorf("rules validation failed: %w", err)
	}

	lib.AppLogger.Info(fmt.Sprintf("Successfully parsed rules from: %s", rp.rulesPath))
	return &rules, nil
}

// validateRules performs validation on the parsed rules to ensure they are correct
func (rp *RulesParser) validateRules(rules *types.GtfsRules) error {
	// Validate severity values across all rule configurations
	validationErrors := []string{}

	// Helper function to validate a single rule config
	validateRuleConfig := func(config types.RuleConfig, fieldName string) {
		if !isValidSeverity(config.Severity) {
			validationErrors = append(validationErrors, fmt.Sprintf("Invalid severity '%s' for field '%s'", config.Severity, fieldName))
		}
	}

	// Validate agency rules
	if rules.Agency.File != types.SEVERITY_FORBIDDEN {
		validateRuleConfig(rules.Agency.AgencyId, "agency.agency_id")
		validateRuleConfig(rules.Agency.AgencyName, "agency.agency_name")
		validateRuleConfig(rules.Agency.AgencyUrl, "agency.agency_url")
		validateRuleConfig(rules.Agency.AgencyTimezone, "agency.agency_timezone")
		validateRuleConfig(rules.Agency.AgencyLang, "agency.agency_lang")
		validateRuleConfig(rules.Agency.AgencyPhone, "agency.agency_phone")
		validateRuleConfig(rules.Agency.AgencyFare, "agency.agency_fare_url")
		validateRuleConfig(rules.Agency.AgencyEmail, "agency.agency_email")
	}

	// Validate stops rules
	if rules.Stops.File != types.SEVERITY_FORBIDDEN {
		validateRuleConfig(rules.Stops.StopId, "stops.stop_id")
		validateRuleConfig(rules.Stops.StopCode, "stops.stop_code")
		validateRuleConfig(rules.Stops.StopName, "stops.stop_name")
		validateRuleConfig(rules.Stops.StopShortName, "stops.stop_short_name")
		validateRuleConfig(rules.Stops.TtsStopName, "stops.tts_stop_name")
		validateRuleConfig(rules.Stops.StopDesc, "stops.stop_desc")
		validateRuleConfig(rules.Stops.StopLat, "stops.stop_lat")
		validateRuleConfig(rules.Stops.StopLon, "stops.stop_lon")
		validateRuleConfig(rules.Stops.ZoneId, "stops.zone_id")
		validateRuleConfig(rules.Stops.StopUrl, "stops.stop_url")
		validateRuleConfig(rules.Stops.LocationType, "stops.location_type")
		validateRuleConfig(rules.Stops.ParentStation, "stops.parent_station")
		validateRuleConfig(rules.Stops.StopTimezone, "stops.stop_timezone")
		validateRuleConfig(rules.Stops.WheelchairBoarding, "stops.wheelchair_boarding")
		validateRuleConfig(rules.Stops.LevelId, "stops.level_id")
		validateRuleConfig(rules.Stops.PlatformCode, "stops.platform_code")
		validateRuleConfig(rules.Stops.PublicVisible, "stops.public_visible")
		validateRuleConfig(rules.Stops.HasStopSign, "stops.has_stop_sign")
		validateRuleConfig(rules.Stops.HasShelter, "stops.has_shelter")
		validateRuleConfig(rules.Stops.ShelterCode, "stops.shelter_code")
		validateRuleConfig(rules.Stops.ShelterMaintainer, "stops.shelter_maintainer")
		validateRuleConfig(rules.Stops.HasBench, "stops.has_bench")
		validateRuleConfig(rules.Stops.HasNetworkMap, "stops.has_network_map")
		validateRuleConfig(rules.Stops.HasSchedules, "stops.has_schedules")
		validateRuleConfig(rules.Stops.HasPipRealTime, "stops.has_pip_real_time")
		validateRuleConfig(rules.Stops.HasTariffsInformation, "stops.has_tariffs_information")
		validateRuleConfig(rules.Stops.RegionId, "stops.region_id")
		validateRuleConfig(rules.Stops.MunicipalityId, "stops.municipality_id")
		validateRuleConfig(rules.Stops.ParishId, "stops.parish_id")
	}

	// Validate routes rules
	if rules.Routes.File != types.SEVERITY_FORBIDDEN {
		validateRuleConfig(rules.Routes.LineId, "routes.line_id")
		validateRuleConfig(rules.Routes.LineShortName, "routes.line_short_name")
		validateRuleConfig(rules.Routes.LineLongName, "routes.line_long_name")
		validateRuleConfig(rules.Routes.RouteId, "routes.route_id")
		validateRuleConfig(rules.Routes.AgencyId, "routes.agency_id")
		validateRuleConfig(rules.Routes.RouteShortName, "routes.route_short_name")
		validateRuleConfig(rules.Routes.RouteLongName, "routes.route_long_name")
		validateRuleConfig(rules.Routes.RouteDesc, "routes.route_desc")
		validateRuleConfig(rules.Routes.RouteRemarks, "routes.route_remarks")
		validateRuleConfig(rules.Routes.RouteType, "routes.route_type")
		validateRuleConfig(rules.Routes.PathType, "routes.path_type")
		validateRuleConfig(rules.Routes.Circular, "routes.circular")
		validateRuleConfig(rules.Routes.School, "routes.school")
		validateRuleConfig(rules.Routes.RouteUrl, "routes.route_url")
		validateRuleConfig(rules.Routes.RouteColor, "routes.route_color")
		validateRuleConfig(rules.Routes.RouteTextColor, "routes.route_text_color")
		validateRuleConfig(rules.Routes.ContinuousPickup, "routes.continuous_pickup")
		validateRuleConfig(rules.Routes.ContinuousDropOff, "routes.continuous_drop_off")
	}

	// Add validation for other rule types (trips, stop_times, calendar, etc.)
	if rules.Trips.File != types.SEVERITY_FORBIDDEN {
		validateRuleConfig(rules.Trips.RouteId, "trips.route_id")
		validateRuleConfig(rules.Trips.PatternId, "trips.pattern_id")
		validateRuleConfig(rules.Trips.ServiceId, "trips.service_id")
		validateRuleConfig(rules.Trips.TripId, "trips.trip_id")
		validateRuleConfig(rules.Trips.TripHeadsign, "trips.trip_headsign")
		validateRuleConfig(rules.Trips.TripShortName, "trips.trip_short_name")
		validateRuleConfig(rules.Trips.DirectionId, "trips.direction_id")
		validateRuleConfig(rules.Trips.BlockId, "trips.block_id")
		validateRuleConfig(rules.Trips.ShapeId, "trips.shape_id")
		validateRuleConfig(rules.Trips.WheelchairAccessible, "trips.wheelchair_accessible")
		validateRuleConfig(rules.Trips.BikesAllowed, "trips.bikes_allowed")
	}

	if rules.StopTimes.File != types.SEVERITY_FORBIDDEN {
		validateRuleConfig(rules.StopTimes.TripId, "stop_times.trip_id")
		validateRuleConfig(rules.StopTimes.ArrivalTime, "stop_times.arrival_time")
		validateRuleConfig(rules.StopTimes.DepartureTime, "stop_times.departure_time")
		validateRuleConfig(rules.StopTimes.StopId, "stop_times.stop_id")
		validateRuleConfig(rules.StopTimes.StopSequence, "stop_times.stop_sequence")
		validateRuleConfig(rules.StopTimes.StopHeadsign, "stop_times.stop_headsign")
		validateRuleConfig(rules.StopTimes.PickupType, "stop_times.pickup_type")
		validateRuleConfig(rules.StopTimes.DropOffType, "stop_times.drop_off_type")
		validateRuleConfig(rules.StopTimes.ContinuousPickup, "stop_times.continuous_pickup")
		validateRuleConfig(rules.StopTimes.ContinuousDropOff, "stop_times.continuous_drop_off")
		validateRuleConfig(rules.StopTimes.ShapeDistTraveled, "stop_times.shape_dist_traveled")
		validateRuleConfig(rules.StopTimes.Timepoint, "stop_times.timepoint")
	}

	if rules.Calendar.File != types.SEVERITY_FORBIDDEN {
		validateRuleConfig(rules.Calendar.ServiceId, "calendar.service_id")
		validateRuleConfig(rules.Calendar.Monday, "calendar.monday")
		validateRuleConfig(rules.Calendar.Tuesday, "calendar.tuesday")
		validateRuleConfig(rules.Calendar.Wednesday, "calendar.wednesday")
		validateRuleConfig(rules.Calendar.Thursday, "calendar.thursday")
		validateRuleConfig(rules.Calendar.Friday, "calendar.friday")
		validateRuleConfig(rules.Calendar.Saturday, "calendar.saturday")
		validateRuleConfig(rules.Calendar.Sunday, "calendar.sunday")
		validateRuleConfig(rules.Calendar.StartDate, "calendar.start_date")
		validateRuleConfig(rules.Calendar.EndDate, "calendar.end_date")
	}

	if rules.CalendarDates.File != types.SEVERITY_FORBIDDEN {
		validateRuleConfig(rules.CalendarDates.ServiceId, "calendar_dates.service_id")
		validateRuleConfig(rules.CalendarDates.Date, "calendar_dates.date")
		validateRuleConfig(rules.CalendarDates.ExceptionType, "calendar_dates.exception_type")
	}

	// Return validation errors if any
	if len(validationErrors) > 0 {
		return fmt.Errorf("validation errors: %v", validationErrors)
	}

	return nil
}

// isValidSeverity checks if the given severity is valid
func isValidSeverity(severity types.Severity) bool {
	switch severity {
	case types.SEVERITY_IGNORE, types.SEVERITY_ERROR, types.SEVERITY_WARNING:
		return true
	default:
		return false
	}
}

// ParseRulesFromFile is a convenience function to parse rules from a file path
func ParseRulesFromFile(rulesPath string) (*types.GtfsRules, error) {
	if rulesPath == "" {
		return nil, nil
	}

	parser := NewRulesParser(rulesPath)
	return parser.ParseRules()
}

func (rp *RulesParser) GetRequiredFiles(rules *types.GtfsRules) []string {
	if rules == nil {
		return []string{}
	}

	allFiles := lib.GetAllStructTagValues(types.GtfsRules{}, "json")
	requiredFiles := make([]string, 0)

	for _, file := range allFiles {
		if lib.GetFieldByTag(rules, file, "_file") == "error" {
			requiredFiles = append(requiredFiles, file)
		}
	}

	return requiredFiles
}

func (rp *RulesParser) GetForbiddenFiles(rules *types.GtfsRules) []string {
	if rules == nil {
		return []string{}
	}

	allFiles := lib.GetAllStructTagValues(types.GtfsRules{}, "json")
	requiredFiles := make([]string, 0)

	for _, file := range allFiles {
		if lib.GetFieldByTag(rules, file, "_file") == "forbidden" {
			requiredFiles = append(requiredFiles, file)
		}
	}

	return requiredFiles
}
