package types

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

/* AGENCY */
type AgencyRaw struct {
	AgencyEmail    string `gtfs:"agency_email"`
	AgencyFareUrl  string `gtfs:"agency_fare_url"`
	AgencyId       string `gtfs:"agency_id"`
	AgencyLang     string `gtfs:"agency_lang"`
	AgencyName     string  `gtfs:"agency_name"`
	AgencyPhone    string `gtfs:"agency_phone"`
	AgencyTimezone string  `gtfs:"agency_timezone"`
	AgencyUrl      string  `gtfs:"agency_url"`
}


/* STOP */
type StopRaw struct {
	LevelId            string  `gtfs:"level_id,omitempty"`
	LocationType       string     `gtfs:"location_type,omitempty"`
	ParentStation      string  `gtfs:"parent_station,omitempty"`
	PlatformCode       string  `gtfs:"platform_code,omitempty"`
	StopCode           string  `gtfs:"stop_code,omitempty"`
	StopName           string  `gtfs:"stop_name,omitempty"`
	TtsStopName        string  `gtfs:"tts_stop_name,omitempty"`
	StopDesc           string  `gtfs:"stop_desc,omitempty"`
	StopId             string   `gtfs:"stop_id"`
	StopLat            string `gtfs:"stop_lat,omitempty"`
	StopLon            string `gtfs:"stop_lon,omitempty"`
	StopTimezone       string  `gtfs:"stop_timezone,omitempty"`
	StopUrl            string  `gtfs:"stop_url,omitempty"`
	WheelchairBoarding string     `gtfs:"wheelchair_boarding,omitempty"`
	ZoneId             string  `gtfs:"zone_id,omitempty"`
}



/* ROUTE */
type RouteRaw struct {
	// Required fields
	RouteId           string  `gtfs:"route_id"`
	RouteType         string     `gtfs:"route_type"`

	// Optional fields
	AgencyId          string `gtfs:"agency_id"`
	ContinuousDropOff string `gtfs:"continuous_drop_off"`
	ContinuousPickup  string `gtfs:"continuous_pickup"`
	RouteColor        string `gtfs:"route_color"`
	RouteDesc         string `gtfs:"route_desc"`
	RouteLongName     string `gtfs:"route_long_name"`
	RouteShortName    string `gtfs:"route_short_name"`
	RouteSortOrder    string    `gtfs:"route_sort_order"`
	RouteTextColor    string `gtfs:"route_text_color"`
	RouteUrl          string `gtfs:"route_url"`
	NetworkId         string `gtfs:"network_id"`
}

/* TRIP */

type TripRaw struct {
	BikesAllowed         string    `gtfs:"bikes_allowed"`
	BlockId              string `gtfs:"block_id"`
	CalendarDesc         string  `gtfs:"calendar_desc"`
	DirectionId          string   `gtfs:"direction_id"`
	PatternId            string `gtfs:"pattern_id"`
	RouteId              string  `gtfs:"route_id"`
	ServiceId            string  `gtfs:"service_id"`
	ShapeId              string `gtfs:"shape_id"`
	TripHeadsign         string `gtfs:"trip_headsign"`
	TripId               string  `gtfs:"trip_id"`
	TripShortName        string `gtfs:"trip_short_name"`
	WheelchairAccessible string `gtfs:"wheelchair_accessible"`
}

/* STOP TIME */
type StopTimeRaw struct {
	TripId 					 string  `gtfs:"trip_id"`
	ArrivalTime 			 string  `gtfs:"arrival_time"`
	DepartureTime 			 string  `gtfs:"departure_time"`
	StopId 					 string  `gtfs:"stop_id"`
	LocationGroupId 		 string  `gtfs:"location_group_id"`
	LocationId 				 string  `gtfs:"location_id"`
	StopSequence 			 string     `gtfs:"stop_sequence"`
	StopHeadsign 			 string  `gtfs:"stop_headsign"`
	StartPickupDropOffWindow string  `gtfs:"start_pickup_drop_off_window"`
	EndPickupDropOffWindow 	 string  `gtfs:"end_pickup_drop_off_window"`
	PickupType 				 string 	  `gtfs:"pickup_type"`
	DropOffType 			 string 	  `gtfs:"drop_off_type"`
	ContinuousPickup 		 string 	  `gtfs:"continuous_pickup"`
	ContinuousDropOff 		 string 	  `gtfs:"continuous_drop_off"`
	ShapeDistTraveled 		 string `gtfs:"shape_dist_traveled"`
	Timepoint 				 string 	  `gtfs:"timepoint"`
	PickupBookingRuleId 	 string  `gtfs:"pickup_booking_rule_id"`
	DropOffBookingRuleId 	 string  `gtfs:"drop_off_booking_rule_id"`
}

/* CALENDAR */
type CalendarRaw struct {
	EndDate   string `gtfs:"end_date"`
	Friday    string   `gtfs:"friday"`
	Monday    string   `gtfs:"monday"`
	Saturday  string   `gtfs:"saturday"`
	ServiceId string `gtfs:"service_id"`
	StartDate string `gtfs:"start_date"`
	Sunday    string   `gtfs:"sunday"`
	Thursday  string   `gtfs:"thursday"`
	Tuesday   string   `gtfs:"tuesday"`
	Wednesday string   `gtfs:"wednesday"`
}

/* CALENDAR DATES */
type CalendarDatesRaw struct {
	Date          string `gtfs:"date"`
	ExceptionType string    `gtfs:"exception_type"`
	ServiceId     string `gtfs:"service_id"`
}

/* FARE ATTRIBUTES */
type FareAttributeRaw struct {
	FareId           string  `gtfs:"fare_id"`           // Identifies a fare class
	Price            string `gtfs:"price"`             // Fare price, in the unit specified by currency_type
	CurrencyType     string  `gtfs:"currency_type"`     // Currency used to pay the fare
	PaymentMethod    string     `gtfs:"payment_method"`    // When the fare must be paid (0: on board, 1: before boarding)
	Transfers        string     `gtfs:"transfers"`         // Number of transfers permitted (0: none, 1: once, 2: twice, empty: unlimited)
	AgencyId         string  `gtfs:"agency_id"`         // Agency associated with the fare (required if multiple agencies)
	TransferDuration string     `gtfs:"transfer_duration"` // Length of time in seconds before a transfer expires
}

/* FARE RULES */
type FareRuleRaw struct {
	FareId 		  string `gtfs:"fare_id"` // Identifies a fare class
	RouteId       string `gtfs:"route_id"`       // Identifies a route associated with the fare class
	OriginId      string `gtfs:"origin_id"`      // Identifies an origin zone
	DestinationId string `gtfs:"destination_id"` // Identifies a destination zone
	ContainsId    string `gtfs:"contains_id"`    // Identifies zones that a rider will enter while using a given fare class
}

/* SHAPES */
type ShapeRaw struct {
	ShapeId           string  `gtfs:"shape_id"`
	ShapePtLat        string `gtfs:"shape_pt_lat"`
	ShapePtLon        string `gtfs:"shape_pt_lon"`
	ShapePtSequence   string     `gtfs:"shape_pt_sequence"`
	ShapeDistTraveled string `gtfs:"shape_dist_traveled"`
}

/* FREQUENCIES */
type FrequenciesRaw struct {
	EndTime     string  `gtfs:"end_time"`
	ExactTimes  string    `gtfs:"exact_times"`
	HeadwaySecs string `gtfs:"headway_secs"`
	StartTime   string  `gtfs:"start_time"`
	TripId      string  `gtfs:"trip_id"`
}

/* TRANSFERS */
type TransfersRaw struct {
	FromRouteId     string `gtfs:"from_route_id"`
	FromStopId      string  `gtfs:"from_stop_id"`
	FromTripId      string `gtfs:"from_trip_id"`
	MinTransferTime string `gtfs:"min_transfer_time"`
	ToRouteId       string `gtfs:"to_route_id"`
	ToStopId        string  `gtfs:"to_stop_id"`
	ToTripId        string `gtfs:"to_trip_id"`
	TransferType    string  `gtfs:"transfer_type"`
}

/* PATHWAYS */
type PathwaysRaw struct {
	FromStopId           string  `gtfs:"from_stop_id"`
	IsBidirectional      string     `gtfs:"is_bidirectional"`
	Length               string `gtfs:"length"`
	MaxSlope             string  `gtfs:"max_slope"`
	MinWidth             string  `gtfs:"min_width"`
	PathwayId            string   `gtfs:"pathway_id"`
	PathwayMode          string   `gtfs:"pathway_mode"`
	ReversedSignpostedAs string  `gtfs:"reversed_signposted_as"`
	SignpostedAs         string  `gtfs:"signposted_as"`
	StairCount           string  `gtfs:"stair_count"`
	ToStopId             string  `gtfs:"to_stop_id"`
	TraversalTime        string `gtfs:"traversal_time"`
}

/* LEVELS */
type LevelsRaw struct {
	LevelId    string  `gtfs:"level_id"`
	LevelIndex uint16  `gtfs:"level_index"`
	LevelName  string `gtfs:"level_name"`
}

/* FEED INFO */
type FeedInfoRaw struct {
	// Required fields
	FeedLang          string `gtfs:"feed_lang"`
	FeedPublisherName string `gtfs:"feed_publisher_name"`
	FeedPublisherUrl  string `gtfs:"feed_publisher_url"`

	// Optional fields
	DefaultLang      string `gtfs:"default_lang"`
	FeedContactEmail string `gtfs:"feed_contact_email"`
	FeedContactUrl   string `gtfs:"feed_contact_url"`
	FeedEndDate      string `gtfs:"feed_end_date"`
	FeedStartDate    string `gtfs:"feed_start_date"`
	FeedVersion      string `gtfs:"feed_version"`
}

/* TRANSLATIONS */
type TranslationsRaw struct {
	FieldName   string  `gtfs:"field_name"`
	FieldValue  string `gtfs:"field_value"`
	Language    string  `gtfs:"language"`
	RecordId    string `gtfs:"record_id"`
	RecordSubId string `gtfs:"record_sub_id"`
	TableName   string  `gtfs:"table_name"`
	Translation string  `gtfs:"translation"`
}

/* ATTRIBUTIONS */
type AttributionsRaw struct {
	AgencyId         string `gtfs:"agency_id"`
	AttributionEmail string `gtfs:"attribution_email"`
	AttributionId    string `gtfs:"attribution_id"`
	AttributionPhone string `gtfs:"attribution_phone"`
	AttributionUrl   string `gtfs:"attribution_url"`
	IsAuthority      string   `gtfs:"is_authority"`
	IsOperator       string   `gtfs:"is_operator"`
	IsProducer       string   `gtfs:"is_producer"`
	OrganizationName string  `gtfs:"organization_name"`
	RouteId          string `gtfs:"route_id"`
	TripId           string `gtfs:"trip_id"`
}

/* TIMEFRAME */
type TimeframeRaw struct {
	EndTime          string `gtfs:"end_time"`
	ServiceId        string  `gtfs:"service_id"`
	StartTime        string `gtfs:"start_time"`
	TimeframeGroupId string  `gtfs:"timeframe_group_id"`
}

/* RIDER CATEGORY*/
type RiderCategoryRaw struct {
	EligibilityUrl        string `gtfs:"eligibility_url"`
	IsDefaultFareCategory string   `gtfs:"is_default_fare_category"`
	RiderCategoryId       string `gtfs:"rider_category_id"`
	RiderCategoryName     string `gtfs:"rider_category_name"`
}

/* FARE MEDIA */
type FareMediaRaw struct {
	FareMediaId   string `gtfs:"fare_media_id"`
	FareMediaName string `gtfs:"fare_media_name"`
	FareMediaType string `gtfs:"fare_media_type"`
}

/* FARE PRODUCT */
type FareProductRaw struct {
	Ammount         string `gtfs:"ammount"`
	Currency        string  `gtfs:"currency"`
	FareMediaId     string `gtfs:"fare_media_id"`
	FareProductId   string  `gtfs:"fare_product_id"`
	FareProductName string `gtfs:"fare_product_name"`
	RiderCategoryId string `gtfs:"rider_category_id"`
}

/* FARE LEG RULE */
type FareLegRuleRaw struct {
	FareProductId        string   `gtfs:"fare_product_id"`
	FromAreaId           string  `gtfs:"from_area_id"`
	FromTimeframeGroupId string  `gtfs:"from_timeframe_group_id"`
	LegGroupId           string  `gtfs:"leg_group_id"`
	NetworkId            string  `gtfs:"network_id"`
	RulePriority         string `gtfs:"rule_priority"`
	ToAreaId             string  `gtfs:"to_area_id"`
	ToTimeframeGroupId   string  `gtfs:"to_timeframe_group_id"`
}

/* FARE LEG JOIN RULE */
type FareLegJoinRuleRaw struct {
	FromNetworkId string  `gtfs:"from_network_id"`
	FromStopId    string `gtfs:"from_stop_id"`
	ToNetworkId   string  `gtfs:"to_network_id"`
	ToStopId      string `gtfs:"to_stop_id"`
}

/* FARETRANSFERRULE */
type FareTransferRuleRaw struct {
	DurationLimit     string `gtfs:"duration_limit"`
	DurationLimitType string     `gtfs:"duration_limit_type"`
	FareProductId     string  `gtfs:"fare_product_id"`
	FareTransferType  string     `gtfs:"fare_transfer_type"`
	FromLegGroupId    string  `gtfs:"from_leg_group_id"`
	ToLegGroupId      string  `gtfs:"to_leg_group_id"`
	TransferCount     string `gtfs:"transfer_count"`
}

/* AREA */
type AreaRaw struct {
	AreaId   string  `gtfs:"area_id"`
	AreaName string `gtfs:"area_name"`
}

/* STOPAREA */
type StopAreaRaw struct {
	AreaId string `gtfs:"area_id"`
	StopId string `gtfs:"stop_id"`
}

/* NETWORK */
type NetworkRaw struct {
	NetworkId   string  `gtfs:"network_id"`
	NetworkName string `gtfs:"network_name"`
}

/* ROUTENETWORK */
type RouteNetworkRaw struct {
	NetworkId string `gtfs:"network_id"`
	RouteId   string `gtfs:"route_id"`
}

/* LOCATIONGROUP */
type LocationGroupRaw struct {
	LocationGroupId   string  `gtfs:"location_group_id"`
	LocationGroupName string `gtfs:"location_group_name"`
}

/* LOCATIONGROUPSTOP */
type LocationGroupStopRaw struct {
	LocationGroupId string `gtfs:"location_group_id"`
	StopId          string `gtfs:"stop_id"`
}

/* BOOKINGRULE */
type BookingRuleRaw struct {
	BookingRuleId          string   `gtfs:"booking_rule_id"`
	BookingType            string   `gtfs:"booking_type"`
	BookingUrl             string  `gtfs:"booking_url"`
	DropOffMessage         string  `gtfs:"drop_off_message"`
	InfoUrl                string  `gtfs:"info_url"`
	Message                string  `gtfs:"message"`
	PhoneNumber            string  `gtfs:"phone_number"`
	PickupMessage          string  `gtfs:"pickup_message"`
	PriorNoticeDurationMax string  `gtfs:"prior_notice_duration_max"`
	PriorNoticeDurationMin string  `gtfs:"prior_notice_duration_min"`
	PriorNoticeLastDay     string `gtfs:"prior_notice_last_day"`
	PriorNoticeLastTime    string  `gtfs:"prior_notice_last_time"`
	PriorNoticeServiceId   string  `gtfs:"prior_notice_service_id"`
	PriorNoticeStartDay    string `gtfs:"prior_notice_start_day"`
	PriorNoticeStartTime   string  `gtfs:"prior_notice_start_time"`
}

/* ARCHIVE */
type ArchiveRaw struct {
	ArchiveEndDate   string `gtfs:"archive_end_date"`
	ArchiveId        string `gtfs:"archive_id"`
	ArchiveStartDate string `gtfs:"archive_start_date"`
	OperatorId       string `gtfs:"operator_id"`
}

/* MUNICIPALITY */
type MunicipalityRaw struct {
	DistrictId         string `gtfs:"district_id"`
	DistrictName       string `gtfs:"district_name"`
	MunicipalityId     string `gtfs:"municipality_id"`
	MunicipalityName   string `gtfs:"municipality_name"`
	MunicipalityPrefix string `gtfs:"municipality_prefix"`
	RegionId           string `gtfs:"region_id"`
	RegionName         string `gtfs:"region_name"`
}

/* AFETACAO */
type AfetacaoRaw struct {
	AcceptedZoneCodes string  `gtfs:"accepted_zone_codes"`
	AcceptedZoneNames string  `gtfs:"accepted_zone_names"`
	Interchange       string  `gtfs:"interchange"`
	LineId            string  `gtfs:"line_id"`
	LineType          string  `gtfs:"line_type"`
	OnboardFares      string  `gtfs:"onboard_fares"`
	OperatorId        string  `gtfs:"operator_id"`
	PatternId         string  `gtfs:"pattern_id"`
	PrepaidFare       string  `gtfs:"prepaid_fare"`
	PrepaidFarePrice  string  `gtfs:"prepaid_fare_price"`
	StopId            string  `gtfs:"stop_id"`
	StopName          string  `gtfs:"stop_name"`
	StopSequence      string `gtfs:"stop_sequence"`
}

/* PERIOD */
type PeriodRaw struct {
	PeriodId   string `gtfs:"period_id"`
	PeriodName string `gtfs:"period_name"`
}

// Gtfs represents a collection of parsed GTFS data files where the key is the filename (without  extension)
// and the value is a slice of maps containing the CSV data with column headers as keys.
type GtfsFiles map[string][]map[string]string
type GtfsIdMap map[string]map[string][]int // key is the filename, value is a map of ids and their row number

type Gtfs struct {
	Agency []AgencyRaw  `gtfs:"agency"`
	Stop []StopRaw  `gtfs:"stop"`
	Route []RouteRaw  `gtfs:"route"`
	Trip []TripRaw  `gtfs:"trip"`
	StopTime []StopTimeRaw  `gtfs:"stop_time"`
	Calendar []CalendarRaw  `gtfs:"calendar"`
	CalendarDates []CalendarDatesRaw  `gtfs:"calendar_dates"`
	FareAttribute []FareAttributeRaw  `gtfs:"fare_attribute"`
	FareRule []FareRuleRaw  `gtfs:"fare_rule"`
	Shape []ShapeRaw  `gtfs:"shape"`
	Frequencies []FrequenciesRaw  `gtfs:"frequencies"`
	Transfers []TransfersRaw  `gtfs:"transfers"`
	Pathways []PathwaysRaw  `gtfs:"pathways"`
	Levels []LevelsRaw  `gtfs:"levels"`
	FeedInfo []FeedInfoRaw  `gtfs:"feed_info"`
	Translations []TranslationsRaw  `gtfs:"translations"`
	Attributions []AttributionsRaw  `gtfs:"attributions"`
	Timeframe []TimeframeRaw  `gtfs:"timeframe"`
	RiderCategory []RiderCategoryRaw  `gtfs:"rider_category"`
	FareMedia []FareMediaRaw  `gtfs:"fare_media"`
	FareProduct []FareProductRaw  `gtfs:"fare_product"`
	FareLegRule []FareLegRuleRaw  `gtfs:"fare_leg_rule"`
	FareLegJoinRule []FareLegJoinRuleRaw  `gtfs:"fare_leg_join_rule"`
	FareTransferRule []FareTransferRuleRaw  `gtfs:"fare_transfer_rule"`
	Area []AreaRaw  `gtfs:"area"`
	StopArea []StopAreaRaw  `gtfs:"stop_area"`
	Network []NetworkRaw  `gtfs:"network"`
	RouteNetwork []RouteNetworkRaw  `gtfs:"route_network"`
	LocationGroup []LocationGroupRaw  `gtfs:"location_group"`
	LocationGroupStop []LocationGroupStopRaw  `gtfs:"location_group_stop"`
	BookingRule []BookingRuleRaw  `gtfs:"booking_rule"`
	Archive []ArchiveRaw  `gtfs:"archive"`
	Municipality []MunicipalityRaw  `gtfs:"municipality"`
	Afetacao []AfetacaoRaw  `gtfs:"afetacao"`
	Period []PeriodRaw  `gtfs:"period"`
	
	IdMap        map[string]map[string][]int // key is the filename, value is a map of ids and their row number
}

func NewGtfs() *Gtfs {
	return &Gtfs{
		Agency: make([]AgencyRaw, 0),
		Stop: make([]StopRaw, 0),
		Route: make([]RouteRaw, 0),
		Trip: make([]TripRaw, 0),
		StopTime: make([]StopTimeRaw, 0),
		Calendar: make([]CalendarRaw, 0),
		CalendarDates: make([]CalendarDatesRaw, 0),
		FareAttribute: make([]FareAttributeRaw, 0),
		FareRule: make([]FareRuleRaw, 0),
		Shape: make([]ShapeRaw, 0),
		Frequencies: make([]FrequenciesRaw, 0),
		Transfers: make([]TransfersRaw, 0),
		Pathways: make([]PathwaysRaw, 0),
		Levels: make([]LevelsRaw, 0),
		FeedInfo: make([]FeedInfoRaw, 0),
		Translations: make([]TranslationsRaw, 0),
		Attributions: make([]AttributionsRaw, 0),
		Timeframe: make([]TimeframeRaw, 0),
		RiderCategory: make([]RiderCategoryRaw, 0),
		FareMedia: make([]FareMediaRaw, 0),
		FareProduct: make([]FareProductRaw, 0),
		FareLegRule: make([]FareLegRuleRaw, 0),
		FareLegJoinRule: make([]FareLegJoinRuleRaw, 0),
		FareTransferRule: make([]FareTransferRuleRaw, 0),
		Area: make([]AreaRaw, 0),
		StopArea: make([]StopAreaRaw, 0),
		Network: make([]NetworkRaw, 0),
		RouteNetwork: make([]RouteNetworkRaw, 0),
		LocationGroup: make([]LocationGroupRaw, 0),
		LocationGroupStop: make([]LocationGroupStopRaw, 0),
		BookingRule: make([]BookingRuleRaw, 0),
		Archive: make([]ArchiveRaw, 0),
		Municipality: make([]MunicipalityRaw, 0),
		Afetacao: make([]AfetacaoRaw, 0),
		Period: make([]PeriodRaw, 0),

		IdMap: make(map[string]map[string][]int),
	}
}

// findFieldByPath traverses the struct based on a dot-separated path and tag name.
// It returns the reflect.Value of the final field.
func findFieldByPath(obj any, path string) (reflect.Value, error) {
	// Start with the reflection value of the object.
	// We need a pointer to be able to set values, so we start by getting the element.
	v := reflect.ValueOf(obj).Elem()

	// Split the path into parts (e.g., "socials.facebook" -> ["socials", "facebook"])
	parts := strings.Split(path, ".")

	// Traverse the path
	for _, part := range parts {
		// If we are not on a struct, we can't go deeper.
		if v.Kind() != reflect.Struct {
			return reflect.Value{}, fmt.Errorf("cannot access field %s on a non-struct type %s", part, v.Kind())
		}

		// Find the field in the current struct that matches the tag.
		found := false
		for i := 0; i < v.NumField(); i++ {
			fieldTag := v.Type().Field(i).Tag.Get("gtfs")
			if fieldTag == part {
				v = v.Field(i)
				found = true
				break
			}
		}

		if !found {
			return reflect.Value{}, fmt.Errorf("field with tag '%s' not found in path '%s'", part, path)
		}
	}
	return v, nil
}


// SetField sets a field on the user struct using a path and a string value.
// It handles basic type conversion (e.g., string to int).
func (u *Gtfs) SetField(path string, value string) error {
	fieldVal, err := findFieldByPath(u, path)
	if err != nil {
		return err
	}

	// Check if the field is actually settable (i.e., it's exported).
	if !fieldVal.CanSet() {
		return fmt.Errorf("field at path '%s' is not settable", path)
	}

	// Perform type conversion based on the field's kind.
	switch fieldVal.Kind() {
	case reflect.String:
		fieldVal.SetString(value)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intValue, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fmt.Errorf("could not parse '%s' as an integer for path '%s': %w", value, path, err)
		}
		fieldVal.SetInt(intValue)
	// Add other types like bool, float64 as needed
	// case reflect.Bool: ...
	default:
		return fmt.Errorf("unsupported field type %s at path '%s'", fieldVal.Kind(), path)
	}

	return nil
}

// GetField retrieves a field's value from the user struct as a string.
func (u *Gtfs) GetField(path string) (string, error) {
	fieldVal, err := findFieldByPath(u, path)
	if err != nil {
		return "", err
	}

	// Use fmt.Sprintf to reliably convert any value to its string representation.
	return fmt.Sprintf("%v", fieldVal.Interface()), nil
}

// SetFieldData converts CSV map data to the appropriate struct slice and sets it on the Gtfs struct
func (u *Gtfs) SetFieldData(filename string, data []map[string]string) error {
	switch filename {
	case "agency":
		u.Agency = convertToStructSlice[AgencyRaw](data)
	case "stops":
		u.Stop = convertToStructSlice[StopRaw](data)
	case "routes":
		u.Route = convertToStructSlice[RouteRaw](data)
	case "trips":
		u.Trip = convertToStructSlice[TripRaw](data)
	case "stop_times":
		u.StopTime = convertToStructSlice[StopTimeRaw](data)
	case "calendar":
		u.Calendar = convertToStructSlice[CalendarRaw](data)
	case "calendar_dates":
		u.CalendarDates = convertToStructSlice[CalendarDatesRaw](data)
	case "fare_attributes":
		u.FareAttribute = convertToStructSlice[FareAttributeRaw](data)
	case "fare_rules":
		u.FareRule = convertToStructSlice[FareRuleRaw](data)
	case "shapes":
		u.Shape = convertToStructSlice[ShapeRaw](data)
	case "frequencies":
		u.Frequencies = convertToStructSlice[FrequenciesRaw](data)
	case "transfers":
		u.Transfers = convertToStructSlice[TransfersRaw](data)
	case "pathways":
		u.Pathways = convertToStructSlice[PathwaysRaw](data)
	case "levels":
		u.Levels = convertToStructSlice[LevelsRaw](data)
	case "feed_info":
		u.FeedInfo = convertToStructSlice[FeedInfoRaw](data)
	case "translations":
		u.Translations = convertToStructSlice[TranslationsRaw](data)
	case "attributions":
		u.Attributions = convertToStructSlice[AttributionsRaw](data)
	case "archives":
		u.Archive = convertToStructSlice[ArchiveRaw](data)
	case "municipalities":
		u.Municipality = convertToStructSlice[MunicipalityRaw](data)
	case "periods":
		u.Period = convertToStructSlice[PeriodRaw](data)
	default:
		return fmt.Errorf("unknown GTFS file: %s", filename)
	}
	return nil
}

// convertToStructSlice converts CSV map data to a slice of structs using reflection
func convertToStructSlice[T any](data []map[string]string) []T {
	result := make([]T, len(data))
	for i, row := range data {
		var item T
		v := reflect.ValueOf(&item).Elem()
		t := v.Type()
		
		for j := 0; j < v.NumField(); j++ {
			field := v.Field(j)
			fieldType := t.Field(j)
			tag := fieldType.Tag.Get("gtfs")
			
			if tag != "" && tag != "-" {
				if value, exists := row[tag]; exists && field.CanSet() {
					field.SetString(value)
				}
			}
		}
		result[i] = item
	}
	return result
}

// GetFieldByTag retrieves a field's value by its GTFS tag name from any struct
func GetFieldByTag[T any](obj *T, tagName string) string {
	v := reflect.ValueOf(obj).Elem()
	t := v.Type()
	
	for i := range v.NumField() {
		field := v.Field(i)
		fieldType := t.Field(i)
		tag := fieldType.Tag.Get("gtfs")
		
		if tag == tagName {
			return field.String()
		}
	}
	return ""
}