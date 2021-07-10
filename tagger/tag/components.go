package tag

import (
	"fmt"
	"github.com/diivix/pathfinder-models"
)

func Components(spell models.Spell) []string {
	fmt.Println(" |_ Building tags for Components")
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
