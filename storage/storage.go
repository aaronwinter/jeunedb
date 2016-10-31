package storage

type Block struct {
	Key    []byte
	KeyL   uint32
	Value  []byte
	ValueL uint32
}

func NewBlock(k []byte, v []byte) *Block {
	lenK := uint32(len(k))
	lenV := uint32(len(v))
	b := &Block{
		Key:    k,
		KeyL:   lenK,
		Value:  v,
		ValueL: lenV,
	}
	return b
}
