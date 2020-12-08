package fileReader

import (
	"bufio"
	"os"
	"strconv"
	"sync"
)

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
