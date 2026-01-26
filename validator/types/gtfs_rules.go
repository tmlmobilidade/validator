package types

const ALL_OPTIONS = "all_options"

type Compare struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type RuleConfig struct {
	Severity Severity   `json:"severity"`
	Options  *[]string  `json:"options,omitempty"`
	Compare  *[]Compare `json:"compare,omitempty"`
}

type AgencyRules struct {
	File              Severity   `json:"_file"`
	AgencyId          RuleConfig `json:"agency_id"`
	AgencyNameIdMatch RuleConfig `json:"agency_name_id_match"`
	AgencyName        RuleConfig `json:"agency_name"`
	AgencyUrl         RuleConfig `json:"agency_url"`
	AgencyTimezone    RuleConfig `json:"agency_timezone"`
	AgencyLang        RuleConfig `json:"agency_lang"`
	AgencyPhone       RuleConfig `json:"agency_phone"`
	AgencyFare        RuleConfig `json:"agency_fare_url"`
	AgencyEmail       RuleConfig `json:"agency_email"`
}

type StopsRules struct {
	File                  Severity   `json:"_file"`
	StopId                RuleConfig `json:"stop_id"`
	StopCode              RuleConfig `json:"stop_code"`
	StopName              RuleConfig `json:"stop_name"`
	StopShortName         RuleConfig `json:"stop_short_name"`
	TtsStopName           RuleConfig `json:"tts_stop_name"`
	StopDesc              RuleConfig `json:"stop_desc"`
	StopLat               RuleConfig `json:"stop_lat"`
	StopLon               RuleConfig `json:"stop_lon"`
	ZoneId                RuleConfig `json:"zone_id"`
	StopUrl               RuleConfig `json:"stop_url"`
	LocationType          RuleConfig `json:"location_type"`
	ParentStation         RuleConfig `json:"parent_station"`
	StopTimezone          RuleConfig `json:"stop_timezone"`
	WheelchairBoarding    RuleConfig `json:"wheelchair_boarding"`
	LevelId               RuleConfig `json:"level_id"`
	PlatformCode          RuleConfig `json:"platform_code"`
	PublicVisible         RuleConfig `json:"public_visible"`
	HasStopSign           RuleConfig `json:"has_stop_sign"`
	HasShelter            RuleConfig `json:"has_shelter"`
	ShelterCode           RuleConfig `json:"shelter_code"`
	ShelterMaintainer     RuleConfig `json:"shelter_maintainer"`
	HasBench              RuleConfig `json:"has_bench"`
	HasNetworkMap         RuleConfig `json:"has_network_map"`
	HasSchedules          RuleConfig `json:"has_schedules"`
	HasPipRealTime        RuleConfig `json:"has_pip_real_time"`
	HasTariffsInformation RuleConfig `json:"has_tariffs_information"`
	RegionId              RuleConfig `json:"region_id"`
	MunicipalityId        RuleConfig `json:"municipality_id"`
	ParishId              RuleConfig `json:"parish_id"`
}

type RoutesRules struct {
	File              Severity   `json:"_file"`
	LineId            RuleConfig `json:"line_id"`
	LineShortName     RuleConfig `json:"line_short_name"`
	LineLongName      RuleConfig `json:"line_long_name"`
	RouteId           RuleConfig `json:"route_id"`
	AgencyId          RuleConfig `json:"agency_id"`
	RouteShortName    RuleConfig `json:"route_short_name"`
	RouteLongName     RuleConfig `json:"route_long_name"`
	RouteDesc         RuleConfig `json:"route_desc"`
	RouteSortOrder    RuleConfig `json:"route_sort_order"`
	RouteRemarks      RuleConfig `json:"route_remarks"`
	NetworkId         RuleConfig `json:"network_id"`
	RouteType         RuleConfig `json:"route_type"`
	PathType          RuleConfig `json:"path_type"`
	Circular          RuleConfig `json:"circular"`
	School            RuleConfig `json:"school"`
	RouteUrl          RuleConfig `json:"route_url"`
	RouteColor        RuleConfig `json:"route_color"`
	RouteTextColor    RuleConfig `json:"route_text_color"`
	ContinuousPickup  RuleConfig `json:"continuous_pickup"`
	ContinuousDropOff RuleConfig `json:"continuous_drop_off"`
}

type TripsRules struct {
	File                    Severity   `json:"_file"`
	RouteId                 RuleConfig `json:"route_id"`
	PatternId               RuleConfig `json:"pattern_id"`
	ServiceId               RuleConfig `json:"service_id"`
	TripId                  RuleConfig `json:"trip_id"`
	TripHeadsign            RuleConfig `json:"trip_headsign"`
	TripShortName           RuleConfig `json:"trip_short_name"`
	DirectionId             RuleConfig `json:"direction_id"`
	BlockId                 RuleConfig `json:"block_id"`
	ShapeId                 RuleConfig `json:"shape_id"`
	WheelchairAccessible    RuleConfig `json:"wheelchair_accessible"`
	BikesAllowed            RuleConfig `json:"bikes_allowed"`
	StopSequence            RuleConfig `json:"stop_sequence"`
	DirectionPatternIdMatch RuleConfig `json:"direction_pattern_id_match"`
}

type StopTimesRules struct {
	File                     Severity   `json:"_file"`
	TripId                   RuleConfig `json:"trip_id"`
	ArrivalTime              RuleConfig `json:"arrival_time"`
	DepartureTime            RuleConfig `json:"departure_time"`
	StopId                   RuleConfig `json:"stop_id"`
	StopSequence             RuleConfig `json:"stop_sequence"`
	StopHeadsign             RuleConfig `json:"stop_headsign"`
	PickupType               RuleConfig `json:"pickup_type"`
	DropOffType              RuleConfig `json:"drop_off_type"`
	ContinuousPickup         RuleConfig `json:"continuous_pickup"`
	ContinuousDropOff        RuleConfig `json:"continuous_drop_off"`
	ShapeDistTraveled        RuleConfig `json:"shape_dist_traveled"`
	StartPickupDropOffWindow RuleConfig `json:"start_pickup_drop_off_window"`
	EndPickupDropOffWindow   RuleConfig `json:"end_pickup_drop_off_window"`
	Timepoint                RuleConfig `json:"timepoint"`
	PickupBookingRuleId      RuleConfig `json:"pickup_booking_rule_id"`
	DropOffBookingRuleId     RuleConfig `json:"drop_off_booking_rule_id"`
}

type CalendarRules struct {
	File      Severity   `json:"_file"`
	ServiceId RuleConfig `json:"service_id"`
	Monday    RuleConfig `json:"monday"`
	Tuesday   RuleConfig `json:"tuesday"`
	Wednesday RuleConfig `json:"wednesday"`
	Thursday  RuleConfig `json:"thursday"`
	Friday    RuleConfig `json:"friday"`
	Saturday  RuleConfig `json:"saturday"`
	Sunday    RuleConfig `json:"sunday"`
	StartDate RuleConfig `json:"start_date"`
	EndDate   RuleConfig `json:"end_date"`
}

type CalendarDatesRules struct {
	File          Severity   `json:"_file"`
	ServiceId     RuleConfig `json:"service_id"`
	Date          RuleConfig `json:"date"`
	ExceptionType RuleConfig `json:"exception_type"`
}

type VehiclesRules struct {
	File              Severity   `json:"_file"`
	VehicleId         RuleConfig `json:"vehicle_id"`
	AgencyId          RuleConfig `json:"agency_id"`
	LicensePlate      RuleConfig `json:"license_plate"`
	Make              RuleConfig `json:"make"`
	Model             RuleConfig `json:"model"`
	Owner             RuleConfig `json:"owner"`
	RegistrationDate  RuleConfig `json:"registration_date"`
	AvailableSeats    RuleConfig `json:"available_seats"`
	AvailableStanding RuleConfig `json:"available_standing"`
	Typology          RuleConfig `json:"typology"`
	Propulsion        RuleConfig `json:"propulsion"`
	Emission          RuleConfig `json:"emission"`
	Climatization     RuleConfig `json:"climatization"`
	Wheelchair        RuleConfig `json:"wheelchair"`
	LoweredFloor      RuleConfig `json:"lowered_floor"`
	Ramp              RuleConfig `json:"ramp"`
	Kneeling          RuleConfig `json:"kneeling"`
	StaticInformation RuleConfig `json:"static_information"`
	OnboardMonitor    RuleConfig `json:"onboard_monitor"`
	FrontDisplay      RuleConfig `json:"front_display"`
	RearDisplay       RuleConfig `json:"rear_display"`
	SideDisplay       RuleConfig `json:"side_display"`
	InternalSound     RuleConfig `json:"internal_sound"`
	ExternalSound     RuleConfig `json:"external_sound"`
	ConsumptionMeter  RuleConfig `json:"consumption_meter"`
	Bicycles          RuleConfig `json:"bicycles"`
	PassengerCounting RuleConfig `json:"passenger_counting"`
	VideoSurveillance RuleConfig `json:"video_surveillance"`
}

type FareAttributesRules struct {
	File             Severity   `json:"_file"`
	FareId           RuleConfig `json:"fare_id"`
	Price            RuleConfig `json:"price"`
	CurrencyType     RuleConfig `json:"currency_type"`
	PaymentMethod    RuleConfig `json:"payment_method"`
	Transfers        RuleConfig `json:"transfers"`
	AgencyId         RuleConfig `json:"agency_id"`
	TransferDuration RuleConfig `json:"transfer_duration"`
}

type FareRulesRules struct {
	File          Severity   `json:"_file"`
	FareId        RuleConfig `json:"fare_id"`
	RouteId       RuleConfig `json:"route_id"`
	OriginId      RuleConfig `json:"origin_id"`
	DestinationId RuleConfig `json:"destination_id"`
	ContainsId    RuleConfig `json:"contains_id"`
}

type FareMediaRules struct {
	File     Severity   `json:"_file"`
	FareId   RuleConfig `json:"fare_id"`
	FareName RuleConfig `json:"fare_Name"`
	FareType RuleConfig `json:"fare_type"`
}

type FareProductRules struct {
	File            Severity   `json:"_file"`
	FareProductId   RuleConfig `json:"fare_product_id"`
	FareProductName RuleConfig `json:"fare_product_name"`
	RiderCategoryId RuleConfig `json:"rider_category_id"`
	Currency        RuleConfig `json:"currency"`
	Ammount         RuleConfig `json:"ammount"`
}

type ShapesRules struct {
	File              Severity   `json:"_file"`
	ShapeId           RuleConfig `json:"shape_id"`
	ShapePtLat        RuleConfig `json:"shape_pt_lat"`
	ShapePtLon        RuleConfig `json:"shape_pt_lon"`
	ShapePtSequence   RuleConfig `json:"shape_pt_sequence"`
	ShapeDistTraveled RuleConfig `json:"shape_dist_traveled"`
}

type FrequenciesRules struct {
	File        Severity   `json:"_file"`
	TripId      RuleConfig `json:"trip_id"`
	StartTime   RuleConfig `json:"start_time"`
	EndTime     RuleConfig `json:"end_time"`
	HeadwaySecs RuleConfig `json:"headway_secs"`
	ExactTimes  RuleConfig `json:"exact_times"`
}

type TransfersRules struct {
	File            Severity   `json:"_file"`
	FromStopId      RuleConfig `json:"from_stop_id"`
	ToStopId        RuleConfig `json:"to_stop_id"`
	TransferType    RuleConfig `json:"transfer_type"`
	MinTransferTime RuleConfig `json:"min_transfer_time"`
}

type PathwaysRules struct {
	File                 Severity   `json:"_file"`
	PathwayId            RuleConfig `json:"pathway_id"`
	FromStopId           RuleConfig `json:"from_stop_id"`
	ToStopId             RuleConfig `json:"to_stop_id"`
	PathwayMode          RuleConfig `json:"pathway_mode"`
	IsBidirectional      RuleConfig `json:"is_bidirectional"`
	Length               RuleConfig `json:"length"`
	TraversalTime        RuleConfig `json:"traversal_time"`
	StairCount           RuleConfig `json:"stair_count"`
	MaxSlope             RuleConfig `json:"max_slope"`
	MinWidth             RuleConfig `json:"min_width"`
	SignpostedAs         RuleConfig `json:"signposted_as"`
	ReversedSignpostedAs RuleConfig `json:"reversed_signposted_as"`
}

type LevelsRules struct {
	File       Severity   `json:"_file"`
	LevelId    RuleConfig `json:"level_id"`
	LevelIndex RuleConfig `json:"level_index"`
	LevelName  RuleConfig `json:"level_name"`
}

type FeedInfoRules struct {
	File              Severity   `json:"_file"`
	FeedType          RuleConfig `json:"feed_type"`
	FeedPublisherName RuleConfig `json:"feed_publisher_name"`
	FeedPublisherUrl  RuleConfig `json:"feed_publisher_url"`
	FeedLang          RuleConfig `json:"feed_lang"`
	DefaultLang       RuleConfig `json:"default_lang"`
	FeedStartDate     RuleConfig `json:"feed_start_date"`
	FeedEndDate       RuleConfig `json:"feed_end_date"`
	FeedVersion       RuleConfig `json:"feed_version"`
	FeedRemarks       RuleConfig `json:"feed_remarks"`
	FeedContactEmail  RuleConfig `json:"feed_contact_email"`
	FeedContactUrl    RuleConfig `json:"feed_contact_url"`
}

type TranslationsRules struct {
	File        Severity   `json:"_file"`
	TableName   RuleConfig `json:"table_name"`
	FieldName   RuleConfig `json:"field_name"`
	Language    RuleConfig `json:"language"`
	Translation RuleConfig `json:"translation"`
	RecordId    RuleConfig `json:"record_id"`
	RecordSubId RuleConfig `json:"record_sub_id"`
	FieldValue  RuleConfig `json:"field_value"`
}

type AttributionsRules struct {
	File             Severity   `json:"_file"`
	AttributionId    RuleConfig `json:"attribution_id"`
	AgencyId         RuleConfig `json:"agency_id"`
	RouteId          RuleConfig `json:"route_id"`
	TripId           RuleConfig `json:"trip_id"`
	OrganizationName RuleConfig `json:"organization_name"`
	IsProducer       RuleConfig `json:"is_producer"`
	IsOperator       RuleConfig `json:"is_operator"`
	IsAuthority      RuleConfig `json:"is_authority"`
	AttributionUrl   RuleConfig `json:"attribution_url"`
	AttributionEmail RuleConfig `json:"attribution_email"`
	AttributionPhone RuleConfig `json:"attribution_phone"`
}

type GtfsRules struct {
	Agency         AgencyRules         `json:"agency"`
	Stops          StopsRules          `json:"stops"`
	Routes         RoutesRules         `json:"routes"`
	Trips          TripsRules          `json:"trips"`
	StopTimes      StopTimesRules      `json:"stop_times"`
	Calendar       CalendarRules       `json:"calendar"`
	CalendarDates  CalendarDatesRules  `json:"calendar_dates"`
	Vehicles       VehiclesRules       `json:"vehicles"`
	FareAttributes FareAttributesRules `json:"fare_attributes"`
	FareRules      FareRulesRules      `json:"fare_rules"`
	Shapes         ShapesRules         `json:"shapes"`
	Frequencies    FrequenciesRules    `json:"frequencies"`
	Transfers      TransfersRules      `json:"transfers"`
	Pathways       PathwaysRules       `json:"pathways"`
	Levels         LevelsRules         `json:"levels"`
	FeedInfo       FeedInfoRules       `json:"feed_info"`
	Translations   TranslationsRules   `json:"translations"`
	Attributions   AttributionsRules   `json:"attributions"`
}
