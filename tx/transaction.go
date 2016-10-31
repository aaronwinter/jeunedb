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

func NewOperation(cmd string, k []byte, v []byte) Operation {
	op := Operation{
		Key:   k,
		Value: v,
		Cmd:   cmd,
	}
	return op
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

func (t *Transaction) _AddOperation(cmd string, key []byte, argV ...[]byte) {
	value := argV[0]
	op := NewOperation(cmd, key, value)
	t.OpSeq[t.numOps+1] = op
	t.numOps = t.numOps + 1
}

func (t *Transaction) Put(key []byte, value []byte) {
	t._AddOperation("PUT", key, value)
}

func (t *Transaction) Get(key []byte) {
	t._AddOperation("GET", key)
}
