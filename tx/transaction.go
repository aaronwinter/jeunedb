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

type Queue struct {
	queue map[uint32]Transaction
	size  uint32
}

func (q *Queue) push(tx *Transaction) {
	q.queue[q.size] = *tx
	q.size = q.size + 1
}

func (q *Queue) pop() {
	delete(q.queue, q.size)
	q.size = q.size - 1
}


func New() *Transaction {
	t := &Transaction{
		OpSeq:   make(map[uint32]Operation),
		numOps:  0,
		Status:  "init",
		Journal: Journal{},
	}
	return t
}
