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
- Field: feed_version
- Presence: Recommended
- Type: String

# Description

String that indicates the current version of their GTFS dataset. GTFS-consuming applications can display this value to help dataset publishers determine whether the latest dataset has been incorporated.

[feed_info.txt]: https://gtfs.org/schedule/reference/#feed_infotxt
*/
func FeedVersionValidation(severity *types.Severity, feedInfo *types.FeedInfo, row int) {
	s := types.SEVERITY_WARNING
	if severity != nil {
		s = *severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "feed_version",
			FileName:     "feed_info.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     severity,
			ValidationID: "feed_version_validation",
		})
	}

	if feedInfo.FeedVersion == nil || *feedInfo.FeedVersion == "" {
		if s == types.SEVERITY_IGNORE {
			return
		}
		
		warn := lib.IfThenElse(s == types.SEVERITY_ERROR, "required", "recommended")
		addMessage(fmt.Sprintf("Feed version is %s", warn), s)
		return
	}
} 