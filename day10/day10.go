package main

import (
	// "fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Bot struct {
	chips                  []int
	give_low_to_bot        bool
	give_low_to_bot_id     int
	give_low_to_output_id  int
	give_high_to_bot       bool
	give_high_to_bot_id    int
	give_high_to_output_id int
}

func main() {
	content, err := os.ReadFile("input10.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(content), "\n")

	bots := make([]Bot, 255)
	for i := 0; i < len(bots); i++ {
		bots[i] = Bot{}
	}
	outputs := make([]int, 255) // we should use a struct to represent outputs instead of mixing up "0" output from the nil one

	values_goes := regexp.MustCompile(`value (\d+) goes to bot (\d+)`)
	gives := regexp.MustCompile(`bot (\d+) gives low to (output|bot) (\d+) and high to (output|bot) (\d+)`)

	// parse
	for _, line := range lines {
		if line == "" {
			continue
		}
		if match := values_goes.FindStringSubmatch(line); match != nil {
			value, _ := strconv.Atoi(match[1])
			bot_id, _ := strconv.Atoi(match[2])
			bots[bot_id].chips = append(bots[bot_id].chips, value)
		} else if match := gives.FindStringSubmatch(line); match != nil {
			bot_id, _ := strconv.Atoi(match[1])
			low_target_id, _ := strconv.Atoi(match[3])
			high_target_id, _ := strconv.Atoi(match[5])
			if match[2] == "output" {
				bots[bot_id].give_low_to_output_id = low_target_id
			} else {
				bots[bot_id].give_low_to_bot = true
				bots[bot_id].give_low_to_bot_id = low_target_id
			}
			if match[4] == "output" {
				bots[bot_id].give_high_to_output_id = high_target_id
			} else {
				bots[bot_id].give_high_to_bot = true
				bots[bot_id].give_high_to_bot_id = high_target_id
			}

		} else {
			log.Fatal("Unknown line type: ", line)
		}
	}

	// simulate

	simulate := true
	round := 1
	for simulate {
		simulate = false
		//log.Print("Round ", round)
		round++
		for bot_id := 0; bot_id < len(bots); bot_id++ {
			bot := bots[bot_id]
			if len(bot.chips) == 2 {
				simulate = true // at least one bot did something
				// it's so sad golang has no decent library for ints
				low := bot.chips[0]
				high := bot.chips[1]
				if high < low {
					low = bot.chips[1]
					high = bot.chips[0]
				}
				if low == 17 && high == 61 {
					log.Print("Part1: Bot ", bot_id, " is responsible for the relevant comparison")
					//simulate = false
					//break
				}

				if bot.give_low_to_bot {
					// log.Print(fmt.Sprintf("Bot %d gives %d to bot %d", bot_id, low, bot.give_low_to_bot_id))
					bots[bot.give_low_to_bot_id].chips = append(bots[bot.give_low_to_bot_id].chips, low)
				} else {
					outputs[bot.give_low_to_output_id] = low
				}
				if bot.give_high_to_bot {
					// log.Print(fmt.Sprintf("Bot %d gives %d to bot %d", bot_id, high, bot.give_high_to_bot_id))
					bots[bot.give_high_to_bot_id].chips = append(bots[bot.give_high_to_bot_id].chips, high)
				} else {
					outputs[bot.give_high_to_output_id] = high
				}
				bots[bot_id].chips = make([]int, 0)
			}
		}
	}
	part2 := outputs[0] * outputs[1] * outputs[2]
	log.Print("Part2: ", part2)
}
