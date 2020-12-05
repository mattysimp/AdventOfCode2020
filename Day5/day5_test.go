package main

import (
	"testing"
)

func TestExamplePart1(t *testing.T) {
	Result, _ := parts("inputexample.txt")

	if Result != 820 {
		t.Errorf("Expected Result to be '2'. Got %d", Result)
	}
}

func TestExamplePart2(t *testing.T) {
	_, Result := parts("input.txt")
	if Result != 640 {
		t.Errorf("Expected Result to be '640'. Got %d", Result)
	}
}

func BenchmarkParts(b *testing.B) {
	for x := 1; x <= b.N; x++ {
		parts("input.txt")
	}
}