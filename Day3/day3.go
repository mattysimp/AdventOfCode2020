package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func main() {

	input, MaxWidth, err := readLines("Day3/input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(part1(input, MaxWidth, 3, 1))

	fmt.Println(part2(input, MaxWidth))

}

func part2(input [][]int, MaxWidth int) int {
	MultTreesHit := 1
	Slopes := [][]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	wg := new(sync.WaitGroup)
	results := make(chan int, len(Slopes))

	for _, Slope := range Slopes {
		wg.Add(1)
		go func(Slope []int) {
			defer wg.Done()
			Res := part1(input, MaxWidth, Slope[0], Slope[1])
			results <- Res
		}(Slope)
	}

	wg.Wait()
	close(results)

	for Res := range results {
		MultTreesHit *= Res
	}

	return MultTreesHit
}

func part1(input [][]int, MaxWidth int, right int, down int) (TreesHit int) {
	j := 0

	for i := 0; i < len(input); i += down {
		TreesHit += input[i][j]
		j = (j + right) % (MaxWidth + 1)
	}
	return TreesHit
}

func readLines(path string) ([][]int, int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, 0, err
	}
	defer file.Close()

	var lines [][]int
	var MaxWidth int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()

		var line []int
		for i, symbol := range text {
			var tree int
			if string(symbol) == "#" {
				tree = 1
			} else {
				tree = 0
			}
			line = append(line, tree)
			if i > MaxWidth {
				MaxWidth = i
			}
		}
		lines = append(lines, line)
	}
	return lines, MaxWidth, scanner.Err()
}
