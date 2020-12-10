package fileReader

import (
	"bufio"
	"os"
	"strconv"
	"sync"
)

type XMAS struct {
	Num   int
	Place int
}

func ConvertToXmas(lines []int) []XMAS {
	var output []XMAS
	for i, val := range lines {
		lineXMAS := XMAS{Num: val, Place: i}
		output = append(output, lineXMAS)
	}
	return output
}

func ReadLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var output []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		liner := scanner.Text()
		output = append(output, liner)
	}
	return output
}

func ReadLinesInt(path string) []int {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var output []int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		liner := scanner.Text()
		lineInt, _ := strconv.Atoi(liner)
		output = append(output, lineInt)
	}
	return output
}

func ReadLinesXMAS(path string) []XMAS {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var output []int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		liner := scanner.Text()
		lineInt, _ := strconv.Atoi(liner)
		output = append(output, lineInt)
	}
	xmasOuput := ConvertToXmas(output)
	return xmasOuput
}
func ReadLinesAsync(path string, jobs chan<- string, wg *sync.WaitGroup) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	defer close(jobs)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		liner := scanner.Text()
		wg.Add(1)
		jobs <- liner

	}
}

func ReadLinesGroupsAsync(path string, jobs chan<- string, wg *sync.WaitGroup) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	defer close(jobs)

	scanner := bufio.NewScanner(file)
	var pp string
	var lineCount int
	for scanner.Scan() {
		liner := scanner.Text()
		if liner == "" {
			wg.Add(1)
			jobs <- pp + "@" + strconv.Itoa(lineCount)
			pp = ""
			lineCount = 0
		} else {
			pp = pp + " " + liner
			lineCount++
		}

	}
	wg.Add(1)
	jobs <- pp + "@" + strconv.Itoa(lineCount)
}
