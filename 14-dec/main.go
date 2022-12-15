package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func debugPrint(toPrint string, debug bool) {
	if debug {
		fmt.Println(toPrint)
	}
}

func printMap(toPrint [][]string) {
	i := 0
	for _, line := range toPrint {
		fmt.Printf("%v:%v\n", line, i)
		i++
	}
	fmt.Println("------------")
}

func countSand(scanMap [][]string) int {
	count := 0
	for _, line := range scanMap {
		for _, l := range line {
			if l == "O" {
				count++
			}
		}
	}
	return count
}

func drawLine(x1 int, y1 int, x2 int, y2 int, scanMap [][]string) [][]string {
	startVal := 0
	endVal := 0
	if x1 == x2 {
		// y change
		if y1 < y2 {
			startVal = y1
			endVal = y2
		} else {
			startVal = y2
			endVal = y1
		}
		for i := startVal; i <= endVal; i++ {
			scanMap[i][x1] = "#"
		}
	} else {
		// x change
		if x1 < x2 {
			startVal = x1
			endVal = x2
		} else {
			startVal = x2
			endVal = x1
		}
		for i := startVal; i <= endVal; i++ {
			scanMap[y1][i] = "#"
		}
	}
	return scanMap
}

func simSand(scanMap [][]string, sandLoc int) ([][]string, bool) {

	xLen := len(scanMap[0])
	yLen := len(scanMap)
	finished := false

	for i := 0; i < yLen; i++ {
		// fmt.Println(i, sandLoc)
		// end := false
		if i+1 == len(scanMap) && scanMap[i][sandLoc] == "." {
			// falling off bottom
			finished = true
			break
			// fmt.Println("end")
		}

		//check next one
		if scanMap[i+1][sandLoc] == "#" || scanMap[i+1][sandLoc] == "O" {
			// fmt.Println("Blocked next")

			if sandLoc > 0 {
				// fmt.Println("Sandloc > 0")
				if scanMap[i+1][sandLoc-1] == "#" || scanMap[i+1][sandLoc-1] == "O" {
					// test diag left
					// fmt.Println("Diag left blocked")
				} else {
					sandLoc--
					continue
				}
			} else {
				//falling off the edge
				finished = true
				break
			}
			if sandLoc < xLen-1 {
				// fmt.Println("Sandloc <", xLen-1)
				if scanMap[i+1][sandLoc+1] == "#" || scanMap[i+1][sandLoc+1] == "O" {
					// test diag right
					// fmt.Println("Diag right blocked")

				} else {
					sandLoc++
					continue
				}
			} else {
				//falling off the edge
				// fmt.Println("falling off edge")
				finished = true
				break
			}

			// else settle

			scanMap[i][sandLoc] = "O"
			break

		}

	}

	return scanMap, finished
}

func main() {
	// dat, err := os.Open("test-input.txt")
	dat, err := os.Open("input.txt")
	check(err)

	fileScanner := bufio.NewScanner(dat)
	fileScanner.Split(bufio.ScanLines)

	// debug := true

	minX := 500
	maxX := 500
	maxY := 0

	var lines []string

	// this is so that we can build the array
	for fileScanner.Scan() {
		line := fileScanner.Text()
		lines = append(lines, line)
		coords := strings.Split(line, " -> ")
		// fmt.Printf("%v\n", coords)
		for _, coord := range coords {
			c := strings.Split(coord, ",")
			// test x vals
			tmp, err := strconv.Atoi(c[0])
			check(err)
			if tmp < minX {
				minX = tmp
			} else if tmp > maxX {
				maxX = tmp
			}

			//test y vals
			tmp, err = strconv.Atoi(c[1])
			check(err)
			if tmp > maxY {
				maxY = tmp
			}
		}
	}

	// remove the following lines for part 1
	minX -= 400
	maxX += 400

	arrLength := maxX - minX + 1

	// make this maxY+1 for part 1
	scanMap := make([][]string, maxY+3)
	for i := 0; i < maxY+3; i++ {
		scanMap[i] = make([]string, arrLength)
		for j := 0; j < arrLength; j++ {
			if i == maxY+2 {
				// remove this for part 1
				scanMap[i][j] = "#"
			} else {

				scanMap[i][j] = "."
			}
		}
		// fmt.Printf("%v\n", scanMap[i])

	}
	sandLoc := 500 - minX
	scanMap[0][sandLoc] = "+"
	printMap(scanMap)

	part_one_val := 0
	part_two_val := 0

	// build the rock layout
	for _, line := range lines {

		coords := strings.Split(line, " -> ")
		// fmt.Printf("%v\n", coords)
		for i := 0; i < len(coords)-1; i++ {
			c1 := strings.Split(coords[i], ",")
			c2 := strings.Split(coords[i+1], ",")
			// x1 vals
			x1, err := strconv.Atoi(c1[0])
			check(err)

			// y1 vals
			y1, err := strconv.Atoi(c1[1])
			check(err)

			// x2 vals
			x2, err := strconv.Atoi(c2[0])
			check(err)

			// y1 vals
			y2, err := strconv.Atoi(c2[1])
			check(err)
			scanMap = drawLine(x1-minX, y1, x2-minX, y2, scanMap)
		}

	}

	sandFull := false

	for !sandFull {
		scanMap, sandFull = simSand(scanMap, sandLoc)
		// printMap(scanMap)
		//if just below + == O then break
		if scanMap[0][sandLoc] == "O" {
			sandFull = true
			continue
		}
	}

	printMap(scanMap)
	part_one_val = countSand(scanMap)

	fmt.Println("Part one: ", part_one_val)

	fmt.Println("Part two: ", part_two_val)

	dat.Close()
}
