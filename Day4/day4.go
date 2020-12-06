package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"AdventOfCode2020/fileReader"
)

type passPort struct {
	byr int
	iyr int
	eyr int
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

func main() {
	fmt.Println(Parts("Day4/input.txt"))
}

var count1 int
var count2 int
var eyecolour []string = []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

func Parts(input string) (int, int) {
	count1 = 0
	count2 = 0

	jobs := make(chan string)
	result := make(chan int)
	done := make(chan bool, 1)
	wg := new(sync.WaitGroup)

	for x := 0; x < 10; x++ {
		go processor1(jobs, result, wg)
	}

	go reciever(result, done)

	fileReader.ReadLinesGroupsAsync(input, jobs, wg)

	wg.Wait()
	close(result)
	<-done
	return count1, count2
}

func reciever(result <-chan int, done chan<- bool) {
	for Result := range result {
		if Result == 1 {
			count1++
		} else {
			count2++
		}
	}
	done <- true
}

func processor1(jobs <-chan string, result chan<- int, wg *sync.WaitGroup) {
	// var pp *passPort
	for strpp := range jobs {
		pp1 := new(passPort)
		pp2 := new(passPort)
		lineSplit := strings.Split(strpp[1:], "@")
		split := strings.Split(lineSplit[0], " ")

		for _, kv := range split {
			kvsplit := strings.Split(kv, ":")
			
			set(pp1, kvsplit[0], kvsplit[1])
			setWithVal(pp2, kvsplit[0], kvsplit[1])
		}

		if pp1.byr != 0 && pp1.iyr != 0 && pp1.eyr != 0 && pp1.hgt != "" && pp1.hcl != "" && pp1.ecl != "" && pp1.pid != "" {
			result <- 1
		}
		if pp2.byr != 0 && pp2.iyr != 0 && pp2.eyr != 0 && pp2.hgt != "" && pp2.hcl != "" && pp2.ecl != "" && pp2.pid != "" {
			result <- 2
		}
		wg.Done()
	}

}

func set(pp *passPort, key string, val string) {
	if key == "byr" {
		intval, _ := strconv.Atoi(val)
		pp.byr = intval
	} else if key == "iyr" {
		intval, _ := strconv.Atoi(val)
		pp.iyr = intval
	} else if key == "eyr" {
		intval, _ := strconv.Atoi(val)
		pp.eyr = intval
	} else if key == "hgt" {
		pp.hgt = val
	} else if key == "hcl" {
		pp.hcl = val
	} else if key == "ecl" {
		pp.ecl = val
	} else if key == "pid" {
		pp.pid = val
	} else if key == "cid" {
		pp.cid = val
	}
}
func setWithVal(pp *passPort, key string, val string) {
	if key == "byr" {
		if ok, intval := year(val, 1920, 2002); ok {
			pp.byr = intval
		}
	} else if key == "iyr" {
		if ok, intval := year(val, 2010, 2020); ok {
			pp.iyr = intval
		}
	} else if key == "eyr" {
		if ok, intval := year(val, 2020, 2030); ok {
			pp.eyr = intval
		}
	} else if key == "hgt" {
		intval, _ := strconv.Atoi(val[:len(val)-2])
		if (val[len(val)-2:] == "cm" && intval >= 150 && intval <= 193) || (val[len(val)-2:] == "in" && intval >= 59 && intval <= 76) {
			pp.hgt = val
		}
	} else if key == "hcl" {
		if ok, _ := regexp.MatchString("#[0-9a-f]{6}", val); ok {
			pp.hcl = val
		}
	} else if key == "ecl" {
		if Find(eyecolour, val) {
			pp.ecl = val
		}
	} else if key == "pid" {
		if ok, _ := regexp.MatchString(`^\d{9}$`, val); ok {
			pp.pid = val
		}
	} else if key == "cid" {
		pp.cid = val
	}
}

func Find(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func year(val string, low int, high int) (bool, int) {
	intval, _ := strconv.Atoi(val)
	return intval >= low && intval <= high, intval
}
