package types

// Gtfs represents a collection of parsed GTFS data files where the key is the filename (without  extension)
// and the value is a slice of maps containing the CSV data with column headers as keys.
type GtfsFiles map[string][]map[string]string
type GtfsFieldCount map[string]map[string]int
type GtfsIdMap map[string]map[string]int // key is the filename, value is a map of ids and their row number

type Gtfs struct {
	Files        GtfsFiles
	FieldCounter GtfsFieldCount
	IdMap        GtfsIdMap
}

/* PRIMARY KEYS */
var GTFS_PRIMARY_KEYS = map[string]any{
	"afetacao":             nil,
	"agency":               "agency_id",
	"archives":             "archive_id",
	"areas":                "area_id",
	"attributions":         "attribution_id",
	"booking_rules":        "booking_rule_id",
	"calendar_dates":       []string{"service_id", "date"},
	"calendar":             "service_id",
	"dates":                nil,
	"fare_attributes":      "fare_id",
	"fare_leg_join_rules":  []string{"from_network_id", "to_network_id", "from_stop_id", "to_stop_id"},
	"fare_leg_rules":       []string{"network_id", "from_area_id", "to_area_id", "from_timeframe_group_id", "to_timeframe_group_id", "fare_product_id"},
	"fare_media":           "fare_media_id",
	"fare_products":        "fare_product_id",
	"fare_rules":           nil,
	"fare_transfer_rules":  []string{"from_leg_group_id", "to_leg_group_id", "fare_product_id", "transfer_count", "duration_limit"},
	"feed_info":            nil,
	"frequencies":          []string{"trip_id", "start_time"},
	"levels":               "level_id",
	"location_group_stops": []string{"location_group_id", "stop_id"},
	"location_groups":      "location_group_id",
	"municipalities":       "municipality_id",
	"networks":             "network_id",
	"pathways":             "pathway_id",
	"periods":              "period_id",
	"rider_categories":     "rider_category_id",
	"route_networks":       "route_id",
	"routes":               "route_id",
	"shapes":               []string{"shape_id", "shape_pt_sequence"},
	"stop_areas":           []string{"area_id", "stop_id"},
	"stop_times":           []string{"trip_id", "stop_sequence"},
	"stops":                "stop_id",
	"timeframes":           nil,
	"transfers":            []string{"from_stop_id", "to_stop_id", "from_trip_id", "to_trip_id", "from_route_id", "to_route_id"},
	"translations":         []string{"table_name", "field_name", "language", "record_id", "record_sub_id", "field_value"},
	"trips":                "trip_id",
}

/* AGENCY */
type Agency struct {
	AgencyEmail    *string `json:"agency_email"`
	AgencyFareUrl  *string `json:"agency_fare_url"`
	AgencyId       *string `json:"agency_id"`
	AgencyLang     *string `json:"agency_lang"`
	AgencyName     string  `json:"agency_name"`
	AgencyPhone    *string `json:"agency_phone"`
	AgencyTimezone string  `json:"agency_timezone"`
	AgencyUrl      string  `json:"agency_url"`
}

/* STOP */
type Stop struct {
	LevelId            *string  `json:"level_id,omitempty"`
	LocationType       *int     `json:"location_type,omitempty"`
	ParentStation      *string  `json:"parent_station,omitempty"`
	PlatformCode       *string  `json:"platform_code,omitempty"`
	StopCode           *string  `json:"stop_code,omitempty"`
	StopDesc           *string  `json:"stop_desc,omitempty"`
	StopId             string   `json:"stop_id"`
	StopLat            *float32 `json:"stop_lat,omitempty"`
	StopLon            *float32 `json:"stop_lon,omitempty"`
	StopName           *string  `json:"stop_name,omitempty"`
	StopTimezone       *string  `json:"stop_timezone,omitempty"`
	StopUrl            *string  `json:"stop_url,omitempty"`
	WheelchairBoarding *int     `json:"wheelchair_boarding,omitempty"`
	ZoneId             *string  `json:"zone_id,omitempty"`
}

/* ROUTE */
type Route struct {
	AgencyId          *string `json:"agency_id"`
	ContinuousDropOff *string `json:"continuous_drop_off"`
	ContinuousPickup  *string `json:"continuous_pickup"`
	RouteColor        *string `json:"route_color"`
	RouteDesc         *string `json:"route_desc"`
	RouteId           string  `json:"route_id"`
	RouteLongName     *string `json:"route_long_name"`
	RouteShortName    *string `json:"route_short_name"`
	RouteSortOrder    *int    `json:"route_sort_order"`
	RouteTextColor    *string `json:"route_text_color"`
	RouteType         int     `json:"route_type"`
	RouteUrl          *string `json:"route_url"`
}

/* TRIP */

type Trip struct {
	BikesAllowed         *int    `json:"bikes_allowed"`
	BlockId              *string `json:"block_id"`
	CalendarDesc         string  `json:"calendar_desc"`
	DirectionId          *bool   `json:"direction_id"`
	PatternId            *string `json:"pattern_id"`
	RouteId              string  `json:"route_id"`
	ServiceId            string  `json:"service_id"`
	ShapeId              *string `json:"shape_id"`
	TripHeadsign         *string `json:"trip_headsign"`
	TripId               string  `json:"trip_id"`
	TripShortName        *string `json:"trip_short_name"`
	WheelchairAccessible *string `json:"wheelchair_accessible"`
}

/* STOP TIME */
type StopTime struct {
	Airport                       *string  `json:"airport"`
	Areas                         *string  `json:"areas"`
	ArrivalTime                   *string  `json:"arrival_time"`
	BikeParking                   *string  `json:"bike_parking"`
	BikeSharing                   *string  `json:"bike_sharing"`
	Boat                          *string  `json:"boat"`
	CarParking                    *string  `json:"car_parking"`
	ContinuousDropOff             *int     `json:"continuous_drop_off"`
	ContinuousPickup              *int     `json:"continuous_pickup"`
	DepartureTime                 *string  `json:"departure_time"`
	DistrictId                    *string  `json:"district_id"`
	DistrictName                  *string  `json:"district_name"`
	DockingBayType                *string  `json:"docking_bay_type"`
	DropOffType                   *int     `json:"drop_off_type"`
	DropOffBookingRuleId          *string  `json:"drop_off_booking_rule_id"`
	EndPickupDropOffWindow        *string  `json:"end_pickup_drop_off_window"`
	FlagMaintainer                *string  `json:"flag_maintainer"`
	HasAbusiveParking             *string  `json:"has_abusive_parking"`
	HasBench                      *string  `json:"has_bench"`
	HasCover                      *string  `json:"has_cover"`
	HasCrossing                   *string  `json:"has_crossing"`
	HasElectricity                *string  `json:"has_electricity"`
	HasFlag                       *string  `json:"has_flag"`
	HasFlatAccess                 *string  `json:"has_flat_access"`
	HasH2oaSignage                *string  `json:"has_h2oa_signage"`
	HasLighting                   *string  `json:"has_lighting"`
	HasMupi                       *string  `json:"has_mupi"`
	HasNetworkMap                 *string  `json:"has_network_map"`
	HasPipAudio                   *string  `json:"has_pip_audio"`
	HasPipRealtime                *string  `json:"has_pip_realtime"`
	HasPipStatic                  *string  `json:"has_pip_static"`
	HasPole                       *string  `json:"has_pole"`
	HasSchedules                  *string  `json:"has_schedules"`
	HasShelter                    *string  `json:"has_shelter"`
	HasSidewalk                   *string  `json:"has_sidewalk"`
	HasTactileAccess              *string  `json:"has_tactile_access"`
	HasTactileSchedules           *string  `json:"has_tactile_schedules"`
	HasTrashBin                   *string  `json:"has_trash_bin"`
	HasWideAccess                 *string  `json:"has_wide_access"`
	Jurisdiction                  *string  `json:"jurisdiction"`
	LastAccessibilityCheck        *string  `json:"last_accessibility_check"`
	LastAccessibilityMaintenance  *string  `json:"last_accessibility_maintenance"`
	LastFlagCheck                 *string  `json:"last_flag_check"`
	LastFlagMaintenance           *string  `json:"last_flag_maintenance"`
	LastInfrastructureCheck       *string  `json:"last_infrastructure_check"`
	LastInfrastructureMaintenance *string  `json:"last_infrastructure_maintenance"`
	LastSchedulesCheck            *string  `json:"last_schedules_check"`
	LastSchedulesMaintenance      *string  `json:"last_schedules_maintenance"`
	LightRail                     *string  `json:"light_rail"`
	Locality                      *string  `json:"locality"`
	LocationGroupId               *string  `json:"location_group_id"`
	LocationId                    *string  `json:"location_id"`
	MunicipalityId                *string  `json:"municipality_id"`
	MunicipalityName              *string  `json:"municipality_name"`
	NearBeach                     *string  `json:"near_beach"`
	NearFireStation               *string  `json:"near_fire_station"`
	NearHealthClinic              *string  `json:"near_health_clinic"`
	NearHistoricBuilding          *string  `json:"near_historic_building"`
	NearHospital                  *string  `json:"near_hospital"`
	NearPoliceStation             *string  `json:"near_police_station"`
	NearSchool                    *string  `json:"near_school"`
	NearShopping                  *string  `json:"near_shopping"`
	NearTransitOffice             *string  `json:"near_transit_office"`
	NearUniversity                *string  `json:"near_university"`
	OperationalStatus             *string  `json:"operational_status"`
	ParishId                      *string  `json:"parish_id"`
	ParishName                    *string  `json:"parish_name"`
	PickupType                    *int     `json:"pickup_type"`
	PickupBookingRuleId           *string  `json:"pickup_booking_rule_id"`
	PipAudioCode                  *string  `json:"pip_audio_code"`
	PipRealtimeCode               *string  `json:"pip_realtime_code"`
	RegionId                      *string  `json:"region_id"`
	RegionName                    *string  `json:"region_name"`
	ShapeDistTraveled             *float32 `json:"shape_dist_traveled"`
	ShelterCode                   *string  `json:"shelter_code"`
	ShelterMaintainer             *string  `json:"shelter_maintainer"`
	SidewalkType                  *string  `json:"sidewalk_type"`
	StartPickupDropOffWindow      *string  `json:"start_pickup_drop_off_window"`
	StopHeadsign                  *string  `json:"stop_headsign"`
	StopId                        string   `json:"stop_id"`
	StopNameNew                   *string  `json:"stop_name_new"`
	StopSequence                  int      `json:"stop_sequence"`
	StopShortName                 *string  `json:"stop_short_name"`
	Subway                        *string  `json:"subway"`
	Timepoint                     *bool    `json:"timepoint"`
	Train                         *string  `json:"train"`
	TripId                        string   `json:"trip_id"`
	TtsStopName                   *string  `json:"tts_stop_name"`
}

/* CALENDAR */
type Calendar struct {
	EndDate   string `json:"end_date"`
	Friday    bool   `json:"friday"`
	Monday    bool   `json:"monday"`
	Saturday  bool   `json:"saturday"`
	ServiceId string `json:"service_id"`
	StartDate string `json:"start_date"`
	Sunday    bool   `json:"sunday"`
	Thursday  bool   `json:"thursday"`
	Tuesday   bool   `json:"tuesday"`
	Wednesday bool   `json:"wednesday"`
}

/* CALENDAR DATES */
type CalendarDates struct {
	Date          string `json:"date"`
	ExceptionType int    `json:"exception_type"`
	ServiceId     string `json:"service_id"`
}

/* FARE ATTRIBUTES */
type FareAttribute struct {
	// Required fields
	FareId        string   `json:"fare_id"`        // Identifies a fare class
	Price         *float64 `json:"price"`          // Fare price, in the unit specified by currency_type
	CurrencyType  string   `json:"currency_type"`  // Currency used to pay the fare
	PaymentMethod *int     `json:"payment_method"` // When the fare must be paid (0: on board, 1: before boarding)

	// Optional fields
	Transfers        *int    `json:"transfers"`         // Number of transfers permitted (0: none, 1: once, 2: twice, empty: unlimited)
	AgencyId         *string `json:"agency_id"`         // Agency associated with the fare (required if multiple agencies)
	TransferDuration *int    `json:"transfer_duration"` // Length of time in seconds before a transfer expires
}

/* FARE RULES */
type FareRule struct {
	// Required fields
	FareId string `json:"fare_id"` // Identifies a fare class

	// Optional fields
	RouteId       *string `json:"route_id"`       // Identifies a route associated with the fare class
	OriginId      *string `json:"origin_id"`      // Identifies an origin zone
	DestinationId *string `json:"destination_id"` // Identifies a destination zone
	ContainsId    *string `json:"contains_id"`    // Identifies zones that a rider will enter while using a given fare class
}

/* SHAPES */
type Shapes struct {
	ShapeDistTraveled *float32 `json:"shape_dist_traveled"`
	ShapeId           string   `json:"shape_id"`
	ShapePtLat        float32  `json:"shape_pt_lat"`
	ShapePtLon        float32  `json:"shape_pt_lon"`
	ShapePtSequence   float32  `json:"shape_pt_sequence"`
}

/* FREQUENCIES */
type Frequencies struct {
	EndTime     string  `json:"end_time"`
	ExactTimes  *int    `json:"exact_times"`
	HeadwaySecs float32 `json:"headway_secs"`
	StartTime   string  `json:"start_time"`
	TripId      string  `json:"trip_id"`
}

/* TRANSFERS */
type Transfers struct {
	FromRouteId     *string `json:"from_route_id"`
	FromStopId      string  `json:"from_stop_id"`
	FromTripId      *string `json:"from_trip_id"`
	MinTransferTime float32 `json:"min_transfer_time"`
	ToRouteId       *string `json:"to_route_id"`
	ToStopId        string  `json:"to_stop_id"`
	ToTripId        *string `json:"to_trip_id"`
	TransferType    int     `json:"transfer_type"`
}

/* PATHWAYS */
type Pathways struct {
	FromStopId           *string  `json:"from_stop_id"`
	IsBidirectional      bool     `json:"is_bidirectional"`
	Length               *float32 `json:"length"`
	MaxSlope             *string  `json:"max_slope"`
	MinWidth             *string  `json:"min_width"`
	PathwayId            string   `json:"pathway_id"`
	PathwayMode          int      `json:"pathway_mode"`
	ReversedSignpostedAs *string  `json:"reversed_signposted_as"`
	SignpostedAs         *string  `json:"signposted_as"`
	StairCount           *uint16  `json:"stair_count"`
	ToStopId             *string  `json:"to_stop_id"`
	TraversalTime        *float32 `json:"traversal_time"`
}

/* LEVELS */
type Levels struct {
	LevelId    string  `json:"level_id"`
	LevelIndex uint16  `json:"level_index"`
	LevelName  *string `json:"level_name"`
}

/* FEED INFO */
type FeedInfo struct {
	DefaultLang       string `json:"default_lang"`
	FeedContactEmail  string `json:"feed_contact_email"`
	FeedContactUrl    string `json:"feed_contact_url"`
	FeedEndDate       string `json:"feed_end_date"`
	FeedLang          string `json:"feed_lang"`
	FeedPublisherName string `json:"feed_publisher_name"`
	FeedPublisherUrl  string `json:"feed_publisher_url"`
	FeedStartDate     string `json:"feed_start_date"`
	FeedVersion       string `json:"feed_version"`
}

/* TRANSLATIONS */
type Translations struct {
	FieldName   string  `json:"field_name"`
	FieldValue  *string `json:"field_value"`
	Language    string  `json:"language"`
	RecordId    *string `json:"record_id"`
	RecordSubId *string `json:"record_sub_id"`
	TableName   string  `json:"table_name"`
	Translation string  `json:"translation"`
}

/* ATTRIBUTIONS */
type Attributions struct {
	AgencyId         *string `json:"agency_id"`
	AttributionEmail *string `json:"attribution_email"`
	AttributionId    *string `json:"attribution_id"`
	AttributionPhone *string `json:"attribution_phone"`
	AttributionUrl   *string `json:"attribution_url"`
	IsAuthority      *bool   `json:"is_authority"`
	IsOperator       *bool   `json:"is_operator"`
	IsProducer       *bool   `json:"is_producer"`
	OrganizationName string  `json:"organization_name"`
	RouteId          *string `json:"route_id"`
	TripId           *string `json:"trip_id"`
}

/* TIMEFRAME */
type Timeframe struct {
	EndTime          *string `json:"end_time"`
	ServiceId        string  `json:"service_id"`
	StartTime        *string `json:"start_time"`
	TimeframeGroupId string  `json:"timeframe_group_id"`
}

/* RIDER CATEGORY*/
type RiderCategory struct {
	EligibilityUrl        string `json:"eligibility_url"`
	IsDefaultFareCategory bool   `json:"is_default_fare_category"`
	RiderCategoryId       string `json:"rider_category_id"`
	RiderCategoryName     string `json:"rider_category_name"`
}

/* FARE MEDIA */
type FareMedia struct {
	FareMediaId   string `json:"fare_media_id"`
	FareMediaName string `json:"fare_media_name"`
	FareMediaType string `json:"fare_media_type"`
}

/* FARE PRODUCT */
type FareProduct struct {
	Ammount         float32 `json:"ammount"`
	Currency        string  `json:"currency"`
	FareMediaId     *string `json:"fare_media_id"`
	FareProductId   string  `json:"fare_product_id"`
	FareProductName *string `json:"fare_product_name"`
	RiderCategoryId *string `json:"rider_category_id"`
}

/* FARE LEG RULE */
type FareLegRule struct {
	FareProductId        string   `json:"fare_product_id"`
	FromAreaId           *string  `json:"from_area_id"`
	FromTimeframeGroupId *string  `json:"from_timeframe_group_id"`
	LegGroupId           *string  `json:"leg_group_id"`
	NetworkId            *string  `json:"network_id"`
	RulePriority         *float32 `json:"rule_priority"`
	ToAreaId             *string  `json:"to_area_id"`
	ToTimeframeGroupId   *string  `json:"to_timeframe_group_id"`
}

/* FARE LEG JOIN RULE */
type FareLegJoinRule struct {
	FromNetworkId string  `json:"from_network_id"`
	FromStopId    *string `json:"from_stop_id"`
	ToNetworkId   string  `json:"to_network_id"`
	ToStopId      *string `json:"to_stop_id"`
}

/* FARETRANSFERRULE */
type FareTransferRule struct {
	DurationLimit     *float32 `json:"duration_limit"`
	DurationLimitType *int     `json:"duration_limit_type"`
	FareProductId     *string  `json:"fare_product_id"`
	FareTransferType  *int     `json:"fare_transfer_type"`
	FromLegGroupId    *string  `json:"from_leg_group_id"`
	ToLegGroupId      *string  `json:"to_leg_group_id"`
	TransferCount     *float32 `json:"transfer_count"`
}

/* AREA */
type Area struct {
	AreaId   string  `json:"area_id"`
	AreaName *string `json:"area_name"`
}

/* STOPAREA */
type StopArea struct {
	AreaId string `json:"area_id"`
	StopId string `json:"stop_id"`
}

/* NETWORK */
type Network struct {
	NetworkId   string  `json:"network_id"`
	NetworkName *string `json:"network_name"`
}

/* ROUTENETWORK */
type RouteNetwork struct {
	NetworkId string `json:"network_id"`
	RouteId   string `json:"route_id"`
}

/* LOCATIONGROUP */
type LocationGroup struct {
	LocationGroupId   string  `json:"location_group_id"`
	LocationGroupName *string `json:"location_group_name"`
}

/* LOCATIONGROUPSTOP */
type LocationGroupStop struct {
	LocationGroupId string `json:"location_group_id"`
	StopId          string `json:"stop_id"`
}

/* BOOKINGRULE */
type BookingRule struct {
	BookingRuleId          string   `json:"booking_rule_id"`
	BookingType            int      `json:"booking_type"`
	BookingUrl             *string  `json:"booking_url"`
	DropOffMessage         *string  `json:"drop_off_message"`
	InfoUrl                *string  `json:"info_url"`
	Message                *string  `json:"message"`
	PhoneNumber            *string  `json:"phone_number"`
	PickupMessage          *string  `json:"pickup_message"`
	PriorNoticeDurationMax float32  `json:"prior_notice_duration_max"`
	PriorNoticeDurationMin float32  `json:"prior_notice_duration_min"`
	PriorNoticeLastDay     *float32 `json:"prior_notice_last_day"`
	PriorNoticeLastTime    *string  `json:"prior_notice_last_time"`
	PriorNoticeServiceId   *string  `json:"prior_notice_service_id"`
	PriorNoticeStartDay    *float32 `json:"prior_notice_start_day"`
	PriorNoticeStartTime   *string  `json:"prior_notice_start_time"`
}

/* ARCHIVE */
type Archive struct {
	ArchiveEndDate   string `json:"archive_end_date"`
	ArchiveId        string `json:"archive_id"`
	ArchiveStartDate string `json:"archive_start_date"`
	OperatorId       string `json:"operator_id"`
}

/* MUNICIPALITY */
type Municipality struct {
	DistrictId         string `json:"district_id"`
	DistrictName       string `json:"district_name"`
	MunicipalityId     string `json:"municipality_id"`
	MunicipalityName   string `json:"municipality_name"`
	MunicipalityPrefix string `json:"municipality_prefix"`
	RegionId           string `json:"region_id"`
	RegionName         string `json:"region_name"`
}

/* AFETACAO */
type Afetacao struct {
	AcceptedZoneCodes string  `json:"accepted_zone_codes"`
	AcceptedZoneNames string  `json:"accepted_zone_names"`
	Interchange       string  `json:"interchange"`
	LineId            string  `json:"line_id"`
	LineType          string  `json:"line_type"`
	OnboardFares      string  `json:"onboard_fares"`
	OperatorId        string  `json:"operator_id"`
	PatternId         string  `json:"pattern_id"`
	PrepaidFare       string  `json:"prepaid_fare"`
	PrepaidFarePrice  string  `json:"prepaid_fare_price"`
	StopId            string  `json:"stop_id"`
	StopName          string  `json:"stop_name"`
	StopSequence      float32 `json:"stop_sequence"`
}

/* PERIOD */
type Period struct {
	PeriodId   string `json:"period_id"`
	PeriodName string `json:"period_name"`
}
