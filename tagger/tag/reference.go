package tag

import (
	"strings"

	"github.com/diivix/pathfinder-models"
)

func Reference(spell models.Spell) []string {
	var tags []string

	switch {
	case strings.Contains(spell.Reference, "players handbook"):
		tags = append(tags, "duration-players_handbook")
	case strings.Contains(spell.Reference, "instantaneous"):
		tags = append(tags, "duration-instantaneous")
	case strings.Contains(spell.Reference, "xanathar"):
		tags = append(tags, "duration-xanathars_guide")
	case strings.Contains(spell.Reference, "tasha"):
		tags = append(tags, "duration-tashas_guide")
	case strings.Contains(spell.Reference, "ee players companion"):
		tags = append(tags, "duration-ee_players_companion")
	default:
		tags = append(tags, "duration-long")
	}

	return tags
}
