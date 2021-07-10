package tag

import (
	"fmt"
	"github.com/diivix/pathfinder-models"
)

func ClassTypes(spell models.Spell) []string {
	fmt.Println(" |_ Building tags for ClassTypes")
	var tags []string

	for _, element := range spell.ClassTypes {
		tags = append(tags, "classtype-"+element)
	}

	return tags
}
