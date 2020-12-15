package main

import (
	"testing"
)

func TestExamplePart1(t *testing.T) {
	input := []int{1, 3, 2}

	Result := parts(input, 2020)

	if Result != 1 {
		t.Errorf("Expected Result to be '1'. Got %d", Result)
	}
}

func TestExamplePart2(t *testing.T) {
	input := []int{1, 3, 2}

	Result := parts(input, 30000000)

	if Result != 2578 {
		t.Errorf("Expected Result to be '2578'. Got %d", Result)
	}
}

func BenchmarkPart1(b *testing.B) {
	for x := 1; x <= b.N; x++ {
		input := []int{16, 12, 1, 0, 15, 7, 11}
		parts(input, 2020)
	}
}
func BenchmarkPart2(b *testing.B) {
	for x := 1; x <= b.N; x++ {
		input := []int{16, 12, 1, 0, 15, 7, 11}
		parts(input, 30000000)
	}
}
