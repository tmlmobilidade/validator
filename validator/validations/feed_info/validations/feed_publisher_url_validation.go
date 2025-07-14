package feed_info

import (
	"main/i18n"
	"main/lib"
	"main/services"
	"main/types"
)

/*
# Attributes

- File: [feed_info.txt]
- Field: feed_publisher_url
- Presence: Required
- Type: URL

# Description

URL of the dataset publishing organization's website. This may be the same as one of the agency.agency_url values.

[feed_info.txt]: https://gtfs.org/schedule/reference/#feed_infotxt
*/
func FeedPublisherUrlValidation(feedInfo *types.FeedInfo, row int) {
	addMessage := func(msg string) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "feed_publisher_url",
			FileName:     "feed_info.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     types.SEVERITY_ERROR,
			ValidationID: "feed_publisher_url_validation",
		})
	}

	if feedInfo.FeedPublisherUrl == nil || *feedInfo.FeedPublisherUrl == "" {
		addMessage(i18n.AppTranslator.Get("feed_publisher_url_validation.required"))
		return
	}

	if valid := lib.ValidateUrl(*feedInfo.FeedPublisherUrl); !valid {
		addMessage(i18n.AppTranslator.Get("feed_publisher_url_validation.invalid"))
		return
	}
}
