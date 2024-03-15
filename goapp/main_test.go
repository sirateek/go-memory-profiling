package main_test

import (
	"os"
	"runtime/pprof"
	"testing"
)

type TestBool bool

func TestHelloWorld(m *testing.T) {
	f, _ := os.Create("heap.out")
	defer func() {
		f.Close()
	}()

	a := make([]TestBool, 100000000)
	for i := 0; i < 100000000; i++ {
		a[i] = true
	}

	pprof.WriteHeapProfile(f)
}


fun
