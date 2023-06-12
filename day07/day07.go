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
	part2 := 0

	for _, line := range lines {
		ip, err := parseIP(line)
		if err != nil {
			log.Print("Could not parse", line, err)
		} else {
			if ip.SupportTLS() {
				part1 += 1
			}
			if ip.SupportSSL() {
				part2 += 1
			}
		}
	}
	log.Print("Part1:", part1)
	log.Print("Part2:", part2)
}

type Part struct {
	hypernet bool
	value    string
}

type IP struct {
	parts []Part
}

func (ip IP) SupportSSL() bool {
	// list all ABA sequences in supernet sequences
	abas := make([]string, 0)
	for _, part := range ip.parts {
		if part.hypernet {
			continue
		}

		for i := 0; i < len(part.value)-2; i++ {
			if part.value[i+0] == part.value[i+2] && part.value[i+0] != part.value[i+1] {
				abas = append(abas, part.value[i:i+3])
			}
		}
	}

	// search them in hypernet parts
	for _, part := range ip.parts {
		if part.hypernet {
			for _, aba := range abas {
				bab := []byte{aba[1], aba[0], aba[1]}
				if strings.Contains(part.value, string(bab)) {
					return true
				}
			}
		}
	}

	return false
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
