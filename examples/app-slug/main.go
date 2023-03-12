package main

import (
	"log"

	"github.com/baconYao/toolkit"
)

func main() {
	toSlug := "now is the time 123"

	var tools toolkit.Tools
	slugfied, err := tools.Slugify(toSlug)
	if err != nil {
		log.Println(err)
	}
	log.Println(slugfied)
}
