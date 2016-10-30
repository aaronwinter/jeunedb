package tx

type Operation struct {
	key   []byte
	value []byte
	cmd   string // cmd := "Get" | "Put" | "Delete" | "Snapshot"
}

type Journal struct {
}

type Transaction struct {
	operations map[uint16][]Operation
	status     bool
	Journal
}
