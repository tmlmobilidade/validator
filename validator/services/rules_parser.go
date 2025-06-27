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
			validationErrors = append(validationErrors, 
				fmt.Sprintf("invalid severity '%s' for field '%s'", config.Severity, fieldName))
		}
	}

	// Validate agency rules
	if rules.Agency.File {
		validateRuleConfig(rules.Agency.AgencyId, "agency.agency_id")
		validateRuleConfig(rules.Agency.AgencyName, "agency.agency_name")
		validateRuleConfig(rules.Agency.AgencyUrl, "agency.agency_url")
		validateRuleConfig(rules.Agency.AgencyTz, "agency.agency_timezone")
		validateRuleConfig(rules.Agency.AgencyLang, "agency.agency_lang")
		validateRuleConfig(rules.Agency.AgencyPhone, "agency.agency_phone")
		validateRuleConfig(rules.Agency.AgencyFare, "agency.agency_fare_url")
		validateRuleConfig(rules.Agency.AgencyEmail, "agency.agency_email")
	}

	// Validate stops rules
	if rules.Stops.File {
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
	if rules.Routes.File {
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
	if rules.Trips.File {
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

	if rules.StopTimes.File {
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
		validateRuleConfig(rules.StopTimes.Zone1, "stop_times.zone_1")
		validateRuleConfig(rules.StopTimes.Zone2, "stop_times.zone_2")
		validateRuleConfig(rules.StopTimes.Zone3, "stop_times.zone_3")
	}

	if rules.Calendar.File {
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

	if rules.CalendarDates.File {
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
	parser := NewRulesParser(rulesPath)
	return parser.ParseRules()
}

// GetRuleConfig gets the rule configuration for a specific file and field
func (rp *RulesParser) GetRuleConfig(rules *types.GtfsRules, fileName, fieldName string) (*types.RuleConfig, error) {
	switch fileName {
	case "agency":
		return rp.getAgencyRuleConfig(&rules.Agency, fieldName)
	case "stops":
		return rp.getStopsRuleConfig(&rules.Stops, fieldName)
	case "routes":
		return rp.getRoutesRuleConfig(&rules.Routes, fieldName)
	case "trips":
		return rp.getTripsRuleConfig(&rules.Trips, fieldName)
	case "stop_times":
		return rp.getStopTimesRuleConfig(&rules.StopTimes, fieldName)
	case "calendar":
		return rp.getCalendarRuleConfig(&rules.Calendar, fieldName)
	case "calendar_dates":
		return rp.getCalendarDatesRuleConfig(&rules.CalendarDates, fieldName)
	case "fare_attributes":
		return rp.getFareAttributesRuleConfig(&rules.FareAttributes, fieldName)
	case "fare_rules":
		return rp.getFareRulesRuleConfig(&rules.FareRules, fieldName)
	case "shapes":
		return rp.getShapesRuleConfig(&rules.Shapes, fieldName)
	case "frequencies":
		return rp.getFrequenciesRuleConfig(&rules.Frequencies, fieldName)
	case "transfers":
		return rp.getTransfersRuleConfig(&rules.Transfers, fieldName)
	case "pathways":
		return rp.getPathwaysRuleConfig(&rules.Pathways, fieldName)
	case "levels":
		return rp.getLevelsRuleConfig(&rules.Levels, fieldName)
	case "feed_info":
		return rp.getFeedInfoRuleConfig(&rules.FeedInfo, fieldName)
	case "translations":
		return rp.getTranslationsRuleConfig(&rules.Translations, fieldName)
	case "attributions":
		return rp.getAttributionsRuleConfig(&rules.Attributions, fieldName)
	default:
		return nil, fmt.Errorf("unknown file name: %s", fileName)
	}
}

// Helper functions to get rule configurations for specific files
func (rp *RulesParser) getAgencyRuleConfig(rules *types.AgencyRules, fieldName string) (*types.RuleConfig, error) {
	switch fieldName {
	case "agency_id":
		return &rules.AgencyId, nil
	case "agency_name":
		return &rules.AgencyName, nil
	case "agency_url":
		return &rules.AgencyUrl, nil
	case "agency_timezone":
		return &rules.AgencyTz, nil
	case "agency_lang":
		return &rules.AgencyLang, nil
	case "agency_phone":
		return &rules.AgencyPhone, nil
	case "agency_fare_url":
		return &rules.AgencyFare, nil
	case "agency_email":
		return &rules.AgencyEmail, nil
	default:
		return nil, fmt.Errorf("unknown field name: %s", fieldName)
	}
}

func (rp *RulesParser) getStopsRuleConfig(rules *types.StopsRules, fieldName string) (*types.RuleConfig, error) {
	switch fieldName {
	case "stop_id":
		return &rules.StopId, nil
	case "stop_code":
		return &rules.StopCode, nil
	case "stop_name":
		return &rules.StopName, nil
	case "stop_short_name":
		return &rules.StopShortName, nil
	case "tts_stop_name":
		return &rules.TtsStopName, nil
	case "stop_desc":
		return &rules.StopDesc, nil
	case "stop_lat":
		return &rules.StopLat, nil
	case "stop_lon":
		return &rules.StopLon, nil
	case "zone_id":
		return &rules.ZoneId, nil
	case "stop_url":
		return &rules.StopUrl, nil
	case "location_type":
		return &rules.LocationType, nil
	case "parent_station":
		return &rules.ParentStation, nil
	case "stop_timezone":
		return &rules.StopTimezone, nil
	case "wheelchair_boarding":
		return &rules.WheelchairBoarding, nil
	case "level_id":
		return &rules.LevelId, nil
	case "platform_code":
		return &rules.PlatformCode, nil
	default:
		return nil, fmt.Errorf("unknown field name: %s", fieldName)
	}
}

func (rp *RulesParser) getRoutesRuleConfig(rules *types.RoutesRules, fieldName string) (*types.RuleConfig, error) {
	switch fieldName {
	case "route_id":
		return &rules.RouteId, nil
	case "agency_id":
		return &rules.AgencyId, nil
	case "route_short_name":
		return &rules.RouteShortName, nil
	case "route_long_name":
		return &rules.RouteLongName, nil
	case "route_desc":
		return &rules.RouteDesc, nil
	case "route_type":
		return &rules.RouteType, nil
	case "route_url":
		return &rules.RouteUrl, nil
	case "route_color":
		return &rules.RouteColor, nil
	case "route_text_color":
		return &rules.RouteTextColor, nil
	case "continuous_pickup":
		return &rules.ContinuousPickup, nil
	case "continuous_drop_off":
		return &rules.ContinuousDropOff, nil
	default:
		return nil, fmt.Errorf("unknown field name: %s", fieldName)
	}
}

func (rp *RulesParser) getTripsRuleConfig(rules *types.TripsRules, fieldName string) (*types.RuleConfig, error) {
	switch fieldName {
	case "route_id":
		return &rules.RouteId, nil
	case "service_id":
		return &rules.ServiceId, nil
	case "trip_id":
		return &rules.TripId, nil
	case "trip_headsign":
		return &rules.TripHeadsign, nil
	case "trip_short_name":
		return &rules.TripShortName, nil
	case "direction_id":
		return &rules.DirectionId, nil
	case "block_id":
		return &rules.BlockId, nil
	case "shape_id":
		return &rules.ShapeId, nil
	case "wheelchair_accessible":
		return &rules.WheelchairAccessible, nil
	case "bikes_allowed":
		return &rules.BikesAllowed, nil
	default:
		return nil, fmt.Errorf("unknown field name: %s", fieldName)
	}
}

func (rp *RulesParser) getStopTimesRuleConfig(rules *types.StopTimesRules, fieldName string) (*types.RuleConfig, error) {
	switch fieldName {
	case "trip_id":
		return &rules.TripId, nil
	case "arrival_time":
		return &rules.ArrivalTime, nil
	case "departure_time":
		return &rules.DepartureTime, nil
	case "stop_id":
		return &rules.StopId, nil
	case "stop_sequence":
		return &rules.StopSequence, nil
	case "stop_headsign":
		return &rules.StopHeadsign, nil
	case "pickup_type":
		return &rules.PickupType, nil
	case "drop_off_type":
		return &rules.DropOffType, nil
	case "continuous_pickup":
		return &rules.ContinuousPickup, nil
	case "continuous_drop_off":
		return &rules.ContinuousDropOff, nil
	case "shape_dist_traveled":
		return &rules.ShapeDistTraveled, nil
	case "timepoint":
		return &rules.Timepoint, nil
	default:
		return nil, fmt.Errorf("unknown field name: %s", fieldName)
	}
}

func (rp *RulesParser) getCalendarRuleConfig(rules *types.CalendarRules, fieldName string) (*types.RuleConfig, error) {
	switch fieldName {
	case "service_id":
		return &rules.ServiceId, nil
	case "monday":
		return &rules.Monday, nil
	case "tuesday":
		return &rules.Tuesday, nil
	case "wednesday":
		return &rules.Wednesday, nil
	case "thursday":
		return &rules.Thursday, nil
	case "friday":
		return &rules.Friday, nil
	case "saturday":
		return &rules.Saturday, nil
	case "sunday":
		return &rules.Sunday, nil
	case "start_date":
		return &rules.StartDate, nil
	case "end_date":
		return &rules.EndDate, nil
	default:
		return nil, fmt.Errorf("unknown field name: %s", fieldName)
	}
}

func (rp *RulesParser) getCalendarDatesRuleConfig(rules *types.CalendarDatesRules, fieldName string) (*types.RuleConfig, error) {
	switch fieldName {
	case "service_id":
		return &rules.ServiceId, nil
	case "date":
		return &rules.Date, nil
	case "exception_type":
		return &rules.ExceptionType, nil
	default:
		return nil, fmt.Errorf("unknown field name: %s", fieldName)
	}
}

// Placeholder implementations for other rule types
func (rp *RulesParser) getFareAttributesRuleConfig(rules *types.FareAttributesRules, fieldName string) (*types.RuleConfig, error) {
	switch fieldName {
	case "fare_id":
		return &rules.FareId, nil
	case "price":
		return &rules.Price, nil
	case "currency_type":
		return &rules.CurrencyType, nil
	case "payment_method":
		return &rules.PaymentMethod, nil
	case "transfers":
		return &rules.Transfers, nil
	case "agency_id":
		return &rules.AgencyId, nil
	case "transfer_duration":
		return &rules.TransferDuration, nil
	default:
		return nil, fmt.Errorf("unknown field name: %s", fieldName)
	}
}

func (rp *RulesParser) getFareRulesRuleConfig(rules *types.FareRulesRules, fieldName string) (*types.RuleConfig, error) {
	switch fieldName {
	case "fare_id":
		return &rules.FareId, nil
	case "route_id":
		return &rules.RouteId, nil
	case "origin_id":
		return &rules.OriginId, nil
	case "destination_id":
		return &rules.DestinationId, nil
	case "contains_id":
		return &rules.ContainsId, nil
	default:
		return nil, fmt.Errorf("unknown field name: %s", fieldName)
	}
}

func (rp *RulesParser) getShapesRuleConfig(rules *types.ShapesRules, fieldName string) (*types.RuleConfig, error) {
	switch fieldName {
	case "shape_id":
		return &rules.ShapeId, nil
	case "shape_pt_lat":
		return &rules.ShapePtLat, nil
	case "shape_pt_lon":
		return &rules.ShapePtLon, nil
	case "shape_pt_sequence":
		return &rules.ShapePtSequence, nil
	case "shape_dist_traveled":
		return &rules.ShapeDistTraveled, nil
	default:
		return nil, fmt.Errorf("unknown field name: %s", fieldName)
	}
}

func (rp *RulesParser) getFrequenciesRuleConfig(rules *types.FrequenciesRules, fieldName string) (*types.RuleConfig, error) {
	switch fieldName {
	case "trip_id":
		return &rules.TripId, nil
	case "start_time":
		return &rules.StartTime, nil
	case "end_time":
		return &rules.EndTime, nil
	case "headway_secs":
		return &rules.HeadwaySecs, nil
	case "exact_times":
		return &rules.ExactTimes, nil
	default:
		return nil, fmt.Errorf("unknown field name: %s", fieldName)
	}
}

func (rp *RulesParser) getTransfersRuleConfig(rules *types.TransfersRules, fieldName string) (*types.RuleConfig, error) {
	switch fieldName {
	case "from_stop_id":
		return &rules.FromStopId, nil
	case "to_stop_id":
		return &rules.ToStopId, nil
	case "transfer_type":
		return &rules.TransferType, nil
	case "min_transfer_time":
		return &rules.MinTransferTime, nil
	default:
		return nil, fmt.Errorf("unknown field name: %s", fieldName)
	}
}

func (rp *RulesParser) getPathwaysRuleConfig(rules *types.PathwaysRules, fieldName string) (*types.RuleConfig, error) {
	switch fieldName {
	case "pathway_id":
		return &rules.PathwayId, nil
	case "from_stop_id":
		return &rules.FromStopId, nil
	case "to_stop_id":
		return &rules.ToStopId, nil
	case "pathway_mode":
		return &rules.PathwayMode, nil
	case "is_bidirectional":
		return &rules.IsBidirectional, nil
	case "length":
		return &rules.Length, nil
	case "traversal_time":
		return &rules.TraversalTime, nil
	case "stair_count":
		return &rules.StairCount, nil
	case "max_slope":
		return &rules.MaxSlope, nil
	case "min_width":
		return &rules.MinWidth, nil
	case "signposted_as":
		return &rules.SignpostedAs, nil
	case "reversed_signposted_as":
		return &rules.ReversedSignpostedAs, nil
	default:
		return nil, fmt.Errorf("unknown field name: %s", fieldName)
	}
}

func (rp *RulesParser) getLevelsRuleConfig(rules *types.LevelsRules, fieldName string) (*types.RuleConfig, error) {
	switch fieldName {
	case "level_id":
		return &rules.LevelId, nil
	case "level_index":
		return &rules.LevelIndex, nil
	case "level_name":
		return &rules.LevelName, nil
	default:
		return nil, fmt.Errorf("unknown field name: %s", fieldName)
	}
}

func (rp *RulesParser) getFeedInfoRuleConfig(rules *types.FeedInfoRules, fieldName string) (*types.RuleConfig, error) {
	switch fieldName {
	case "feed_publisher_name":
		return &rules.FeedPublisherName, nil
	case "feed_publisher_url":
		return &rules.FeedPublisherUrl, nil
	case "feed_lang":
		return &rules.FeedLang, nil
	case "default_lang":
		return &rules.DefaultLang, nil
	case "feed_start_date":
		return &rules.FeedStartDate, nil
	case "feed_end_date":
		return &rules.FeedEndDate, nil
	case "feed_version":
		return &rules.FeedVersion, nil
	case "feed_contact_email":
		return &rules.FeedContactEmail, nil
	case "feed_contact_url":
		return &rules.FeedContactUrl, nil
	default:
		return nil, fmt.Errorf("unknown field name: %s", fieldName)
	}
}

func (rp *RulesParser) getTranslationsRuleConfig(rules *types.TranslationsRules, fieldName string) (*types.RuleConfig, error) {
	switch fieldName {
	case "table_name":
		return &rules.TableName, nil
	case "field_name":
		return &rules.FieldName, nil
	case "language":
		return &rules.Language, nil
	case "translation":
		return &rules.Translation, nil
	case "record_id":
		return &rules.RecordId, nil
	case "record_sub_id":
		return &rules.RecordSubId, nil
	case "field_value":
		return &rules.FieldValue, nil
	default:
		return nil, fmt.Errorf("unknown field name: %s", fieldName)
	}
}

func (rp *RulesParser) getAttributionsRuleConfig(rules *types.AttributionsRules, fieldName string) (*types.RuleConfig, error) {
	switch fieldName {
	case "attribution_id":
		return &rules.AttributionId, nil
	case "agency_id":
		return &rules.AgencyId, nil
	case "route_id":
		return &rules.RouteId, nil
	case "trip_id":
		return &rules.TripId, nil
	case "organization_name":
		return &rules.OrganizationName, nil
	case "is_producer":
		return &rules.IsProducer, nil
	case "is_operator":
		return &rules.IsOperator, nil
	case "is_authority":
		return &rules.IsAuthority, nil
	case "attribution_url":
		return &rules.AttributionUrl, nil
	case "attribution_email":
		return &rules.AttributionEmail, nil
	case "attribution_phone":
		return &rules.AttributionPhone, nil
	default:
		return nil, fmt.Errorf("unknown field name: %s", fieldName)
	}
}

func (rp *RulesParser) GetRequiredFiles(rules *types.GtfsRules) []string {
	allFiles := lib.GetAllStructTagValues(types.GtfsRules{}, "json")
	requiredFiles := make([]string, 0)

	for _, file := range allFiles {
		if lib.GetFieldByTag(rules, file, "_file") == "true" {
			requiredFiles = append(requiredFiles, file)
		}
	}

	return requiredFiles
}