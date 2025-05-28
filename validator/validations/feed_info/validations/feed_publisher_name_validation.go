package feed_info

import (
	"main/services"
	"main/types"
)

/*
# Attributes

- File: [feed_info.txt]
- Field: feed_publisher_name
- Presence: Required
- Type: String

# Description

Full name of the organization that publishes the dataset. This may be the same as one of the agency.agency_name values.

[feed_info.txt]: https://gtfs.org/schedule/reference/#feed_infotxt
*/
func FeedPublisherNameValidation(feedInfo *types.FeedInfo, row int) {
	addMessage := func(msg string) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "feed_publisher_name",
			FileName:     "feed_info.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     types.SEVERITY_ERROR,
			ValidationID: "feed_publisher_name_validation",
		})
	}

	if feedInfo.FeedPublisherName == nil || *feedInfo.FeedPublisherName == "" {
		addMessage("feed_publisher_name is required")
		return
	}
} 