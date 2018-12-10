package cache

import (
	"bytes"
	"encoding/gob"
	"log"
	"sync"
	"time"

	bolt "go.etcd.io/bbolt"
)

var defaultCacheExpiration = 60 * time.Second

// Cache is an interface for caching values.
type Cache interface {
	// Get retrieves the value from cache. Returns the value if found and true,
	// or a nil array and false if not cached.
	Get(key string) (value []byte, ok bool)

	// GetString is like Get, but returns a string value.
	GetString(key string) (value string, ok bool)

	// Put attempts to save the value in cache.
	Put(key string, value []byte)

	// PutString is like Put, but stores a string value.
	PutString(key, value string)
}

// Item represents a cached item value
type Item struct {
	Value     []byte
	Timestamp time.Time
}

func (i *Item) encode() ([]byte, error) {
	var buff bytes.Buffer
	err := gob.NewEncoder(&buff).Encode(i)
	return buff.Bytes(), err
}

func (i *Item) decode(b []byte) error {
	buff := bytes.NewBuffer(b)
	err := gob.NewDecoder(buff).Decode(i)
	return err
}

var cacheFiles = make(map[string]Cache)
var cacheFielsMu sync.Mutex

// NewCache always returns a valid cache implementation.
// First, it tries to return a file cache using bolt database.
// If there is an error loading the file or creating it, it returns
// a no-op cache implementation.
func NewCache(filename string, expiration time.Duration) Cache {
	if c, ok := cacheFiles[filename]; ok {
		return c
	}

	cacheFielsMu.Lock()
	defer cacheFielsMu.Unlock()

	db, err := bolt.Open(filename, 0600, &bolt.Options{
		Timeout: 1 * time.Second,
	})
	if err != nil {
		log.Printf("cache: error opening database: %v", err)
		return &noopCache{}
	}
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(defaultBucket)
		return err
	})
	if err != nil {
		log.Printf("cache: error initializing database: %v", err)
		return &noopCache{}
	}
	cache := &fileCache{
		db:         db,
		expiration: expiration,
	}
	cacheFiles[filename] = cache
	return cache
}

// fileCache implements a cache using the embedded bolt database
type fileCache struct {
	db         *bolt.DB
	expiration time.Duration
}

var defaultBucket = []byte("swgoh")

func (f *fileCache) Get(key string) (value []byte, ok bool) {
	f.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(defaultBucket)
		v := b.Get([]byte(key))
		if v != nil {
			i := &Item{}
			if err := i.decode(v); err != nil {
				return err
			}
			// If item is expired, ignore it
			if time.Since(i.Timestamp) < f.expiration {
				value = append(value, i.Value...)
			}
		}
		return nil
	})
	return value, value != nil
}

func (f *fileCache) GetString(key string) (value string, ok bool) {
	v, ok := f.Get(key)
	if ok {
		return string(v), true
	}
	return "", false
}

func (f *fileCache) Put(key string, value []byte) {
	err := f.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(defaultBucket)
		i := &Item{Value: value, Timestamp: time.Now()}
		v, err := i.encode()
		if err != nil {
			return err
		}
		return b.Put([]byte(key), v)
	})
	if err != nil {
		log.Printf("cache: error saving value: %v", err)
	}
}

func (f *fileCache) PutString(key, value string) {
	f.Put(key, []byte(value))
}

// noopCache implements a void cache that does not fail to execute but also
// keeps no values in memory or disk.
type noopCache struct{}

func (n *noopCache) Get(key string) (value []byte, ok bool)       { return nil, false }
func (n *noopCache) GetString(key string) (value string, ok bool) { return "", false }
func (n *noopCache) Put(key string, value []byte)                 {}
func (n *noopCache) PutString(key, value string)                  {}
