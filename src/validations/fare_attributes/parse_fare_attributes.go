package fare_attributes

import (
	"main/src/lib"
	"main/src/types"
	"strconv"
)

type parseFareAttributesValidation struct {
	*types.Validation
}

func NewParseFareAttributesValidation(severity *types.Severity) *parseFareAttributesValidation {
	s := types.SEVERITY_ERROR
	if severity != nil {
		s = *severity
	}

	return &parseFareAttributesValidation{
		Validation: &types.Validation{
			ID:          "parse_fare_attributes",
			Description: "Validate fare attributes data",
			Severity:    s,
		},
	}
}

func (v *parseFareAttributesValidation) Validate(gtfs types.Gtfs) (fareAttributes []types.FareAttribute, messages []types.Message) {
	fareIds := make(map[string]bool)

	for i, fare := range gtfs.Files["fare_attributes"] {
		fareAttribute, fareMessages := parseFareAttribute(fare, gtfs.IdMap["agency"])
		fareAttributes = append(fareAttributes, fareAttribute)

		// Check for duplicate fare IDs
		if fareAttribute.FareId != "" {
			if fareIds[fareAttribute.FareId] {
				messages = append(messages, types.Message{
					Field:        "fare_id",
					FileName:     "fare_attributes.txt",
					Message:      "Duplicate fare_id found. Fare IDs must be unique.",
					Row:          i + 1,
					Severity:     v.Severity,
					ValidationID: v.ID,
				})
			}
			fareIds[fareAttribute.FareId] = true
		}

		// Update row number and other fields for each message
		for _, msg := range fareMessages {
			msg.Row = i + 1
			msg.FileName = "fare_attributes.txt"
			msg.Severity = v.Severity
			msg.ValidationID = v.ID
			messages = append(messages, msg)
		}
	}
	return fareAttributes, messages
}

func parseFareAttribute(m map[string]string, agencyIds map[string]int) (fareAttribute types.FareAttribute, messages []types.Message) {
	var parsingErrors []string

	// Parse Required Fields
	lib.ParseStringToPrimitive(m["fare_id"], &fareAttribute.FareId, &parsingErrors)
	lib.ParseStringToPrimitive(m["currency_type"], &fareAttribute.CurrencyType, &parsingErrors)

	// Parse price as float
	if m["price"] != "" {
		price, err := strconv.ParseFloat(m["price"], 64)
		if err != nil {
			messages = append(messages, types.Message{
				Field:   "price",
				Message: "Price must be a valid non-negative float value.",
			})
		} else if price < 0 {
			messages = append(messages, types.Message{
				Field:   "price",
				Message: "Price must be non-negative.",
			})
		} else {
			fareAttribute.Price = &price
		}
	}

	// Parse payment_method as int
	if m["payment_method"] != "" {
		paymentMethod, err := strconv.Atoi(m["payment_method"])
		if err != nil {
			messages = append(messages, types.Message{
				Field:   "payment_method",
				Message: "Payment method must be a valid integer (0 or 1).",
			})
		} else {
			fareAttribute.PaymentMethod = &paymentMethod
		}
	}

	// Parse transfers
	if m["transfers"] != "" {
		transfers, err := strconv.Atoi(m["transfers"])
		if err != nil {
			messages = append(messages, types.Message{
				Field:   "transfers",
				Message: "Transfers must be a valid integer (0, 1, or 2) or empty for unlimited transfers.",
			})
		} else {
			fareAttribute.Transfers = &transfers
		}
	}

	// Parse Optional Fields
	if m["agency_id"] != "" {
		agencyId := m["agency_id"]
		fareAttribute.AgencyId = &agencyId
	}

	if m["transfer_duration"] != "" {
		transferDuration, err := strconv.Atoi(m["transfer_duration"])
		if err != nil {
			messages = append(messages, types.Message{
				Field:   "transfer_duration",
				Message: "Transfer duration must be a valid non-negative integer.",
			})
		} else if transferDuration < 0 {
			messages = append(messages, types.Message{
				Field:   "transfer_duration",
				Message: "Transfer duration must be non-negative.",
			})
		} else {
			fareAttribute.TransferDuration = &transferDuration
		}
	}

	// Validate Required Fields
	if fareAttribute.FareId == "" {
		messages = append(messages, types.Message{
			Field:   "fare_id",
			Message: "Fare ID is required and must be unique.",
		})
	}

	if fareAttribute.Price == nil {
		messages = append(messages, types.Message{
			Field:   "price",
			Message: "Price is required.",
		})
	}

	if fareAttribute.CurrencyType == "" {
		messages = append(messages, types.Message{
			Field:   "currency_type",
			Message: "Currency type is required.",
		})
	}

	if fareAttribute.PaymentMethod == nil {
		messages = append(messages, types.Message{
			Field:   "payment_method",
			Message: "Payment method is required.",
		})
	} else {
		// Validate payment_method enum values
		validPaymentMethods := map[int]bool{0: true, 1: true}
		if !validPaymentMethods[*fareAttribute.PaymentMethod] {
			messages = append(messages, types.Message{
				Field:   "payment_method",
				Message: "Invalid payment_method value. Valid values are 0 (paid on board) or 1 (paid before boarding).",
			})
		}
	}

	if fareAttribute.Transfers != nil {
		// Validate transfers enum values
		validTransfers := map[int]bool{0: true, 1: true, 2: true}
		if !validTransfers[*fareAttribute.Transfers] {
			messages = append(messages, types.Message{
				Field:   "transfers",
				Message: "Invalid transfers value. Valid values are 0 (no transfers), 1 (one transfer), 2 (two transfers), or empty (unlimited).",
			})
		}
	}

	// Validate agency_id if provided
	if fareAttribute.AgencyId != nil {
		_, ok := agencyIds[*fareAttribute.AgencyId]
		if !ok {
			messages = append(messages, types.Message{
				Field:   "agency_id",
				Message: "Agency ID must reference a valid agency_id from agency.txt.",
			})
		}
	}

	// Check if agency_id is required (when multiple agencies exist)
	if len(agencyIds) > 1 && fareAttribute.AgencyId == nil {
		messages = append(messages, types.Message{
			Field:   "agency_id",
			Message: "Agency ID is required when multiple agencies are defined in agency.txt.",
		})
	}

	return fareAttribute, messages
}
