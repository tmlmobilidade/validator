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
- Field: route_color
- Presence: Optional
- Type: Color

# Description

Route color designation that matches public facing material. Defaults to white (FFFFFF) when omitted or left empty. The color difference between route_color and route_text_color should provide sufficient contrast when viewed on a black and white screen.

[routes.txt]: https://gtfs.org/schedule/reference/#routestxt
*/
func RouteColorValidation(severity *types.Severity, route *types.Route, row int) {

	s := types.SEVERITY_IGNORE
	if severity != nil {
		s = *severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "route_color",
			FileName:     "routes.txt",
			ValidationID: "route_color_validation",
			Message:      msg,
			Rows:         []int{row},
			Severity:     severity,
		})
	}

	if route.RouteColor == nil || *route.RouteColor == "" {
		if s == types.SEVERITY_IGNORE {
			return
		}

		warn := lib.IfThenElse(s == types.SEVERITY_WARNING, "route_color is recommended", "route_color is required")
		addMessage(warn, s)
		return
	}

	color := strings.ToUpper(*route.RouteColor)
	matched, _ := regexp.MatchString(`^[0-9A-F]{6}$`, color)
	if !matched {
		addMessage("route_color must be a valid 6-character hexadecimal color (e.g., FFFFFF)", types.SEVERITY_ERROR)
		return
	}
} 