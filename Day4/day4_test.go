package main

import (
	"testing"
)

func TestExamplePart1(t *testing.T) {
	Result, _ := Parts("inputexample.txt")

	if Result != 2 {
		t.Errorf("Expected Result to be '2'. Got %d", Result)
	}
}

func TestExamplePart2(t *testing.T) {
	_, Result := Parts("inputexample2.txt")
	if Result != 4 {
		t.Errorf("Expected Result to be '4'. Got %d", Result)
	}
}

func BenchmarkParts(b *testing.B) {
	for x := 1; x <= b.N; x++ {
		Parts("input.txt")
	}
}
func BenchmarkPartsStress(b *testing.B) {
	for x := 1; x <= b.N; x++ {
		Parts("inputstress.txt")
	}
}
