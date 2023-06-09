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

	password := make([]string, 0)
	index := 0

	for len(password) < 8 {
		candidate := []byte(input + strconv.Itoa(index))
		hash := md5.Sum(candidate)
		stringed_hash := hex.EncodeToString(hash[:])
		truncated, valid := strings.CutPrefix(stringed_hash, "00000")
		if valid {
			password = append(password, string(truncated[0]))
			log.Printf("Found one character of the password: %s", string(truncated[0]))
		}
		index++
		if index%10000 == 0 {
			log.Printf("iteration %d", index)
		}
	}
	log.Printf("Part 1: %s", password)
}
