package main

import (
	"github.com/thoas/go-funk"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	content, err := os.ReadFile("input06.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(content), "\n")

	freqs := make([]map[rune]int, len(string(lines[0])))
	for i := 0; i < len(freqs); i++ {
		freqs[i] = make(map[rune]int)
	}

	// count frequencies
	for _, line := range lines {
		for col, char := range line {
			freqs[col][char] += 1
		}
	}

	// find most common letter per column
	part1 := make([]rune, 0)
	for i := 0; i < len(freqs); i++ {
		keys := funk.Keys(freqs[i]).([]rune)
		sort.Slice(keys, func(a, b int) bool {
			return freqs[i][keys[a]] > freqs[i][keys[b]]
		})
		part1 = append(part1, keys[0])
	}
	log.Print("Part1: ", string(part1))
}
