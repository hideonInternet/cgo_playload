# Introduction
introduce how to use cgo-tech to call c/c++ function or import c/c++ library in golang program

The top directory is listed below. The `code` directory contains an example that demonstrates how to use cgo with raw C/C++ code, which is the most fundamental and important method.

The `extlib` directory stores the shared library and static library used to illustrate how to link Golang with C/C++ libraries. In most cases, we can only obtain an SDK composed of header files and library files.

The `shared` and `static` directories are used to introduce how to link Golang with existing libraries.

The `variables` directory is used to introduce the variable-convertion between golang and c/c++.

The `performance` directory is used to compare the perfomance difference while using cgo to reach the same goal.

```
├── code
├── extlib
├── LICENSE
├── Makefile
├── performance
├── readme.md
├── shared
├── static
└── variables
```

# Usage

First of all, execute `make build` to compile. then you can run `make run` to see the result. if necessary, you
can run `make clean` to clear all cache.


# Environment

```
system: ubuntu

gcc: 11.2.0
g++: 11.2.0

golang: 1.21.5
```