# Atlas

This is an experiment and programming exercise in [Go](https://golang.org) to geocode street addresses with Google Maps.

### Getting started:

1. You need [Go](https://golang.org/). Seriously you do, it's awesome. Follow these steps to [install Go](https://golang.org/doc/install) on your system.
2. Make sure you have added Go to your `PATH` (see instructions above) so you can run Go commands from your command line. 
3. Once you have set the `$GOPATH` use `go install github.com/unfulvio/atlas` to fetch this program without cloning elsewhere via Git.
4. Now you can type `atlas` from anywhere in your command line to run this program.
5. Before you do so you need to configure Atlas.

### Configuration:

Place a file `config.toml` in your `$HOME` path `/~`. It must contain the following information:

```toml
[google]
api = "<enter your Google Maps API key here>"
```

Follow [these instructions](https://developers.google.com/maps/documentation/geocoding/get-api-key) to set up a Google Console project and obtain a Google Maps API key.  

### Usage:

Ensure you have Go installed and then have installed also this package:

```shell
$ go install github.com/unfulvio/atlas
```

Then, type:
 
```shell
$ atlas
```

This will start a webserver answering to port 8080. From console you should be able to see an output while the process is running. You can query Atlas by going to the following address from your web browser:

```shell
http://localhost:8080/geocode?address=London
```

You can replace `London` with any address, Atlas will respond with geocode data such as the following:

```json
[{"address_components":[{"long_name":"London","short_name":"London","types":["locality","political"]},{"long_name":"London","short_name":"London","types":["postal_town"]},{"long_name":"Waltham Abbey","short_name":"Waltham Abbey","types":["administrative_area_level_4","political"]},{"long_name":"Greater London","short_name":"Greater London","types":["administrative_area_level_2","political"]},{"long_name":"England","short_name":"England","types":["administrative_area_level_1","political"]},{"long_name":"United Kingdom","short_name":"GB","types":["country","political"]}],"formatted_address":"London, UK","geometry":{"location":{"lat":51.5073509,"lng":-0.1277583},"location_type":"APPROXIMATE","viewport":{"northeast":{"lat":51.6723432,"lng":0.1482711},"southwest":{"lat":51.38494009999999,"lng":-0.3514683}},"types":null},"types":["locality","political"],"place_id":"ChIJdd4hrwug2EcRmSrV3Vo6llI"}]
```
