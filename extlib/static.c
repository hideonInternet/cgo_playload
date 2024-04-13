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