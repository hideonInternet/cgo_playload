#ifdef __cplusplus
extern "C" {
#include "file_c.h"
}
#endif
#include <iostream>
#include <stdio.h>

void CallCPP(){
    std::cout << "hello cgo i am cpp" << std::endl;
}

void Call(){
    printf("hello cgo\n");
}