package cache

type Cache struct {
	c    map[string][]byte
	size uint64
}
