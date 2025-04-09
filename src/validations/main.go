package validations

import (
	"main/src/types"
	"main/src/validations/afetacao"
	"main/src/validations/agency"
	"main/src/validations/archives"
	"main/src/validations/areas"
	"main/src/validations/attributions"
	"main/src/validations/booking_rules"
	"main/src/validations/calendar"
	"main/src/validations/calendar_dates"
	"main/src/validations/fare_attributes"
	"main/src/validations/fare_leg_join_rules"
	"main/src/validations/fare_leg_rules"
	"main/src/validations/fare_media"
	"main/src/validations/fare_products"
	"main/src/validations/fare_rules"
	"main/src/validations/fare_transfer_rules"
	"main/src/validations/feed_info"
	"main/src/validations/frequencies"
	"main/src/validations/levels"
	"main/src/validations/location_group_stops"
	"main/src/validations/location_groups"
	"main/src/validations/municipalities"
	"main/src/validations/networks"
	"main/src/validations/pathways"
	"main/src/validations/periods"
	"main/src/validations/rider_categories"
	"main/src/validations/route_networks"
	"main/src/validations/routes"
	"main/src/validations/shapes"
	"main/src/validations/stop_areas"
	"main/src/validations/stop_times"
	"main/src/validations/stops"
	"main/src/validations/timeframes"
	"main/src/validations/transfers"
	"main/src/validations/translations"
	"main/src/validations/trips"
)

var GTFS_FILE_RULES_MAP = map[string]func(gtfs types.Gtfs){
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
