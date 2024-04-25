package hash

import "github.com/twmb/murmur3"

func SeedStringSum32(data string, seed1 uint32) Uint32 {
	v := murmur3.SeedStringSum32(seed1, data)
	return Uint32(v)
}

func Sum32String(data string) Uint32 {
	v := murmur3.StringSum32(data)
	return Uint32(v)
}

func Sum32(b []byte) Uint32 {
	v := murmur3.Sum32(b)
	return Uint32(v)
}
