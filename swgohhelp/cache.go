package swgohhelp

import (
	"fmt"
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
	GameDataCacheFile       = "gamedata.db"
	GameDataCacheExpiration = 7 * 24 * time.Hour

	PlayerCacheFile       = "players.db"
	PlayerCacheExpiration = 24 * time.Hour

	GuildCacheFile       = "guilds.db"
	GuildCacheExpiration = 20 * time.Hour
)

// CacheDirectory calculates and if necessary creates the directory for
// storing cache data.
func cacheDirectory() (string, error) {
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
