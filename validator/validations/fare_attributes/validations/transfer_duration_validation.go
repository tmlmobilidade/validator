package fare_attributes

import (
	"main/i18n"
	"main/services"
	"main/types"
)

/*
# Attributes
  - File: [fare_attributes.txt]
  - Field: transfer_duration
  - Presence: Optional
  - Type: Non-negative integer

# Description

Length of time in seconds before a transfer expires. When transfers=0 this field may be used to indicate how long a ticket is valid for or it may be left empty.

[fare_attributes.txt]: https://gtfs.org/schedule/reference/#fare_attributestxt
*/
func TransferDurationValidation(severity *types.Severity, fareAttribute *types.FareAttribute, row int, gtfs *types.Gtfs) {
	s := types.SEVERITY_IGNORE
	if severity != nil {
		s = *severity
	}

	addMessage := func(msg string) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "transfer_duration",
			FileName:     "fare_attributes.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     s,
			ValidationID: "transfer_duration_validation",
		})
	}

	if fareAttribute.TransferDuration == nil {
		if s != types.SEVERITY_IGNORE {
			addMessage(i18n.AppTranslator.Get("transfers_validation.required"))
		}
		return
	}

	if *fareAttribute.TransferDuration < 0 {
		addMessage(i18n.AppTranslator.Get("transfers_validation.invalid", *fareAttribute.TransferDuration))
	}
}
