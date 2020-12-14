package main

import (
	"AdventOfCode2020/fileReader"
	"image"
	"testing"
)

func TestExamplePart1(t *testing.T) {
	input := fileReader.ReadLinesCoordMap("inputexample.txt")

	Result := parts(input, 4, func(p1 image.Point, p2 image.Point) image.Point { return p1.Add(p2) })

	if Result != 37 {
		t.Errorf("Expected Result to be '37'. Got %d", Result)
	}
}

func TestExamplePart2(t *testing.T) {
	input := fileReader.ReadLinesCoordMap("inputexample.txt")

	Result := parts(input, 5, func(p1 image.Point, p2 image.Point) image.Point {
		for input[p1.Add(p2)] == '.' {
			p1 = p1.Add(p2)
		}
		return p1.Add(p2)
	})

	if Result != 26 {
		t.Errorf("Expected Result to be '26'. Got %d", Result)
	}
}

func BenchmarkPart1(b *testing.B) {
	for x := 1; x <= b.N; x++ {
		input := fileReader.ReadLinesCoordMap("input.txt")
		parts(input, 4, func(p1 image.Point, p2 image.Point) image.Point { return p1.Add(p2) })
	}
}
func BenchmarkPart2(b *testing.B) {
	for x := 1; x <= b.N; x++ {
		input := fileReader.ReadLinesCoordMap("input.txt")
		parts(input, 5, func(p1 image.Point, p2 image.Point) image.Point {
			for input[p1.Add(p2)] == '.' {
				p1 = p1.Add(p2)
			}
			return p1.Add(p2)
		})
	}
}
