package main

import (
	"AdventOfCode2020/fileReader"
	"fmt"
	"sync"
)

func main() {
	input := fileReader.ReadLinesXMAS("Day9/input.txt")
	Part1 := part1(input, 25)
	fmt.Println(Part1)
	fmt.Println(part2(input, Part1))

}

func part2(input []fileReader.XMAS, invalidNo fileReader.XMAS) int {

	jobs := make(chan []fileReader.XMAS)
	results := make(chan int, 1)

	for x := 0; x < 10; x++ {
		go processor2(invalidNo, jobs, results)
	}
	go func() {
		for x := 0; x < invalidNo.Place; x++ {
			jobs <- input[x:invalidNo.Place]
		}
	}()

	return <-results

}

func processor2(invalidNo fileReader.XMAS, jobs <-chan []fileReader.XMAS, result chan<- int) {
	for jobs := range jobs {
		var summer int
		var smallest int
		var largest int
		for i, job := range jobs {
			if job.Num > largest {
				largest = job.Num
			}
			if job.Num < smallest || smallest == 0 {
				smallest = job.Num
			}
			summer += job.Num
			if summer == invalidNo.Num && i >= 2 {
				result <- smallest + largest
			}
		}
	}
}

func part1(input []fileReader.XMAS, preamble int) fileReader.XMAS {

	wg := new(sync.WaitGroup)
	jobs := make(chan []fileReader.XMAS)
	results := make(chan fileReader.XMAS)

	noOfChecks := len(input) - preamble
	wg.Add(noOfChecks)
	for x := 0; x < 10; x++ {
		go processor(jobs, results, wg)
	}

	var firstNum *fileReader.XMAS = &fileReader.XMAS{}

	go reciever(firstNum, results, wg)

	for x := 0; x < noOfChecks; x++ {
		jobs <- input[x : x+preamble+1]
	}

	wg.Wait()
	return *firstNum
}

func processor(jobs <-chan []fileReader.XMAS, results chan<- fileReader.XMAS, wg *sync.WaitGroup) {
	for jobs := range jobs {
		var valid bool = false
		numHash := make(map[int]bool)

		sumTo := jobs[len(jobs)-1]
		for _, num := range jobs {
			if _, ok := numHash[sumTo.Num-num.Num]; ok {
				valid = true
				break
			} else {
				numHash[num.Num] = true
			}
		}
		if !valid {
			wg.Add(1)
			results <- sumTo
		}
		wg.Done()
	}
}

func reciever(firstNum *fileReader.XMAS, result <-chan fileReader.XMAS, wg *sync.WaitGroup) {
	for Result := range result {
		if Result.Num < firstNum.Num || firstNum.Num == 0 {
			*firstNum = Result
		}
		wg.Done()
	}
}
