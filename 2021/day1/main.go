package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readln() []string {
	file, err := os.Open("input.txt")
	defer file.Close()

	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)

	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	return text
}

func main() {
	data := readln()
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
