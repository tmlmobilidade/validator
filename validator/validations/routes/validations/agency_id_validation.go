package routes

import (
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
func AgencyIdValidation(severity *types.Severity, route *types.Route, row int, gtfs types.Gtfs) {
	s := types.SEVERITY_WARNING
	if severity != nil {
		s = *severity
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

	numAgencies := len(gtfs.Agency)
	
	// Check if agency_id exists and is valid
	if route.AgencyId != nil && *route.AgencyId != "" {
		// Check Foreign Key
		if !lib.GtfsIdMapKeyExists(&gtfs, "agency", *route.AgencyId) {
			addMessage("agency_id '" + *route.AgencyId + "' is not a valid agency_id from agency.txt.", types.SEVERITY_ERROR)
			return
		}
	}

	// Handle required vs recommended cases
	if numAgencies > 1 {
		if route.AgencyId == nil || *route.AgencyId == "" {
			addMessage("agency_id is required when multiple agencies are defined in agency.txt.", types.SEVERITY_ERROR)
		}
	} else if route.AgencyId == nil && s != types.SEVERITY_IGNORE {
		warn := lib.IfThenElse(s == types.SEVERITY_WARNING, "recommended", "required")
		addMessage("agency_id is " + warn + " even if only one agency is defined in agency.txt.", s)
	}
} 