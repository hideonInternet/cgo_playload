package main

/*
#include <stdio.h>

int Call(){
	printf("hello cgo\n");
	return 0;
}
*/
import "C"

func main() {
	C.Call()
}
