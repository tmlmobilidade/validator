package lib

import (
	"fmt"
	"net/mail"
	"net/url"
	"regexp"
	"slices"
	"time"
)

func ValidateTimezone(timezone string) []string {
	_, err := time.LoadLocation(timezone)
	if err != nil {
		return []string{fmt.Sprintf("Invalid timezone, expected format: Europe/Lisbon, got: %s", timezone)}
	}
	return nil
}

func ValidateUrl(u string) []string {
	_, err := url.ParseRequestURI(u)
	if err != nil {
		return []string{fmt.Sprintf("Invalid URL, expected format: https://example.com, got: %s", u)}
	}
	return nil
}

func ValidateEmail(e string) []string {
	_, err := mail.ParseAddress(e)
	if err != nil {
		return []string{fmt.Sprintf("Invalid email, expected format: name@domain.com, got: %s", e)}
	}
	return nil
}

func ValidatePhone(p string) []string {
	re := regexp.MustCompile(`^(?:\+[1-9])?[0-9]{1,14}$`)
	if !re.MatchString(p) {
		return []string{fmt.Sprintf("Invalid phone number, expected format: [+1]234567890 (optional country code), got: %s", p)}
	}
	return nil
}

func ValidateLongitude(lon float64) []string {
	if lon < -180 || lon > 180 {
		return []string{fmt.Sprintf("Invalid longitude, expected range: -180 to 180, got: %f", lon)}
	}
	return nil
}

func ValidateLatitude(lat float64) []string {
	if lat < -90 || lat > 90 {
		return []string{fmt.Sprintf("Invalid latitude, expected range: -90 to 90, got: %f", lat)}
	}
	return nil
}

func ValidateLanguage(lang string) []string {
	validISO6391 := []string{"ab", "aa", "af", "ak", "sq", "am", "ar", "an", "hy", "as", "av", "ae", "ay", "az", "bm", "ba", "eu", "be", "bn", "bh", "bi", "bs", "br", "bg", "my", "ca", "ch", "ce", "ny", "zh", "cv", "kw", "co", "cr", "hr", "cs", "da", "dv", "nl", "dz", "en", "eo", "et", "ee", "fo", "fj", "fi", "fr", "ff", "gl", "ka", "de", "el", "gn", "gu", "ht", "ha", "he", "hz", "hi", "ho", "hu", "ia", "id", "ie", "ga", "ig", "ik", "io", "is", "it", "iu", "ja", "jv", "kl", "kn", "kr", "ks", "kk", "km", "ki", "rw", "ky", "kv", "kg", "ko", "ku", "kj", "la", "lb", "lg", "li", "ln", "lo", "lt", "lu", "lv", "gv", "mk", "mg", "ms", "ml", "mt", "mi", "mr", "mh", "mn", "na", "nv", "nd", "ne", "ng", "nb", "nn", "no", "ii", "nr", "oc", "oj", "cu", "om", "or", "os", "pa", "pi", "fa", "pl", "ps", "pt", "qu", "rm", "rn", "ro", "ru", "sa", "sc", "sd", "se", "sm", "sg", "sr", "gd", "sn", "si", "sk", "sl", "so", "st", "es", "su", "sw", "ss", "sv", "ta", "te", "tg", "th", "ti", "bo", "tk", "tl", "tn", "to", "tr", "ts", "tt", "tw", "ty", "ug", "uk", "ur", "uz", "ve", "vi", "vo", "wa", "cy", "wo", "fy", "xh", "yi", "yo", "za", "zu"}
	if !slices.Contains(validISO6391, lang) {
		return []string{fmt.Sprintf("'%s' is not a valid ISO 639-1 code, see https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes", lang)}
	}
	return nil
}
