package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"

	models "github.com/diivix/pathfinder-models"
	"github.com/diivix/pathfinder-tagger/tag"
)

func main() {
	importSpells("../data/spells.json")
}

func importSpells(path string) {
	var startTime = time.Now()

	byteValue, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}

	var spells []models.Spell

	unmarchelErr := json.Unmarshal(byteValue, &spells)
	if unmarchelErr != nil || spells == nil {
		fmt.Println(unmarchelErr)
	}

	for i := 0; i < len(spells); i++ {
		log.Println("Generating tags for: " + spells[i].Name)

		spells[i].Tags = append(spells[i].Tags, tag.ClassTypes(spells[i])...)
		spells[i].Tags = append(spells[i].Tags, tag.Components(spells[i])...)
		spells[i].Tags = append(spells[i].Tags, tag.School(spells[i])...)
		spells[i].Tags = append(spells[i].Tags, tag.Level(spells[i])...)
		spells[i].Tags = append(spells[i].Tags, tag.CastingTime(spells[i])...)
		spells[i].Tags = append(spells[i].Tags, tag.Range(spells[i])...)
		spells[i].Tags = append(spells[i].Tags, tag.Duration(spells[i])...)
		spells[i].Tags = append(spells[i].Tags, tag.Reference(spells[i])...)

		fmt.Println(" |_ Tags are: " + strings.Join(spells[i].Tags, ", "))
	}

	log.Println("End tagging. Start time was " + startTime.Local().String())
}

func exportSpells() {

}
