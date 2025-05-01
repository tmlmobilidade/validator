package agency

import (
	"main/types"
)

// DuplicateAgenciesValidation checks if an agency ID is already used by another agency.
// If a duplicate is found, it returns a validation message.
//
// Parameters:
//   - severity: Optional severity level for validation messages. Defaults to ERROR if nil
//   - agency: The agency being validated
//   - row: The row number of this agency in the GTFS file
//   - gtfs: The GTFS data containing the ID map
//
// Returns:
//   - A validation message if a duplicate agency ID is found, nil otherwise
func DuplicateAgenciesValidation(severity *types.Severity, agency *types.Agency, row int, gtfs *types.Gtfs) *types.Message {

	s := types.SEVERITY_ERROR
	if severity != nil {
		s = *severity
	}

	// Check if agency_id is nil or empty
	if agency.AgencyId == nil || *agency.AgencyId == "" {
		return nil
	}

	// Check if agency_id is already in the map and row is different
	if _, ok := gtfs.IdMap["agency"][*agency.AgencyId]; ok && row != gtfs.IdMap["agency"][*agency.AgencyId] {
		return &types.Message{
			Field: "agency_id",
			FileName: "agency.txt",
			Message: "Duplicate agency_id found. Agency IDs must be unique.",
			Rows: []int{row},
			Severity: s,
		}
	}
	
	return nil
}