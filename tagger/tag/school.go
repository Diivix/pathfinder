package tag

import (
	"fmt"
	"github.com/diivix/pathfinder-models"
)

func School(spell models.Spell) []string {
	fmt.Println(" |_ Building tags for School")
	var tags []string

	tags = append(tags, "school-"+spell.School)

	return tags
}
