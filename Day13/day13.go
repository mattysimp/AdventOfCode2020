package main

import (
	"AdventOfCode2020/fileReader"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := fileReader.ReadLines("Day13/input.txt")

	fmt.Println(part1(input))
	fmt.Println(part2(input[1]))
}

func part2(input string) int {
	Buses := strings.Split(input, ",")
	part2, step := 0, 1

	for i, Bus := range Buses {
		if Bus != "x" {
			intBus, _ := strconv.Atoi(Bus)
			for (part2+i)%intBus != 0 {
				part2 += step
			}
			step *= intBus
		}
	}
	return part2
}

func part1(input []string) int {
	StartTime, _ := strconv.Atoi(input[0])
	Buses := strings.Split(input[1],",")
	var MinWait int
	var Result int

	for _, Bus := range Buses {
		if Bus != "x" {
			intBus, _ := strconv.Atoi(Bus)
			WaitTime := intBus - (StartTime % intBus)
			if WaitTime < MinWait || MinWait == 0 {
				MinWait = WaitTime
				Result = WaitTime * intBus
			}
		}
	}

	return Result
}

func checkResult(t int, Buses []string) bool {
	var BusTime []int

	
	for i, Bus := range Buses {
		if Bus != "x" && i != 0 {
			intBus, _ := strconv.Atoi(Bus)
			WaitTime := intBus - (t % intBus)
			if WaitTime != i {
				return false
			}
			BusTime = append(BusTime, WaitTime)
		}
	}
	return true

}