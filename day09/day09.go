package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var marker_re = regexp.MustCompile(`^\((\d+)x(\d+)\)`)

func main() {
	content, err := os.ReadFile("input09.txt")
	if err != nil {
		log.Fatal(err)
	}
	line := strings.Trim(string(content), "\n")

	decompressed_data := make([]string, 0)

	i := 0
	for {
		match := marker_re.FindStringSubmatch(line[i:])
		if match != nil {
			size, _ := strconv.Atoi(match[1])
			count, _ := strconv.Atoi(match[2])
			i += len(match[0])
			for c := 0; c < count; c++ {
				decompressed_data = append(decompressed_data, line[i:i+size])
			}
			i += size
		} else {
			// we could be much more efficient by finding next occurence of marker and append all data until then
			decompressed_data = append(decompressed_data, string(line[i]))
			i++
		}
		if i >= len(line) {
			break
		}
	}
	full_output := strings.Join(decompressed_data, "")
	log.Print("Part1: ", len(full_output))

	part2 := length([]byte(line))
	log.Print("Part2: ", part2)
}

func length(input []byte) int {
	if len(input) == 0 {
		return 0
	}
	match := marker_re.FindSubmatch(input)
	if match != nil {
		size, _ := strconv.Atoi(string(match[1]))
		count, _ := strconv.Atoi(string(match[2]))
		start := len(match[0])
		stop := start + size
		return count*length(input[start:stop]) + length(input[stop:])
	} else {
		return 1 + length(input[1:])
	}
}
