package tag

import (
	"github.com/diivix/pathfinder-models"
)

func Range(spell models.Spell) []string {
	var tags []string

	switch spell.Range {
	case "self":
		tags = append(tags, "range-self")
	case "touch":
		tags = append(tags, "range-touch")
	default:
		tags = append(tags, "range-ranged")
	}

	return tags
}
