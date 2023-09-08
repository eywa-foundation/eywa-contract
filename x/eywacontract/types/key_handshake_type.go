package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// HandshakeTypeKeyPrefix is the prefix to retrieve all HandshakeType
	HandshakeTypeKeyPrefix = "HandshakeType/value/"
)

// HandshakeTypeKey returns the store key to retrieve a HandshakeType from the index fields
func HandshakeTypeKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
