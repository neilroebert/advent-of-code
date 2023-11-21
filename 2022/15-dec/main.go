package main

import (
	"bufio"
	"fmt"
	"math"
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

func printMap(toPrint [][]string, i int) {
	j := 0
	for _, line := range toPrint {
		fmt.Printf("%v:%v\t%v\n", line, i, j)
		i++
		j++
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

func drawScanOpt(sX int, sY int, shortDist int, scanMap []string, testVal int) []string {
	// [2, 18, 7]
	fmt.Println("drawScan: ", sX, sY, shortDist)
	add := true
	x := 1
	for y := 0; y < shortDist*2+1; y++ {
		// fmt.Println("sY-shortDist+y: ", sY-shortDist+y)

		if add {

			if (sY - shortDist + y) == testVal {
				for t := 0; t < x; t++ {
					xVal := sX - (x - 1) + (((x - 1) / 2) + t)
					// fmt.Println("xVal: ", (((x - 1) / 2) + t))

					if scanMap[xVal] == "." {
						scanMap[xVal] = "#"
					}

				}
			}
			if (sY - shortDist + y) == sY {
				add = false
				x -= 2
			} else {
				x += 2
			}

		} else {
			if (sY - shortDist + y) == testVal {
				for t := 0; t < x; t++ {

					xVal := sX - (x - 1) + (((x - 1) / 2) + t)
					// fmt.Println("xVal: ", xVal)

					if scanMap[xVal] == "." {
						scanMap[xVal] = "#"
					}

				}
			}
			x -= 2

		}
		// printMap(scanMap, -10)
	}
	return scanMap
}

func drawScan(sX int, sY int, shortDist int, scanMap [][]string) [][]string {
	// [2, 18, 7]
	fmt.Println("drawScan: ", sX, sY, shortDist)
	add := true
	x := 1
	for y := 0; y < shortDist*2+1; y++ {
		// fmt.Println("sY-shortDist+y: ", sY-shortDist+y)

		if add {

			for t := 0; t < x; t++ {
				xVal := sX - (x - 1) + (((x - 1) / 2) + t)
				// fmt.Println("xVal: ", (((x - 1) / 2) + t))
				if scanMap[sY-shortDist+y][xVal] == "." {
					scanMap[sY-shortDist+y][xVal] = "#"
				}
			}
			if (sY - shortDist + y) == sY {
				add = false
				x -= 2
			} else {
				x += 2
			}

		} else {
			for t := 0; t < x; t++ {

				xVal := sX - (x - 1) + (((x - 1) / 2) + t)
				// fmt.Println("xVal: ", xVal)
				if scanMap[sY-shortDist+y][xVal] == "." {
					scanMap[sY-shortDist+y][xVal] = "#"
				}
			}
			x -= 2

		}
		// printMap(scanMap, -10)
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

//transform coord
func tC(c int, start int) int {
	return int(math.Abs(float64(start - c)))
}

func main() {
	// dat, err := os.Open("test-input.txt")
	dat, err := os.Open("input.txt")
	check(err)

	fileScanner := bufio.NewScanner(dat)
	fileScanner.Split(bufio.ScanLines)

	// debug := true

	var coords [][]int

	// Arrays to contain min and max vals - [minX, maxX, minY, maxY]
	minmaxArray := make([]int, 4)
	first := true
	yTestVal := 2000000
	// yTestVal = 10

	// this is so that we can build the array
	for fileScanner.Scan() {
		line := strings.Replace(fileScanner.Text(), "Sensor at ", "", 1)
		line = strings.Replace(line, " closest beacon is at ", "", 1)
		line = strings.Replace(line, "x=", "", 2)
		line = strings.Replace(line, " y=", "", 2)

		co := strings.Split(line, ":")
		sb := make([]int, 5)
		// fmt.Printf("%v\n", line)
		c := strings.Split(co[0], ",")
		sb[0], err = strconv.Atoi(c[0])
		check(err)
		sb[1], err = strconv.Atoi(c[1])
		check(err)
		c = strings.Split(co[1], ",")
		sb[2], err = strconv.Atoi(c[0])
		check(err)
		sb[3], err = strconv.Atoi(c[1])
		check(err)

		// shortest path distance
		sb[4] = int(math.Abs(float64(sb[0]-sb[2])) + math.Abs(float64(sb[1]-sb[3])))
		// fmt.Printf("%v\n", sb)
		coords = append(coords, sb)

		if first {
			// init min and max vals
			minmaxArray[0] = sb[0]
			minmaxArray[1] = sb[0]
			minmaxArray[2] = sb[1]
			minmaxArray[3] = sb[1]

			// beaconArray[0] = sb[2]
			// beaconArray[1] = sb[2]
			// beaconArray[2] = sb[3]
			// beaconArray[3] = sb[3]

			first = false
		}

		// // Sensor min max checks
		// // X checks
		// if sb[0] < minmaxArray[0] {
		// 	minmaxArray[0] = sb[0]
		// }
		// if sb[0] > minmaxArray[1] {
		// 	minmaxArray[1] = sb[0]
		// }

		// // Y checks
		// if sb[1] < minmaxArray[2] {
		// 	minmaxArray[2] = sb[1]
		// }
		// if sb[1] > minmaxArray[3] {
		// 	minmaxArray[3] = sb[1]
		// }
		// // fmt.Printf("%v\n\n", minmaxArray)

		// // Beacon checks
		// // X checks
		// if sb[2] < minmaxArray[0] {
		// 	minmaxArray[0] = sb[2]
		// }
		// if sb[2] > minmaxArray[1] {
		// 	minmaxArray[1] = sb[2]
		// }

		// // Y checks
		// if sb[3] < minmaxArray[2] {
		// 	minmaxArray[2] = sb[3]
		// }
		// if sb[3] > minmaxArray[3] {
		// 	minmaxArray[3] = sb[3]
		// }

		// Furthest distance checks
		// X checks
		if (sb[0] - sb[4]) < minmaxArray[0] {
			minmaxArray[0] = sb[0] - sb[4] - 2
		}
		if (sb[0] + sb[4]) > minmaxArray[1] {
			minmaxArray[1] = sb[0] + sb[4] + 2
		}

		// Y checks
		if (sb[1] - sb[4]) < minmaxArray[2] {
			minmaxArray[2] = (sb[1] - sb[4] - 2)
		}
		if (sb[1] + sb[4]) > minmaxArray[3] {
			minmaxArray[3] = (sb[1] + sb[4]) + 2
		}

		// fmt.Printf("%v\n\n", minmaxArray)
	}

	sXLength := minmaxArray[1] - minmaxArray[0]
	sYLength := minmaxArray[3] - minmaxArray[2]

	fmt.Println(sXLength, sYLength)

	// sbArray := make([][]string, sYLength)

	// This didn't work because of memory issues. Very big arrays
	// for i := 0; i < sYLength; i++ {
	// 	sbArray[i] = make([]string, sXLength)
	// 	for j := 0; j < sXLength; j++ {
	// 		sbArray[i][j] = "."
	// 	}
	// }

	smallArray := make([]string, sXLength)
	for j := 0; j < sXLength; j++ {
		smallArray[j] = "."
	}

	// printMap(sbArray, minmaxArray[2])

	part_one_val := 0
	part_two_val := 0

	// build the sensor and beacon layout
	for _, line := range coords {
		fmt.Printf("Line: %v\n", line)

		// sbArray[tC(minmaxArray[2], line[1])][tC(minmaxArray[0], line[0])] = "S"
		// sbArray[tC(minmaxArray[2], line[3])][tC(minmaxArray[0], line[2])] = "B"
		if line[1] == yTestVal {
			smallArray[tC(minmaxArray[0], line[0])] = "S"
		}
		if line[3] == yTestVal {
			smallArray[tC(minmaxArray[0], line[2])] = "B"
		}
		smallArray = drawScanOpt(tC(minmaxArray[0], line[0]), tC(minmaxArray[2], line[1]), line[4], smallArray, tC(minmaxArray[2], yTestVal))
		// printMap(sbArray, minmaxArray[2])

	}
	// printMap(sbArray, minmaxArray[2])
	// fmt.Printf("%v\n", minmaxArray)
	// yTestVal := 20000
	// for _, testLine := range sbArray[tC(minmaxArray[2], yTestVal)] {
	// 	if testLine == "#" {
	// 		part_one_val++
	// 	}
	// }

	for _, testLine := range smallArray {
		if testLine == "#" {
			part_one_val++
		}
	}

	fmt.Println("Part one: ", part_one_val)

	fmt.Println("Part two: ", part_two_val)

	dat.Close()
}
