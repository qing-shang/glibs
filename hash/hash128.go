package hash

import (
	"github.com/twmb/murmur3"
)

func SeedStringSum128(data string, seed1, seed2 uint64) Uint128 {
	h1, h2 := murmur3.SeedStringSum128(seed1, seed2, data)
	return Uint128{
		Hi: h1,
		Lo: h2,
	}
}

func Sum128String(data string) Uint128 {
	h1, h2 := murmur3.StringSum128(data)
	return Uint128{
		Hi: h1,
		Lo: h2,
	}
}

func Sum128(b []byte) Uint128 {
	h1, h2 := murmur3.Sum128(b)
	return Uint128{
		Hi: h1,
		Lo: h2,
	}
}
