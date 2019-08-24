package operations

import (
	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/shaileshpandey/jokebyname/entities"
	"github.com/shaileshpandey/jokebyname/models"
)

func FetchJoke(params ListJokeParams) middleware.Responder {
	inv := entities.IndividualWithJoke{CurrentJoke: entities.Joke{}}
	inv.TellRandomJoke("http://uinames.com/api/", "http://api.icndb.com/jokes/random?firstName={0}&lastName={1}&limitTo=[nerdy]")
	payload := &models.IndividualWithJoke{
		Individual: models.Individual{
			Name:    inv.Name,
			Gender:  inv.Gender,
			Region:  inv.Region,
			Surname: inv.SurName,
			Joke: &models.Joke{
				Type: inv.CurrentJoke.Type,
				Value: &models.JokeDetail{
					Category: inv.CurrentJoke.Value.Category,
					ID:       inv.CurrentJoke.Value.ID,
					Joke:     inv.CurrentJoke.Value.Text,
				},
			},
		},
	}
	return NewListJokeOK().WithPayload(payload)
}
