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
	status     string // status := "init" | "queued" | "processing" | "over"
	Journal
}

func New() Transaction {
	t := Transaction{
		operations: make(map[uint16][]Operation),
		status:     "init",
		Journal:    Journal{},
	}
	return t
}
