package lib

import (
	"main/src/models"
	"main/src/validations"
)

/* GTFS FILE-MODEL MAP */
var GTFS_FILE_RULES_MAP = map[string]func(gtfsData models.Gtfs){
	"afetacao":             validations.AfetacaoValidations,
	"agency":               validations.AgencyValidations,
	"archives":             validations.ArchivesValidations,
	"areas":                validations.AreasValidations,
	"attributions":         validations.AttributionsValidations,
	"booking_rules":        validations.BookingRulesValidations,
	"calendar":             validations.CalendarValidations,
	"calendar_dates":       validations.CalendarDatesValidations,
	"fare_attributes":      validations.FareAttributesValidations,
	"fare_leg_join_rules":  validations.FareLegJoinRulesValidations,
	"fare_leg_rules":       validations.FareLegRulesValidations,
	"fare_media":           validations.FareMediaValidations,
	"fare_products":        validations.FareProductsValidations,
	"fare_rules":           validations.FareRulesValidations,
	"fare_transfer_rules":  validations.FareTransferRulesValidations,
	"feed_info":            validations.FeedInfoValidations,
	"frequencies":          validations.FrequenciesValidations,
	"levels":               validations.LevelsValidations,
	"location_group_stops": validations.LocationGroupStopsValidations,
	"location_groups":      validations.LocationGroupsValidations,
	"municipalities":       validations.MunicipalitiesValidations,
	"networks":             validations.NetworksValidations,
	"pathways":             validations.PathwaysValidations,
	"periods":              validations.PeriodsValidations,
	"rider_categories":     validations.RiderCategoriesValidations,
	"route_networks":       validations.RouteNetworksValidations,
	"routes":               validations.RoutesValidations,
	"shapes":               validations.ShapesValidations,
	"stop_areas":           validations.StopAreasValidations,
	"stop_times":           validations.StopTimesValidations,
	"stops":                validations.StopsValidations,
	"timeframes":           validations.TimeFramesValidations,
	"transfers":            validations.TransfersValidations,
	"translations":         validations.TranslationsValidations,
	"trips":                validations.TripsValidations,
}
