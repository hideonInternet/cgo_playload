package main

/*
#cgo CFLAGS: -I../../extlib

#cgo LDFLAGS: ${SRCDIR}/../../extlib/libcomm.a

#include "static.h"
*/
import "C"
import (
	"fmt"
	"os"
	"strconv"
	"time"
	"unsafe"
)

func CFibonacci(n int) int64 {
	res := C.Fibonacci(C.int(n))
	return int64(res)
}

// CwriteAll is a cgo call for write
func CwriteAll(fd int, buf []byte) error {
	_, err := C.write_all(C.int(fd), unsafe.Pointer(&buf[0]), C.size_t(len(buf)))
	return err
}

// Cread is a cgo call for read
func Cread(fd int, buf []byte) (int, error) {
	ret, err := C.read_call(C.int(fd), unsafe.Pointer(&buf[0]), C.size_t(len(buf)))
	return int(ret), err
}

func main() {
	start := time.Now().UnixMilli()
	n, _ := strconv.ParseInt(os.Args[1], 10, 64)
	res := CFibonacci(int(n))
	fmt.Printf("cgo res: %d cost: %d\n", res, time.Now().UnixMilli()-start)
}
