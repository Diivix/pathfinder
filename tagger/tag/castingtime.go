package tag

import (
	"github.com/diivix/pathfinder-models"
)

func CastingTime(spell models.Spell) []string {
	var tags []string

	switch spell.CastingTime {
	case "1 action":
		tags = append(tags, "castingtime-action")
	case "1 bonus action":
		tags = append(tags, "castingtime-bonus_action")
	case "1 minute":
		tags = append(tags, "castingtime-short_duration")
	case "10 minutes":
		tags = append(tags, "castingtime-medium_duration")
	default:
		tags = append(tags, "castingtime-long_duration")
	}

	return tags
}
