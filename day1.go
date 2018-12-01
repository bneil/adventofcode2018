package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}


// just iterate and see the final culmative answer
func part1(){
	ctr := 0
	frequencies, err := readLines("day1input.txt")
	check(err)

	for _, num := range frequencies {
		i, err := strconv.Atoi(num)
		check(err)
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

	frequencies, err := readLines("day1input.txt")
	check(err)

	for found == false {
		iteration += 1
		for _, num := range frequencies {
			next, err := strconv.Atoi(num)
			check(err)

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
