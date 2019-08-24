package main

import (
	"log"
	"os"

	"github.com/shaileshpandey/jokebyname/restapi/operations"

	loads "github.com/go-openapi/loads"
	flags "github.com/jessevdk/go-flags"
	"github.com/shaileshpandey/jokebyname/restapi"
)

func main() {
	// inv := entities.IndividualWithJoke{CurrentJoke: entities.Joke{}}
	// inv.TellRandomJoke("http://uinames.com/api/", "http://api.icndb.com/jokes/random?firstName={0}&lastName={1}&limitTo=[nerdy]")

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewIndividualJokeTellerAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "A To Do list application"
	parser.LongDescription = "The product of a tutorial on goswagger.io"

	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
