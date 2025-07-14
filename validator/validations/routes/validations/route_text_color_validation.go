package routes

import (
	"main/i18n"
	"main/lib"
	"main/services"
	"main/types"
	"regexp"
	"slices"
	"strings"
)

/*
# Attributes

- File: [routes.txt]
- Field: route_text_color
- Presence: Optional
- Type: Color

# Description

Route color designation that matches public facing material. Defaults to white (FFFFFF) when omitted or left empty. The color difference between route_text_color and route_text_color should provide sufficient contrast when viewed on a black and white screen.

[routes.txt]: https://gtfs.org/schedule/reference/#routestxt
*/
func RouteTextColorValidation(route *types.Route, row int, rules *types.RoutesRules) {

	s := types.SEVERITY_IGNORE
	if rules != nil && rules.RouteTextColor.Severity != "" {
		s = rules.RouteTextColor.Severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "route_text_color",
			FileName:     "routes.txt",
			ValidationID: "route_text_color_validation",
			Message:      msg,
			Rows:         []int{row},
			Severity:     severity,
		})
	}

	if route.RouteTextColor == nil || *route.RouteTextColor == "" {
		if s == types.SEVERITY_IGNORE {
			return
		}

		warn := lib.IfThenElse(s == types.SEVERITY_WARNING, i18n.AppTranslator.Get("route_text_color_validation.recommended"), i18n.AppTranslator.Get("route_text_color_validation.required"))
		addMessage(warn, s)
		return
	}

	if s == types.SEVERITY_FORBIDDEN {
		addMessage(i18n.AppTranslator.Get("route_text_color_validation.forbidden"), s)
		return
	}

	color := strings.ToUpper(*route.RouteTextColor)
	matched, _ := regexp.MatchString(`^[0-9A-F]{6}$`, color)
	if !matched {
		addMessage(i18n.AppTranslator.Get("route_text_color_validation.invalid"), types.SEVERITY_ERROR)
		return
	}

	// Validate rules
	if rules != nil && rules.RouteTextColor.Options != nil {
		if slices.Contains(*rules.RouteTextColor.Options, types.ALL_OPTIONS) {
			return
		}

		if !slices.Contains(*rules.RouteTextColor.Options, *route.RouteTextColor) {
			addMessage(i18n.AppTranslator.Get("route_text_color_validation.not_allowed", map[string]interface{}{"value": *route.RouteTextColor}), s)
			return
		}
	}
}
