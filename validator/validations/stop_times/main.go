package stop_times

import (
	"fmt"
	"main/config"
	"main/lib"
	"main/types"
	registry "main/validations"
	validations "main/validations/stop_times/validations"
	"strconv"
)

func init() {
	registry.Register("stop_times", RunValidations)
}

func RunValidations(gtfs types.Gtfs, rules *types.GtfsRules) {
	lib.AppLogger.Debug("Running StopTimes Validations...")

	// Pre-compute min/max stop sequences per trip_id for performance
	// This avoids N+1 queries in arrival_time validation
	lib.AppLogger.Debug("Pre-computing trip stop sequences...")
	tripStopSequences := make(map[string]types.TripStopSequence)

	err := gtfs.IterateStopTimes(func(i int, rawStopTime types.StopTimeRaw) error {
		// Parse stop_sequence from raw data for pre-computation
		if rawStopTime.TripId == "" || rawStopTime.StopSequence == "" {
			return nil
		}

		tripId := rawStopTime.TripId
		stopSeq, err := strconv.Atoi(rawStopTime.StopSequence)
		if err != nil {
			return nil // Skip invalid sequences
		}

		if seq, exists := tripStopSequences[tripId]; exists {
			if stopSeq < seq.Min {
				seq.Min = stopSeq
			}
			if stopSeq > seq.Max {
				seq.Max = stopSeq
			}
			tripStopSequences[tripId] = seq
		} else {
			tripStopSequences[tripId] = types.TripStopSequence{Min: stopSeq, Max: stopSeq}
		}
		return nil
	})
	if err != nil {
		lib.AppLogger.Error(fmt.Sprintf("Error pre-computing trip stop sequences: %v", err))
	}
	lib.AppLogger.Debug(fmt.Sprintf("Pre-computed stop sequences for %d trips", len(tripStopSequences)))

	// Create progress tracker
	tracker := lib.CreateProgressTracker(gtfs, "stop_times.txt", config.ProgressThresholdLarge)

	// Track previous stop_id per trip_id for consecutive stop_id validation
	previousStopIdByTrip := make(map[string]*string)

	err = gtfs.IterateStopTimes(func(i int, rawStopTimes types.StopTimeRaw) error {
		tracker.Track()
		stopTime := validations.ParseStopTimes(rawStopTimes, i)

		if stopTime == (types.StopTime{}) {
			return nil
		}

		var stopTimesRules *types.StopTimesRules
		if rules != nil {
			stopTimesRules = &rules.StopTimes
		}

		// Validate trip_id
		validations.TripIdValidation(&stopTime, i, &gtfs)

		// Validate arrival_time (pass cached trip stop sequences)
		validations.ArrivalTimeValidation(&stopTime, i, &gtfs, stopTimesRules, tripStopSequences)

		// Validate departure_time
		validations.DepartureTimeValidation(&stopTime, i, &gtfs, stopTimesRules)

		// Validate stop_id
		validations.StopIdValidation(&stopTime, i, &gtfs)

		// Validate consecutive stop_ids
		validations.ConsecutiveStopIdValidation(&stopTime, i, previousStopIdByTrip)

		// Validate location_group_id
		validations.LocationGroupIdValidation(&stopTime, i, &gtfs)

		// Validate start_pickup_drop_off_window
		validations.StartPickupDropOffWindowValidation(&stopTime, i, stopTimesRules)

		// Validate end_pickup_drop_off_window
		validations.EndPickupDropOffWindowValidation(&stopTime, i, stopTimesRules)

		// Validate pickup_type
		validations.PickupTypeValidation(&stopTime, i, stopTimesRules)

		// Validate stop_headsign
		validations.StopHeadsignValidation(&stopTime, i, stopTimesRules)

		// Validate continuous_drop_off
		validations.ContinuousDropOffValidation(&stopTime, i, stopTimesRules)

		// Validate continuous_pickup
		validations.ContinuousPickupValidation(&stopTime, i, stopTimesRules)

		// Validate drop_off_type
		validations.DropOffTypeValidation(&stopTime, i, stopTimesRules)

		// Validate shape_dist_traveled
		validations.ShapeDistTraveledValidation(&stopTime, i, stopTimesRules)

		// Validate timepoint
		validations.TimepointValidation(&stopTime, i, stopTimesRules)

		// Validate pickup_booking_rule_id
		validations.PickupBookingRuleIdValidation(&stopTime, i, &gtfs, stopTimesRules)

		// Validate drop_off_booking_rule_id
		validations.DropOffBookingRuleIdValidation(&stopTime, i, &gtfs, stopTimesRules)

		return nil
	})

	if err != nil {
		lib.AppLogger.Error(fmt.Sprintf("Error iterating stop times: %v", err))
	} else {
		lib.AppLogger.Debug(fmt.Sprintf("Completed stop_times.txt validation: %d rows processed", tracker.GetProcessedCount()))
	}
}
