package main

import (
	"os"
	"fmt"
	"googlemaps.github.io/maps"
	"log"
	"golang.org/x/net/context"
	"github.com/kr/pretty"
	"errors"
)

func main() {

	fmt.Println( "Welcome to Atlas! " )

	params := os.Args

	if len(os.Args) < 3 {
		err := errors.New( "Insufficient number of arguments: you need to specify an address and an API key!")
		log.Fatal( err )
	}

	address := params[1]
	apiKey  := params[2]

	fmt.Printf("Fetching coordinates for '%s' using '%s' for authentication...\n", address, apiKey)

	client, err := maps.NewClient(maps.WithAPIKey(apiKey))

	if err != nil {
		log.Fatalf("Fatal error: %s", err)
	}

	request := &maps.GeocodingRequest{Address:address}

	response, err := client.Geocode( context.Background(), request )

	if err != nil {
		log.Fatalf("Fatal error: %s", err)
	}

	pretty.Println( response )
}

