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
- Field: feed_start_date
- Presence: Recommended
- Type: Date

# Description

The dataset provides complete and reliable schedule information for service in the period from the beginning of the feed_start_date day to the end of the feed_end_date day.

Both days may be left empty if unavailable.

The feed_end_date date must not precede the feed_start_date date if both are given.

It is recommended that dataset providers give schedule data outside this period to advise of likely future service, but dataset consumers should treat it mindful of its non-authoritative status.

If feed_start_date or feed_end_date extend beyond the active calendar dates defined in calendar.txt and calendar_dates.txt, the dataset is making an explicit assertion that there is no service for dates within the feed_start_date or feed_end_date range but not included in the active calendar dates.

[feed_info.txt]: https://gtfs.org/schedule/reference/#feed_infotxt
*/
func FeedStartDateValidation(severity *types.Severity, feedInfo *types.FeedInfo, row int) {
	s := types.SEVERITY_WARNING
	if severity != nil {
		s = *severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "feed_start_date",
			FileName:     "feed_info.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     severity,
			ValidationID: "feed_start_date_validation",
		})
	}

	if feedInfo.FeedStartDate == nil {
		if s == types.SEVERITY_IGNORE || s == types.SEVERITY_FORBIDDEN {
			return
		}

		warn := lib.IfThenElse(s == types.SEVERITY_ERROR, i18n.AppTranslator.Get("feed_start_date_validation.required"), i18n.AppTranslator.Get("feed_start_date_validation.recommended"))
		addMessage(warn, s)
		return
	}

	if feedInfo.FeedStartDate != nil && *feedInfo.FeedStartDate != "" {
		if !lib.IsValidServiceDate(*feedInfo.FeedStartDate) {
			addMessage(i18n.AppTranslator.Get("feed_start_date_validation.invalid"), types.SEVERITY_ERROR)
			return
		}
	}
}
