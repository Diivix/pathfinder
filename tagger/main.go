package main

import (
	"encoding/json"
	"fmt"
	"hash/fnv"
	"io/ioutil"
	"log"
	"time"

	"github.com/diivix/pathfinder-models"
	"github.com/diivix/pathfinder-tagger/tag"
)

func main() {
	var startTime = time.Now()
	spells := importSpells("../data/spells.original.json")
	tagSpells(spells)
	exportSpells("../data/spells.json", spells)

	log.Printf("End tagging at %s. Start time was %s", time.Now().Format(time.StampMilli), startTime.Format(time.StampMilli))
}

func importSpells(path string) []models.Spell {
	byteValue, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}

	var spells []models.Spell

	unmarchelErr := json.Unmarshal(byteValue, &spells)
	if unmarchelErr != nil || spells == nil {
		fmt.Println(unmarchelErr)
	}

	return spells
}

func exportSpells(path string, spells []models.Spell) {
	bytes, jsonErr := json.Marshal(spells)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fileErr := ioutil.WriteFile(path, bytes, 0644)
	if fileErr != nil {
		log.Fatal(fileErr)
	}
}

func tagSpells(spells []models.Spell) {
	for i := 0; i < len(spells); i++ {
		spells[i].Id = hash(spells[i].Name)

		log.Printf("Generating tags for: %s (%d)", spells[i].Name, spells[i].Id)
		spells[i].Tags = append(spells[i].Tags, tag.ClassTypes(spells[i])...)
		spells[i].Tags = append(spells[i].Tags, tag.Components(spells[i])...)
		spells[i].Tags = append(spells[i].Tags, tag.School(spells[i])...)
		spells[i].Tags = append(spells[i].Tags, tag.Level(spells[i])...)
		spells[i].Tags = append(spells[i].Tags, tag.CastingTime(spells[i])...)
		spells[i].Tags = append(spells[i].Tags, tag.Range(spells[i])...)
		spells[i].Tags = append(spells[i].Tags, tag.Duration(spells[i])...)
		spells[i].Tags = append(spells[i].Tags, tag.Reference(spells[i])...)
		spells[i].Tags = append(spells[i].Tags, tag.Theme(spells[i])...)
	}
}

func hash(value string) uint64 {
	hash := fnv.New64()
	hash.Write([]byte(value))
	return hash.Sum64()
}
