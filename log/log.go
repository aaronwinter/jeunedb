package log

import "sync"

type Log struct {
	mu       sync.Mutex
	buffer   []byte
	prefix   string
	BasePath string
}

func New(pathToLog string, prefix string) Log {
	log := Log{
		buffer:   make([]byte, 1024),
		BasePath: pathToLog,
		prefix:   prefix,
	}
	return log
}
