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

func main() {
	content, err := os.ReadFile("input01.txt")
	if err != nil {
		log.Fatal(err)
	}
	// part 1
	instructions := strings.Split(strings.TrimSuffix(string(content), "\n"), ", ")

	var direction int = 0
	var x, y int = 0, 0

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
		fmt.Printf("Instruction was: rotate %s and then walk %d steps\n", string(instruction[0]), steps)
		switch direction {
		case 0:
			y -= steps
		case 2:
			y += steps
		case 1:
			x += steps
		case 3:
			x -= steps
		default:
			log.Fatal(fmt.Sprintf("Unknown direction %d, something is really really wrong", direction))
		}
	}
	manhattan := abs(x) + abs(y)
	fmt.Printf("Part1: %d\n", manhattan)
}
