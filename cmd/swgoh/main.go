package main

import (
	"flag"
	"fmt"
	"log"
	"ronoaldo.gopkg.net/swgoh/data"
)

var (
	profile   string
	starLevel int
)

func init() {
	flag.StringVar(&profile, "profile", "", "The user `profile` on https://swgoh.gg/")
	flag.IntVar(&starLevel, "stars", 0, "The minimal `character stars` to filter")
}

func main() {
	flag.Parse()

	roster, err := data.Roster(profile)
	if err != nil {
		log.Fatal(err)
	}
	for _, char := range roster {
		if char.Stars >= starLevel {
			fmt.Println(char)
		}
	}
}
