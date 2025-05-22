package fare_attributes

import (
	"main/lib"
	"main/services"
	"main/types"
)

/*
# Attributes

	- File: [calendar_dates.txt]
	- Field: agency_id
	- Presence: Conditionally Required
	- Type: Foreign ID referencing agency.agency_id

# Description

Identifies the relevant agency for a fare.

Conditionally Required:
	- Required when the dataset contains data for multiple transit [agencies.txt].
	- Recommended otherwise.

[agencies.txt]: https://gtfs.org/schedule/reference/#agencytxt
[calendar_dates.txt]: https://gtfs.org/schedule/reference/#calendar_datestxt
*/
func AgencyIdValidation(severity *types.Severity, fareAttribute *types.FareAttribute, row int, gtfs *types.Gtfs) {
	s := types.SEVERITY_WARNING
	if severity != nil {
		s = *severity
	}

	addMessage := func(msg string) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "agency_id",
			FileName:     "fare_attributes.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     s,
			ValidationID: "agency_id_validation",
		})
	}

	//  Check if agency_id is required
	if fareAttribute.AgencyId == nil && len(gtfs.Files["agency"]) > 1 {
		s = types.SEVERITY_ERROR
		addMessage("Agency ID is required when there is more than one agency")
		return;
	}

	if s != types.SEVERITY_IGNORE && fareAttribute.AgencyId == nil {
		addMessage("Agency ID is " + lib.IfThenElse(s == types.SEVERITY_ERROR, "required", "recommended"))
		return;
	}

	// Check if agency_id exists in agencies.txt
	if fareAttribute.AgencyId != nil {
		if _, ok := gtfs.IdMap["agency"][*fareAttribute.AgencyId]; !ok {
			s = types.SEVERITY_ERROR
			addMessage("Agency ID does not exist in agencies.txt")
		}
	}
}