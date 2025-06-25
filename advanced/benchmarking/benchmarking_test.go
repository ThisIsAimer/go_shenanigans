package main

import (
	"math/rand"
	"testing"
)

// for bechmarking
// go test -bench=. -benchmem ./benchmarking_test.go

// functions should start with Benchmark

func add(a, b int) int {
	return  a+b
}

func BenchmarkAddSmallInput(b *testing.B){
 for range b.N{
	add(9,10)
 }
}

func BenchmarkAddMediumInput(b *testing.B){
 for range b.N{
	add(200,300)
 }
}


func BenchmarkAddLargeInput(b *testing.B){
 for range b.N{
	add(4000,5000)
 }
}

//------------------------------------------------------------

func GenerateRandomSlice(size int) [] int{
	slice := make([]int,size)
	for i := range slice{
		slice[i] = rand.Intn(100)+1
	}
	return slice
}


func SumSlice(slice []int ) int{
	sum := 0
	for _, value := range slice{
		sum += value
	}
	return sum
}


func TestGenerateRandomSlice(t *testing.T) {
	size := 100
	slice := GenerateRandomSlice(size)
	if len(slice) != size {
		t.Errorf("expected slice size %d, received %d", size, len(slice))
	}
}

func BenchmarkGenerateRandomSlice(b *testing.B) {
	for range b.N {
		GenerateRandomSlice(1000)
	}
}

func BenchmarkSumSlice(b *testing.B) {
	slice := GenerateRandomSlice(1000)
	b.ResetTimer()
	for range b.N {
		SumSlice(slice)
	}
}