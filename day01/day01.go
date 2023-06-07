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

func manhattan(p Pos) int {
	return abs(p.x) + abs(p.y)
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
	var cur_pos = Pos{0, 0}

	var positions = make(map[Pos]bool)
	positions[cur_pos] = true

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
		var mover func(Pos) Pos
		switch direction {
		case 0:
			mover = func(p Pos) Pos {
				return Pos{p.x, p.y - 1}
			}
		case 2:
			mover = func(p Pos) Pos {
				return Pos{p.x, p.y + 1}
			}
		case 1:
			mover = func(p Pos) Pos {
				return Pos{p.x + 1, p.y}
			}
		case 3:
			mover = func(p Pos) Pos {
				return Pos{p.x - 1, p.y}
			}
		default:
			log.Fatal(fmt.Sprintf("Unknown direction %d, something is really really wrong", direction))
		}
		for i := 0; i < steps; i++ {
			cur_pos = mover(cur_pos)
			if checkDuplicate(positions, cur_pos) && !part2_found {
				part2_found = true
				fmt.Printf("Part2: %d\n", manhattan(cur_pos))
			}
		}
	}
	fmt.Printf("Part1: %d\n", manhattan(cur_pos))
}
