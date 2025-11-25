package routes

import (
	"main/i18n"
	"main/lib"
	"main/services"
	"main/types"
)

/*
# Attributes

- File: [routes.txt]
- Field: agency_id
- Presence: Conditionally Required
- Type: Foreign ID referencing agency.txt

# Description

Agency for the specified route.

Conditionally Required:
- Required if multiple agencies are defined in agency.txt.
- Recommended otherwise.

[routes.txt]: https://gtfs.org/schedule/reference/#routestxt
*/
func AgencyIdValidation(route *types.Route, row int, gtfs types.Gtfs, rules *types.RoutesRules) {
	s := types.SEVERITY_WARNING
	if rules != nil && rules.AgencyId.Severity != "" {
		s = rules.AgencyId.Severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "agency_id",
			FileName:     "routes.txt",
			ValidationID: "agency_id_validation",
			Message:      msg,
			Rows:         []int{row},
			Severity:     severity,
		})
	}

	numAgencies, _ := gtfs.GetTableCount("agency")

	// Check if agency_id exists and is valid
	if route.AgencyId != nil && *route.AgencyId != "" {
		// Check Foreign Key
		if !lib.GtfsIdMapKeyExists(&gtfs, "agency", *route.AgencyId) {
			addMessage(i18n.AppTranslator.Get("agency_id_validation.not_found", map[string]interface{}{"agency_id": *route.AgencyId}), types.SEVERITY_ERROR)
			return
		}
	}

	// Handle required vs recommended cases
	if numAgencies > 1 {
		if route.AgencyId == nil || *route.AgencyId == "" {
			addMessage(i18n.AppTranslator.Get("agency_id_validation.required_multiple_agencies"), types.SEVERITY_ERROR)
		}
	} else if route.AgencyId == nil && s != types.SEVERITY_IGNORE {
		warn := lib.IfThenElse(s == types.SEVERITY_WARNING, i18n.AppTranslator.Get("agency_id_validation.recommended"), i18n.AppTranslator.Get("agency_id_validation.required"))
		addMessage(warn, s)
	}
}
