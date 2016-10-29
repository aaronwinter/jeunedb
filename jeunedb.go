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

func New(c Config) *JeuneDB {
	cache := Cache.New()
	log := Log.New()

	db := &JeuneDB{
		Config: c,
		Cache:  cache,
		Log:    log,
	}
	return db
}
