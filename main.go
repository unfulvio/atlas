package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pelletier/go-toml"
	"golang.org/x/net/context"
	"googlemaps.github.io/maps"
	"log"
	"net/http"
	"os"
	"io/ioutil"
)

type configFile struct {
	google googleConfig
}

type googleConfig struct {
	key string
}

func main() {

	configFileName    := os.Getenv("HOME") + "/atlas.toml"
	fileContents, err := ioutil.ReadFile(configFileName)

	if err != nil {
		log.Fatalf("Fatal error: %s", err)
	}

	config, err := toml.Load(string(fileContents))

	if err != nil {
		log.Fatalf("Fatal error, could not load %s configuration file: %s\n", configFileName, err.Error())
	}

	key := config.Get("google.key").(string)

	client, err := maps.NewClient(maps.WithAPIKey(key))

	if err != nil {
		log.Fatalf("Fatal error, could not connect to Google Maps API: %s", err)
	}

	router := gin.Default()

	// The request responds to a url matching: /geocode?address=London
	router.GET("/geocode", func(c *gin.Context) {

		address := c.Query("address")
		request := &maps.GeocodingRequest{Address:address}

		response, err := client.Geocode( context.Background(), request )

		if err != nil {
			log.Fatalf("Fatal error: %s", err)
		}

		c.JSON(http.StatusOK, response )
	})

	router.Run(":8080")
}
