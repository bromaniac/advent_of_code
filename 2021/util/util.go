package input

import (
	"bufio"
	"log"
	"os"
)

func Readln() []string {
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
