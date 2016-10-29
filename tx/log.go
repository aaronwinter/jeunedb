package transaction

import "sync"

type Log struct {
	mu       sync.Mutex
	buffer   []byte
	prefix   string
	BasePath string
}
