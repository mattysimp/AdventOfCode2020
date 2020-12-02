package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

var part1 int
var part2 int

func main() {

	fmt.Println(Parts("Day2/input.txt"))

}

func Parts(input string) (int, int) {
	part1 = 0
	part2 = 0
	jobs := make(chan []string)
	result := make(chan int)
	wg := new(sync.WaitGroup)
	readwg := new(sync.WaitGroup)

	for x := 0; x < 10; x++ {
		go processor(jobs, result, wg)
	}

	go reciever(result, wg)

	readLinesAsync(input, jobs, wg, readwg)

	return part1, part2
}

func reciever(result <-chan int, wg *sync.WaitGroup) {
	for Result := range result {
		if Result == 1 {
			part1++
		} else {
			part2++
		}
		wg.Done()
	}

}

func processor(jobs <-chan []string, result chan<- int, wg *sync.WaitGroup) {
	for line := range jobs {
		day2(line, result, wg)
	}

}

func day2(line []string, result chan<- int, wg *sync.WaitGroup) {
	Min, _ := strconv.Atoi(line[0])
	Max, _ := strconv.Atoi(line[1])
	Letter := line[2]
	Password := line[3]
	Amount := strings.Count(Password, Letter)
	if Min <= Amount && Amount <= Max {
		result <- 1
	} else {
		defer wg.Done()
	}
	if (string(Password[Min-1]) == Letter || string(Password[Max-1]) == Letter) && !(string(Password[Min-1]) == Letter && string(Password[Max-1]) == Letter) {
		result <- 2
	} else {
		defer wg.Done()
	}

}

func readLinesAsync(path string, jobs chan<- []string, wg *sync.WaitGroup, readwg *sync.WaitGroup) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	defer close(jobs)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		wg.Add(2)
		readwg.Add(1)
		liner := scanner.Text()
		go func() {
			defer readwg.Done()
			split := strings.Split(liner, " ")
			minmax := strings.Split(split[0], "-")
			line := []string{minmax[0], minmax[1], split[1][:1], split[2]}
			jobs <- line
		}()

	}
	readwg.Wait()
}
