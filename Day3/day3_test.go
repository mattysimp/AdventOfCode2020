package main

import (
	"testing"
)

func TestExamplePart1(t *testing.T) {
	Input, MaxWidth, _ := readLines("inputexample.txt")

	Result := part1(Input, MaxWidth, 3, 1)

	if Result != 7 {
		t.Errorf("Expected Part1 to be '7'. Got %d", Result)
	}
}

func TestExamplePart2(t *testing.T) {
	Input, MaxWidth, _ := readLines("inputexample.txt")

	Result := part2(Input, MaxWidth)

	if Result != 336 {
		t.Errorf("Expected Part1 to be '336'. Got %d", Result)
	}
}

func BenchmarkParts(b *testing.B) {
	for x := 1; x <= b.N; x++ {
		Input, MaxWidth, _ := readLines("input.txt")
		part1(Input, MaxWidth, 3, 1)
	}
}
func BenchmarkPartsStress(b *testing.B) {
	for x := 1; x <= b.N; x++ {
		Input, MaxWidth, _ := readLines("input.txt")
		part2(Input, MaxWidth)
	}
}
