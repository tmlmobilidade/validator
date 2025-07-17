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

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "default_lang",
			FileName:     "feed_info.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     severity,
			ValidationID: "default_lang_validation",
		})
	}

	if feedInfo.DefaultLang == nil {
		if s == types.SEVERITY_IGNORE || s == types.SEVERITY_FORBIDDEN {
			return
		}

		warn := lib.IfThenElse(s == types.SEVERITY_ERROR, i18n.AppTranslator.Get("default_lang_validation.required"), i18n.AppTranslator.Get("default_lang_validation.recommended"))
		addMessage(warn, s)
		return
	}

	if s == types.SEVERITY_FORBIDDEN {
		addMessage(i18n.AppTranslator.Get("default_lang_validation.forbidden"), s)
		return
	}

	if feedInfo.DefaultLang != nil && *feedInfo.DefaultLang != "" {
		if valid := lib.ValidateLanguage(*feedInfo.DefaultLang); !valid {
			addMessage(i18n.AppTranslator.Get("default_lang_validation.invalid"), types.SEVERITY_ERROR)
			return
		}
	}
}
