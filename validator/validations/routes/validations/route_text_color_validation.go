package routes

import (
	"main/lib"
	"main/services"
	"main/types"
	"regexp"
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
func RouteTextColorValidation(severity *types.Severity, route *types.Route, row int) {

	s := types.SEVERITY_IGNORE
	if severity != nil {
		s = *severity
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

		warn := lib.IfThenElse(s == types.SEVERITY_WARNING, "route_text_color is recommended", "route_text_color is required")
		addMessage(warn, s)
		return
	}

	color := strings.ToUpper(*route.RouteTextColor)
	matched, _ := regexp.MatchString(`^[0-9A-F]{6}$`, color)
	if !matched {
		addMessage("route_text_color must be a valid 6-character hexadecimal color (e.g., FFFFFF)", types.SEVERITY_ERROR)
		return
	}
} 