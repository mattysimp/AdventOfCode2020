package main

import (
	"AdventOfCode2020/fileReader"
	"fmt"
	"strconv"
	"strings"
)

func main() {

	fmt.Println(part1("Day8/input.txt"))
	fmt.Println(part2("Day8/input.txt"))

}

func part1(filePath string) int {
	input := fileReader.ReadLines(filePath)
	acc, _ := boot(input)
	return acc
}

func part2(filePath string) int {
	done := make(chan int, 1)

	input := fileReader.ReadLines(filePath)
	for i, line := range input {
		lineSplit := strings.Split(line, " ")
		switch lineSplit[0] {
		case "nop":
			go bootAsync(input, "nop", "jmp", i, lineSplit[1], done)
		case "jmp":

			go bootAsync(input, "jmp", "nop", i, lineSplit[1], done)
		}
	}
	return <-done
}

func bootAsync(input []string, changeFrom string, changeTo string, i int, number string, done chan<- int) {
	cInput := make([]string, len(input))
	copy(cInput, input)

	cInput[i] = strings.ReplaceAll(cInput[i], changeFrom, changeTo)

	acc, booted := boot(cInput)
	if booted {
		done <- acc
	}
}

func boot(input []string) (int, bool) {
	var accumulator int
	var i int
	var booted bool = false

	iHash := make(map[int]bool)

	for {
		i, accumulator = bootLine(input, i, accumulator)
		if i >= len(input) {
			booted = true
			break
		}
		if _, ok := iHash[i]; ok {
			break
		} else {
			iHash[i] = true
		}

	}
	return accumulator, booted
}

func bootLine(input []string, i int, accumulator int) (int, int) {
	lineSplit := strings.Split(input[i], " ")
	number, _ := strconv.Atoi(lineSplit[1])
	switch lineSplit[0] {
	case "nop":
		i++
	case "acc":
		accumulator += number
		i++
	case "jmp":
		i += number
	}
	return i, accumulator
}
