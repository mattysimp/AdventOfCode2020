package main

import (
	"testing"
)

func TestExamplePart1(t *testing.T) {
	Result, _ := parts("inputexample.txt", "shiny gold")

	if Result != 4 {
		t.Errorf("Expected Result to be '4'. Got %d", Result)
	}
}

func TestExamplePart2(t *testing.T) {
	_, Result := parts("inputexample2.txt", "shiny gold")
	if Result != 126 {
		t.Errorf("Expected Result to be '126'. Got %d", Result)
	}
}

func BenchmarkParts(b *testing.B) {
	for x := 1; x <= b.N; x++ {
		parts("input.txt", "shiny gold")
	}
}
