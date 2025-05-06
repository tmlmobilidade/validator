package agency

import (
	"main/lib"
	"main/services"
	"main/types"
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
func AgencyIdValidation(severity *types.Severity, agency *types.Agency, row int, gtfs *types.Gtfs) {
	s := types.SEVERITY_WARNING
	if severity != nil {
		s = *severity
	}

	message := types.Message{
		Field: "agency_id",
		FileName: "agency.txt",
		Message: lib.IfThenElse(s == types.SEVERITY_ERROR, "Agency ID is required", "Agency ID is recomended"),
		Rows: []int{row},
		Severity: s,
		ValidationID: "agency_id_validation",
	}

	//  Check if agency_id is required
	if agency.AgencyId == nil && len(gtfs.Files["agency"]) > 1 {
		message.Message = "Agency ID is required when there is more than one agency"
		message.Severity = types.SEVERITY_ERROR
	}

	if agency.AgencyId != nil {
		// Check if agency_id is Unique ID
		if _, ok := gtfs.IdMap["agency"][*agency.AgencyId]; ok && row != gtfs.IdMap["agency"][*agency.AgencyId] {
			services.AppMessageService.AddMessage(types.Message{
				Field: "agency_id",
				FileName: "agency.txt",
				ValidationID: "duplicate_agencies_validation",
				Message: "Duplicate agency_id found. Agency IDs must be unique.",
				Rows: []int{row},
				Severity: types.SEVERITY_ERROR,
			})
		}
	}

	services.AppMessageService.AddMessage(message)
}