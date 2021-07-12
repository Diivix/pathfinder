package tag

import (
	"strings"

	"github.com/diivix/pathfinder-models"
)

func Reference(spell models.Spell) []string {
	var tags []string

	switch {
	case strings.Contains(spell.Reference, "players handbook"):
		tags = append(tags, "reference-players_handbook")
	case strings.Contains(spell.Reference, "instantaneous"):
		tags = append(tags, "reference-instantaneous")
	case strings.Contains(spell.Reference, "xanathar"):
		tags = append(tags, "reference-xanathars_guide")
	case strings.Contains(spell.Reference, "tasha"):
		tags = append(tags, "reference-tashas_guide")
	case strings.Contains(spell.Reference, "ee players companion"):
		tags = append(tags, "reference-ee_players_companion")
	default:
		tags = append(tags, "reference-unknown")
	}

	return tags
}
