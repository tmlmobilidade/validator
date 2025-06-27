package validations

import (
	"main/types"
	"main/validations/afetacao"
	"main/validations/agency"
	"main/validations/archives"
	"main/validations/areas"
	"main/validations/attributions"
	"main/validations/booking_rules"
	"main/validations/calendar"
	"main/validations/calendar_dates"
	"main/validations/fare_attributes"
	"main/validations/fare_leg_join_rules"
	"main/validations/fare_leg_rules"
	"main/validations/fare_media"
	"main/validations/fare_products"
	"main/validations/fare_rules"
	"main/validations/fare_transfer_rules"
	"main/validations/feed_info"
	"main/validations/frequencies"
	"main/validations/levels"
	"main/validations/location_group_stops"
	"main/validations/location_groups"
	"main/validations/municipalities"
	"main/validations/networks"
	"main/validations/pathways"
	"main/validations/periods"
	"main/validations/rider_categories"
	"main/validations/route_networks"
	"main/validations/routes"
	"main/validations/shapes"
	"main/validations/stop_areas"
	"main/validations/stop_times"
	"main/validations/stops"
	"main/validations/timeframes"
	"main/validations/transfers"
	"main/validations/translations"
	"main/validations/trips"
)

var GTFS_FILE_RULES_MAP = map[string]func(gtfs types.Gtfs, rules *types.GtfsRules){
	"afetacao":             afetacao.RunValidations,
	"agency":               agency.RunValidations,
	"archives":             archives.RunValidations,
	"areas":                areas.RunValidations,
	"attributions":         attributions.RunValidations,
	"booking_rules":        booking_rules.RunValidations,
	"calendar":             calendar.RunValidations,
	"calendar_dates":       calendar_dates.RunValidations,
	"fare_attributes":      fare_attributes.RunValidations,
	"fare_leg_join_rules":  fare_leg_join_rules.RunValidations,
	"fare_leg_rules":       fare_leg_rules.RunValidations,
	"fare_media":           fare_media.RunValidations,
	"fare_products":        fare_products.RunValidations,
	"fare_rules":           fare_rules.RunValidations,
	"fare_transfer_rules":  fare_transfer_rules.RunValidations,
	"feed_info":            feed_info.RunValidations,
	"frequencies":          frequencies.RunValidations,
	"levels":               levels.RunValidations,
	"location_group_stops": location_group_stops.RunValidations,
	"location_groups":      location_groups.RunValidations,
	"municipalities":       municipalities.RunValidations,
	"networks":             networks.RunValidations,
	"pathways":             pathways.RunValidations,
	"periods":              periods.RunValidations,
	"rider_categories":     rider_categories.RunValidations,
	"route_networks":       route_networks.RunValidations,
	"routes":               routes.RunValidations,
	"shapes":               shapes.RunValidations,
	"stop_areas":           stop_areas.RunValidations,
	"stop_times":           stop_times.RunValidations,
	"stops":                stops.RunValidations,
	"timeframes":           timeframes.RunValidations,
	"transfers":            transfers.RunValidations,
	"translations":         translations.RunValidations,
	"trips":                trips.RunValidations,
}
