package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func fibonacci(n int64) int64 {
	if n <= 1 {
		return n
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

func main() {
	start := time.Now().UnixMilli()

	n, _ := strconv.ParseInt(os.Args[1], 10, 64)

	res := fibonacci(n)

	fmt.Printf("golang res: %d cost: %d\n", res, time.Now().UnixMilli()-start)
}
