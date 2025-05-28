package feed_info

import (
	"fmt"
	"main/lib"
	"main/services"
	"main/types"
)

/*
# Attributes

- File: [feed_info.txt]
- Field: feed_contact_email
- Presence: Optional
- Type: Email

# Description

Email address for communication regarding the GTFS dataset and data publishing practices. feed_contact_email is a technical contact for GTFS-consuming applications. Provide customer service contact information through [agency.txt]. It's recommended that at least one of feed_contact_email or feed_contact_url are provided.

[feed_info.txt]: https://gtfs.org/schedule/reference/#feed_infotxt
[agency.txt]: https://gtfs.org/schedule/reference/#agencytxt
*/
func FeedContactEmailValidation(severity *types.Severity, feedInfo *types.FeedInfo, row int) {
	s := types.SEVERITY_WARNING
	if severity != nil {
		s = *severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "feed_contact_email",
			FileName:     "feed_info.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     severity,
			ValidationID: "feed_contact_email_validation",
		})
	}

	if feedInfo.FeedContactEmail == nil {
		if s == types.SEVERITY_IGNORE {
			return
		}
		
		warn := lib.IfThenElse(s == types.SEVERITY_ERROR, "required", "recommended")
		addMessage(fmt.Sprintf("Feed contact email is %s", warn), s)
		return
	}

	if feedInfo.FeedContactEmail != nil && *feedInfo.FeedContactEmail != "" {
		if err := lib.ValidateEmail(*feedInfo.FeedContactEmail); err != "" {
			addMessage(err, types.SEVERITY_ERROR)
			return
		}
	}
} 