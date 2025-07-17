package stops

import (
	"main/i18n"
	"main/lib"
	"main/services"
	"main/types"
	"slices"
)

/*
# Attributes

  - File: [stops.txt]
  - Field: level_id
  - Presence: Optional
  - Type: Foreign ID referencing levels.level_id

# Description

Level of the location. The same level may be used by multiple unlinked stations.

[stops.txt]: https://gtfs.org/schedule/reference/#stopstxt
*/
func LevelIdValidation(stop *types.Stop, row int, gtfs types.Gtfs, rules *types.StopsRules) {
	s := types.SEVERITY_IGNORE
	if rules != nil && rules.LevelId.Severity != "" {
		s = rules.LevelId.Severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "level_id",
			FileName:     "stops.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     severity,
			ValidationID: "level_id_validation",
		})
	}

	if stop.LevelId == nil || *stop.LevelId == "" {
		if s == types.SEVERITY_IGNORE || s == types.SEVERITY_FORBIDDEN {
			return
		}

		message := i18n.AppTranslator.Get(
			lib.IfThenElse(s == types.SEVERITY_ERROR,
				"level_id_validation.required",
				"level_id_validation.recommended",
			),
		)
		addMessage(message, s)
		return
	}

	if s == types.SEVERITY_FORBIDDEN {
		addMessage(i18n.AppTranslator.Get("level_id_validation.forbidden"), s)
		return
	}

	// Check Foreign Key
	if !lib.GtfsIdMapKeyExists(&gtfs, "levels", *stop.LevelId) {
		addMessage(i18n.AppTranslator.Get("level_id_validation.not_found", *stop.LevelId), types.SEVERITY_ERROR)
		return
	}

	// Validate rules
	if rules != nil && rules.LevelId.Options != nil {
		if slices.Contains(*rules.LevelId.Options, types.ALL_OPTIONS) {
			return
		}

		if !slices.Contains(*rules.LevelId.Options, *stop.LevelId) {
			addMessage(i18n.AppTranslator.Get("level_id_validation.not_allowed", *stop.LevelId), s)
			return
		}
	}
}
