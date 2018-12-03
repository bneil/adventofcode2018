package main

import (
	"fmt"
	"github.com/bneil/adventofcode2018/shared"
	"strconv"
	"strings"
)

//#1 @ 1,3: 4x4
type Claim struct {
	id int
	x int
	y int
	width int
	height int
}

// if the overlap is greater then 2 mark the claim
// so we need to setup a huge map or just determine areas of ownership
// and mark it when the claim is higher
func day3_part1(){
	lines, err := shared.ReadLines("day3input.txt")
	shared.Check(err)

	var fabric map[string]int
	fabric = make(map[string]int)

	for _, line := range lines {
		claimRaw := strings.Split(line, " ")
		_, err := strconv.Atoi(claimRaw[0][1:])
		shared.Check(err)

		coords := strings.Split(claimRaw[2], ",")

		x, err := strconv.Atoi(coords[0])
		shared.Check(err)

		meh := strings.Split(coords[1], ":")
		y, err := strconv.Atoi(meh[0])
		shared.Check(err)

		dimensions := strings.Split(claimRaw[3], "x")

		width, err := strconv.Atoi(dimensions[0])
		shared.Check(err)

		height, err := strconv.Atoi(dimensions[1])
		shared.Check(err)

		// lets create a map of claims of owned areas
		xClaim := x + width
		yClaim := y + height
		for posX := x; posX < xClaim; posX++ {
			for posY := y; posY < yClaim; posY++ {
				position := string(posX)+":"+string(posY)
				fabric[position] += 1
			}
		}
	}

	var twoOrMore = 0
	for _, value := range fabric {
		if(value >= 2 ){
			twoOrMore += 1
		}
	}
	fmt.Printf("How many square inches of fabric are within two or more claims? %d\n", twoOrMore)
}

func GenerateClaim(claimRaw []string) Claim {
		id, err := strconv.Atoi(claimRaw[0][1:])
		shared.Check(err)

		coords := strings.Split(claimRaw[2], ",")

		x, err := strconv.Atoi(coords[0])
		shared.Check(err)

		meh := strings.Split(coords[1], ":")
		y, err := strconv.Atoi(meh[0])
		shared.Check(err)

		dimensions := strings.Split(claimRaw[3], "x")

		width, err := strconv.Atoi(dimensions[0])
		shared.Check(err)

		height, err := strconv.Atoi(dimensions[1])
		shared.Check(err)

		return Claim{
			id: id,
			x: x,
			y: y,
			width: width,
			height: height,
		}
}

// dodgeball
// first register all folks, then if any of there claims intersect they are out!
// its worth mentioning id => positions means that if any intersection occurs
// then position invalidates id
func day3_part2() {
	lines, err := shared.ReadLines("day3input.txt")
	shared.Check(err)

	var fabric map[string]int
	fabric = make(map[string]int)

	var erase map[int]int
	erase = make(map[int]int)

	for _, line := range lines {
		claimRaw := strings.Split(line, " ")
		claim := GenerateClaim(claimRaw)

		// lets create a map of claims of owned areas
		xClaim := claim.x + claim.width
		yClaim := claim.y + claim.height
		//fmt.Printf("xClaim: %d\n",xClaim)
		//fmt.Printf("yClaim: %d\n",yClaim)
		for posX := claim.x; posX < xClaim; posX++ {
			//fmt.Printf("posX: %d\n",posX)
			for posY := claim.y; posY < yClaim; posY++ {
				//fmt.Printf("posY: %d\n",posY)
				position := fmt.Sprintf("%d:%d", posX,posY)
				if(fabric[position] == 0) {
					fabric[position] = claim.id
				}else{
					// erradicate both claim.id's
					erase[claim.id] = 1
					erase[fabric[position]] = 1
					fabric[position] = 0

				}
			}
		}
	}
	fmt.Printf("removing all empties\n")
	for pos, id := range fabric {
		if(id == 0) {
			//fmt.Printf("deleting %s\n", pos)
			delete(fabric, pos)
		}
	}
	fmt.Printf("finished going over positions: now removing (%d) claims\n", len(erase))
	for eid, _ := range erase {
		for pos, id := range fabric {
			if(eid == id){
				delete(fabric, pos)
			}
		}
	}

	fmt.Printf("done removing erased ids\n")
	for _, value := range fabric {
		fmt.Printf("final answer: %d\n", value)
		break
	}
}

// checksums and bad loops
func main() {
	//fmt.Println("Part 1")
	//day3_part1()

	fmt.Println("Part 2")
	day3_part2()
}

