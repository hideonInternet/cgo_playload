#include "static.h"

int Add(int a, int b) {
    return a + b;
}

int Del(int a, int b){
    return a - b;
}

unsigned long long Fibonacci(int n) {
    if (n <= 1) {
        return n;
    } else {
        return Fibonacci(n - 1) + Fibonacci(n - 2);
    }
}

int write_all(int fd, void* buffer, size_t length) {
    while (length > 0) {
        int written = write(fd, buffer, length);
        if (written < 0)
            return -1;
        length -= written;
        buffer += written;
    }
    return length;
}

int read_call(int fd, void *buffer, size_t length) {
	return read(fd, buffer, length);
}