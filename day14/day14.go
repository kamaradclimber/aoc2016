package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
)

func Hash(salt string, i int, stretches int) string {
	input := fmt.Sprintf("%s%d", salt, i)
	for i := 0; i <= stretches; i++ {
		h := md5.New()
		io.WriteString(h, input)
		input = fmt.Sprintf("%x", h.Sum(nil))
	}
	return input
}

type Hashes struct {
	hashes              [][1000]string
	index_of_quintuples map[rune]([]int)
}

func NewHashes() Hashes {
	h := Hashes{}
	h.index_of_quintuples = make(map[rune]([]int))
	return h
}

func (h *Hashes) get(i int, stretches int) string {
	for j := len(h.hashes); j < 1+i/1000; j++ {
		var block [1000]string
		for k := 0; k < 1000; k++ {
			block[k] = Hash("qzyelonm", j*1000+k, stretches)

			for _, rune := range ExtractChunks(block[k], 5) {
				h.index_of_quintuples[rune] = append(h.index_of_quintuples[rune], j*1000+k)
			}
		}

		h.hashes = append(h.hashes, block)
	}
	return h.hashes[i/1000][i%1000]
}

func ExtractChunks(s string, n int) []rune {
	var chunks []rune
	i := 0
	j := 0
	for i < len(s)-n+1 {
		if i+j < len(s) && s[i] == s[i+j] {
			j++
		} else {
			if j >= n {
				chunks = append(chunks, rune(s[i]))
			}
			i = i + j
			j = 0
		}
	}
	return chunks
}

func main() {
	log.Printf("Part1: %d", buildOneTimePad(0))
	log.Printf("Part2: %d", buildOneTimePad(2016))
}

func buildOneTimePad(stretches int) int {
	h := NewHashes()

	found_keys := 0
	i := 0
	for {
		// invoke i-th hash to make sure it exist. They are lazyly computed and treated
		h.get(i+1000, stretches)

		for _, rune := range ExtractChunks(h.get(i, stretches), 3) {
			for _, index := range h.index_of_quintuples[rune] {
				if index <= i {
					continue
				}
				if index > i+1000 {
					break
				}
				found_keys += 1
				log.Printf("Found %d-th key %s at index %d (%s) because at index %d we have %s", found_keys, string(rune), i, h.get(i, stretches), index, h.get(index, stretches))
				if found_keys >= 64 {
					return i
				}
				break // ideally we should break out of the 2 loops at once but it's likely not that frequent to have 2 triples in the same hash
			}
			break // we should only consider the first triplet of the hash
		}
		i++
	}
}
