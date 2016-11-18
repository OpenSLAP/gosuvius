package main

// #cgo CFLAGS: -std=c11  -D__CGO
// #cgo CXXFLAGS: -std=c++11  -D__CGO
// #include "structures.hpp"
import "C"

type ActionFrame struct {
	raw C.ActionFrame
}

type RequestFrame struct {
	raw C.RequestFrame
}

type SyncFrame struct {
	raw C.SyncFrame
}
