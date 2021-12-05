package main

import (
	"fmt"
	"github.com/bromaniac/advent_of_code/2021/util"
	"strconv"
	"strings"
)

func main() {
	data := input.Readln()
	gamma := ""
	epsilon := ""
	ones := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for _, j := range data {
		s := strings.Split(j, "")
		for i, j := range s {
			bit, _ := strconv.Atoi(j)
			if bit == 1 {
				ones[i] += 1
			}
		}
	}
	for _, j := range ones {
		if j > 500 {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}
	g, _ := strconv.ParseInt(gamma, 2, 64)
	e, _ := strconv.ParseInt(epsilon, 2, 64)
	fmt.Println("gamma * epsilon = ", g*e)
}
