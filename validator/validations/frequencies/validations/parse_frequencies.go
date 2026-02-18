package frequencies

import (
	"main/lib"
	"main/services"
	"main/types"
)

func ParseFrequencies(frequencies *types.FrequenciesRaw, row int) *types.Frequencies {
	var (
		frequency                  types.Frequencies = types.Frequencies{}
		tripId, endTime, startTime string
		exactTimes                 string
		headwaySecs                float64
		messages                   []types.Message
	)

	stringFields := map[string]*string{
		"trip_id":     &tripId,
		"end_time":    &endTime,
		"start_time":  &startTime,
		"exact_times": &exactTimes,
	}

	floatFields := map[string]*float64{
		"headway_secs": &headwaySecs,
	}

	addMessage := func(field, msg string) {
		messages = append(messages, types.Message{
			Field:        field,
			FileName:     "frequencies.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     types.SEVERITY_ERROR,
			ValidationID: "frequencies_parse",
		})
	}

	for field, target := range stringFields {
		if errMsg := lib.ParseStringToPrimitive(lib.GetFieldByTag(&frequencies, "gtfs", field), target); errMsg != "" {
			addMessage(field, errMsg)
		}
	}

	for field, target := range floatFields {
		if errMsg := lib.ParseStringToPrimitive(lib.GetFieldByTag(&frequencies, "gtfs", field), target); errMsg != "" {
			addMessage(field, errMsg)
		}
	}

	if len(messages) > 0 {
		services.AppMessageService.AddMessages(messages)
		return &frequency
	}

	frequency.TripId = lib.IfThenElse(frequencies.TripId != "", &tripId, nil)
	frequency.ExactTimes = exactTimes
	frequency.HeadwaySecs = float32(headwaySecs)
	frequency.StartTime = startTime
	frequency.EndTime = endTime

	return &frequency
}
