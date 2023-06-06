package main

import "os"
import "log"
import "strings"
import "fmt"
import "strconv"

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

type Pos struct {
	x int
	y int
}

func checkDuplicate(positions map[Pos]bool, pos Pos) bool {
	if positions[pos] {
		return true
	}
	positions[pos] = true
	return false
}

func main() {
	content, err := os.ReadFile("input01.txt")
	if err != nil {
		log.Fatal(err)
	}
	// part 1
	instructions := strings.Split(strings.TrimSuffix(string(content), "\n"), ", ")

	var direction int = 0
	var x, y int = 0, 0

	var positions = make(map[Pos]bool)
	positions[Pos{x, y}] = true

	var part2_found = false

	for _, instruction := range instructions {
		switch instruction[0] {
		case 'R':
			direction += 1
		case 'L':
			direction -= 1
		default:
			log.Fatal("Unknown direction: " + string(instruction[0]))
		}
		direction = (direction + 4) % 4
		steps, err := strconv.Atoi(instruction[1:])
		if err != nil {
			log.Fatal(fmt.Sprintf("Impossible to parse %s as a int", instruction[1:]))
		}
		// fmt.Printf("Instruction was: rotate %s and then walk %d steps\n", string(instruction[0]), steps)
		switch direction {
		case 0:
			for i := 0; i < steps; i++ {
				y -= 1
				if checkDuplicate(positions, Pos{x, y}) && !part2_found {
					part2_found = true
					fmt.Printf("Part2: %d\n", abs(x)+abs(y))
				}
			}
		case 2:
			for i := 0; i < steps; i++ {
				y += 1
				if checkDuplicate(positions, Pos{x, y}) && !part2_found {
					part2_found = true
					fmt.Printf("Part2: %d\n", abs(x)+abs(y))
				}
			}
		case 1:
			for i := 0; i < steps; i++ {
				x += 1
				if checkDuplicate(positions, Pos{x, y}) && !part2_found {
					part2_found = true
					fmt.Printf("Part2: %d\n", abs(x)+abs(y))
				}
			}
		case 3:
			for i := 0; i < steps; i++ {
				x -= 1
				if checkDuplicate(positions, Pos{x, y}) && !part2_found {
					part2_found = true
					fmt.Printf("Part2: %d\n", abs(x)+abs(y))
				}
			}
		default:
			log.Fatal(fmt.Sprintf("Unknown direction %d, something is really really wrong", direction))
		}
	}
	manhattan := abs(x) + abs(y)
	fmt.Printf("Part1: %d\n", manhattan)
}
