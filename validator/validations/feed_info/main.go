package feed_info

import (
	"main/lib"
	"main/types"
	validations "main/validations/feed_info/validations"
)

func RunValidations(gtfs types.Gtfs) {
	lib.AppLogger.Debug("Running FeedInfo Validations...")

	for i, feedInfo := range gtfs.Files["feed_info"] {
		feedInfo := validations.ParseFeedInfo(feedInfo, i, &gtfs)

		if feedInfo == (types.FeedInfo{}) {
			continue
		}

		// Validate feed_lang
		validations.FeedLangValidation(&feedInfo, i)

		// Validate feed_publisher_name
		validations.FeedPublisherNameValidation(&feedInfo, i)

		// Validate feed_publisher_url
		validations.FeedPublisherUrlValidation(&feedInfo, i)
		
		// Validate feed_contact_email
		validations.FeedContactEmailValidation(nil, &feedInfo, i)

		// Validate feed_contact_url
		validations.FeedContactUrlValidation(nil, &feedInfo, i)

		// Validate feed_end_date
		validations.FeedEndDateValidation(nil, &feedInfo, i)

		// Validate feed_start_date
		validations.FeedStartDateValidation(nil, &feedInfo, i)

		// Validate feed_version
		validations.FeedVersionValidation(nil, &feedInfo, i)

		// Validate default_lang
		validations.DefaultLangValidation(nil, &feedInfo, i)
	}
}
