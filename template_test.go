// fib_test.go
package main

import (
	"bytes"
	"testing"
)

func BenchmarkT1(b *testing.B) {
	wr := bytes.NewBuffer([]byte{})
	for n := 0; n < b.N; n++ {
		T1(wr)
	}
}

func BenchmarkT2(b *testing.B) {
	wr := bytes.NewBuffer([]byte{})
	for n := 0; n < b.N; n++ {
		T2(wr)
	}
}
