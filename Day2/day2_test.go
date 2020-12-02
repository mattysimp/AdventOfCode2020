package main

import (
	"testing"
)

func TestExample(t *testing.T) {
	Part1, Part2 := Parts("inputexample.txt")
	if Part1 != 2 {
		t.Errorf("Expected Part1 to be '2'. Got %d", Part1)
	}
	if Part2 != 1 {
		t.Errorf("Expected Part2 to be '1'. Got %d", Part2)
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
