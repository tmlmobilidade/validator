package routes

import (
	"main/lib"
	"main/services"
	"main/types"
)

/*
# Attributes

- File: [routes.txt]
- Field: line_short_name
- Presence: Conditionally Required
- Type: string

# Description

Line Short Name for the specified route.

Conditionally Required:
  - Required if line_id column contains a value.
  - Ignored if line_id column is empty.

[routes.txt]: https://gtfs.org/schedule/reference/#routestxt
*/
func LineShortNameValidation(route *types.Route, row int, gtfs *types.Gtfs, rules *types.RoutesRules) {
	ctx := lib.NewValidationContext("line_short_name", "routes.txt", "line_short_name_present_when_line_id_present", row, services.AppMessageService)
	if rules != nil && rules.LineShortName.Severity != "" {
		ctx.WithSeverity(rules.LineShortName.Severity)
	}

	// Check if line_id is present
	if route.LineId == nil || *route.LineId == "" {
		if ctx.ShouldSkip() {
			return
		}
		return
	}

	// Check if line_short_name is present
	if route.LineShortName == nil || *route.LineShortName == "" {
		if ctx.ShouldSkip() {
			return
		}

		ctx.AddError(ctx.GetTranslatedMessage("line_short_name_validation.required"))
	}

	// Validate are different from line_long_name
	if route.LineId != nil && *route.LineId != "" && route.LineLongName != nil && *route.LineLongName == *route.LineShortName {
		ctx.AddError(ctx.GetTranslatedMessage("line_short_name_validation.same_as_line_long_name"))
	}
}
