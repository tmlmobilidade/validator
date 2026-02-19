package pathways

import (
	"main/lib"
	"main/services"
	"main/types"
)

/*
# Attributes

- File: [pathways.txt]
- Field: to_stop_id
- Presence: Required
- Type: Foreign ID referencing [stops.stop_id]

# Description

Identifies the stop where the pathway ends.

Must contain a stop_id that identifies a platform (location_type=0 or empty), entrance/exit (location_type=2), generic node (location_type=3) or boarding area (location_type=4).

Values for stop_id that identify stations (location_type=1), or stops (location_type=0 or empty) with stop_access=1, are forbidden.

[pathways.txt]: https://gtfs.org/schedule/reference/#pathwaystxt
[stops.stop_id]: https://gtfs.org/schedule/reference/#stopstxt
*/

func ToStopIdValidation(pathways *types.Pathways, row int, gtfs *types.Gtfs, rules *types.PathwaysRules) {
	ctx := lib.NewValidationContext("to_stop_id", "pathways.txt", "to_stop_id_validation", row, services.AppMessageService)
	if rules != nil && rules.ToStopId.Severity != "" {
		ctx.WithSeverity(rules.ToStopId.Severity)
	}

	if pathways.ToStopId == nil {
		ctx.AddMessageWithSeverity(ctx.GetTranslatedMessage("to_stop_id_validation.required"))
		return
	}

	// Check Foreign Key
	if !lib.GtfsIdMapKeyExists(gtfs, "stops", *pathways.ToStopId) {
		ctx.AddError(ctx.GetTranslatedMessage("to_stop_id_validation.not_found", map[string]any{"to_stop_id": *pathways.ToStopId}))
		return
	}
}
