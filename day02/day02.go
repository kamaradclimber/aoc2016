package main

import (
	//	"fmt"
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

	keys := solve(1, lines)
	log.Print("Part 1: ", keys)
	keys = solve(2, lines)
	log.Print("Part 2: ", keys)
}

func solve(part int, lines []string) []int {
	starting_key := 5

	var toCoords func(int) (int, int)
	var toKey func(int, int) int
	var move func(int, int, rune) (int, int)

	if part == 1 {
		toCoords = to_Coords_part1
		toKey = toKey_part1
		move = func(line, column int, direction rune) (int, int) {
			switch direction {
			case 'L':
				return line, max(column-1, 0)
			case 'R':
				return line, min(column+1, 2)
			case 'U':
				return max(line-1, 0), column
			default:
				return min(line+1, 2), column
			}
		}

	} else {
		toCoords = toCoords_part2
		toKey = toKey_part2
		move = func(line, column int, direction rune) (int, int) {
			// really not proud of this! I wish there was a coordinate system that would be
			// more suited to this!
			var all_moves = map[int](map[rune]int){
				1:  {'L': 1, 'R': 1, 'U': 1, 'D': 3},
				2:  {'L': 2, 'R': 3, 'U': 2, 'D': 6},
				3:  {'L': 2, 'R': 4, 'U': 1, 'D': 7},
				4:  {'L': 3, 'R': 4, 'U': 4, 'D': 8},
				5:  {'L': 5, 'R': 6, 'U': 5, 'D': 5},
				6:  {'L': 5, 'R': 7, 'U': 2, 'D': 10},
				7:  {'L': 6, 'R': 8, 'U': 3, 'D': 11},
				8:  {'L': 7, 'R': 9, 'U': 4, 'D': 12},
				9:  {'L': 8, 'R': 9, 'U': 9, 'D': 9},
				10: {'L': 10, 'R': 11, 'U': 6, 'D': 10},
				11: {'L': 10, 'R': 12, 'U': 7, 'D': 13},
				12: {'L': 11, 'R': 12, 'U': 8, 'D': 12},
				13: {'L': 13, 'R': 13, 'U': 11, 'D': 13},
			}
			new_key := all_moves[toKey_part2(line, column)][direction]
			return toCoords_part2(new_key)
		}
	}

	key_line, key_column := toCoords(starting_key)
	keys := make([]int, 0)

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		for _, char := range line {
			key_line, key_column = move(key_line, key_column, char)
			// log.Print(fmt.Sprintf("Dir %s, Coords (%d, %d), key %d", string(char), key_line, key_column, toKey(key_line, key_column)))
		}
		key := toKey(key_line, key_column)
		keys = append(keys, key)
	}
	return keys

}

func to_Coords_part1(n int) (int, int) {
	return (n - 1) / 3, (n - 1) % 3
}

func toKey_part1(line, column int) int {
	return line*3 + column + 1
}

func toCoords_part2(n int) (int, int) {
	if n == 1 {
		return 0, 0
	} else if n < 5 {
		return 1, n - 2
	} else if n < 10 {
		return 2, n - 5
	} else if n < 13 {
		return 3, n - 10
	} else {
		return 4, 0
	}
}

func toKey_part2(line, column int) int {
	starts := []int{1, 2, 5, 10, 13}
	return starts[line] + column
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
