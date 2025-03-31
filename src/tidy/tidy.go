package tidy

/*
#cgo LDFLAGS: -L../ -ltidy
#include "headers/call.c"
#include <stdlib.h>
#include <stdint.h>
*/
import "C"

import (
	"unsafe"
)

func TidyHTML(input []byte) []byte {
	cInput := C.CBytes(input)
	defer C.free(unsafe.Pointer(cInput))

	cOutput := C.tidy_html((*C.char)(cInput), C.uint64_t(len(input)))

	if cOutput == nil {
		return nil
	}
	defer C.free(unsafe.Pointer(cOutput))

	return C.GoBytes(unsafe.Pointer(cOutput), C.int(C.strlen(cOutput)))
}
