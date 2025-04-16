package fare_attributes

import (
	"main/validator/lib"
	"main/validator/types"
)

type parseFareAttributeValidation struct {
	*types.Validation
}

func NewParseFareAttributeValidation(severity *types.Severity) *parseFareAttributeValidation {
	s := types.SEVERITY_ERROR
	if severity != nil {
		s = *severity
	}

	return &parseFareAttributeValidation{
		Validation: &types.Validation{
			ID:          "parse_fare_attribute",
			Description: "Validate fare attribute data",
			Severity:    s,
		},
	}
}

func (v *parseFareAttributeValidation) Validate(gtfs types.Gtfs) (fareAttributes []types.FareAttribute, messages []types.Message) {
	fareIds := make(map[string]bool)

	// Check if multiple agencies exist
	multipleAgencies := len(gtfs.Files["agency"]) > 1

	for i, fareAttribute := range gtfs.Files["fare_attributes"] {
		fareAttr, fareAttrMessages := parseFareAttribute(fareAttribute, multipleAgencies, gtfs.IdMap["agency"])
		fareAttributes = append(fareAttributes, fareAttr)

		// Check for duplicate fare IDs
		if fareAttr.FareId != "" {
			if fareIds[fareAttr.FareId] {
				messages = append(messages, types.Message{
					Field:        "fare_id",
					FileName:     "fare_attributes.txt",
					Message:      "Duplicate fare_id found. Fare IDs must be unique.",
					Rows:         []int{i + 1},
					Severity:     v.Severity,
					ValidationID: v.ID,
				})
			}
			fareIds[fareAttr.FareId] = true
		}

		// Update row number and other fields for each message
		for _, msg := range fareAttrMessages {
			msg.Rows = []int{i + 1}
			msg.FileName = "fare_attributes.txt"
			msg.Severity = v.Severity
			msg.ValidationID = v.ID
			messages = append(messages, msg)
		}
	}
	return fareAttributes, messages
}

func parseFareAttribute(m map[string]string, multipleAgencies bool, agencyIdMap map[string]int) (fareAttribute types.FareAttribute, messages []types.Message) {
	var parsingErrors []string

	// Convert Optional Primitive Values
	var agencyId string
	var transferDuration int
	var transfers int

	lib.ParseStringToPrimitive(m["transfers"], &transfers, &parsingErrors)
	lib.ParseStringToPrimitive(m["agency_id"], &agencyId, &parsingErrors)
	lib.ParseStringToPrimitive(m["transfer_duration"], &transferDuration, &parsingErrors)

	fareAttribute.AgencyId = lib.IfThenElse(m["agency_id"] != "", &agencyId, nil)
	fareAttribute.TransferDuration = lib.IfThenElse(m["transfer_duration"] != "", &transferDuration, nil)

	// Convert Required Values

	lib.ParseStringToPrimitive(m["fare_id"], &fareAttribute.FareId, &parsingErrors)
	lib.ParseStringToPrimitive(m["currency_type"], &fareAttribute.CurrencyType, &parsingErrors)

	var price float64
	lib.ParseStringToPrimitive(m["price"], &price, &parsingErrors)
	fareAttribute.Price = lib.IfThenElse(m["price"] != "", &price, nil)

	// Convert Required Enums
	var paymentMethod int
	lib.ParseStringToPrimitive(m["payment_method"], &paymentMethod, &parsingErrors)

	fareAttribute.PaymentMethod = lib.IfThenElse(m["payment_method"] != "", &paymentMethod, nil)
	fareAttribute.Transfers = lib.IfThenElse(m["transfers"] != "", &transfers, nil)

	// Handle parsing errors
	if len(parsingErrors) > 0 {
		for _, err := range parsingErrors {
			messages = append(messages, types.Message{
				Field:   "fare_attributes.txt",
				Message: err,
			})
		}
	}

	// Validate required fields
	if fareAttribute.FareId == "" {
		messages = append(messages, types.Message{
			Field:   "fare_id",
			Message: "Fare ID is required.",
		})
	}

	if fareAttribute.Price == nil {
		messages = append(messages, types.Message{
			Field:   "price",
			Message: "Price field is required.",
		})
	} else if *fareAttribute.Price < 0 {
		messages = append(messages, types.Message{
			Field:   "price",
			Message: "Price must be a non-negative float.",
		})
	}

	if fareAttribute.CurrencyType == "" {
		messages = append(messages, types.Message{
			Field:   "currency_type",
			Message: "Currency type is required.",
		})
	}

	if fareAttribute.PaymentMethod != nil {
		if validPaymentMethod := map[int]bool{0: true, 1: true}; !validPaymentMethod[*fareAttribute.PaymentMethod] {
			messages = append(messages, types.Message{
				Field:   "payment_method",
				Message: "Payment method must be 0 or 1.",
			})
		}
	} else {
		messages = append(messages, types.Message{
			Field:   "payment_method",
			Message: "Payment method is required.",
		})
	}

	// Validate transfers field
	if fareAttribute.Transfers != nil {
		if validTransfers := map[int]bool{0: true, 1: true, 2: true}; !validTransfers[*fareAttribute.Transfers] && fareAttribute.Transfers != nil {
			messages = append(messages, types.Message{
				Field:   "transfers",
				Message: "Transfers must be 0, 1, 2, or empty (for unlimited transfers).",
			})
		}
	}

	// Validate transfer_duration
	if fareAttribute.TransferDuration != nil && *fareAttribute.TransferDuration < 0 {
		messages = append(messages, types.Message{
			Field:   "transfer_duration",
			Message: "Transfer duration must be a non-negative integer.",
		})
	}

	// Validate agency_id if multiple agencies exist
	if multipleAgencies {
		if fareAttribute.AgencyId == nil || *fareAttribute.AgencyId == "" {
			messages = append(messages, types.Message{
				Field:   "agency_id",
				Message: "Agency ID is required when the dataset contains multiple agencies.",
			})
		} else if _, exists := agencyIdMap[*fareAttribute.AgencyId]; !exists {
			messages = append(messages, types.Message{
				Field:   "agency_id",
				Message: "Agency ID must reference a valid agency_id from agency.txt.",
			})
		}
	}

	return fareAttribute, messages
}
