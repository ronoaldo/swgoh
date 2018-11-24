package swgohhelp

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"time"
)

// Cache related environment variables
var (
	EnvDataCacheDir = "SWGOH_CACHE_DIR"
)

// Game data related caching parameters
var (
	GameDataCacheFile = "gamedata.json"
	GameDataCacheTTL  = 5 * 24 * time.Hour
)

// GameDataCache game data cache
type GameDataCache struct {
	Titles     map[string]DataPlayerTitle
	Abilities  map[string]DataUnitAbility
	Skills     map[string]DataUnitSkill
	Categories map[string]DataUnitCategory
	Units      map[string]DataUnit
}

func (d *GameDataCache) save() error {
	cacheDir, err := CacheDirectory()
	if err != nil {
		return err
	}
	dataCacheFile := path.Join(cacheDir, GameDataCacheFile)
	log.Printf("swgohhelp: saving cache file %s", dataCacheFile)
	fd, err := os.OpenFile(dataCacheFile, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer fd.Close()
	return json.NewEncoder(fd).Encode(d)
}

func (d *GameDataCache) load() error {
	cacheDir, err := CacheDirectory()
	if err != nil {
		return err
	}
	dataCacheFile := path.Join(cacheDir, GameDataCacheFile)
	// Data is serialized as JSON, so we just open the cache.json file
	info, err := os.Stat(dataCacheFile)
	if os.IsNotExist(err) || info == nil {
		log.Printf("swgohhelp: unable to load cache file: %v", err)
		return nil
	}
	// If cache is too old, avoid using it
	if time.Since(info.ModTime()) > GameDataCacheTTL {
		log.Printf("swgohhelp: cached data is too old (%v) ignoring it.", info.ModTime())
		return nil
	}
	fd, err := os.Open(dataCacheFile)
	if err != nil {
		return err
	}
	defer fd.Close()
	return json.NewDecoder(fd).Decode(d)
}

// CacheDirectory calculates and if necessary creates the directory for
// storing cache data.
func CacheDirectory() (string, error) {
	cacheDir := os.Getenv(EnvDataCacheDir)
	if cacheDir == "" {
		cacheDir = path.Join(os.Getenv("HOME"), ".cache", "api.swgoh.help")
	}
	info, err := os.Stat(cacheDir)
	switch {
	case os.IsNotExist(err):
		if err = os.MkdirAll(cacheDir, 0755); err != nil {
			return cacheDir, err
		}
	case err != nil:
		return "", err
	default:
		if !info.IsDir() {
			return cacheDir, fmt.Errorf("swgohhelp: cache dir %s is not a directory", cacheDir)
		}
		return cacheDir, nil
	}
	return cacheDir, nil
}
