package validations

import (
	"fmt"
	"main/src/lib"
	"main/src/models"
	"main/src/validations/agency"
	"main/src/validations/stops"
)

var GTFS_FILE_RULES_MAP = map[string]func(gtfsData models.Gtfs){
	"afetacao":             AfetacaoValidations,
	"agency":               AgencyValidations,
	"archives":             ArchivesValidations,
	"areas":                AreasValidations,
	"attributions":         AttributionsValidations,
	"booking_rules":        BookingRulesValidations,
	"calendar":             CalendarValidations,
	"calendar_dates":       CalendarDatesValidations,
	"fare_attributes":      FareAttributesValidations,
	"fare_leg_join_rules":  FareLegJoinRulesValidations,
	"fare_leg_rules":       FareLegRulesValidations,
	"fare_media":           FareMediaValidations,
	"fare_products":        FareProductsValidations,
	"fare_rules":           FareRulesValidations,
	"fare_transfer_rules":  FareTransferRulesValidations,
	"feed_info":            FeedInfoValidations,
	"frequencies":          FrequenciesValidations,
	"levels":               LevelsValidations,
	"location_group_stops": LocationGroupStopsValidations,
	"location_groups":      LocationGroupsValidations,
	"municipalities":       MunicipalitiesValidations,
	"networks":             NetworksValidations,
	"pathways":             PathwaysValidations,
	"periods":              PeriodsValidations,
	"rider_categories":     RiderCategoriesValidations,
	"route_networks":       RouteNetworksValidations,
	"routes":               RoutesValidations,
	"shapes":               ShapesValidations,
	"stop_areas":           StopAreasValidations,
	"stop_times":           StopTimesValidations,
	"stops":                StopsValidations,
	"timeframes":           TimeFramesValidations,
	"transfers":            TransfersValidations,
	"translations":         TranslationsValidations,
	"trips":                TripsValidations,
}

func AgencyValidations(gtfsData models.Gtfs) {
	// Parse Agency
	for _, a := range gtfsData["agency"] {
		ag, errors := agency.ParseAgency(a)
		if len(errors) > 0 {
			fmt.Println("Errors:", errors)
		} else {
			lib.PrintMap(ag)
		}
	}
}

func AfetacaoValidations(gtfsData models.Gtfs) {
	fmt.Println("Running Afetacao Validations...")
}

func ArchivesValidations(gtfsData models.Gtfs) {
	fmt.Println("Running Archives Validations...")
}

func AreasValidations(gtfsData models.Gtfs) {
	fmt.Println("Running Areas Validations...")
}

func AttributionsValidations(gtfsData models.Gtfs) {
	fmt.Println("Running Attributions Validations...")
}

func BookingRulesValidations(gtfsData models.Gtfs) {
	fmt.Println("Running Booking Rules Validations...")
}

func CalendarValidations(gtfsData models.Gtfs) {
	fmt.Println("Running Calendar Validations...")
}

func CalendarDatesValidations(gtfsData models.Gtfs) {
	fmt.Println("Running CalendarDates Validations...")
}

func FareAttributesValidations(gtfsData models.Gtfs) {
	fmt.Println("Running FareAttributes Validations...")
}

func FareLegJoinRulesValidations(gtfsData models.Gtfs) {
	fmt.Println("Running FareLegJoinValidations Validations...")
}

func FareLegRulesValidations(gtfsData models.Gtfs) {
	fmt.Println("Running FareLegValidations Validations...")
}

func FareMediaValidations(gtfsData models.Gtfs) {
	fmt.Println("Running FareMedia Validations...")
}

func FareProductsValidations(gtfsData models.Gtfs) {
	fmt.Println("Running FareProducts Validations...")
}

func FareRulesValidations(gtfsData models.Gtfs) {
	fmt.Println("Running FareRules Validations...")
}

func FareTransferRulesValidations(gtfsData models.Gtfs) {
	fmt.Println("Running FareTransfer Rules Validations...")
}

func FeedInfoValidations(gtfsData models.Gtfs) {
	fmt.Println("Running FeedInfo Validations...")
}

func FrequenciesValidations(gtfsData models.Gtfs) {
	fmt.Println("Running Frequencies Validations...")
}

func LevelsValidations(gtfsData models.Gtfs) {
	fmt.Println("Running Levels Validations...")
}

func LocationGroupStopsValidations(gtfsData models.Gtfs) {
	fmt.Println("Running LocationGroupStops Validations...")
}

func LocationGroupsValidations(gtfsData models.Gtfs) {
	fmt.Println("Running LocationGroups Validations...")
}

func MunicipalitiesValidations(gtfsData models.Gtfs) {
	fmt.Println("Running Municipalities Validations...")
}

func NetworksValidations(gtfsData models.Gtfs) {
	fmt.Println("Running Networks Validations...")
}

func PathwaysValidations(gtfsData models.Gtfs) {
	fmt.Println("Running Pathways Validations...")
}

func PeriodsValidations(gtfsData models.Gtfs) {
	fmt.Println("Running Periods Validations...")
}

func RiderCategoriesValidations(gtfsData models.Gtfs) {
	fmt.Println("Running RiderCategories Validations...")
}

func RouteNetworksValidations(gtfsData models.Gtfs) {
	fmt.Println("Running RouteNetworks Validations...")
}

func RoutesValidations(gtfsData models.Gtfs) {
	fmt.Println("Running Routes Validations...")
}

func ShapesValidations(gtfsData models.Gtfs) {
	fmt.Println("Running Shapes Validations...")
}

func StopAreasValidations(gtfsData models.Gtfs) {
	fmt.Println("Running StopAreas Validations...")
}

func StopTimesValidations(gtfsData models.Gtfs) {
	fmt.Println("Running StopTimes Validations...")
}

func StopsValidations(gtfsData models.Gtfs) {
	// Parse Agency
	for _, a := range gtfsData["stops"] {
		st, errors := stops.ParseStop(a)
		if len(errors) > 0 {
			fmt.Println("Errors:", errors)
		} else {
			lib.PrintMap(st)
		}
	}
}

func TimeFramesValidations(gtfsData models.Gtfs) {
	fmt.Println("Running TimeFrames Validations...")
}

func TransfersValidations(gtfsData models.Gtfs) {
	fmt.Println("Running Transfers Validations...")
}

func TranslationsValidations(gtfsData models.Gtfs) {
	fmt.Println("Running Translations Validations...")
}

func TripsValidations(gtfsData models.Gtfs) {
	fmt.Println("Running Trips Validations...")
}
