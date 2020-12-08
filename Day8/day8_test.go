package main

import (
	"testing"
)

func TestExamplePart1(t *testing.T) {
	Result := part1("inputexample.txt")

	if Result != 5 {
		t.Errorf("Expected Result to be '5'. Got %d", Result)
	}
}

func TestExamplePart2(t *testing.T) {
	Result := part2("inputexample.txt")
	if Result != 8 {
		t.Errorf("Expected Result to be '8'. Got %d", Result)
	}
}

func BenchmarkPart1(b *testing.B) {
	for x := 1; x <= b.N; x++ {
		part1("input.txt")
	}
}
func BenchmarkPart2(b *testing.B) {
	for x := 1; x <= b.N; x++ {
		part2("input.txt")
	}
}
