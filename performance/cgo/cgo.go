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
)

func main() {
	start := time.Now().UnixMilli()

	n, _ := strconv.ParseInt(os.Args[1], 10, 64)

	res := C.Fibonacci(C.int(n))

	fmt.Printf("cgo res: %d cost: %d\n", res, time.Now().UnixMilli()-start)
}
