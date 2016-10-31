package tx

type Operation struct {
	Key   []byte
	Value []byte
	Cmd   string // cmd := "Get" | "Put" | "Delete" | "Snapshot"
}

type Journal struct {
}

type Transaction struct {
	OpSeq  map[uint32]Operation
	numOps uint32
	Status string // := "init" | "queued" | "processing "| "success" | "failure"
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
