package levels

import (
	"main/lib"
	"main/services"
	"main/types"
)

/*
# Attributes

- File: [levels.txt]
- Field: level_id
- Presence: Required
- Type: Unique ID

# Description

Id of the level that can be referenced from stops.txt.

[levels.txt]: https://gtfs.org/schedule/reference/#levelstxt
*/
func LevelIdValidation(level *types.Levels, row int, gtfs types.Gtfs, rules *types.LevelsRules) {
	ctx := lib.NewValidationContext("level_id", "levels.txt", "level_id_validation", row, services.AppMessageService)
	if rules != nil && rules.LevelId.Severity != "" {
		ctx.WithSeverity(rules.LevelId.Severity)
	}

	if level.LevelId == nil || *level.LevelId == "" {
		if ctx.ShouldSkip() {
			lib.AppLogger.Accent("level_id_validation.required: should skip")
			return
		}
		lib.AppLogger.Accent("level_id_validation.required: should not skip")
		message := ctx.GetTranslatedMessage("level_id_validation.required")
		ctx.AddError(message)
		return
	}

	// Check Foreign Key
	if !lib.GtfsIdMapKeyExists(&gtfs, "levels", *level.LevelId) {
		lib.AppLogger.Accent("level_id_validation.not_found: should not skip")
		ctx.AddError(ctx.GetTranslatedMessage("level_id_validation.not_found", *level.LevelId))
		return
	}

	lib.AppLogger.Accent("level_id_validation.should_skip: should skip")
	return
}
