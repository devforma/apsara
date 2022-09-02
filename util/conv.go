package util

import "unsafe"

func StringToBytes(str string) []byte {
	return *(*[]byte)(unsafe.Pointer(&struct {
		Data string
		Cap  int64
	}{
		Data: str,
		Cap:  int64(len(str)),
	}))
}

func BytesToString(bytes []byte) string {
	return *(*string)(unsafe.Pointer(&bytes))
}
