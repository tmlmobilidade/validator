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
- Field: feed_end_date
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
func FeedEndDateValidation(severity *types.Severity, feedInfo *types.FeedInfo, row int) {
	s := types.SEVERITY_WARNING
	if severity != nil {
		s = *severity
	}

	addMessage := func(msg string) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "feed_end_date",
			FileName:     "feed_info.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     types.SEVERITY_ERROR,
			ValidationID: "feed_start_date_validation",
		})
	}

	if feedInfo.FeedEndDate == nil {
		if s == types.SEVERITY_IGNORE {
			return
		}
		
		warn := lib.IfThenElse(s == types.SEVERITY_ERROR, "required", "recommended")
		addMessage(fmt.Sprintf("Feed start date is %s", warn))
		return
	}

	if feedInfo.FeedEndDate != nil && *feedInfo.FeedEndDate != "" {
		if !lib.IsValidServiceDate(*feedInfo.FeedEndDate) {
			addMessage("feed_end_date must be in YYYYMMDD format")
			return
		}
	}
} 