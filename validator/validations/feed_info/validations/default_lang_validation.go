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
- Field: default_lang
- Presence: Optional
- Type: Language Code

# Description

Defines the language that should be used when the data consumer doesn't know the language of the rider. It will often be en (English).

[feed_info.txt]: https://gtfs.org/schedule/reference/#feed_infotxt
*/
func DefaultLangValidation(severity *types.Severity, feedInfo *types.FeedInfo, row int) {
	s := types.SEVERITY_IGNORE
	if severity != nil {
		s = *severity
	}

	addMessage := func(msg string) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "default_lang",
			FileName:     "feed_info.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     types.SEVERITY_ERROR,
			ValidationID: "default_lang_validation",
		})
	}

	if feedInfo.DefaultLang == nil {
		if s == types.SEVERITY_IGNORE {
			return
		}

		warn := lib.IfThenElse(s == types.SEVERITY_ERROR, "required", "recommended")
		addMessage(fmt.Sprintf("Default language is %s", warn))
		return
	}

	if feedInfo.DefaultLang != nil && *feedInfo.DefaultLang != "" {
		if err := lib.ValidateLanguage(*feedInfo.DefaultLang); err != "" {
			addMessage(err)
			return
		}
	}
} 