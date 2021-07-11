package tag

import (
	"github.com/diivix/pathfinder-models"
)

func Level(spell models.Spell) []string {
	var tags []string

	switch {
	case spell.Level > 0 && spell.Level <= 3:
		tags = append(tags, "level-low_level")
	case spell.Level > 3 && spell.Level <= 6:
		tags = append(tags, "level-mid_level")
	case spell.Level > 6:
		tags = append(tags, "level-high_level")
	default:
		tags = append(tags, "level-cantrip")
	}

	return tags
}
