package routes

import (
	"main/lib"
	"main/services"
	"main/types"
)

/*
# Attributes

- File: [routes.txt]
- Field: line_id
- Presence: Conditionally Required
- Type: string

# Description

Line ID for the specified route.

Conditionally Required:
  - Required if line_id column contains a value.
  - Ignored if line_id column is empty.

[routes.txt]: https://gtfs.org/schedule/reference/#routestxt
*/
func LineIdValidation(route *types.Route, row int, gtfs *types.Gtfs, rules *types.RoutesRules) {
	ctx := lib.NewValidationContext("line_id", "routes.txt", "line_id_required", row, services.AppMessageService)
	if rules != nil && rules.LineId.Severity != "" {
		ctx.WithSeverity(rules.LineId.Severity)
	}

	// Check if line_id is required
	if route.LineId == nil || *route.LineId == "" {
		if ctx.ShouldSkip() {
			return
		}

		ctx.AddMessageWithSeverity(ctx.GetRequiredMessage("line_id_required.required", "line_id_required.recommended"))
	}

	// Check if line_id are same route_short_name
	if *route.LineId != *route.RouteShortName {
		ctx.AddError(ctx.GetTranslatedMessage("line_id_required.same_as_route_short_name", *route.LineId, *route.RouteShortName))
	}
}
