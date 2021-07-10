package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/diivix/pathfinder-models"
	"github.com/diivix/pathfinder-tagger/tag"
)

func main() {
	importSpells("../data/spells.json")
}

func importSpells(path string) {
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

		fmt.Println(" |_ Tags are: " + strings.Join(spells[i].Tags, ", "))
	}
}

func exportSpells() {

}
