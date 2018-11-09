package main

import (
	"testing"
	"time"
)

func TestPrin1(t *testing.T) {
	print1()

}

func TestGoPrint1(t *testing.T) {
	goPrint1()
	time.Sleep(1 * time.Millisecond)
}

func TestGoPrint2(t *testing.T) {
	goPrint2()
	time.Sleep(1 * time.Millisecond)
}

func BenchmarkPrint1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		print1()
	}
}

func BenchmarkGoPrint1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goPrint1()
	}
}

// go test -run x -bench .
// goos: darwin
// goarch: amd64
// pkg: github.com/yurakawa/gowebprog/chap09/1goroutine
// BenchmarkPrint1-4       100000000               10.9 ns/op
// BenchmarkGoPrint1-4      5000000               318 ns/op
// PASS
// ok      github.com/yurakawa/gowebprog/chap09/1goroutine 3.020s

func BenchmarkPrint2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		print2()
	}
}

func BenchmarkGoPrint2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goPrint2()
	}
}

// go test -run x -bench .
// goos: darwin
// goarch: amd64
// pkg: github.com/yurakawa/gowebprog/chap09/1goroutine
// BenchmarkPrint1-4       100000000               10.8 ns/op
// BenchmarkGoPrint1-4      5000000               316 ns/op
// BenchmarkPrint2-4          10000            191343 ns/op
// BenchmarkGoPrint2-4       300000              7992 ns/op
// PASS
// ok      github.com/yurakawa/gowebprog/chap09/1goroutine 7.520s
