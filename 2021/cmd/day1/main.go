package main

import (
	"fmt"
	"github.com/bromaniac/advent_of_code/2021/util"
	"strconv"
)

func main() {
	data := input.Readln()
	counter := 0

	for i, j := range data {
		if i == len(data)-1 {
			break
		}

		c, _ := strconv.Atoi(j)         // current
		n, _ := strconv.Atoi(data[i+1]) //next

		if c < n {
			counter++
		}
	}
	fmt.Println("Num of increases: ", counter)
}
