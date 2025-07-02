package agency

import (
	"fmt"
	"main/lib"
	"main/services"
	"main/types"
	"slices"
)

/*
# Attributes

  - File: [agency.txt]
  - Field: agency_id
  - Presence: Conditionally Required
  - Type: Unique ID

# Description

Identifies a transit brand which is often synonymous with a transit agency.
Note that in some cases, such as when a single agency operates multiple separate services, agencies and brands are distinct.

This document uses the term "agency" in place of "brand". A dataset may contain data from multiple agencies.

Conditionally Required:
  - Required when the dataset contains data for multiple transit agencies.
  - Recommended otherwise.

[agency.txt]: https://gtfs.org/schedule/reference/#agencytxt
*/
func AgencyIdValidation(agency *types.Agency, row int, gtfs types.Gtfs, rules *types.AgencyRules) {
	s := types.SEVERITY_WARNING
	if rules != nil {
		s = rules.AgencyId.Severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "agency_id",
			FileName:     "agency.txt",
			Message:      msg,
			Rows:         []int{row},
			Severity:     severity,
			ValidationID: "agency_id_validation",
		})
	}

	//  Check if agency_id is required
	if agency.AgencyId == nil {
		if len(gtfs.Agency) > 1 {
			addMessage("Agency ID is required when there is more than one agency", types.SEVERITY_ERROR)
			return
		}

		if s == types.SEVERITY_IGNORE {
			return
		}

		warn := lib.IfThenElse(s == types.SEVERITY_ERROR, "Agency ID is required", "Agency ID is recommended")
		addMessage(warn, s)
		return
	}

	if agency.AgencyId != nil {
		// Check if agency_id is Unique ID
		if _, ok := gtfs.IdMap["agency"][*agency.AgencyId]; ok && len(gtfs.IdMap["agency"][*agency.AgencyId]) > 1 {
			addMessage("Duplicate agency_id found. Agency IDs must be unique.", types.SEVERITY_ERROR)
		}

		// Validate rules
		if rules != nil && rules.AgencyId.Options != nil {
			if slices.Contains(*rules.AgencyId.Options, types.ALL_OPTIONS) {
				return
			}

			if slices.Contains(*rules.AgencyId.Options, *agency.AgencyId) {
				return
			}

			addMessage(fmt.Sprintf("Agency ID is not allowed: %s", *agency.AgencyId), s)
			return
		}
	}
}
