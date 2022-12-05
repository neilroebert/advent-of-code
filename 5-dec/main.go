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

func printBox(toPrint [][]string) {
	for i := 0; i < len(toPrint); i++ {
		// fmt.Println(len(toPrint[i]))
		fmt.Print(i+1, ": ")
		for j := len(toPrint[i]); j > 0; j-- {

			fmt.Print(toPrint[i][j-1], " ")
		}
		fmt.Println()
	}
}

func main() {
	// dat, err := os.Open("test-input.txt")
	dat, err := os.Open("input.txt")
	check(err)

	fileScanner := bufio.NewScanner(dat)

	fileScanner.Split(bufio.ScanLines)

	const boxLength = 9

	var boxQueue = make([][]string, boxLength)
	var boxQueue2 = make([][]string, boxLength)
	for i := 0; i < boxLength; i++ {
		boxQueue[i] = make([]string, 0)
		boxQueue2[i] = make([]string, 0)
	}

	inCommands := false

	debug := 10
	for fileScanner.Scan() {
		line := fileScanner.Text()

		if line == "" {
			inCommands = true

			continue
		}

		if inCommands {
			//process commands

			if debug < 10 {
				fmt.Println(line)
			}
			//move 5 from 6 to 7
			line = strings.Replace(line, "move ", "", 1)
			s := strings.Split(line, " from ")
			tmps := strings.Split(s[1], " to ")

			from, err := strconv.Atoi(tmps[0])
			check(err)
			to, err := strconv.Atoi(tmps[1])
			check(err)

			moveCount, err := strconv.Atoi(s[0])
			check(err)

			if debug < 10 {
				printBox(boxQueue2)

				fmt.Print("Pop: ")
			}

			// boxqueue 1
			for i := 0; i < moveCount; i++ {
				// pop
				move := boxQueue[from-1][0]
				// if debug < 10 {
				// 	fmt.Print(move)
				// }

				if len(boxQueue[from-1]) > 1 {
					boxQueue[from-1] = boxQueue[from-1][1:]
				} else {
					boxQueue[from-1] = make([]string, 0)
				}

				// push
				boxQueue[to-1] = append([]string{move}, boxQueue[to-1]...)
				// boxQueue[to-1] = append(boxQueue[to-1], move)

			}

			// boxqueue2
			// We
			// pop slice
			move := boxQueue2[from-1][0:moveCount]

			if moveCount < len(boxQueue2[from-1]) {
				if debug < 10 {
					fmt.Printf("Boxqueue [movecount:] %v\n", boxQueue2[from-1][moveCount:])
				}
				boxQueue2[from-1] = boxQueue2[from-1][moveCount:]
				if debug < 10 {
					fmt.Printf("Boxqueue from-1 %v\n", boxQueue2[from-1])
				}
			} else {
				boxQueue2[from-1] = make([]string, 0)
			}

			// push slice
			// we can't just straight up push a slice. Getting into some weird pointer territory then.
			// https://riteeksrivastava.medium.com/how-slices-internally-work-in-golang-a47fcb5d42ce
			// this explains slices and memory quite well

			if debug < 10 {
				fmt.Printf("Boxqueue to-1 %v\n", boxQueue2[to-1])
			}

			for i := moveCount; i > 0; i-- {
				// fmt.Println(move[i-1])
				boxQueue2[to-1] = append([]string{move[i-1]}, boxQueue2[to-1]...)
			}

			// boxQueue2[to-1] = append(move, boxQueue2[to-1]...)

			if debug < 10 {
				fmt.Printf("Boxqueue to-1 %v\n", boxQueue2[to-1])
			}

			if debug < 10 {
				fmt.Printf("Boxqueue from-1 %v\n", boxQueue2[from-1])
			}

			if debug < 10 {
				fmt.Printf("%v", move)
				fmt.Println()
				printBox(boxQueue2)
				fmt.Println("========================")
			}

			debug++

		} else {
			//build arrays
			n := 0
			for i := 1; i < len(line); i += 4 {

				if string(line[i]) == "1" {
					break
				}
				print(string(line[i]), " ")
				if string(line[i]) != " " {
					boxQueue[n] = append(boxQueue[n], string(line[i]))
					boxQueue2[n] = append(boxQueue2[n], string(line[i]))
				}
				n++
			}
			n = 0
			fmt.Println()
		}

	}

	fmt.Print("Part one: ")
	for i := 0; i < boxLength; i += 1 {
		print(boxQueue[i][0])
	}
	fmt.Println()

	fmt.Print("Part two: ")
	for i := 0; i < boxLength; i += 1 {
		print(boxQueue2[i][0])
	}
	fmt.Println()
	dat.Close()
}
