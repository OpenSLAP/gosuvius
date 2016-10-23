package main

// #include "structures.hpp"
// #cgo CFLAGS: -std=c11
// #cgo CXXFLAGS: -std=c++11
import "C"

type Frame struct {
	raw C.Frame
}
