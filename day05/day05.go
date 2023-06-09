package main

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("input05.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := strings.Trim(string(content), "\n")

	password_part1 := make([]string, 0)
	password_part2 := make([]string, 8)
	index := 0

	for {
		candidate := []byte(input + strconv.Itoa(index))
		hash := md5.Sum(candidate)
		stringed_hash := hex.EncodeToString(hash[:])
		truncated, valid := strings.CutPrefix(stringed_hash, "00000")
		if valid {
			password_part1 = append(password_part1, string(truncated[0]))
			log.Printf("Found one character of the password: %s", string(truncated[0]))

			position := truncated[0] - 48
			if position >= 0 && position < 8 && password_part2[position] == "" {
				log.Printf("Found one character of the password for part2: %s", string(truncated[1]))
				password_part2[position] = string(truncated[1])

				must_continue := false
				for _, c := range password_part2 {
					if c == "" {
						must_continue = true
					}
				}
				if !must_continue {
					break
				}
			}
		}
		index++
		if index%100000 == 0 {
			log.Printf("iteration %d", index)
		}
	}
	log.Printf("Part 1: %s", password_part1[0:8])
	log.Printf("Part 2: %s", password_part2)
}
