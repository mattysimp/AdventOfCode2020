package main

import (
	"AdventOfCode2020/fileReader"
	"fmt"
	"image"
	"math"
	"strconv"
)

type ferry struct {
	position image.Point
	angle    float64
}

func (f *ferry) turn(degrees float64) {
	rad := degrees * math.Pi / 180
	f.angle += rad
}

func (f *ferry) move(amount int, w ferry) {
	for x := 0; x < amount; x++ {
		f.position = f.position.Add(w.position)
	}
}

func (f *ferry) rotate(degrees float64) {
	rad := float64(degrees) * math.Pi / 180
	X := float64(f.position.X)*math.Cos(rad) - float64(f.position.Y)*math.Sin(rad)
	Y := float64(f.position.X)*math.Sin(rad) + float64(f.position.Y)*math.Cos(rad)

	f.position = image.Point{round(X), round(Y)}
}

func main() {
	input := fileReader.ReadLines("Day12/inputinput.txt")

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part2(inputArr []string) int {

	Waypoint := &ferry{position: image.Point{1, 10}}
	Ferry := &ferry{position: image.Point{0, 0}, angle: math.Pi / 2}
	// fmt.Println(Waypoint)
	for _, input := range inputArr {

		direc := input[0]
		amount, _ := strconv.Atoi(input[1:])

		if direc == 'N' {
			Waypoint.position = Waypoint.position.Add(image.Point{amount, 0})
		} else if direc == 'E' {
			Waypoint.position = Waypoint.position.Add(image.Point{0, amount})
		} else if direc == 'S' {
			Waypoint.position = Waypoint.position.Add(image.Point{-amount, 0})
		} else if direc == 'W' {
			Waypoint.position = Waypoint.position.Add(image.Point{0, -amount})
		} else if direc == 'F' {
			Ferry.move(amount, *Waypoint)
		} else if direc == 'R' {
			Waypoint.rotate(float64(amount))
		} else if direc == 'L' {
			Waypoint.rotate(float64(-amount))
		} else {
			// fmt.Println("Err")
		}
		// fmt.Println(Waypoint.position, Ferry.position, input)
	}

	return abs(Ferry.position.X) + abs(Ferry.position.Y)
}

func part1(inputArr []string) int {

	Ferry := &ferry{position: image.Point{0, 0}, angle: math.Pi / 2}
	for _, input := range inputArr {

		direc := input[0]
		amount, _ := strconv.Atoi(input[1:])
		amount64 := float64(amount)

		if direc == 'N' {
			Ferry.position = Ferry.position.Add(image.Point{amount, 0})
		} else if direc == 'E' {
			Ferry.position = Ferry.position.Add(image.Point{0, amount})
		} else if direc == 'S' {
			Ferry.position = Ferry.position.Add(image.Point{-amount, 0})
		} else if direc == 'W' {
			Ferry.position = Ferry.position.Add(image.Point{0, -amount})
		} else if direc == 'F' {
			X := int(amount64 * math.Cos(Ferry.angle))
			Y := int(amount64 * math.Sin(Ferry.angle))
			Ferry.position = Ferry.position.Add(image.Point{X, Y})
		} else if direc == 'R' {
			Ferry.turn(amount64)
		} else if direc == 'L' {
			Ferry.turn(-amount64)
		} else {
			fmt.Println("Err")
		}
	}

	return abs(Ferry.position.X) + abs(Ferry.position.Y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func round(val float64) int {
	if val < 0 {
		return int(val - 0.5)
	}
	return int(val + 0.5)
}
