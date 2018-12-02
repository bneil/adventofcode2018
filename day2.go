package main

import (
	"fmt"
	"github.com/bneil/adventofcode2018/shared"
	"strings"
)

// read in the line, then count the occurences of the letter
// next break those results into there own map
func day2_part1(){
	lines, err := shared.ReadLines("day2input.txt")
	shared.Check(err)

	var charOccurance map[string]int
	charOccurance = make(map[string]int)

	var overallOccurance map[int]int
	overallOccurance = make(map[int]int)

	var onlyOnce map[int]int
	onlyOnce = make(map[int]int)

	for _, str := range lines {
		for _, r := range str {
			c := string(r)
			charOccurance[c] = charOccurance[c] + 1
		}

		// c => 3 (char occurance)
		// 3 => 1 (overall occurance)
		for _, occurance := range charOccurance {
			// another gotch rule, the occurance may only be counted once
			// so two occurences of 3 for an id only could as 1 - cost me 10 minutes of debugging
			//fmt.Printf("%s %d \n", c, occurance)
			if(onlyOnce[occurance] == 0) {
				overallOccurance[occurance] = overallOccurance[occurance] + 1
				onlyOnce[occurance] = 1
			}
		}

		charOccurance = make(map[string]int)
		onlyOnce = make(map[int]int)
	}

	var acc = 1
	for occ, occurance := range overallOccurance {
		// reading the rules only 2 and 3 occurances matter
		if(occ == 2 || occ == 3){
			acc = acc * occurance
		}
	}
	println(acc)
}
//discover likeness quickly
func howAlike(w1 string, w2 string) int {
	w1broken := []rune(w1)
	w2broken := []rune(w2)

	var times = 0
	for i, r := range w1broken {
		if(r == w2broken[i]){
			times += 1
		}
	}
	return times
}

func removeDifferences(w1 string, w2 string) string {
	w1broken := []rune(w1)
	w2broken := []rune(w2)
	var final = ""
	for i, r := range w1broken {
		if (r == w2broken[i]) {
			final += string(r)
		}
	}
	return final
}

// figure out the amount of differences between each id, tally and
// finally figure out which ids are closest together
// - the following could have used alot of cleanup
// - having a go channel handle a scatter gather would have been way more efficient
func day2_part2() {
	lines, err := shared.ReadLines("day2input.txt")
	shared.Check(err)

	var overallBestMatches map[int]string
	overallBestMatches = make(map[int]string)

	// used so we dont repeat the words we've already scanned
	var alreadyCompared map[string]int
	alreadyCompared = make(map[string]int)


	for me, currentWord := range lines {

		for me2, word := range lines {
			if(me == me2){
				continue
			}

			if(alreadyCompared[currentWord+":"+word] == 0){
				alreadyCompared[currentWord+":"+word] = 1
				alreadyCompared[word+":"+currentWord] = 1
			}else{
				continue
			}
			numOfMatchingLetters := howAlike(currentWord, word)

			fmt.Printf("%s and %s are %d\n", currentWord, word, numOfMatchingLetters)
			if(numOfMatchingLetters > 1){
				overallBestMatches[numOfMatchingLetters] = currentWord+":"+word
			}
		}

	}

	var highest = 0
	var winner = ""
	for numberMatched, theBestWords := range overallBestMatches {
		if(numberMatched > highest){
			winner = theBestWords
			highest = numberMatched
		}
	}

	possibleWinners := strings.Split(winner, ":")
	finalID := removeDifferences(possibleWinners[0], possibleWinners[1])
	fmt.Println(finalID)
}

// checksums and bad loops
func main() {
	fmt.Println("Part 1")
	day2_part1()

	fmt.Println("Part 2")
	day2_part2()
}
