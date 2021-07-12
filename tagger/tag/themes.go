package tag

import (
	"regexp"
	"strings"

	models "github.com/diivix/pathfinder-models"
)

const (
	THEME_BATTLE        = "theme-battle"
	THEME_CONVERSATION  = "theme-conversation"
	THEME_INVESTIGATION = "theme-investigation"

	EFFECT_AID            = "effect-aid"
	EFFECT_AREA_OF_EFFECT = "effect-area_of_effect"
	EFFECT_BUFF           = "effect-buff"
	EFFECT_CONJURATION    = "effect-conjuration"
	EFFECT_DAMAGE         = "effect-damage"
	EFFECT_DEBUFF         = "effect-debuff"
	EFFECT_DEFENCE        = "effect-defence"
	EFFECT_ENVIRONMENTAL  = "effect-environmental"
	EFFECT_HEALING        = "effect-healing"
	EFFECT_HIT_POINTS     = "effect-hit_points"
	EFFECT_KNOWLEDGE      = "effect-knowledge"
	EFFECT_MOVEMENT       = "effect-movement"
	EFFECT_SPEECH         = "effect-speech"
	EFFECT_STRENGTH       = "effect-strength"
	EFFECT_TIME           = "effect-time"
)

func Theme(spell models.Spell) []string {
	var tags []string

	// Aid
	reg := regexp.MustCompile(`\bwilling creature(s)?\b`)
	if reg.MatchString(spell.Description) {
		tags = append(tags, EFFECT_AID)
	}

	if strings.Contains(spell.Description, "and add the number rolled to one ability check of its choice") {
		tags = append(tags, THEME_CONVERSATION, THEME_INVESTIGATION, EFFECT_AID, EFFECT_BUFF)
	}

	// Area of effect
	reg = regexp.MustCompile(`([0-9]{2}-foot((-| )(radius|diameter))? (cube|cone|sphere|square))|(in the area)`)
	if reg.MatchString(spell.Description) {
		tags = append(tags, EFFECT_AREA_OF_EFFECT)
	}

	return tags
}
