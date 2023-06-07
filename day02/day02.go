package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	content, err := os.ReadFile("input02.txt")
	if err != nil {
		log.Fatal(err)
	}
	var lines []string = strings.Split(string(content), "\n")

	starting_key := 5
	key_line, key_column := (starting_key-1)/3, (starting_key-1)%3
	keys := make([]int, 0)

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		for _, char := range line {
			switch char {
			case 'L':
				key_column = max(key_column-1, 0)
			case 'R':
				key_column = min(key_column+1, 2)
			case 'U':
				key_line = max(key_line-1, 0)
			case 'D':
				key_line = min(key_line+1, 2)
			}
		}
		key := key_line*3 + key_column + 1
		keys = append(keys, key)
	}
	log.Print("Part 1: ", keys)

}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}
