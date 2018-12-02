package main

import (
	"fmt"
	"github.com/bneil/adventofcode2018/shared"
	"strconv"
)

// just iterate and see the final culmative answer
func part1(){
	ctr := 0
	frequencies, err := shared.ReadLines("day1input.txt")
	shared.Check(err)

	for _, num := range frequencies {
		i, err := strconv.Atoi(num)
		shared.Check(err)
		ctr += i
	}

	fmt.Printf("final value: %d\n", ctr)
}

// keep iterating until a duplicate frequency is found
func part2(){
	frequency := 0
	iteration := 0
	firstRepeatedFreq := 0
	var seen map[int]int
	seen = make(map[int]int)
	found := false

	frequencies, err := shared.ReadLines("day1input.txt")
	shared.Check(err)

	for found == false {
		iteration += 1
		for _, num := range frequencies {
			next, err := strconv.Atoi(num)
			shared.Check(err)

			frequency += next

			seen[frequency] = seen[frequency] + 1

			if(seen[frequency] == 2 && firstRepeatedFreq == 0){
				firstRepeatedFreq = frequency
				found = true
				break
			}
		}
	}

	fmt.Printf("the first repeated frequency was %d after %d iterations \n", firstRepeatedFreq, iteration)
}

func main() {
	fmt.Println("Part 1")
	part1()

	fmt.Println("Part 2")
	part2()
}
