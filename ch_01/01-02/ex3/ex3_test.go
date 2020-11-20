package main

import (
	"testing"
)

func BenchmarkEchoInefficientLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EchoInefficientLoop()
	}
}

func BenchmarkEchoEfficientJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EchoEfficientJoin()
	}
}
