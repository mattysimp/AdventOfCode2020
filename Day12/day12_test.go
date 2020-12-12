package main

import (
	"AdventOfCode2020/fileReader"
	"testing"
)

func TestExamplePart1(t *testing.T) {
	input := fileReader.ReadLines("inputexample.txt")
	Result := part1(input)

	if Result != 25 {
		t.Errorf("Expected Result to be '25'. Got %d", Result)
	}
}

func TestExamplePart2(t *testing.T) {
	input := fileReader.ReadLines("inputexample.txt")
	Result := part2(input)
	if Result != 286 {
		t.Errorf("Expected Result to be '286'. Got %d", Result)
	}
}

func BenchmarkPart1(b *testing.B) {
	for x := 1; x <= b.N; x++ {
		input := fileReader.ReadLines("input.txt")
		part1(input)
	}
}
func BenchmarkPart2WithPart1(b *testing.B) {
	for x := 1; x <= b.N; x++ {
		input := fileReader.ReadLines("input.txt")
		part2(input)
	}
}
