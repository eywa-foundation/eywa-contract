package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// SendChatTypeKeyPrefix is the prefix to retrieve all SendChatType
	SendChatTypeKeyPrefix = "SendChatType/value/"
)

// SendChatTypeKey returns the store key to retrieve a SendChatType from the index fields
func SendChatTypeKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
