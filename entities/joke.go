package entities

import (
	"github.com/shaileshpandey/jokebyname/helper"
)

// Joke composes it's type and detail about the joke
type Joke struct {
	Type  string     `json:"type"`
	Value JokeDetail `json:"value"`
}

// JokeDetail will hold joke id, text and summary
type JokeDetail struct {
	ID       int64    `json:"id"`
	Text     string   `json:"joke"`
	Category []string `json:"categories"`
}

// Self will populate itself from JSON
func (joke *Joke) self(url string) error {
	return helper.JSONToStruct(url, joke)
}
