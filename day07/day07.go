package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	content, err := os.ReadFile("input07.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(content), "\n")

	part1 := 0

	for _, line := range lines {
		ip, err := parseIP(line)
		if err != nil {
			log.Print("Could not parse", line, err)
		} else {
			if ip.SupportTLS() {
				part1 += 1
			}
		}
	}
	log.Print("Part1:", part1)
}

type Part struct {
	hypernet bool
	value    string
}

type IP struct {
	parts []Part
}

func (ip IP) SupportTLS() bool {
	at_least_one_abba := false
	for _, part := range ip.parts {

		for i := 0; i < len(part.value)-3; i++ {
			if part.value[i+0] == part.value[i+3] && part.value[i+1] == part.value[i+2] && part.value[i+0] != part.value[i+1] {
				if part.hypernet {
					return false
				} else {
					at_least_one_abba = true
				}
			}
		}
	}
	return at_least_one_abba
}

// TODO(g.seux): implement error in this method
func parseIP(line string) (IP, error) {
	ip := IP{make([]Part, 0)}
	start, stop := 0, 0
	for stop < len(line) {
		if line[stop] == '[' {
			ip.parts = append(ip.parts, Part{false, line[start:stop]})
			start = stop + 1
			stop = start
		} else if line[stop] == ']' {
			ip.parts = append(ip.parts, Part{true, line[start:stop]})
			start = stop + 1
			stop = start
		} else {
			stop += 1
		}
	}
	ip.parts = append(ip.parts, Part{false, line[start:stop]})

	return ip, nil
}
