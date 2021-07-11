package tag

import (
	"fmt"
	"strings"

	"github.com/diivix/pathfinder-models"
)

func Duration(spell models.Spell) []string {
	fmt.Println(" |_ Building tags for Duration")
	var tags []string

	switch {
	case strings.Contains(spell.Duration, "concentration"):
		tags = append(tags, "duration-concentration")
	case strings.Contains(spell.Duration, "instantaneous"):
		tags = append(tags, "duration-instantaneous")
	case strings.Contains(spell.Duration, "until dispelled"):
		tags = append(tags, "duration-until_dispelled")
	case strings.Contains(spell.Duration, "minute") || strings.Contains(spell.Duration, "minutes"):
		tags = append(tags, "duration-short")
	case strings.Contains(spell.Duration, "1 hour"):
		tags = append(tags, "duration-medium")
	default:
		tags = append(tags, "duration-long")
	}

	return tags
}
