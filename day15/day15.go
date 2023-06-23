package main

import (
	"log"
	"math/big"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Disc struct {
	nb_positions   int
	position_at_t0 int
}

func (d Disc) position(t int) int {
	return (d.position_at_t0 + t) % d.nb_positions
}

func main() {
	content, err := os.ReadFile("input15.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(content), "\n")

	discs := make([]Disc, len(lines)-1)

	re := regexp.MustCompile(`Disc #.+ has (\d+) positions; at time=0, it is at position (\d+).`)

	for i := 0; i < len(discs); i++ {
		if match := re.FindStringSubmatch(lines[i]); match != nil {
			nb_pos, err := strconv.Atoi(match[1])
			if err != nil {
				log.Fatal(err)
			}
			initial_pos, err := strconv.Atoi(match[2])
			if err != nil {
				log.Fatal(err)
			}
			discs[i] = Disc{nb_pos, initial_pos}
		}
	}
	log.Printf("We have %d discs", len(discs))

	// this looks like chinese reminder theorem

	n_hats := make([]int, len(discs))
	for i := 0; i < len(n_hats); i++ {
		product := 1
		for j := 0; j < len(n_hats); j++ {
			if i == j {
				continue
			}
			product *= discs[j].nb_positions
		}
		n_hats[i] = product
	}
	reverses := make([]int, len(discs))
	for i := 0; i < len(reverses); i++ {
		bigN := big.NewInt(int64(n_hats[i]))
		bigBase := big.NewInt(int64(discs[i].nb_positions))
		bigGcd := big.NewInt(0)
		bigInverse := big.NewInt(0)
		bigGcd.GCD(bigInverse, nil, bigN, bigBase)
		if bigInverse.IsInt64() {
			reverses[i] = real_modulo(int(bigInverse.Int64()), discs[i].nb_positions)
		} else {
			log.Fatalf("%v cannot be represented as int64", bigInverse)
		}
	}
	solution := 0
	for i := 0; i < len(reverses); i++ {
		solution += (discs[i].nb_positions - discs[i].position_at_t0 - i - 1) * reverses[i] * n_hats[i]
	}
	full_product := 1
	for i := 0; i < len(discs); i++ {
		full_product *= discs[i].nb_positions
	}
	solution = real_modulo(solution, full_product)
	log.Printf("Part1: %d", solution)

}

func real_modulo(value int, base int) int {
	return ((value % base) + base) % base
}
