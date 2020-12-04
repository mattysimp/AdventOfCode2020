package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
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
	fmt.Println(Parts("Day4/inputexample2.txt"))
}

var count1 int
var count2 int
var eyecolour []string = []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

func Parts(input string) (int, int) {
	count1 = 0
	count2 = 0

	jobs1 := make(chan string)
	jobs2 := make(chan string)
	result1 := make(chan int)
	result2 := make(chan int)
	done := make(chan bool, 1)
	wg := new(sync.WaitGroup)

	for x := 0; x < 10; x++ {
		go processor1(jobs1, result1, wg)
		go processor2(jobs2, result2, wg)
	}

	go reciever1(result1, done)
	go reciever2(result2, done)

	readLinesAsync(input, jobs1, jobs2, wg)

	wg.Wait()
	close(result1)
	close(result2)
	<-done
	<-done
	return count1, count2
}

func reciever1(result <-chan int, done chan<- bool) {
	for Result := range result {
		count1 += Result
	}
	done <- true
}
func reciever2(result <-chan int, done chan<- bool) {
	for Result := range result {
		count2 += Result
	}
	done <- true
}

func processor1(jobs <-chan string, result chan<- int, wg *sync.WaitGroup) {
	var pp *passPort
	for strpp := range jobs {
		pp = new(passPort)
		split := strings.Split(strpp, " ")

		for _, kv := range split {
			kvsplit := strings.Split(kv, ":")
			set(pp, kvsplit[0], kvsplit[1])
		}

		if pp.byr == 0 || pp.iyr == 0 || pp.eyr == 0 || pp.hgt == "" || pp.hcl == "" || pp.ecl == "" || pp.pid == "" {
			result <- 0
		} else {
			result <- 1
		}
		wg.Done()
	}

}
func processor2(jobs <-chan string, result chan<- int, wg *sync.WaitGroup) {
	var pp *passPort
	for strpp := range jobs {
		pp = new(passPort)
		split := strings.Split(strpp, " ")

		for _, kv := range split {
			kvsplit := strings.Split(kv, ":")
			setWithVal(pp, kvsplit[0], kvsplit[1])
		}
		if pp.byr == 0 || pp.iyr == 0 || pp.eyr == 0 || pp.hgt == "" || pp.hcl == "" || pp.ecl == "" || pp.pid == "" {
			result <- 0
		} else {
			result <- 1
		}
		wg.Done()
	}

}

func readLinesAsync(path string, jobs1 chan<- string, jobs2 chan<- string, wg *sync.WaitGroup) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	defer close(jobs1)
	defer close(jobs2)

	scanner := bufio.NewScanner(file)
	var pp string
	for scanner.Scan() {
		liner := scanner.Text()

		if liner == "" {
			wg.Add(2)
			jobs1 <- pp[1:]
			jobs2 <- pp[1:]
			pp = ""
		} else {
			pp = pp + " " + liner
		}
	}
	wg.Add(2)
	jobs1 <- pp[1:]
	jobs2 <- pp[1:]
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
