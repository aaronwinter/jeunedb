package table

type Entry struct {
	key    []byte
	value  []byte
	size   uint64
	offset uint16
}

type Table struct {
	entries [100]Entry
	size    uint64
}

type Tablet struct {
	tables [50]Table
	size   uint64
}
