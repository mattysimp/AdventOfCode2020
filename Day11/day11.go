package main

import (
	"AdventOfCode2020/fileReader"
	"fmt"
	"image"
)

func main() {

	input := fileReader.ReadLinesCoordMap("Day11/input.txt")

	fmt.Println(parts(input, 4, func(p1 image.Point, p2 image.Point) image.Point { return p1.Add(p2) }))
	fmt.Println(parts(input, 5, func(p1 image.Point, p2 image.Point) image.Point {
		for input[p1.Add(p2)] == '.' {
			p1 = p1.Add(p2)
		}
		return p1.Add(p2)
	}))

}

func parts(input map[image.Point]rune, tolerance int, addFunc func(image.Point, image.Point) image.Point) int {
	nextInput := make(map[image.Point]rune)
	change := false
	occupiedSeat := 0
	for point, seatVal := range input {
		occupiedAdjacents := 0
		for x := -1; x <= 1; x++ {
			for y := -1; y <= 1; y++ {
				if x != 0 || y != 0 {
					if input[addFunc(point, image.Point{x, y})] == '#' {
						occupiedAdjacents++
					}
				}
			}
		}
		if occupiedAdjacents >= tolerance && seatVal != '.' {
			nextInput[point] = 'L'
		} else if occupiedAdjacents == 0 && seatVal != '.' {
			nextInput[point] = '#'
		} else {
			nextInput[point] = seatVal
		}

		if nextInput[point] == '#' {
			occupiedSeat++
		}
		if seatVal != nextInput[point] {
			change = true
		}
	}

	if change {
		return parts(nextInput, tolerance, addFunc)
	} else {
		return occupiedSeat
	}

}
