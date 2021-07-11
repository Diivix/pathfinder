package tag

import (
	"github.com/diivix/pathfinder-models"
)

func ClassTypes(spell models.Spell) []string {
	var tags []string

	for _, element := range spell.ClassTypes {
		tags = append(tags, "classtype-"+element)
	}

	return tags
}
