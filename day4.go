package main

import (
	"fmt"
	"github.com/bneil/adventofcode2018/shared"
	"strings"
	"time"
)

//[1518-11-01 00:00] Guard #10 begins shift
type Entry struct {
	timestamp time.Time
	state string
}

//
func parseEntry(line string) Entry{
	halfs := strings.Split(line, "]")

	layout := "2006-02-01 00:60"
	rawDate := strings.Split(halfs[0], "[")[1]
	println(rawDate)
	t, err := time.Parse(layout, rawDate)
	shared.Check(err)
	action := halfs[1]

	if(strings.Contains(action, "being shift")){
		return Entry{
			timestamp: t,
			state: "working",
		}

	}else if(strings.Contains(action, "wakes up")){
		return Entry{
			timestamp: t,
			state: "woke",
		}

	}else{
		return Entry{
			timestamp: t,
			state: "asleep",
		}
	}
}

// Strategy 1: Find the guard that has the most minutes asleep.
// What minute does that guard spend asleep the most?
func day4_part1(){
	lines, err := shared.ReadLines("test")
	shared.Check(err)

	for _, line := range lines {
		entry := parseEntry(line)
		fmt.Printf("%+v\n", entry)

	}
}

func day4_part2() {

}

// checksums and bad loops
func main() {
	fmt.Println("Part 1")
	day4_part1()

	fmt.Println("Part 2")
	day4_part2()
}
