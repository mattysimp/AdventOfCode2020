package main

import (
	"AdventOfCode2020/fileReader"
	"fmt"
	"sort"
)

var cache map[fileReader.XMAS]int = make(map[fileReader.XMAS]int)

func main() {

	input := fileReader.ReadLinesInt("Day10/input.txt")
	sort.Ints(input)
	fullInput := append([]int{0}, input...)

	fmt.Println(part1(fullInput))
	fmt.Println(part2(fullInput))
}

func part2(input []int) (validCombinations int) {
	input = append(input, input[len(input)-1]+3)
	xmasInput := fileReader.ConvertToXmas(input)
	return combinator(xmasInput[0], xmasInput)
}

func part1(input []int) int {
	diffMap := make(map[int]int)

	for x := 1; x < len(input); x++ {
		Diff := input[x] - input[x-1]

		diffMap[Diff]++
	}
	return diffMap[1] * (diffMap[3] + 1)
}

func combinator(start fileReader.XMAS, adapters []fileReader.XMAS) (res int) {

	if start.Place == len(adapters)-1 {
		return 1
	}
	if res, ok := cache[start]; ok {
		return res
	}

	validNext := validNextAdapter(start, adapters[start.Place+1:])
	for _, adapter := range validNext {
		res += combinator(adapter, adapters)
	}
	cache[start] = res
	return res

}

func validNextAdapter(thisadapter fileReader.XMAS, adapters []fileReader.XMAS) (validAdapters []fileReader.XMAS) {
	limit := thisadapter.Num + 3
	for _, adapter := range adapters {
		if adapter.Num <= limit {
			validAdapters = append(validAdapters, adapter)
		} else {
			break
		}
	}
	return validAdapters
}
