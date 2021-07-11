package models

type Spell struct {
	Id             uint64
	Name           string
	ClassTypes     []string
	Components     []string
	School         string
	Level          int
	CastingTime    string
	Range          string
	Materials      string
	Duration       string
	Description    string
	AtHigherLevels string
	Reference      string
	Tags           []string
}
