package jeunedb

import "bytes"
import "fmt"
import "sync"

import Log "jeunedb/log"
import Cache "jeunedb/cache"
import Table "jeunedb/table"

type Config struct {
	BasePath string
}

type JeuneDB struct {
	lock sync.RWMutex
	Config
	Cache.Cache
	Log.Log
}

func New(c Config) *JeuneDB {
	cache := Cache.New()
	log := Log.New(c.BasePath, "log_")

	db := &JeuneDB{
		Config: c,
		Cache:  cache,
		Log:    log,
	}
	return db
}
