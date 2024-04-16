package main

import (
	"net"
	"os"
	"syscall"
	"testing"
)

func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacci(30)
	}
}

const message = "hello, world!"

var buffer = make([]byte, 13)

func BenchmarkReadWriteNetCalls(b *testing.B) {
	cs, _ := socketpair()
	for i := 0; i < b.N; i++ {
		cs[0].Write([]byte(message))
		cs[1].Read(buffer)
	}
}

func socketpair() (conns [2]net.Conn, err error) {
	fds, err := syscall.Socketpair(syscall.AF_LOCAL, syscall.SOCK_STREAM, 0)
	if err != nil {
		return
	}
	conns[0], err = fdToFileConn(fds[0])
	if err != nil {
		return
	}
	conns[1], err = fdToFileConn(fds[1])
	if err != nil {
		conns[0].Close()
		return
	}
	return
}

func fdToFileConn(fd int) (net.Conn, error) {
	f := os.NewFile(uintptr(fd), "")
	defer f.Close()
	return net.FileConn(f)
}
