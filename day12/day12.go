package main

import (
	"fmt"
	"golang.org/x/exp/slices"
	"log"
	"os"
	"strconv"
	"strings"
)

func read_number(s string) int {
	j := 0
	for j < len(s) && s[j] != ' ' {
		j++
	}
	value, err := strconv.Atoi(s[0:j])
	if err != nil {
		log.Fatal(err)
	}
	return value
}

func read_value(current_registers []int, s string) int {
	known_registers := []byte{'a', 'b', 'c', 'd'}
	if slices.Contains(known_registers, s[0]) {
		var src_register int = int(s[0] - 97)
		return current_registers[src_register]
	} else {
		return read_number(s)
	}
}

func main() {
	content, error := os.ReadFile("input12.txt")
	if error != nil {
		log.Fatal(error)
	}

	for _, part := range [2]int{1, 2} {

		registers := make([]int, 4)
		if part == 2 {
			registers[2] = 1
		}
		lines := strings.Split(string(content), "\n")
		if lines[len(lines)-1] == "" {
			lines = lines[0 : len(lines)-1]
		}

		current_line := 0

		for current_line < len(lines) {
			line := lines[current_line]
			switch line[0:3] {
			case "cpy":
				src_value := read_value(registers, line[4:])
				dst_register := int(lines[current_line][len(line)-1] - 97)
				registers[dst_register] = src_value

			case "inc":
				dst_register := int(lines[current_line][len(line)-1] - 97)
				registers[dst_register] = registers[dst_register] + 1
			case "dec":
				dst_register := int(lines[current_line][len(line)-1] - 97)
				registers[dst_register] = registers[dst_register] - 1
			case "jnz":
				if read_value(registers, line[4:]) != 0 {
					shift := read_number(line[6:])
					current_line += shift
					continue
				}
			default:
				log.Fatal(fmt.Sprintf("Invalid instruction %s on line %d", line[0:3], current_line))
			}
			current_line++
		}
		log.Printf("Part%d: %d", part, registers[0])
	}

}
