package main

import "testing"

func BenchmarkVersion1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Version1()
	}
}

func BenchmarkVersion2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Version2()
	}
}
