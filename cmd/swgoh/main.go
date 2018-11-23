package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"gopkg.in/yaml.v2"

	"github.com/ronoaldo/swgoh/swgohgg"
	"github.com/ronoaldo/swgoh/swgohhelp"
)

var (
	allyCode       string
	username       string
	password       string
	starLevel      int
	unitFilter     string
	optimizeStat   string
	maxStat        string
	shape          string
	showCharacters bool
	showShips      bool
	showMods       bool
	showStats      bool
	showArena      bool
	useCache       bool
	debug          bool
)

func init() {
	flag.StringVar(&allyCode, "a", "", "The user `ally code` in game")
	flag.StringVar(&username, "u", "", "The `username` to authenticate")
	flag.StringVar(&password, "p", "", "The `password` to authenticate")

	// Operation flags
	flag.BoolVar(&showCharacters, "characters", false, "Show user character collection")
	flag.BoolVar(&showShips, "ships", false, "Show user ships collection")
	flag.BoolVar(&showMods, "mods", false, "Show user mods collection")
	flag.BoolVar(&showStats, "stats", false, "Show a single character stats (requires -char)")
	flag.BoolVar(&showArena, "arena", false, "Show stats for your current arena team")

	// Cache flags
	flag.BoolVar(&useCache, "cache", true, "Use cache to save mod query")
	flag.BoolVar(&debug, "debug", false, "Debug request and response to temporary folder")

	// Filter flags
	flag.IntVar(&starLevel, "stars", 0, "The minimal character or mod `stars` to display")
	flag.StringVar(&unitFilter, "unit", "", "Restrict mods used by this `character`")
	flag.StringVar(&optimizeStat, "optimize-set", "", "Build a set optimized with this `stat` looking up for all combinations")
	flag.StringVar(&maxStat, "max-set", "", "Suggest a set that has the provided `stat` best values")
	flag.StringVar(&shape, "shape", "", "Filter mods by this `shape`")
}

func main() {
	flag.Parse()
	ctx := context.Background()
	swgoh := swgohhelp.New(ctx).SetDebug(debug)
	if _, err := swgoh.SignIn(username, password); err != nil {
		log.Fatalf("swgoh: error authenticating with API backend: %v", err)
	}

	players, err := swgoh.Players(allyCode)
	if err != nil {
		log.Fatalf("swgoh: error fetching player profile from API: %v", err)
	}

	player := players[0]
	unitFilter = swgohgg.CharName(unitFilter)
	if showStats {
		for _, unit := range player.Roster {
			if unit.Name == unitFilter {
				// Unit found dump stats
				b, err := yaml.Marshal(unit)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("Stats for %s's '%s':\n", player.Name, unitFilter)
				fmt.Println(string(b))
				break
			}
		}
	}

	if showCharacters {
		for _, u := range player.Roster {
			if u.Rarity >= starLevel && u.CombatType == swgohhelp.CombatTypeChar {
				fmt.Println(fmt.Sprintf("%s %d* G%d Lvl%d", u.Name, u.Rarity, u.Gear, u.Level))
			}
		}
	}

	if showShips {
		for _, u := range player.Roster {
			if u.Rarity >= starLevel && u.CombatType == swgohhelp.CombatTypeShip {
				fmt.Println(fmt.Sprintf("%s %d* G%d Lvl%d", u.Name, u.Rarity, u.Gear, u.Level))
			}
		}
	}

	/*
		if showMods {
			mods, err := fetchMods(swgg)
			if err != nil {
				log.Fatal(err)
			}
			if optimizeStat != "" {
				set := mods.Optimize(optimizeStat, false)
				for _, shape := range swgohgg.ShapeNames {
					mod := set[shape]
					fmt.Println(mod)
				}
				fmt.Println("---")
				for _, s := range set.StatSummary() {
					fmt.Println(s)
				}
			} else if maxStat != "" {
				set := mods.SetWith(maxStat)
				for _, shape := range swgohgg.ShapeNames {
					mod := set[shape]
					fmt.Println(mod)
				}
				fmt.Println("---")
				for _, s := range set.StatSummary() {
					fmt.Println(s)
				}
			} else {
				filter := swgohgg.ModFilter{
					Char: charFilter,
				}
				mods = mods.Filter(filter)
				if err != nil {
					log.Fatal(err)
				}
				if shape != "" {
					mods = mods.ByShape(shape)
				}
				for _, mod := range mods {
					fmt.Println(mod)
				}
			}
		}
	*/

	if showArena {
		fmt.Printf("%s's Arena Teams (%s)\n", player.Name, player.Titles.Selected)
		fmt.Printf("\nCharacter Arena (Ranking %d)\n\n", player.Arena.Char.Rank)
		for _, unit := range player.Arena.Char.Squad {
			suffix := ""
			if unit.Type.String() != "" {
				suffix = "(" + unit.Type.String() + ")"
			}
			if char, ok := player.Roster.FindByID(unit.UnitID); ok {
				fmt.Printf("- %d* %s G%d Lvl%d %v\n", char.Rarity, char.Name, char.Gear, char.Level, suffix)
			}
		}
		fmt.Printf("\nShip Arena (Ranking %d)\n\n", player.Arena.Ship.Rank)
		for _, unit := range player.Arena.Ship.Squad {
			if ship, ok := player.Roster.FindByID(unit.UnitID); ok {
				suffix := ""
				if unit.Type.String() != "" {
					suffix = "(" + unit.Type.String() + ")"
				}
				fmt.Printf("- %d* %s Lvl%d %v\n", ship.Rarity, ship.Name, ship.Level, suffix)
			}
		}
	}
}
