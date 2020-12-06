package main

import (
	"sync"
	"fmt"
	"strings"
	"strconv"
	"unicode"
	"AdventOfCode2020/fileReader"
)

var part1Total int
var part2Total int

func main() {
	fmt.Println(parts("Day6/inputexample.txt"))
}

func parts(inputFile string) (int, int) {
	part1Total=0
	part2Total=0

	jobs := make(chan string)
	result := make(chan [2]int)
	wg := new(sync.WaitGroup)

	for x := 0; x < 10; x++ {
		go processor(jobs, result, wg)
	}

	go reciever(result, wg)

	fileReader.ReadLinesGroupsAsync(inputFile, jobs, wg)

	wg.Wait()
	
	return part1Total, part2Total
}

func reciever(result <-chan [2]int, wg *sync.WaitGroup) {
	for Result := range result {
		part1Total += Result[0]
		part2Total += Result[1]
		wg.Done()
	}
}
func processor(jobs <-chan string, result chan<- [2]int, wg *sync.WaitGroup) {
	for answers := range jobs {

		answerSplit := strings.Split(answers, "@")
		part1, part2 := countParts(answerSplit[0], answerSplit[1])
		
		result <- [2]int{part1, part2}
	}
}


func countParts(s string, total string) (uniqueCount int, allCount int) {
	totalInt, _ := strconv.Atoi(total)
    a  := make(map[rune]int)
    for _, char := range s {
		if !unicode.IsSpace(char) {
			val, ok := a[char]
			a[char] ++
			if val + 1 == totalInt {
				allCount ++
			}
			if !ok {
				uniqueCount ++
			}
		}
    }
    return uniqueCount, allCount
}

