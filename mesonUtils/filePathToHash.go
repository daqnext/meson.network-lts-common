package mesonUtils

import (
	"encoding/binary"
	"encoding/hex"

	"github.com/segmentio/fasthash/fnv1a"
)

func FilePathToHash(bindName string, filePath string) string {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, fnv1a.HashString64(bindName+filePath))
	return hex.EncodeToString(b)
}
