package main

import (
	"AdventOfCode2020/fileReader"
	"sort"
	"testing"
)

func TestExamplePart1(t *testing.T) {
	input := fileReader.ReadLinesInt("inputexample.txt")
	sort.Ints(input)
	fullInput := append([]int{0}, input...)

	Result := part1(fullInput)

	if Result != 35 {
		t.Errorf("Expected Result to be '35'. Got %d", Result)
	}
}

func TestExamplePart2(t *testing.T) {
	input := fileReader.ReadLinesInt("inputexample.txt")
	sort.Ints(input)
	fullInput := append([]int{0}, input...)

	Result := part2(fullInput)

	if Result != 8 {
		t.Errorf("Expected Result to be '8'. Got %d", Result)
	}
}

func BenchmarkPart1(b *testing.B) {
	for x := 1; x <= b.N; x++ {
		input := fileReader.ReadLinesInt("inputexample.txt")
		sort.Ints(input)
		fullInput := append([]int{0}, input...)
		part1(fullInput)
	}
}
func BenchmarkPart2WithPart1(b *testing.B) {
	for x := 1; x <= b.N; x++ {
		input := fileReader.ReadLinesInt("inputexample.txt")
		sort.Ints(input)
		fullInput := append([]int{0}, input...)
		part2(fullInput)
	}
}
