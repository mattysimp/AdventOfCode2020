package main

import (
	"AdventOfCode2020/fileReader"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	input := fileReader.ReadLines("Day14/input.txt")

	fmt.Println(parts(input, true))
	fmt.Println(parts(input, false))
}

func parts(input []string, p1 bool) int64 {
	var mask string

	mem := make(map[int]int64)
	mem2 := make(map[int64]int64)

	for _, line := range input {
		split := strings.Split(line, " = ")
		if split[0] == "mask" {
			mask = split[1]
		} else {
			memsplit := strings.Split(split[0], "[")
			address, _ := strconv.Atoi(memsplit[1][:len(memsplit[1])-1])
			arg, _ := strconv.Atoi(split[1])
			bit := int64(arg)

			if p1 {
				or, _ := strconv.ParseInt((strings.ReplaceAll(mask, "X", "0")), 2, 64)
				and, _ := strconv.ParseInt((strings.ReplaceAll(mask, "X", "1")), 2, 64)
				bit |= or
				bit &= and
				mem[address] = bit

			} else {
				for _, address := range addAll(int64(address), mask) {
					mem2[address] = int64(arg)
				}
			}
		}

	}

	var sum int64
	if p1 {
		for _, value := range mem {
			sum += value
		}
	} else {
		for _, value := range mem2 {
			sum += value
		}
	}
	return sum
}

func addAll(address int64, mask string) []int64 {
	or, _ := strconv.ParseInt((strings.ReplaceAll(mask, "X", "1")), 2, 64)
	address |= or
	addresses := []int64{int64(address)}
	for i, char := range mask {
		if char == 'X' {
			for _, address := range addresses {
				addresses = append(addresses, address-int64(math.Pow(2, float64(len(mask)-i))/2))
			}
		}
	}
	return addresses
}
