package hash

import (
	"encoding/hex"
	"github.com/qing-shang/glibs"
)

type Uint128 struct {
	Hi uint64
	Lo uint64
}

func (c Uint128) bytes() []byte {
	return []byte{
		byte(c.Hi >> 56), byte(c.Hi >> 48), byte(c.Hi >> 40), byte(c.Hi >> 32),
		byte(c.Hi >> 24), byte(c.Hi >> 16), byte(c.Hi >> 8), byte(c.Hi),

		byte(c.Lo >> 56), byte(c.Lo >> 48), byte(c.Lo >> 40), byte(c.Lo >> 32),
		byte(c.Lo >> 24), byte(c.Lo >> 16), byte(c.Lo >> 8), byte(c.Lo),
	}
}

func (c Uint128) Bytes() []byte {
	hashedValue := c.bytes()
	src := make([]byte, hex.EncodedLen(len(hashedValue)))
	hex.Encode(src, hashedValue)
	return src
}

func (c Uint128) String() string {
	hashedValue := c.bytes()
	return hex.EncodeToString(hashedValue)
}

func (c Uint128) Reset() {
	c.Hi = 0
	c.Lo = 0
}

var PoolUint128 = glibs.NewPool[Uint128]()
