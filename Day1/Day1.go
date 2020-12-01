package main

import (
	"bufio"
	"errors"
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
		for j := 1; j < len(input); j++ {
			if i != j && input[i]+input[j] < 2020 {
				for k := 2; k < len(input); k++ {
					if k != i && k != j {
						if input[i]+input[j]+input[k] == 2020 {
							return input[i] * input[j] * input[k]
						} else if input[i]+input[j]+input[k] > 2020 {
							break
						}
					}
				}
			}
		}
	}
	return 0
}

func day1(input []int) int {
	sort.Ints(input)

	i := 0
	j := len(input) - 1

	for {
		if input[i]+input[j] == 2020 {
			return input[i] * input[j]
		} else if input[i]+input[j] < 2020 {
			i++
		} else {
			j--
		}

		if i == j {
			panic(errors.New("not found"))
		}
	}

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
