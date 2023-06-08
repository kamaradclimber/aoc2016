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
		triangle, err := parseTriangle(line)
		if err != nil {
			log.Fatal(err)
		}
		if triangle.valid() {
			valid_triangle_count += 1
		}
	}
	log.Print(fmt.Sprintf("Part1: %d", valid_triangle_count))
}

func parseTriangle(line string) (Triangle, error) {
	fields := strings.Fields(line)
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
