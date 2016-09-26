package util

import (
	"reflect"
	"unsafe"
)

type Msg struct {
	Type int
	Name [8]byte
}

const (
	REGISTER_NAME = iota
	SEARCH_NAME
)

func Bytes2Msg(bytes []byte) *Msg {
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&bytes))
	return (*Msg)(unsafe.Pointer(sh.Data))
}

func Msg2Bytes(msg *Msg) []byte {
	var bytes []byte
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&bytes))
	sh.Data = uintptr(unsafe.Pointer(msg))
	sh.Len = 16
	sh.Cap = 16
	return bytes
}

func Array2bytes(array [8]byte) []byte {
	t := make([]byte, 8)
	for i := 0; i < 8; i++ {
		t[i] = array[i]
	}
	return t
}

func Bytes2array(bytes []byte) [8]byte {
	var t [8]byte
	for i := 0; i < 8; i++ {
		t[i] = bytes[i]
	}
	return t
}
