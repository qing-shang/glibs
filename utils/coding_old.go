//go:build !go1.20
// +build !go1.20

package utils

import (
	"reflect"
	"unsafe"
)

// B2s 已知有乱码问题
func B2s(b []byte) string {
	/* #nosec G103 */
	return *(*string)(unsafe.Pointer(&b))
}

// S2b string转[]byte 不够安全 运行端值有问题
func S2b(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}
