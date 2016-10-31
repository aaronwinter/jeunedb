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

func (b *Block) Serialize() *bytes.Buffer {
	toBuffer := new(bytes.Buffer)
	toBuffer.Reset()
	binary.Write(toBuffer, binary.BigEndian, b.KeyL)
	binary.Write(toBuffer, binary.BigEndian, b.Key)
	binary.Write(toBuffer, binary.BigEndian, b.ValueL)
	binary.Write(toBuffer, binary.BigEndian, b.Value)
	return toBuffer
}
