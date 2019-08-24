package entities

import "github.com/shaileshpandey/jokebyname/helper"

// Individual struct will keep information about individual person
type Individual struct {
	Name    string `json:"name"`
	SurName string `json:"surname"`
	Gender  string `json:"gender"`
	Region  string `json:"region"`
}

// Self will populate itself from JSON
func (person *Individual) self(url string) error {
	return helper.JSONToStruct(url, person)
}