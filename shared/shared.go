package main

/*
#cgo CFLAGS: -I../extlib

#cgo LDFLAGS: -L${SRCDIR}/../extlib/ -lcomm

#include "static.h"
*/
import "C"
import "fmt"

func main() {
	a, b := 1, 2
	c := C.Add(C.int(a), C.int(b))
	fmt.Printf("%d\n", c)
	c = C.Del(C.int(a), C.int(b))
	fmt.Printf("%d\n", c)
}
