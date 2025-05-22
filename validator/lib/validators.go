package lib

import (
	"fmt"
	"net/mail"
	"net/url"
	"regexp"
	"slices"
	"strings"
	"time"
)

func ValidateTimezone(timezone string) string {
	_, err := time.LoadLocation(timezone)
	if err != nil {
		return fmt.Sprintf("Invalid timezone, expected format: Europe/Lisbon, got: %s", timezone)
	}
	return ""
}

func ValidateUrl(u string) string {
	_, err := url.ParseRequestURI(u)
	if err != nil {
		return fmt.Sprintf("Invalid URL, expected format: https://example.com, got: %s", u)
	}
	return ""
}

func ValidateEmail(e string) string {
	_, err := mail.ParseAddress(e)
	if err != nil {
		return fmt.Sprintf("Invalid email, expected format: name@domain.com, got: %s", e)
	}
	return ""
}

func ValidatePhone(p string) string {
	input := strings.TrimSpace(p)
	if input == "" {
		return fmt.Sprintf("Invalid phone number, expected format: [+1]234567890 (optional country code), got: %s", p)
	}

	// This regex allows:
	// - Digits
	// - Letters (for dialable text like "RIDE")
	// - Punctuation: ()-+ and spaces
	// - But no other characters like @, #, $, descriptive text
	var validPhonePattern = regexp.MustCompile(`^[0-9A-Za-z\s\-\+$begin:math:text$$end:math:text$]+$`)

	// Check if input matches allowed characters
	if !validPhonePattern.MatchString(input) {
		return fmt.Sprintf("Invalid phone number, expected format: [+1]234567890 (optional country code), got: %s", p)
	}

	// (Optional) You could also add rules to ensure at least some digits exist
	digits := regexp.MustCompile(`[0-9]`)
	digitMatches := digits.FindAllString(input, -1)
	if len(digitMatches) < 7 {
		return fmt.Sprintf("Invalid phone number, must contain at least 7 digits, got: %s", p)
	}

	return ""
}

func ValidateLongitude(lon float32) string {
	if lon < -180 || lon > 180 {
		return fmt.Sprintf("Invalid longitude, expected range: -180 to 180, got: %f", lon)
	}
	return ""
}

func ValidateLatitude(lat float32) string {
	if lat < -90 || lat > 90 {
		return fmt.Sprintf("Invalid latitude, expected range: -90 to 90, got: %f", lat)
	}
	return ""
}

func ValidateLanguage(lang string) string {
	validISO6391 := []string{"ab", "aa", "af", "ak", "sq", "am", "ar", "an", "hy", "as", "av", "ae", "ay", "az", "bm", "ba", "eu", "be", "bn", "bh", "bi", "bs", "br", "bg", "my", "ca", "ch", "ce", "ny", "zh", "cv", "kw", "co", "cr", "hr", "cs", "da", "dv", "nl", "dz", "en", "eo", "et", "ee", "fo", "fj", "fi", "fr", "ff", "gl", "ka", "de", "el", "gn", "gu", "ht", "ha", "he", "hz", "hi", "ho", "hu", "ia", "id", "ie", "ga", "ig", "ik", "io", "is", "it", "iu", "ja", "jv", "kl", "kn", "kr", "ks", "kk", "km", "ki", "rw", "ky", "kv", "kg", "ko", "ku", "kj", "la", "lb", "lg", "li", "ln", "lo", "lt", "lu", "lv", "gv", "mk", "mg", "ms", "ml", "mt", "mi", "mr", "mh", "mn", "na", "nv", "nd", "ne", "ng", "nb", "nn", "no", "ii", "nr", "oc", "oj", "cu", "om", "or", "os", "pa", "pi", "fa", "pl", "ps", "pt", "qu", "rm", "rn", "ro", "ru", "sa", "sc", "sd", "se", "sm", "sg", "sr", "gd", "sn", "si", "sk", "sl", "so", "st", "es", "su", "sw", "ss", "sv", "ta", "te", "tg", "th", "ti", "bo", "tk", "tl", "tn", "to", "tr", "ts", "tt", "tw", "ty", "ug", "uk", "ur", "uz", "ve", "vi", "vo", "wa", "cy", "wo", "fy", "xh", "yi", "yo", "za", "zu"}
	if !slices.Contains(validISO6391, lang) {
		return fmt.Sprintf("'%s' is not a valid ISO 639-1 code, see https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes", lang)
	}
	return ""
}

func IsValidServiceDate(date string) bool {
	if len(date) != 8 {
		return false
	}
	
	// Convert the date to the format YYYY-MM-DD for the time.Parse function validation
	newDate := date[:4] + "-" + date[4:6] + "-" + date[6:]

	_, err := time.Parse(time.DateOnly, newDate)
	return err == nil
}
