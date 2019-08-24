package entities

import (
	"fmt"
	"log"
	"strings"
)

// IndividualWithJoke will compose detail of individual and random joke
type IndividualWithJoke struct {
	Individual
	CurrentJoke Joke
}

// TellRandomJoke will be use to print random joke
func (personWithJoke *IndividualWithJoke) TellRandomJoke(personURL, jokeURL string) string {
	err := personWithJoke.self(personURL)
	if err != nil {
		log.Fatal("Unable to tell joke since joke teller is not available.")
	}
	jokeURL = strings.Replace(jokeURL, "{0}", personWithJoke.Name, -1)
	jokeURL = strings.Replace(jokeURL, "{1}", personWithJoke.SurName, -1)
	err = personWithJoke.CurrentJoke.self(jokeURL)
	if err != nil {
		log.Fatal("Unable to tell joke since joke content is not available.")
	}

	result := fmt.Sprintf("%s -> %s", personWithJoke.Name, personWithJoke.CurrentJoke.Value.Text)
	return result
}
