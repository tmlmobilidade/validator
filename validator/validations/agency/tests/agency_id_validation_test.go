package agency

import (
	"main/services"
	"main/types"
	validations "main/validations/agency/validations"
	"testing"
)

func TestAgencyIdValidation_Required(t *testing.T) {
	severity := types.SEVERITY_ERROR
	agency := &types.Agency{AgencyId: nil}
	gtfs := &types.Gtfs{Files: map[string][]map[string]string{"agency": {{}}}}
	gtfs.Files["agency"] = append(gtfs.Files["agency"], map[string]string{}) // Add a second agency for >1
	validations.AgencyIdValidation(&severity, agency, 1, gtfs)
	assertMessage(t, Assertion{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Agency ID is required when there is more than one agency",
	})
	services.AppMessageService.Clear()
}

func TestAgencyIdValidation_Recommended(t *testing.T) {
	severity := types.SEVERITY_WARNING
	agency := &types.Agency{AgencyId: nil}
	gtfs := &types.Gtfs{Files: map[string][]map[string]string{"agency": {{}}}}
	validations.AgencyIdValidation(&severity, agency, 2, gtfs)
	assertMessage(t, Assertion{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalWarnings,
		Message: "Agency ID is recomended",
	})
	services.AppMessageService.Clear()
}

func TestAgencyIdValidation_Unique(t *testing.T) {
	id := "unique"
	agency := &types.Agency{AgencyId: &id}
	gtfs := &types.Gtfs{IdMap: map[string]map[string][]int{"agency": {"unique": {1}}}}
	validations.AgencyIdValidation(nil, agency, 3, gtfs)
	assertMessage(t, Assertion{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalWarnings,
		Message: "Agency ID is recomended",
	})
	services.AppMessageService.Clear()
}

func TestAgencyIdValidation_Duplicate(t *testing.T) {
	id := "duplicate"
	agency := &types.Agency{AgencyId: &id}
	gtfs := &types.Gtfs{IdMap: map[string]map[string][]int{"agency": {"duplicate": {1, 2}}}}
	validations.AgencyIdValidation(nil, agency, 4, gtfs)
	assertMessage(t, Assertion{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Duplicate agency_id found. Agency IDs must be unique.",
	})
	services.AppMessageService.Clear()
} 