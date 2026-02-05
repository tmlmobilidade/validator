package trips

import (
	"main/lib"
	"main/services"
	"main/types"
)

/*
# Attributes

  - File: [trips.txt]
  - Field: trip_id
  - Presence: Required
  - Type: Unique ID

# Description

Identifies a trip.

[trips.txt]: https://gtfs.org/schedule/reference/#trips
*/
func TripIdValidation(trip *types.Trip, row int, gtfs *types.Gtfs) {
	ctx := lib.NewValidationContext("trip_id", "trips.txt", "trip_id_validation", row, services.AppMessageService)

	if trip.TripId == nil {
		ctx.AddError(ctx.GetTranslatedMessage("trip_id_validation.required"))
		return
	}

	// rows, err := gtfs.GetRowsById("trips", *trip.TripId)
	// if err == nil && len(rows) > 1 {
	// 	ctx.AddError(ctx.GetTranslatedMessage("trip_id_validation.duplicate", map[string]interface{}{"trip_id": *trip.TripId}))
	// 	return
	// }

	if trip.TripId != nil {
		// Check if trip_id is Unique ID
		rows, err := gtfs.GetRowsById("trips", *trip.TripId)
		if err != nil {
			// Fallback to in-memory IdMap if database is not available
			if gtfs.IdMap != nil {
				if tripRows, exists := gtfs.IdMap["trips"]; exists {
					if indices, found := tripRows[*trip.TripId]; found {
						if len(indices) > 1 {
							ctx.AddError(ctx.GetTranslatedMessage("trip_id_validation.duplicate", *trip.TripId))
							return
						}
					}
				}
			}
			return
		}
		if len(rows) > 1 {
			ctx.AddError(ctx.GetTranslatedMessage("trip_id_validation.duplicate", *trip.TripId))
			return
		}
	}
}
