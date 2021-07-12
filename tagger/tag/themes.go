package tag

import (
	"log"
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
	log.Print("\tSetting theme tags")
	var tags []string

	if hasEffectAid(spell) {
		tags = append(tags, EFFECT_AID)
	}

	if hasEffectAreaOfEffect(spell) {
		tags = append(tags, EFFECT_AREA_OF_EFFECT)
	}

	if hasEffectBuff(spell) {
		tags = append(tags, EFFECT_BUFF)
	}

	return tags
}

func hasEffectAid(spell models.Spell) bool {
	reg := regexp.MustCompile(`\bwilling creature(s)?\b`)
	if reg.MatchString(spell.Description) {
		log.Printf("\t\thasEffectAid - Rule 1")
		return true
	}

	if strings.Contains(spell.Description, "add the number rolled to one ability check of its choice") {
		log.Printf("\t\thasEffectAid - Rule 2")
		return true
	}

	return false
}

func hasEffectAreaOfEffect(spell models.Spell) bool {
	reg := regexp.MustCompile(`([0-9]{2}-foot((-| )(radius|diameter))? (cube|cone|sphere|square))|(in the area)`)
	if reg.MatchString(spell.Description) {
		log.Printf("\t\thasEffectAreaOfEffect - Rule 1")
		return true
	}

	return false
}

func hasEffectBuff(spell models.Spell) bool {
	if strings.Contains(spell.Description, "add the number rolled to one ability check of its choice") {
		log.Printf("\t\thasEffectBuff - Rule 1")
		return true
	}

	return false
}
