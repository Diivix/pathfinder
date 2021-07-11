package tag

import (
	"github.com/diivix/pathfinder-models"
)

func Components(spell models.Spell) []string {
	var tags []string

	for _, element := range spell.Components {
		switch element {
		case "v":
			tags = append(tags, "component-verbal")
		case "s":
			tags = append(tags, "component-somatic")
		case "m":
			tags = append(tags, "component-material")
		}
	}

	return tags
}
