package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Triangle struct {
	a int
	b int
	c int
}

func main() {
	content, err := os.ReadFile("input03.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(content), "\n")

	valid_triangle_count := 0

	for _, line := range lines {
		if len(line) == 0 {
			break
		}
		fields := strings.Fields(line)
		triangle, err := parseTriangle(fields)
		if err != nil {
			log.Fatal(err)
		}
		if triangle.valid() {
			valid_triangle_count += 1
		}
	}
	log.Print(fmt.Sprintf("Part1: %d", valid_triangle_count))

	// part 2
	valid_triangle_count = 0
	// we assume number of lines is a perfect multiple of 3
	for i := 0; i < len(lines)/3; i++ {
		fields0 := strings.Fields(lines[3*i+0])
		fields1 := strings.Fields(lines[3*i+1])
		fields2 := strings.Fields(lines[3*i+2])
		// we could transpose but golang does not have a transposition method built-in 🤦
		triangle0, err := parseTriangle([]string{fields0[0], fields1[0], fields2[0]})
		triangle1, err := parseTriangle([]string{fields0[1], fields1[1], fields2[1]})
		triangle2, err := parseTriangle([]string{fields0[2], fields1[2], fields2[2]})
		if err != nil {
			log.Fatal(err)
		}
		if triangle0.valid() {
			valid_triangle_count += 1
		}
		if triangle1.valid() {
			valid_triangle_count += 1
		}
		if triangle2.valid() {
			valid_triangle_count += 1
		}
	}
	log.Print(fmt.Sprintf("Part2: %d", valid_triangle_count))
}

func parseTriangle(fields []string) (Triangle, error) {
	if len(fields) != 3 {
		return Triangle{}, errors.New("Incorrect number of side for this triangle")
	}
	a, erra := strconv.Atoi(fields[0])
	if erra != nil {
		return Triangle{}, erra
	}
	b, errb := strconv.Atoi(fields[1])
	if errb != nil {
		return Triangle{}, errb
	}
	c, errc := strconv.Atoi(fields[2])
	if errc != nil {
		return Triangle{}, errc
	}
	return Triangle{a, b, c}, nil
}

func (tri Triangle) valid() bool {
	return tri.a+tri.b > tri.c && tri.a+tri.c > tri.b && tri.b+tri.c > tri.a
}
