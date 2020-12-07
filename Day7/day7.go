package main

import (
	"fmt"
	"regexp"
	"strconv"
	"sync"

	"AdventOfCode2020/fileReader"
)

type bag struct {
	colour            string
	innerBags         []*heldBag
	possibleOuterBags []*bag
}

type heldBag struct {
	count int
	bag   *bag
}

var bagMap map[string]*bag
var recieved map[string]bool
var recievedCount int
var part2Count int

func main() {
	fmt.Println(parts("Day7/input.txt", "shiny gold"))
}

func parts(inputFile string, startBag string) (int, int) {
	bagMap = make(map[string]*bag)
	recieved = make(map[string]bool)
	recievedCount = 0
	part2Count = 0

	jobs := make(chan string)
	result1 := make(chan *bag)
	result2 := make(chan int)
	wg := new(sync.WaitGroup)

	go processor(jobs, wg)

	fileReader.ReadLinesAsync(inputFile, jobs, wg)

	wg.Wait()

	go reciever(result1, wg)
	go reciever2(result2, wg)

	StartingBag := bagMap[startBag]
	wg.Add(len(StartingBag.possibleOuterBags))
	loopUp(StartingBag, result1, wg)

	wg.Add(len(StartingBag.innerBags))
	loopDown(StartingBag, 1, result2, wg)

	wg.Wait()

	return recievedCount, part2Count
}
func loopDown(Bag *bag, amount int, result chan<- int, wg *sync.WaitGroup) {
	for _, downBag := range Bag.innerBags {
		count := downBag.count * amount

		wg.Add(len(downBag.bag.innerBags))
		go loopDown(downBag.bag, count, result, wg)

		result <- count
	}
}
func reciever2(result <-chan int, wg *sync.WaitGroup) {
	for Result := range result {
		part2Count += Result
		wg.Done()
	}
}
func loopUp(Bag *bag, result chan<- *bag, wg *sync.WaitGroup) {
	for _, upBag := range Bag.possibleOuterBags {
		wg.Add(len(upBag.possibleOuterBags))
		go loopUp(upBag, result, wg)
		result <- upBag
	}
}

func reciever(result <-chan *bag, wg *sync.WaitGroup) {
	for Result := range result {
		if _, ok := recieved[Result.colour]; !ok {
			recievedCount++
			recieved[Result.colour] = true
		}
		wg.Done()
	}
}
func processor(jobs <-chan string, wg *sync.WaitGroup) {
	for jobs := range jobs {
		split := splitLine(jobs)
		numberOfSubBags := (len(split) - 4) / 4
		colour := split[0] + " " + split[1]
		thisBag, ok := bagMap[colour]
		if !ok {
			thisBag = &bag{colour: colour}
			bagMap[colour] = thisBag
		}

		for x := 0; x < numberOfSubBags; x++ {
			number, _ := strconv.Atoi(split[x*4+4])
			colour := split[x*4+5] + " " + split[x*4+6]

			innerBag, ok := bagMap[colour]
			if !ok {
				innerBag = &bag{colour: colour}
				bagMap[colour] = innerBag
			}
			thisBag.innerBags = append(thisBag.innerBags, &heldBag{bag: innerBag, count: number})
			// thisBag.innerBagsCount = append(thisBag.innerBagsCount, number)
			innerBag.possibleOuterBags = append(innerBag.possibleOuterBags, thisBag)
		}

		wg.Done()
	}
}

func splitLine(Line string) []string {
	array := regexp.MustCompile("[\\ \\, \\s]+").Split(Line, -1)
	return array
}
