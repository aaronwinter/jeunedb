package storage

import "bufio"
import "bytes"
import "fmt"
import "encoding/binary"

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

func fillBuffer(size uint32, r *bufio.Reader) []byte {
	buff := make([]byte, size)
	_, err := r.Read(buff)
	if err != nil {
		fmt.Println("Error: ", err) // To replace w/ proper error handling
	}
	return buff
}

func toUint32(buff []byte) uint32 {
	return binary.BigEndian.Uint32(buff)
}

func parseHalfBlock(r *bufio.Reader) (uint32, []byte) {
	buffSize := fillBuffer(4, r)
	sizeEntry := toUint32(buffSize)
	bufferEntry := fillBuffer(sizeEntry, r)
	return sizeEntry, bufferEntry
}

func FetchBlockWithKey(targetKey []byte, reader *bufio.Reader) (*Block, error) {
	currentKey := make([]byte, 0)
	for bytes.Equal(targetKey, currentKey) == false {
		byteSizeKey, currentKey := parseHalfBlock(reader)
		byteSizeVal, currentVal := parseHalfBlock(reader)
		if bytes.Equal(targetKey, currentKey) == true {
			return &Block{
				Key:    currentKey,
				KeyL:   byteSizeKey,
				Value:  currentVal,
				ValueL: byteSizeVal,
			}, nil
		}
	}
	return nil, nil // to replace with errKeyNotFound
}
