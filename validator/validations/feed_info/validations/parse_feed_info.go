package feed_info

import (
	"main/lib"
	"main/services"
	"main/types"
)

func ParseFeedInfo(rawFeedInfo map[string]string, row int, gtfs *types.Gtfs) types.FeedInfo {
	var (
		feedInfo types.FeedInfo = types.FeedInfo{}
		feedLang, feedPublisherName, feedPublisherUrl string
		defaultLang, feedContactEmail, feedContactUrl, feedEndDate, feedStartDate, feedVersion string
		messages []types.Message
	)

	stringFields := map[string]*string{
		"feed_lang": &feedLang,
		"feed_publisher_name": &feedPublisherName,
		"feed_publisher_url": &feedPublisherUrl,
		"default_lang": &defaultLang,
		"feed_contact_email": &feedContactEmail,
		"feed_contact_url": &feedContactUrl,
		"feed_end_date": &feedEndDate,
		"feed_start_date": &feedStartDate,
		"feed_version": &feedVersion,
	}

	addMessage := func(field, msg string) {
		messages = append(messages, types.Message{
			Field:        field,
			FileName:     "feed_info.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     types.SEVERITY_ERROR,
			ValidationID: "feed_info_parse",
		})
	}

	// Parse string fields
	for field, target := range stringFields {
		if errMsg := lib.ParseStringToPrimitive(rawFeedInfo[field], target); errMsg != "" {
			addMessage(field, errMsg)
		}
	}

	if len(messages) > 0 {
		services.AppMessageService.AddMessages(messages)
		return feedInfo
	}

	// Assign required fields
	feedInfo.FeedLang = lib.IfThenElse(feedLang != "", &feedLang, nil)
	feedInfo.FeedPublisherName = lib.IfThenElse(feedPublisherName != "", &feedPublisherName, nil)
	feedInfo.FeedPublisherUrl = lib.IfThenElse(feedPublisherUrl != "", &feedPublisherUrl, nil)

	// Assign optional fields
	feedInfo.DefaultLang = lib.IfThenElse(rawFeedInfo["default_lang"] != "", &defaultLang, nil)
	feedInfo.FeedContactEmail = lib.IfThenElse(rawFeedInfo["feed_contact_email"] != "", &feedContactEmail, nil)
	feedInfo.FeedContactUrl = lib.IfThenElse(rawFeedInfo["feed_contact_url"] != "", &feedContactUrl, nil)
	feedInfo.FeedEndDate = lib.IfThenElse(rawFeedInfo["feed_end_date"] != "", &feedEndDate, nil)
	feedInfo.FeedStartDate = lib.IfThenElse(rawFeedInfo["feed_start_date"] != "", &feedStartDate, nil)
	feedInfo.FeedVersion = lib.IfThenElse(rawFeedInfo["feed_version"] != "", &feedVersion, nil)

	return feedInfo
}
