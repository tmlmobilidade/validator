package feed_info

import (
	"main/validator/lib"
	"main/validator/types"
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
			msg.Rows = []int{i + 1}
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
	var defaultLang, feedStartDate, feedEndDate, feedVersion, feedContactEmail, feedContactUrl string

	lib.ParseStringToPrimitive(m["default_lang"], &defaultLang, &parsingErrors)
	lib.ParseStringToPrimitive(m["feed_start_date"], &feedStartDate, &parsingErrors)
	lib.ParseStringToPrimitive(m["feed_end_date"], &feedEndDate, &parsingErrors)
	lib.ParseStringToPrimitive(m["feed_version"], &feedVersion, &parsingErrors)
	lib.ParseStringToPrimitive(m["feed_contact_email"], &feedContactEmail, &parsingErrors)
	lib.ParseStringToPrimitive(m["feed_contact_url"], &feedContactUrl, &parsingErrors)

	info.DefaultLang = lib.IfThenElse(defaultLang != "", &defaultLang, nil)
	info.FeedStartDate = lib.IfThenElse(feedStartDate != "", &feedStartDate, nil)
	info.FeedEndDate = lib.IfThenElse(feedEndDate != "", &feedEndDate, nil)
	info.FeedVersion = lib.IfThenElse(feedVersion != "", &feedVersion, nil)
	info.FeedContactEmail = lib.IfThenElse(feedContactEmail != "", &feedContactEmail, nil)
	info.FeedContactUrl = lib.IfThenElse(feedContactUrl != "", &feedContactUrl, nil)

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
	if info.DefaultLang != nil {
		if err := lib.ValidateLanguage(*info.DefaultLang); err != "" {
			messages = append(messages, types.Message{
				Field:   "default_lang",
				Message: err,
			})
		}
	}

	if info.FeedStartDate != nil && !lib.IsValidServiceDate(*info.FeedStartDate) {
		messages = append(messages, types.Message{
			Field:   "feed_start_date",
			Message: "feed_start_date must be in YYYYMMDD format",
		})
	}

	if info.FeedEndDate != nil && !lib.IsValidServiceDate(*info.FeedEndDate) {
		messages = append(messages, types.Message{
			Field:   "feed_end_date",
			Message: "feed_end_date must be in YYYYMMDD format",
		})
	} else if info.FeedStartDate != nil && info.FeedEndDate != nil && *info.FeedEndDate < *info.FeedStartDate {
		messages = append(messages, types.Message{
			Field:   "feed_end_date",
			Message: "feed_end_date cannot be earlier than feed_start_date",
		})
	}

	if info.FeedContactEmail != nil {
		if err := lib.ValidateEmail(*info.FeedContactEmail); err != "" {
			messages = append(messages, types.Message{
				Field:   "feed_contact_email",
				Message: err,
			})
		}
	}

	if info.FeedContactUrl != nil {
		if err := lib.ValidateUrl(*info.FeedContactUrl); err != "" {
			messages = append(messages, types.Message{
				Field:   "feed_contact_url",
				Message: err,
			})
		}
	}

	// Validate that at least one contact method is provided
	if info.FeedContactEmail == nil && info.FeedContactUrl == nil {
		messages = append(messages, types.Message{
			Field:    "",
			Message:  "It's recommended to provide at least one of feed_contact_email or feed_contact_url",
			Severity: types.SEVERITY_WARNING,
		})
	}

	return info, messages
}
