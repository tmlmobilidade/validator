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
- Field: feed_contact_url
- Presence: Optional
- Type: URL

# Description

URL for contact information, a web-form, support desk, or other tools for communication regarding the GTFS dataset and data publishing practices. feed_contact_url is a technical contact for GTFS-consuming applications. Provide customer service contact information through [agency.txt]. It's recommended that at least one of feed_contact_url or feed_contact_email are provided.

[feed_info.txt]: https://gtfs.org/schedule/reference/#feed_infotxt
[agency.txt]: https://gtfs.org/schedule/reference/#agencytxt
*/
func FeedContactUrlValidation(severity *types.Severity, feedInfo *types.FeedInfo, row int) {
	s := types.SEVERITY_WARNING
	if severity != nil {
		s = *severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "feed_contact_url",
			FileName:     "feed_info.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     severity,
			ValidationID: "feed_contact_url_validation",
		})
	}

	if feedInfo.FeedContactUrl == nil {
		if s == types.SEVERITY_IGNORE {
			return
		}

		warn := lib.IfThenElse(s == types.SEVERITY_ERROR, i18n.AppTranslator.Get("feed_contact_url_validation.required"), i18n.AppTranslator.Get("feed_contact_url_validation.recommended"))
		addMessage(warn, s)
		return
	}

	if feedInfo.FeedContactUrl != nil && *feedInfo.FeedContactUrl != "" {
		if valid := lib.ValidateUrl(*feedInfo.FeedContactUrl); !valid {
			addMessage(i18n.AppTranslator.Get("feed_contact_url_validation.invalid"), types.SEVERITY_ERROR)
			return
		}
	}
}
