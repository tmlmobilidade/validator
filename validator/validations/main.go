package validations

import (
	"main/validator/types"
	"main/validator/validations/afetacao"
	"main/validator/validations/agency"
	"main/validator/validations/archives"
	"main/validator/validations/areas"
	"main/validator/validations/attributions"
	"main/validator/validations/booking_rules"
	"main/validator/validations/calendar"
	"main/validator/validations/calendar_dates"
	"main/validator/validations/fare_attributes"
	"main/validator/validations/fare_leg_join_rules"
	"main/validator/validations/fare_leg_rules"
	"main/validator/validations/fare_media"
	"main/validator/validations/fare_products"
	"main/validator/validations/fare_rules"
	"main/validator/validations/fare_transfer_rules"
	"main/validator/validations/feed_info"
	"main/validator/validations/frequencies"
	"main/validator/validations/levels"
	"main/validator/validations/location_group_stops"
	"main/validator/validations/location_groups"
	"main/validator/validations/municipalities"
	"main/validator/validations/networks"
	"main/validator/validations/pathways"
	"main/validator/validations/periods"
	"main/validator/validations/rider_categories"
	"main/validator/validations/route_networks"
	"main/validator/validations/routes"
	"main/validator/validations/shapes"
	"main/validator/validations/stop_areas"
	"main/validator/validations/stop_times"
	"main/validator/validations/stops"
	"main/validator/validations/timeframes"
	"main/validator/validations/transfers"
	"main/validator/validations/translations"
	"main/validator/validations/trips"
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
