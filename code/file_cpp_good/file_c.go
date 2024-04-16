package main

/*
#include "file_c.h"
*/
import "C"

func main() {
	C.Call()
	C.CallCPP()
}
