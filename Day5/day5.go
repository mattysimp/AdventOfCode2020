package main

import (
	"bufio"
	"strconv"
	"os"
	"sync"
	"fmt"
	"strings"
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
	str = strings.ReplaceAll(str, "F", "0")
	str = strings.ReplaceAll(str, "L", "0")
	str = strings.ReplaceAll(str, "B", "1")
	str = strings.ReplaceAll(str, "R", "1")
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

	readLinesAsync(inputFile, jobs, wg)

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

func readLinesAsync(path string, jobs chan<- string, wg *sync.WaitGroup) {
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
