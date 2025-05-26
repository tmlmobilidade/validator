package stops

import (
	"main/lib"
	"main/services"
	"main/types"
)

func TtsStopNameValidation(severity *types.Severity, stop *types.Stop, row int) {
	s := types.SEVERITY_IGNORE
	if severity != nil {
		s = *severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "tts_stop_name",
			FileName:     "stops.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     severity,
			ValidationID: "tts_stop_name_validation",
		})
	}


	if stop.TtsStopName == nil || *stop.TtsStopName == "" {
		if s == types.SEVERITY_IGNORE {
			return
		}

		warn := lib.IfThenElse(s == types.SEVERITY_ERROR, "tts_stop_name is required", "tts_stop_name is recommended")
		addMessage(warn, s)
		return
	}
	
}