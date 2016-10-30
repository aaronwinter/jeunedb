package cache

type Cache struct {
	c    map[byte][]byte
	size uint64
}

func New() Cache {
	cache := Cache{
		c:    make(map[byte][]byte, 0),
		size: 0,
	}
	return cache
}
