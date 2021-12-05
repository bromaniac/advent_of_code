package main

import (
	"fmt"
	"github.com/bromaniac/advent_of_code/2021/util"
	"strconv"
	"strings"
)

func main() {
	data := input.Readln()
	pos := 0
	depth := 0

	for _, j := range data {
		s := strings.Fields(j)

		switch s[0] {
		case "forward":
			i, _ := strconv.Atoi(s[1])
			pos += i
		case "down":
			i, _ := strconv.Atoi(s[1])
			depth += i
		case "up":
			i, _ := strconv.Atoi(s[1])
			depth -= i
		}
	}
	fmt.Println("Pos * depth ", pos*depth)
}
