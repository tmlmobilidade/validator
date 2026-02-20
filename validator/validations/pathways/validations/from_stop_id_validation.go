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
		ctx.AddError(ctx.GetTranslatedMessage("from_stop_id_validation.not_found", *pathways.FromStopId))
		return
	}

	// Get the stop to check location_type
	stopRows, err := gtfs.GetRowsById("stops", *pathways.FromStopId)
	if err != nil || len(stopRows) == 0 {
		return // Already handled by foreign key check above
	}

	stop, err := gtfs.GetStop(stopRows[0])
	if err != nil {
		return
	}

	// Parse location_type
	var locationType int
	locationTypeStr := stop.LocationType
	if locationTypeStr == "" {
		locationType = 0 // Empty defaults to 0 (platform/stop)
	} else {
		if errMsg := lib.ParseStringToPrimitive(locationTypeStr, &locationType); errMsg != "" {
			// If location_type cannot be parsed, skip this validation
			return
		}
	}

	// Validate location_type
	// Allowed: platform (0 or empty), entrance/exit (2), generic node (3), boarding area (4)
	// Forbidden: station (1)
	if locationType == 1 {
		ctx.AddError(ctx.GetTranslatedMessage("from_stop_id_validation.invalid_location_type_station", *pathways.FromStopId, locationType))
		return
	}

	if locationType != 0 && locationType != 2 && locationType != 3 && locationType != 4 {
		ctx.AddError(ctx.GetTranslatedMessage("from_stop_id_validation.invalid_location_type", *pathways.FromStopId, locationType))
		return
	}

	// Check stop_access only for platforms/stops (location_type=0 or empty)
	// According to the spec: "stops (location_type=0 or empty) with stop_access=1, are forbidden"
	if locationType == 0 {
		var stopAccess int
		stopAccessStr := stop.StopAccess
		if stopAccessStr == "" {
			stopAccess = 0 // Empty defaults to 0
		} else {
			if errMsg := lib.ParseStringToPrimitive(stopAccessStr, &stopAccess); errMsg != "" {
				// If stop_access cannot be parsed, skip this check
				return
			}
		}

		if stopAccess == 1 {
			ctx.AddError(ctx.GetTranslatedMessage("from_stop_id_validation.forbidden_stop_access_1", *pathways.FromStopId))
			return
		}
	}
}
