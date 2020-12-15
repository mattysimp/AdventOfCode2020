package main

import (
	"fmt"
)

func main() {
	input := []int{16, 12, 1, 0, 15, 7, 11}

	fmt.Println(parts(input, 2020))
	fmt.Println(parts(input, 30000000))

}

func parts(input []int, target int) int {

	lastSeen := [30000000]int{}
	var nextNum int

	for i := 1; i < target; i++ {
		if i <= len(input) {
			lastSeen[input[i-1]] = i
			nextNum = 0
		} else if last := lastSeen[nextNum]; last != 0 {
			lastSeen[nextNum] = i
			nextNum = i - last
		} else {
			lastSeen[nextNum] = i
			nextNum = 0
		}
	}
	return nextNum
}
