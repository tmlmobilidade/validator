package trips

import (
	"main/i18n"
	"main/lib"
	"main/services"
	"main/types"
)

/*
# Attributes

  - File: trips.txt
  - Field: pattern_id
  - Presence: Optional (Required for "Transportes Metropolitanos de Lisboa")
  - Type: ID

# Description

Pattern to which the trips belongs.
Patterns correspond to the unfolding of the routes by the directions, if more than one (round trip).

Trips with the same pattern_id must have the same route_id, trip_headsign, direction_id, shape_id and the same stop sequence.
*/
func PatternIdValidation(trip *types.Trip, row int, gtfs *types.Gtfs, rules *types.TripsRules) bool {
	s := types.SEVERITY_IGNORE
	if rules != nil && rules.PatternId.Severity != "" {
		s = rules.PatternId.Severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "pattern_id",
			FileName:     "trips.txt",
			Message:      msg,
			Rows:         []int{row},
			Severity:     severity,
			ValidationID: "pattern_id_validation",
		})
	}

	if trip.PatternId == nil {
		if s == types.SEVERITY_IGNORE {
			return false
		}

		warn := lib.IfThenElse(s == types.SEVERITY_WARNING, i18n.AppTranslator.Get("pattern_id_validation.recommended"), i18n.AppTranslator.Get("pattern_id_validation.required"))
		addMessage(warn, s)
		return false
	}

	return true
}
