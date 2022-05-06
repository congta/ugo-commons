package ubytes

import "encoding/binary"

func IntToBytes(i int) []byte {
	bytes := make([]byte, 4)
	binary.BigEndian.PutUint32(bytes, uint32(i))
	return bytes
}

func Int16ToBytes(i uint16) []byte {
	bytes := make([]byte, 2)
	binary.BigEndian.PutUint16(bytes, i)
	return bytes
}

// ShortToBytes alias of Int8ToBytes
func ShortToBytes(i uint16) []byte {
	return Int16ToBytes(i)
}

func BytesToInt16(data []byte) uint16 {
	return binary.BigEndian.Uint16(data)
}
