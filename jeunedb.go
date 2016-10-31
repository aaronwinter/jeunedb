package jeunedb

//import "io"
//import Error "jeunedb/errors"
//import Table "jeunedb/table"
//import "fmt"
import "sync"
import "errors"
import "os"
import "bytes"
import "io/ioutil"
import "encoding/binary"

import Log "jeunedb/log"
import Cache "jeunedb/cache"
import Tx "jeunedb/tx"

var invalidCommand error = errors.New("Invalid command")

type Config struct {
	BasePath string
	permDir  os.FileMode
	permFile os.FileMode
}

type JeuneDB struct {
	mutex sync.RWMutex
	Config
	Cache.Cache
	Tx.Queue
	Log.Log
}

func New(c Config) *JeuneDB {
	if c.permDir <= 0 || c.permDir > 777 {
		c.permDir = 744
	}

	if c.permFile <= 0 || c.permFile > 777 {
		c.permFile = 777
	}

	if c.BasePath == "" {
		c.BasePath = "./tmpdb"
	}

	cache := Cache.New()
	log := Log.New(c.BasePath, "log_")

	db := &JeuneDB{
		Config: c,
		Cache:  cache,
		Log:    log,
	}
	return db
}

func WriteBatch() *Tx.Transaction {
	tx := Tx.New()
	return &tx
}
