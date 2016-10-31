package jeunedb

//import Error "jeunedb/errors"
//import Table "jeunedb/table"
//import "io"
import "fmt"
import "sync"
import "errors"
import "os"
import "bufio"

import Log "jeunedb/log"
import Cache "jeunedb/cache"
import Tx "jeunedb/tx"
import Storage "jeunedb/storage"

var invalidCommand error = errors.New("Invalid command")

const osWriteFlag int = os.O_CREATE | os.O_WRONLY | os.O_APPEND
const osReadFlag int = os.O_RDONLY

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
		c.permFile = 0777
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

func (db *JeuneDB) _Put(block *Storage.Block) ([]byte, error) {
	buffWrite := block.Serialize()
	f, err := os.OpenFile(db.Config.BasePath, osWriteFlag, db.Config.permFile)
	defer f.Close()
	if err != nil {
		fmt.Println("Error: ", err) // to replace w/ proper error handling
	}
	writer := bufio.NewWriter(f)
	writer.Write(buffWrite.Bytes())
	writer.Flush()
	return make([]byte, 0), nil
}

func (db *JeuneDB) _Get(key []byte) ([]byte, error) {
	fmt.Println("_Get")
	// Todo: Bloom filter
	f, _ := os.OpenFile(db.Config.BasePath, osReadFlag, db.Config.permFile)
	defer f.Close()
	reader := bufio.NewReader(f)
	entry, err := Storage.FetchBlockWithKey(key, reader)
	if err != nil {
		return make([]byte, 0), err
	} else {
		fmt.Println("Found key = ", key)
		fmt.Println("Key -value->", entry.Value)
		return entry.Value, nil
	}
}

func (db *JeuneDB) _Snapshot() ([]byte, error) {
	return make([]byte, 0), nil
}

func (db *JeuneDB) _Exec(o Tx.Operation) ([]byte, error) {
	switch o.Cmd {
	case "GET":
		return db._Get(o.Key)
	case "PUT":
		b := Storage.NewBlock(o.Key, o.Value)
		return db._Put(b)
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
	for _, op := range t.OpSeq {
		// res, err := db._Exec(v)
		db._Exec(op)
	}
	t.Status = "success"
	return make([]byte, 0), nil
}
