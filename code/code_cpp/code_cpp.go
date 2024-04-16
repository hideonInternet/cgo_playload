package main

/*
#cgo CXXFLAGS: -I/usr/include/c++/11

extern "C"{
#include <iostream.h>
}

int Call(){
	std::cout << "hellp cgo" << std::endl;
	return 0;
}

*/
import "C"

func main() {
	C.Call()
}
