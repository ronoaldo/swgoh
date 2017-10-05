package main

import (
	"flag"
	"fmt"
	"log"

	"gopkg.in/yaml.v2"

	"ronoaldo.gopkg.net/swgoh/swgohgg"
)

var (
	profile        string
	starLevel      int
	charFilter     string
	optimizeStat   string
	maxStat        string
	shape          string
	showCollection bool
	showShips      bool
	showMods       bool
	showStats      bool
	showArena      bool
	useCache       bool
)

func init() {
	flag.StringVar(&profile, "profile", "", "The user `profile` on https://swgoh.gg/")

	// Operation flags
	flag.BoolVar(&showCollection, "collection", false, "Show user character collection")
	flag.BoolVar(&showShips, "ships", false, "Show user ships collection")
	flag.BoolVar(&showMods, "mods", false, "Show user mods collection")
	flag.BoolVar(&showStats, "stats", false, "Show a single character stats (requires -char)")
	flag.BoolVar(&showArena, "arena", false, "Show stats for your current arena team")

	// Cache flags
	flag.BoolVar(&useCache, "cache", true, "Use cache to save mod query")

	// Filter flags
	flag.IntVar(&starLevel, "stars", 0, "The minimal character or mod `stars` to display")
	flag.StringVar(&charFilter, "char", "", "Restrict mods used by this `character`")
	flag.StringVar(&optimizeStat, "optimize-set", "", "Build a set optimized with this `stat` looking up for all combinations")
	flag.StringVar(&maxStat, "max-set", "", "Suggest a set that has the provided `stat` best values")
	flag.StringVar(&shape, "shape", "", "Filter mods by this `shape`")
}

func fetchCollection(swgg *swgohgg.Client) (collection swgohgg.Collection, err error) {
	log.Printf("Fetching collection ...")
	collection = make(swgohgg.Collection, 0)
	err = loadCache("collection", &collection)
	if err != nil {
		log.Printf("Data not cached, loading from website (%v)", err)
		collection, err = swgg.Collection()
		if err != nil {
			log.Fatal(err)
		}
		if useCache {
			if err = saveCache("collection", &collection); err != nil {
				log.Printf("Can't save to cache: %v", err)
			}
		}
	}
	return collection, nil
}

func fetchShips(swgg *swgohgg.Client) (ships swgohgg.Ships, err error) {
	log.Printf("Fetching ships ...")
	ships = make(swgohgg.Ships, 0)
	err = loadCache("ships", &ships)
	if err != nil {
		log.Printf("Data not cached, loading from website (%v)", err)
		ships, err = swgg.Ships()
		if err != nil {
			log.Fatal(err)
		}
		if useCache {
			if err = saveCache("ships", &ships); err != nil {
				log.Printf("Can't save to cache: %v", err)
			}
		}
	}
	return ships, nil
}

var modFilterAll = swgohgg.ModFilter{}

func fetchMods(swgg *swgohgg.Client) (mods swgohgg.ModCollection, err error) {
	mods = make(swgohgg.ModCollection, 0)
	err = loadCache("mods", &mods)
	if err != nil || !useCache {
		log.Printf("Not using cache (%v)", err)
		mods, err = swgg.Mods(modFilterAll)
		if err != nil {
			log.Fatal(err)
		}
		if useCache {
			if err = saveCache("mods", &mods); err != nil {
				log.Printf("Can't save to cache: %v", err)
			}
		}
	}
	return mods, nil
}

func fetchStats(swgg *swgohgg.Client) (stats *swgohgg.CharacterStats, err error) {
	// TODO(ronoaldo) add cache support for stats
	return swgg.CharacterStats(charFilter)
}

func main() {
	flag.Parse()
	swgg := swgohgg.NewClient(profile)

	if showStats {
		stats, err := fetchStats(swgg)
		if err != nil {
			log.Fatal(err)
		}
		b, err := yaml.Marshal(stats)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Stats for %s's '%s':\n", profile, charFilter)
		fmt.Println(string(b))
	}

	if showCollection {
		collection, err := fetchCollection(swgg)
		if err != nil {
			log.Fatal(err)
		}
		for _, char := range collection {
			if char.Stars >= starLevel {
				fmt.Println(char)
			}
		}
	}

	if showShips {
		ships, err := fetchShips(swgg)
		if err != nil {
			log.Fatal(err)
		}
		for _, ship := range ships {
			if ship.Stars >= starLevel {
				fmt.Println(ship)
			}
		}
	}

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

	if showArena {
		team, _, err := swgg.Arena()
		if err != nil {
			log.Fatal(err)
		}
		for _, char := range team {
			b, err := yaml.Marshal(char)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(b))
		}
	}
}
