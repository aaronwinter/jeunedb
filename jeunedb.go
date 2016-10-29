package jeunedb

import "bytes"
import "fmt"
import "sync"

import Log "jeunedb/tx"
import Cache "jeunedb/cache"

type Config struct {
	BasePath string
}

type JeuneDB struct {
	lock sync.RWMutex
	Config
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
