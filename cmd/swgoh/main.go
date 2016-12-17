package main

import (
	"flag"
	"fmt"
	"log"
	"ronoaldo.gopkg.net/swgoh/swgohgg"
)

var (
	profile    string
	starLevel  int
	showRoster bool
	showMods   bool
)

func init() {
	flag.BoolVar(&showRoster, "roster", false, "Show user character collection")
	flag.BoolVar(&showMods, "mods", false, "Show user mods collection")
	flag.StringVar(&profile, "profile", "", "The user `profile` on https://swgoh.gg/")
	flag.IntVar(&starLevel, "stars", 0, "The minimal `character stars` to filter")
}

func main() {
	flag.Parse()
	swgg := swgohgg.NewClient(profile)
	if showRoster {
		roster, err := swgg.Roster()
		if err != nil {
			log.Fatal(err)
		}
		for _, char := range roster {
			if char.Stars >= starLevel {
				fmt.Println(char)
			}
		}
	}

	if showMods {
		mods, err := swgg.Mods()
		if err != nil {
			log.Fatal(err)
		}
		for _, mod := range mods {
			fmt.Println(mod)
		}
	}
}
