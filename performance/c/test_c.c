#include <stdio.h>
#include <stdlib.h>
#include "static.h"
#include <sys/time.h>

long long ts();

int main(int argc, char *argv[]) {
    if (argc < 2) {
        printf("Usage: %s num ...\n", argv[0]);
        return 1;
    }

    int n = atoi(argv[1]);

    long long start = ts();
    unsigned long long res = Fibonacci(n);

    printf("pure c res: %lld, cost: %lld\n", res, ts() - start);
}

long long ts() {
    struct timeval tv;
    gettimeofday(&tv, NULL);

    long long timestamp_ms = (long long)tv.tv_sec * 1000 + (long long)tv.tv_usec / 1000;

    return timestamp_ms;
}