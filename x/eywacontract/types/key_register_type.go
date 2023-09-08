package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// RegisterTypeKeyPrefix is the prefix to retrieve all RegisterType
	RegisterTypeKeyPrefix = "RegisterType/value/"
)

// RegisterTypeKey returns the store key to retrieve a RegisterType from the index fields
func RegisterTypeKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
