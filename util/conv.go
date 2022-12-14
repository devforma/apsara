package util

import (
	"encoding/json"
	"strconv"
	"unsafe"
)

func Int64ToString(num int64) string {
	return strconv.FormatInt(num, 10)
}

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

func JsonMarshal(data any) string {
	d, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return BytesToString(d)
}
