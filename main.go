package main

import (
	"flag"
	"songsort/cmd/getartist"
)

var artist = flag.String("a", " ", "Get artist info")

func main() {
	flag.Parse()

	kanye := getartist.Artist{Name: *artist}

	kanye.Printer()
}
