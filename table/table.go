package table

type Entry struct {
	size uint64
}

type Table struct {
	entries [100]Entry
	size    uint64
}

type Tablet struct {
	tables [50]Table
	size   uint64
}
