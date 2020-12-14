package main

import (
	"AdventOfCode2020/fileReader"
	"testing"
)

func TestExamplePart1(t *testing.T) {
	input := fileReader.ReadLines("inputexample.txt")

	Result := parts(input, true)

	if Result != 165 {
		t.Errorf("Expected Result to be '165'. Got %d", Result)
	}
}

func TestExamplePart2(t *testing.T) {
	input := fileReader.ReadLines("inputexample2.txt")
	Result := parts(input, false)
	if Result != 208 {
		t.Errorf("Expected Result to be '208'. Got %d", Result)
	}
}

func BenchmarkPart1(b *testing.B) {
	for x := 1; x <= b.N; x++ {
		input := fileReader.ReadLines("input.txt")
		parts(input, true)
	}
}
func BenchmarkPart2(b *testing.B) {
	for x := 1; x <= b.N; x++ {
		input := fileReader.ReadLines("input.txt")
		parts(input, false)
	}
}
