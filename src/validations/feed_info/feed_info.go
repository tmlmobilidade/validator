package feed_info

import (
	"main/src/lib"
	"main/src/types"
)

type parseFeedInfoValidation struct {
	*types.Validation
}

func NewParseFeedInfoValidation(severity *types.Severity) *parseFeedInfoValidation {
	s := types.SEVERITY_ERROR
	if severity != nil {
		s = *severity
	}

	return &parseFeedInfoValidation{
		Validation: &types.Validation{
			ID:          "parse_feed_info",
			Description: "Validate feed info data",
			Severity:    s,
		},
	}
}

func (v *parseFeedInfoValidation) Validate(gtfs types.Gtfs) (feedInfos []types.FeedInfo, messages []types.Message) {
	// Check if feed_info.txt is present when translations.txt exists
	if _, hasTranslations := gtfs.Files["translations"]; hasTranslations {
		if _, hasFeedInfo := gtfs.Files["feed_info"]; !hasFeedInfo {
			messages = append(messages, types.Message{
				Field:        "",
				FileName:     "feed_info.txt",
				Message:      "feed_info.txt is required when translations.txt is present",
				Severity:     v.Severity,
				ValidationID: v.ID,
			})
			return
		}
	}

	for i, feedInfo := range gtfs.Files["feed_info"] {
		info, infoMessages := parseFeedInfo(feedInfo)
		feedInfos = append(feedInfos, info)

		// Update row number and other fields for each message
		for _, msg := range infoMessages {
			msg.Row = i + 1
			msg.FileName = "feed_info.txt"
			msg.Severity = v.Severity
			msg.ValidationID = v.ID
			messages = append(messages, msg)
		}
	}

	return feedInfos, messages
}

func parseFeedInfo(m map[string]string) (info types.FeedInfo, messages []types.Message) {
	var parsingErrors []string

	// Parse required fields
	lib.ParseStringToPrimitive(m["feed_publisher_name"], &info.FeedPublisherName, &parsingErrors)
	lib.ParseStringToPrimitive(m["feed_publisher_url"], &info.FeedPublisherUrl, &parsingErrors)
	lib.ParseStringToPrimitive(m["feed_lang"], &info.FeedLang, &parsingErrors)

	// Parse optional fields
	lib.ParseStringToPrimitive(m["default_lang"], &info.DefaultLang, &parsingErrors)
	lib.ParseStringToPrimitive(m["feed_start_date"], &info.FeedStartDate, &parsingErrors)
	lib.ParseStringToPrimitive(m["feed_end_date"], &info.FeedEndDate, &parsingErrors)
	lib.ParseStringToPrimitive(m["feed_version"], &info.FeedVersion, &parsingErrors)
	lib.ParseStringToPrimitive(m["feed_contact_email"], &info.FeedContactEmail, &parsingErrors)
	lib.ParseStringToPrimitive(m["feed_contact_url"], &info.FeedContactUrl, &parsingErrors)

	if len(parsingErrors) > 0 {
		for _, err := range parsingErrors {
			messages = append(messages, types.Message{
				Field:   "N/A",
				Message: err,
			})
		}
	}

	// Validate required fields
	if info.FeedPublisherName == "" {
		messages = append(messages, types.Message{
			Field:   "feed_publisher_name",
			Message: "feed_publisher_name is required",
		})
	}

	if info.FeedPublisherUrl == "" {
		messages = append(messages, types.Message{
			Field:   "feed_publisher_url",
			Message: "feed_publisher_url is required",
		})
	} else if err := lib.ValidateUrl(info.FeedPublisherUrl); err != "" {
		messages = append(messages, types.Message{
			Field:   "feed_publisher_url",
			Message: err,
		})
	}

	if info.FeedLang == "" {
		messages = append(messages, types.Message{
			Field:   "feed_lang",
			Message: "feed_lang is required",
		})
	} else if info.FeedLang != "mul" {
		if err := lib.ValidateLanguage(info.FeedLang); err != "" {
			messages = append(messages, types.Message{
				Field:   "feed_lang",
				Message: err,
			})
		}
	}

	// Validate optional fields
	if info.DefaultLang != "" {
		if err := lib.ValidateLanguage(info.DefaultLang); err != "" {
			messages = append(messages, types.Message{
				Field:   "default_lang",
				Message: err,
			})
		}
	}

	if info.FeedStartDate != "" && !lib.IsValidServiceDate(info.FeedStartDate) {
		messages = append(messages, types.Message{
			Field:   "feed_start_date",
			Message: "feed_start_date must be in YYYYMMDD format",
		})
	}

	if info.FeedEndDate != "" {
		if !lib.IsValidServiceDate(info.FeedEndDate) {
			messages = append(messages, types.Message{
				Field:   "feed_end_date",
				Message: "feed_end_date must be in YYYYMMDD format",
			})
		} else if info.FeedStartDate != "" && info.FeedEndDate < info.FeedStartDate {
			messages = append(messages, types.Message{
				Field:   "feed_end_date",
				Message: "feed_end_date cannot be earlier than feed_start_date",
			})
		}
	}

	if info.FeedContactEmail != "" {
		if err := lib.ValidateEmail(info.FeedContactEmail); err != "" {
			messages = append(messages, types.Message{
				Field:   "feed_contact_email",
				Message: err,
			})
		}
	}

	if info.FeedContactUrl != "" {
		if err := lib.ValidateUrl(info.FeedContactUrl); err != "" {
			messages = append(messages, types.Message{
				Field:   "feed_contact_url",
				Message: err,
			})
		}
	}

	// Validate that at least one contact method is provided
	if info.FeedContactEmail == "" && info.FeedContactUrl == "" {
		messages = append(messages, types.Message{
			Field:   "",
			Message: "At least one of feed_contact_email or feed_contact_url should be provided",
		})
	}

	return info, messages
}
