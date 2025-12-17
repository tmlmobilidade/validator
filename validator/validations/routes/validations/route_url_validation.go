package routes

import (
	"main/lib"
	"main/services"
	"main/types"
	"slices"
)

/*
# Attributes

- File: [routes.txt]
- Field: route_url
- Presence: Optional
- Type: URL

# Description

URL of a web page about the particular route. Should be different from the agency.agency_url value.

[routes.txt]: https://gtfs.org/schedule/reference/#routestxt
*/
func RouteUrlValidation(route *types.Route, row int, gtfs *types.Gtfs, rules *types.RoutesRules) {
	ctx := lib.NewValidationContext("route_url", "routes.txt", "route_url_validation", row, services.AppMessageService)
	if rules != nil && rules.RouteUrl.Severity != "" {
		ctx.WithSeverity(rules.RouteUrl.Severity)
	}

	if route.RouteUrl == nil || *route.RouteUrl == "" {
		if ctx.ShouldSkip() {
			return
		}

		message := ctx.GetRequiredMessage("route_url_validation.required", "route_url_validation.recommended")
		ctx.AddMessageWithSeverity(message)
		return
	}

	if ctx.IsForbidden() {
		ctx.AddMessageWithSeverity(ctx.GetTranslatedMessage("route_url_validation.forbidden"))
		return
	}

	if valid := lib.ValidateUrl(*route.RouteUrl); !valid {
		ctx.AddError(ctx.GetTranslatedMessage("route_url_validation.invalid"))
		return
	}

	// Check if route_url is the same as agency_url
	if route.AgencyId != nil {
		agencyId := *route.AgencyId
		agencyRows, err := gtfs.GetRowsById("agency", agencyId)
		var agencyUrl string

		if err != nil {
			// Fallback to in-memory data if database is not available
			if gtfs.IdMap != nil {
				if agencyMap, exists := gtfs.IdMap["agency"]; exists {
					if indices, found := agencyMap[agencyId]; found && len(indices) > 0 {
						// Find agency in in-memory slice
						for _, agency := range gtfs.Agency {
							if agency.AgencyId == agencyId {
								agencyUrl = agency.AgencyUrl
								break
							}
						}
					}
				}
			}
		} else if len(agencyRows) > 0 {
			agencyRaw, err := gtfs.GetAgency(agencyRows[0])
			if err == nil {
				agencyUrl = agencyRaw.AgencyUrl
			}
		}

		if agencyUrl != "" && *route.RouteUrl == agencyUrl {
			ctx.AddWarning(ctx.GetTranslatedMessage("route_url_validation.same_as_agency_url"))
		}
	}

	// Validate rules
	if rules != nil && rules.RouteUrl.Options != nil {
		if slices.Contains(*rules.RouteUrl.Options, types.ALL_OPTIONS) {
			return
		}

		if !slices.Contains(*rules.RouteUrl.Options, *route.RouteUrl) {
			ctx.AddMessageWithSeverity(ctx.GetTranslatedMessage("route_url_validation.not_allowed", map[string]interface{}{"value": *route.RouteUrl}))
			return
		}
	}
}
