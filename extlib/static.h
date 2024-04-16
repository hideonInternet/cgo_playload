#include <stdio.h>
#include <unistd.h>

int Add(int a, int b);

int Del(int a, int b);

unsigned long long Fibonacci(int n);

int write_all(int fd, void* buffer, size_t length);

int read_call(int fd, void *buffer, size_t length);