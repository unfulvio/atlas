# Atlas

This is an experiment and programming exercise in [Go](https://golang.org) to geocode street addresses with Google Maps.

### Getting started:

1. You need go installed in your local system. Follow these steps to [install Go](https://golang.org/doc/install) on your system.
2. Make sure you have added Go to your `PATH` (see instructions above) so you can run Go commands from your command line. 
3. Once you have set the `$GOPATH` use `go install github.com/unfulvio/atlas` to fetch this program without cloning elsewhere via Git.
4. Now you can type `atlas` from anywhere in your command line to run this program.

### Usage:

The program _needs_ two arguments. 

```shell
$ atlas "<the address to geocode>" <your API key here>
```

#### For example*:

```shell
$ atlas "221 Baker Street, London, United Kingdom" someGibberish123ApiKeyCodeBlah
```

\* spoiler: this address actually works :)

The program will return the geocode response in JSON format.
