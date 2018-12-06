package main

import (
	"fmt"
	"github.com/bneil/adventofcode2018/shared"
	"sort"
	"strings"
	"time"
)

//[1518-11-01 00:00] Guard #10 begins shift
type Entry struct {
	timestamp time.Time
	state string
}

func parseEntry(line string) Entry{
	halfs := strings.Split(line, "]")

	layout := "2006-01-02 15:04"
	rawDate := strings.Split(halfs[0], "[")[1]
	t, err := time.Parse(layout, rawDate)
	shared.Check(err)
	action := halfs[1]

	if(strings.Contains(action, "begins shift")){
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

	var entries []Entry
	for _, line := range lines {
		entry := parseEntry(line)
		fmt.Printf("%+v\n", entry)
		entries = append(entries, entry)
	}
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].timestamp.Before(entries[j].timestamp)
	})
	//for entry := range entries {
	//}

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
