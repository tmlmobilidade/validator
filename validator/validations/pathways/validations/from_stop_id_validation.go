package pathways

import (
	"main/lib"
	"main/services"
	"main/types"
)

/*
# Attributes

- File: [pathways.txt]
- Field: from_stop_id
- Presence: Required
- Type: Foreign ID referencing [stops.stop_id]

# Description

Identifies the stop where the pathway begins.

Must contain a stop_id that identifies a platform (location_type=0 or empty), entrance/exit (location_type=2), generic node (location_type=3) or boarding area (location_type=4).

Values for stop_id that identify stations (location_type=1), or stops (location_type=0 or empty) with stop_access=1, are forbidden.

[pathways.txt]: https://gtfs.org/schedule/reference/#pathwaystxt
[stops.stop_id]: https://gtfs.org/schedule/reference/#stopstxt
*/
func FromStopIdValidation(pathways *types.Pathways, row int, gtfs *types.Gtfs, rules *types.PathwaysRules) {
	ctx := lib.NewValidationContext("from_stop_id", "pathways.txt", "from_stop_id_validation", row, services.AppMessageService)
	if rules != nil && rules.FromStopId.Severity != "" {
		ctx.WithSeverity(rules.FromStopId.Severity)
	}

	if pathways.FromStopId == nil {
		ctx.AddMessageWithSeverity(ctx.GetTranslatedMessage("from_stop_id_validation.required"))
		return
	}

	// Check Foreign Key
	if !lib.GtfsIdMapKeyExists(gtfs, "stops", *pathways.FromStopId) {
		ctx.AddError(ctx.GetTranslatedMessage("from_stop_id_validation.not_found", map[string]any{"from_stop_id": *pathways.FromStopId}))
		return
	}
}
