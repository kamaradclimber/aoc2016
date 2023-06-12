package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var rect_r = regexp.MustCompile(`rect (\d+)x(\d+)`)
var rotate_row_r = regexp.MustCompile(`rotate row y=(\d+) by (\d+)`)
var rotate_column_r = regexp.MustCompile(`rotate column x=(\d+) by (\d+)`)

var width = 50
var height = 6

func main() {
	content, err := os.ReadFile("input08.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(content), "\n")

	grid := make([][]bool, height)
	for i := 0; i < height; i++ {
		grid[i] = make([]bool, width)
	}

	for _, line := range lines {
		switch {
		case line == "\n":
			continue
		case rect_r.MatchString(line):
			res := rect_r.FindStringSubmatch(line)
			wide, _ := strconv.Atoi(res[1]) // assume success
			tall, _ := strconv.Atoi(res[2]) // assume success
			for line := 0; line < tall; line++ {
				for column := 0; column < wide; column++ {
					grid[line][column] = true
				}
			}
		case rotate_row_r.MatchString(line):
			res := rotate_row_r.FindStringSubmatch(line)
			line, _ := strconv.Atoi(res[1])  // assume success
			shift, _ := strconv.Atoi(res[2]) // assume success
			old_line := make([]bool, width)
			for c := 0; c < width; c++ {
				old_line[c] = grid[line][c]
			}
			for c := 0; c < width; c++ {
				grid[line][c] = old_line[(c-shift+2*width)%width]
			}
		case rotate_column_r.MatchString(line):
			res := rotate_column_r.FindStringSubmatch(line)
			column, _ := strconv.Atoi(res[1]) // assume success
			shift, _ := strconv.Atoi(res[2])  // assume success
			old_column := make([]bool, height)
			for l := 0; l < height; l++ {
				old_column[l] = grid[l][column]
			}
			for l := 0; l < height; l++ {
				grid[l][column] = old_column[(l-shift+2*height)%height]
			}
		}
		log.Print(line)
		printGrid(grid)

		part1 := 0
		for l := 0; l < len(grid); l++ {
			for c := 0; c < len(grid[l]); c++ {
				if grid[l][c] {
					part1 += 1
				}
			}
		}
		log.Print("Part1: ", part1)

	}
}

func printGrid(grid [][]bool) {
	for l := 0; l < len(grid); l++ {
		for c := 0; c < len(grid[l]); c++ {
			if grid[l][c] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
	log.Print("---------------------------")
}
