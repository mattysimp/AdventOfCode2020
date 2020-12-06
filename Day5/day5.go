package main

import (
	"strconv"
	"sync"
	"fmt"
	"strings"
	"AdventOfCode2020/fileReader"
)

type Seat struct {
	bsp    string
	row    int
	column int
	id     int
}

var maxID int
var seatMap  = make(map[int]bool)

func (s *Seat) getID() {
	s.row = binary(s.bsp[:7])
	s.column = binary(s.bsp[7:])
	s.id = s.row*8 + s.column
}

func binary(str string) int {
	replacer := strings.NewReplacer("F", "0", "L", "0", "B", "1", "R", "1")
	str = replacer.Replace(str)
	i, _ := strconv.ParseInt(str, 2, 64)
	return int(i)
}


func main() {
	fmt.Println(parts("Day5/input.txt"))
}


func parts(inputFile string) (int, int) {

	jobs := make(chan string)
	result := make(chan *Seat)
	wg := new(sync.WaitGroup)

	for x := 0; x < 10; x++ {
		go processor(jobs, result, wg)
	}

	go reciever(result, wg)

	fileReader.ReadLinesAsync(inputFile, jobs, wg)

	wg.Wait()
	
	freeSeatID := checkSeatMap()
	return maxID, freeSeatID
}

func reciever(result <-chan *Seat, wg *sync.WaitGroup) {
	for Result := range result {
		if Result.id > maxID {
			maxID = Result.id
		}
		seatMap[Result.id] = true
		wg.Done()
	}
}
func processor(jobs <-chan string, result chan<- *Seat, wg *sync.WaitGroup) {
	for seatStr := range jobs {
		seat := &Seat{bsp: seatStr}
		seat.getID()
		result <- seat
	}
}


func checkSeatMap() (int) {
	var lastok bool
	var nextok bool
	for row := 0; row <128; row ++ {
		for col := 0; col<8; col++ {
			id := row * 8 + col
			_, ok := seatMap[id+1]
			if !nextok && ok && lastok {
				return row * 8 + col	
			}
			lastok=nextok
			nextok=ok
		}
	}
	return 0
}
