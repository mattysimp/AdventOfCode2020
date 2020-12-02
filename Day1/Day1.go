package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {

	Input, err := readLines("Day1/input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(day1(Input))
	fmt.Println(day2(Input))

}

func day2(input []int) int {
	sort.Ints(input)

	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			if i != j && input[i]+input[j] < 2020 {
				kBot := i
				kTop := j
				for {
					k := (kBot + kTop) / 2
					if k != kBot && k != kTop {
						if input[i]+input[j]+input[k] == 2020 {
							return input[i] * input[j] * input[k]
						} else if input[i]+input[j]+input[k] > 2020 {
							kTop = k
							k = (k + kBot) / 2
						} else {
							kBot = k
							k = (k + kTop) / 2
						}
					} else {
						break
					}
				}
			}
		}
	}
	return 0
}

func day1(input []int) int {
	intMap := make(map[int]bool)

	for _, i := range input {
		lookup := 2020 - i

		if _, ok := intMap[lookup]; ok {
			return i * lookup
		}
		intMap[i] = true
	}
	return 0

}

func readLines(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, line)
	}
	return lines, scanner.Err()
}
