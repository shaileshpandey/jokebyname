package entities

import (
	"log"
	"strings"

	netUrl "net/url"

	"github.com/shaileshpandey/jokebyname/helper"
)

// Individual struct will keep information about individual person
type Individual struct {
	Name    string `json:"name"`
	SurName string `json:"surname"`
	Gender  string `json:"gender"`
	Region  string `json:"region"`
}

// Self will populate itself from JSON
func (person *Individual) self(url string) error {
	if strings.TrimSpace(url) == "" {
		log.Panic("Unable to fetch random person")
	}

	_, err := netUrl.Parse(url)
	if err != nil {
		log.Panic("inavlid Url passed to fetch individual")
	}
	return helper.JSONToStruct(url, person)
}
