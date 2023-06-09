package main

import (
	"fmt"
	"github.com/thoas/go-funk"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type InvalidRoom struct {
	Line string
}

func (e InvalidRoom) Error() string {
	return fmt.Sprintf("%s cannot be parsed as a room", e.Line)
}

func main() {
	content, err := os.ReadFile("input04.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")

	part1 := 0

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		room, err := parseRoom(line)
		if err != nil {
			log.Fatal(err)
		}
		if room.valid() {
			part1 += room.sector_id
			if room.decrypt() == "northpole object storage" {
				log.Print(fmt.Sprintf("Part 2: %d", room.sector_id))
			}
		} else {
			// log.Print(room, "is not a valid room")
		}
	}
	log.Print(fmt.Sprintf("Part 1: %d", part1))
}

func parseRoom(line string) (Room, error) {
	re := regexp.MustCompile(`([a-z-]+)-([0-9]+)\[([a-z]+)\]`)
	if match := re.FindSubmatch([]byte(line)); match != nil {
		id, err := strconv.Atoi(string(match[2]))
		if err != nil {
			return Room{}, InvalidRoom{line} // TODO(g.seux): we could add a reason for invalid parsing
		}
		return Room{
			letters:   string(match[1]),
			sector_id: id,
			checksum:  string(match[3])}, nil
	} else {
		return Room{}, InvalidRoom{line}
	}
}

func (room Room) decrypt() string {
	var decrypted = make([]byte, 0)
	for i := 0; i < len(room.letters); i++ {
		if room.letters[i] == '-' {
			decrypted = append(decrypted, ' ')
			continue
		}
		decrypted_letter := (int(room.letters[i])-97+room.sector_id)%26 + 97
		decrypted = append(decrypted, byte(decrypted_letter))
	}
	return string(decrypted)
}

func (room Room) valid() bool {
	// count letters
	var counts = make(map[byte]int)
	for i := 0; i < len(room.letters); i++ {
		if room.letters[i] == '-' {
			continue
		}
		counts[room.letters[i]] += 1
	}
	keys := funk.Keys(counts).([]byte)
	sort.Slice(keys, func(i, j int) bool {
		counti, countj := counts[keys[i]], counts[keys[j]]
		if counti == countj {
			return keys[i] < keys[j] // alphabetical order
		}
		return counti > countj
	})
	expected_checksum := string(keys[0:5])
	return expected_checksum == room.checksum
}

type Room struct {
	letters   string
	sector_id int
	checksum  string
}
