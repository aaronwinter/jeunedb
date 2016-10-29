package jeunedb

import "bytes"
import "fmt"
import "sync"

type Config struct {
	BasePath     string
	CacheSizeMax uint64
}

type JeuneDB struct {
	Config
	lock      sync.RWMutex
	cache     map[string][]byte
	cacheSize uint64
}

func Create(c Config) *JeuneDB {
	db := &JeuneDB{
		Config:    c,
		cache:     map[string][]byte{},
		cacheSize: 0,
	}
	return db
}
