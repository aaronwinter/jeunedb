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

func (db *JeuneDB) Put(key []byte, value []byte) error {
	tx := Tx.New()
	tx.Put(key, value)
	_, err := db.Commit(tx)
	return err
}

func (db *JeuneDB) Get(key []byte) ([]byte, error) {
	tx := Tx.New()
	tx.Get(key)
	res, err := db.Commit(tx)
	return res, err
}


func (db *JeuneDB) _Put(key []byte, value []byte) ([]byte, error) {
}

func (db *JeuneDB) _Get(key []byte) ([]byte, error) {
	return make([]byte, 0), nil
}

func (db *JeuneDB) _Snapshot() ([]byte, error) {
	return make([]byte, 0), nil
}

func (db *JeuneDB) _Exec(o Tx.Operation) ([]byte, error) {
	switch o.Cmd {
	case "GET":
		return db._Get(o.Key)
	case "PUT":
		return db._Put(o.Key, o.Value)
	case "SNAPSHOT":
		return db._Snapshot()
	default:
		return make([]byte, 0), invalidCommand
	}
}

func (db *JeuneDB) Commit(t *Tx.Transaction) ([]byte, error) {
	t.Status = "processing"
	db.mutex.Lock()
	defer db.mutex.Unlock()
	for _, v := range t.OpSeq {
		// res, err := db._Exec(v)
		db._Exec(v)
	}
	t.Status = "success"
	return make([]byte, 0), nil
}
