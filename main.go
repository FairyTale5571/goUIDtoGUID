package main

/*
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
*/
import "C"

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"unsafe"
)

func convert(uidInput int64) string {
	h := md5.New()
	h.Write([]byte("BE"))

	for i := 0; i < 8; i++ {
		h.Write([]byte{byte(uidInput & 0xFF)})
		uidInput >>= 8
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}

//export RVExtensionVersion
func RVExtensionVersion(output *C.char, outputsize C.size_t) {
	result := C.CString("uid to guid by FairyTale5571 v1.0")
	defer C.free(unsafe.Pointer(result))
	var size = C.strlen(result) + 1
	if size > outputsize {
		size = outputsize
	}
	C.memmove(unsafe.Pointer(output), unsafe.Pointer(result), size)
}

//export RVExtension
func RVExtension(output *C.char, outputsize C.size_t, input *C.char) {
	toConv, _ := strconv.ParseInt(C.GoString(input), 0, 64)
	result := C.CString(convert(toConv))
	defer C.free(unsafe.Pointer(result))
	var size = C.strlen(result) + 1
	if size > outputsize {
		size = outputsize
	}
	C.memmove(unsafe.Pointer(output), unsafe.Pointer(result), size)
}

func main() {}
