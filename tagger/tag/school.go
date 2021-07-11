package tag

import (
	"github.com/diivix/pathfinder-models"
)

func School(spell models.Spell) []string {
	var tags []string

	tags = append(tags, "school-"+spell.School)

	return tags
}
