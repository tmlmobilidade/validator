package routes

import (
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
func AgencyIdValidation(route *types.Route, row int, gtfs *types.Gtfs) {
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

	agencyMap, hasAgencies := gtfs.IdMap["agency"]
	numAgencies := 0
	if hasAgencies {
		numAgencies = len(agencyMap)
	}

	if numAgencies > 1 {
		// Required
		if route.AgencyId == nil || *route.AgencyId == "" {
			addMessage("agency_id is required when multiple agencies are defined in agency.txt.", types.SEVERITY_ERROR)
			return
		}
		if _, ok := agencyMap[*route.AgencyId]; !ok {
			addMessage("agency_id must reference a valid agency_id from agency.txt.", types.SEVERITY_ERROR)
			return
		}
	} else {
		// Recommended
		if route.AgencyId == nil || *route.AgencyId == "" {
			addMessage("agency_id is recommended even if only one agency is defined in agency.txt.", types.SEVERITY_WARNING)
			return
		}
		if _, ok := agencyMap[*route.AgencyId]; !ok {
			addMessage("agency_id must reference a valid agency_id from agency.txt.", types.SEVERITY_ERROR)
			return
		}
	}
} 