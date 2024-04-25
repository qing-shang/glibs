package hash

import (
	"encoding/hex"
)

type Uint32 uint64

func (c Uint32) Bytes() []byte {
	return []byte{
		byte(c >> 24), byte(c >> 16), byte(c >> 8), byte(c),
	}
}

func (c Uint32) HexBytes() []byte {
	hashedValue := c.Bytes()
	src := make([]byte, hex.EncodedLen(len(hashedValue)))
	hex.Encode(src, hashedValue)
	return src
}

func (c Uint32) String() string {
	hashedValue := c.Bytes()
	return hex.EncodeToString(hashedValue)
}
