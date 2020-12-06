package main

import (
	"testing"
)

func TestExamplePart1(t *testing.T) {
	Result, _ := parts("inputexample.txt")

	if Result != 11 {
		t.Errorf("Expected Result to be '11'. Got %d", Result)
	}
}

func TestExamplePart2(t *testing.T) {
	_, Result := parts("inputexample.txt")
	if Result != 6 {
		t.Errorf("Expected Result to be '6'. Got %d", Result)
	}
}

func BenchmarkParts(b *testing.B) {
	for x := 1; x <= b.N; x++ {
		parts("input.txt")
	}
}
func BenchmarkPartsStress(b *testing.B) {
	for x := 1; x <= b.N; x++ {
		parts("inputstress.txt")
	}
}
