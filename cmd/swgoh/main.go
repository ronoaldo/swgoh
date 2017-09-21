package main

import (
	"flag"
	"fmt"
	"log"

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
	showMods       bool
	useCache       bool
)

func init() {
	flag.StringVar(&profile, "profile", "", "The user `profile` on https://swgoh.gg/")

	// Operation flags
	flag.BoolVar(&showCollection, "collection", false, "Show user character collection")
	flag.BoolVar(&showMods, "mods", false, "Show user mods collection")

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

func main() {
	flag.Parse()
	swgg := swgohgg.NewClient(profile)

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
}
