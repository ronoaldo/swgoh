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
	charFilter string
	showRoster bool
	showMods   bool
)

func init() {
	flag.StringVar(&profile, "profile", "", "The user `profile` on https://swgoh.gg/")

	// Operation flags
	flag.BoolVar(&showRoster, "roster", false, "Show user character collection")
	flag.BoolVar(&showMods, "mods", false, "Show user mods collection")

	// Filter flags
	flag.IntVar(&starLevel, "stars", 0, "The minimal character or mod `stars` to display")
	flag.StringVar(&charFilter, "char", "", "Restrict mods used by this `character`")
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
		filter := swgohgg.ModFilter{
			Char: charFilter,
		}
		mods, err := swgg.Mods(filter)
		if err != nil {
			log.Fatal(err)
		}
		for _, mod := range mods {
			fmt.Println(mod)
		}
	}
}
