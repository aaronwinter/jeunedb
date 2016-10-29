package jeunedb

import "bytes"
import "fmt"
import "sync"

import Log "jeunedb/tx/log.go"
import Cache "jeunedb/cache/cache.go"

type Config struct {
	BasePath string
}

type JeuneDB struct {
	Config
	lock sync.RWMutex
	Cache
	Log
}

func Create(c Config) *JeuneDB {
	cache := Cache.Create()
	log := Log.Create()

	db := &JeuneDB{
		Config: c,
		Cache:  cache,
		Log:    log,
	}
	return db
}
