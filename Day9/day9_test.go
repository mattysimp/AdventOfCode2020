package main

import (
	"AdventOfCode2020/fileReader"
	"testing"
)

func TestExamplePart1(t *testing.T) {
	input := fileReader.ReadLinesXMAS("inputexample.txt")
	Result := part1(input, 5)

	if Result.Num != 127 {
		t.Errorf("Expected Result to be '127'. Got %d", Result)
	}
}

func TestExamplePart2(t *testing.T) {
	input := fileReader.ReadLinesXMAS("inputexample.txt")
	Result := part2(input, fileReader.XMAS{Num: 127, Place: 14})
	if Result != 62 {
		t.Errorf("Expected Result to be '62'. Got %d", Result)
	}
}

func BenchmarkPart1(b *testing.B) {
	for x := 1; x <= b.N; x++ {
		input := fileReader.ReadLinesXMAS("input.txt")
		part1(input, 25)
	}
}
func BenchmarkPart2WithPart1(b *testing.B) {
	for x := 1; x <= b.N; x++ {
		input := fileReader.ReadLinesXMAS("input.txt")
		part2(input, fileReader.XMAS{Num: 258585477, Place: 593})
	}
}
