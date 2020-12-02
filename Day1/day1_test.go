package main

import (
	"testing"
)

func TestExamplePart1(t *testing.T) {

	Result := day1([]int{1721,
		979,
		366,
		299,
		675,
		1456})

	if Result != 514579 {
		t.Errorf("Expected Part1 to be '514579'. Got %d", Result)
	}
}
func TestExamplePart2(t *testing.T) {

	Result := day2([]int{1721,
		979,
		366,
		299,
		675,
		1456})

	if Result != 241861950 {
		t.Errorf("Expected Part1 to be '241861950'. Got %d", Result)
	}
}

func BenchmarkPart1(b *testing.B) {
	for x := 1; x <= b.N; x++ {
		Input, err := readLines("input.txt")
		if err != nil {
			panic(err)
		}
		day1(Input)
	}
}
func BenchmarkParts2(b *testing.B) {
	for x := 1; x <= b.N; x++ {
		Input, err := readLines("input.txt")
		if err != nil {
			panic(err)
		}
		day2(Input)
	}
}
