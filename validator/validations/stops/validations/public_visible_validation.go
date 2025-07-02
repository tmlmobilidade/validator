package stops

import (
	"main/lib"
	"main/services"
	"main/types"
)

/*
# Attributes

  - File: [stops.txt]
  - Field: public_visible
  - Presence: Optional
  - Type: Boolean

# Description

Describes if the stop is visible to the public.

[stops.txt]: https://gtfs.org/schedule/reference/#stopstxt
*/
func PublicVisibleValidation(stop *types.Stop, row int, rules *types.StopsRules) {
	s := types.SEVERITY_IGNORE
	if rules != nil && rules.PublicVisible.Severity != "" {
		s = rules.PublicVisible.Severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "public_visible",
			FileName:     "stops.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     severity,
			ValidationID: "public_visible_validation",
		})
	}

	if stop.PublicVisible == nil {
		if s == types.SEVERITY_IGNORE {
			return
		}

		warn := lib.IfThenElse(s == types.SEVERITY_ERROR, "public_visible is required", "public_visible is recommended")
		addMessage(warn, s)
		return
	}
}
