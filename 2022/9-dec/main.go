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

func testTail(tailX int, tailY int, headX int, headY int) (int, int) {
	fmt.Println("Test Vals: ", tailX, tailY, headX, headY)
	if (headX-tailX) == 2 && headY != tailY {
		//diagonal h
		tailY += headY - tailY
		tailX++
	} else if (headX-tailX) == -2 && headY != tailY {
		//diagonal h
		tailY += headY - tailY
		tailX--
	} else if (headY-tailY) == 2 && headX != tailX {
		//diagonal y
		tailX += headX - tailX
		tailY++

	} else if (headY-tailY) == -2 && headX != tailX {
		//diagonal y
		tailX += headX - tailX
		tailY--

	} else if (headX-tailX) == 2 && headY == tailY {
		tailX++

	} else if (headX-tailX) == -2 && headY == tailY {
		tailX--

	} else if (headY-tailY) == 2 && headX == tailX {
		tailY++

	} else if (headY-tailY) == -2 && headX == tailX {
		tailY--
	}
	fmt.Println("Retn Vals: ", tailX, tailY, headX, headY)
	return tailX, tailY
}

func main() {
	// dat, err := os.Open("test-input.txt")
	dat, err := os.Open("input.txt")
	check(err)

	fileScanner := bufio.NewScanner(dat)

	fileScanner.Split(bufio.ScanLines)

	// debug := true

	part_one_val := 0
	part_two_val := 0

	tailX := 0
	tailY := 0
	headX := 0
	headY := 0
	visited := make(map[string]int)

	for fileScanner.Scan() {

		line := fileScanner.Text()
		cmd := strings.Split(line, " ")

		move, err := strconv.Atoi(cmd[1])
		check(err)
		for i := 0; i < move; i++ {

			if cmd[0] == "U" {
				//up
				fmt.Println("U")
				headY++
				tailX, tailY = testTail(tailX, tailY, headX, headY)
				visitedString := fmt.Sprintf("%d,%d", tailX, tailY)
				fmt.Println(visitedString)
				visited[visitedString] = 1

			} else if cmd[0] == "D" {
				//down
				headY--
				tailX, tailY = testTail(tailX, tailY, headX, headY)
				visitedString := fmt.Sprintf("%d,%d", tailX, tailY)
				visited[visitedString] = 1

			} else if cmd[0] == "L" {
				//left
				headX--
				tailX, tailY = testTail(tailX, tailY, headX, headY)
				visitedString := fmt.Sprintf("%d,%d", tailX, tailY)
				visited[visitedString] = 1

			} else {
				//right
				fmt.Println("R")
				headX++
				tailX, tailY = testTail(tailX, tailY, headX, headY)

				visitedString := fmt.Sprintf("%d,%d", tailX, tailY)
				fmt.Println(visitedString)
				visited[visitedString] = 1

			}
		}
		fmt.Println()

	}

	for k, _ := range visited {
		fmt.Println(k)
	}

	part_one_val = len(visited)

	fmt.Println("Part one: ", part_one_val)

	fmt.Println("Part two: ", part_two_val)

	dat.Close()
}
