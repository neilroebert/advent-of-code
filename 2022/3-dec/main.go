package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func round2(first string, second string, third string) int {

	part_one_sum := 0
	fmt.Println("Round 2")

	// Can't find a filter function in go, so doing it the hard way :joy:
	for _, char := range first {
		for _, char2 := range second {
			for _, char3 := range third {

				if (char == char2) && (char == char3) {
					fmt.Println(string(char))
					fmt.Println(first, " ", second, " ", third)

					if unicode.IsUpper(char) {
						return int(char) - (65 - 27)
					} else {
						return int(char) - (96)
					}

				}
			}

		}

	}
	fmt.Println()

	return part_one_sum
}

func round1(first string, second string) int {

	part_one_sum := 0
	for _, char := range first {
		for _, char2 := range second {
			if char == char2 {
				// fmt.Println(string(char))
				// fmt.Println(first, " ", second)

				if unicode.IsUpper(char) {
					return int(char) - (65 - 27)
				} else {
					return int(char) - (96)
				}
			}
		}
	}
	// fmt.Println()
	return part_one_sum
}

// Day 3 of learning go, getting there ;)
func main() {
	// dat, err := os.Open("test-input.txt")
	dat, err := os.Open("input.txt")
	check(err)

	fileScanner := bufio.NewScanner(dat)

	fileScanner.Split(bufio.ScanLines)

	var part_one_sum int = 0
	var part_two_sum int = 0

	s1 := ""
	s2 := ""

	for fileScanner.Scan() {
		line := fileScanner.Text()

		first := line[:len(line)/2]
		second := line[len(line)/2:]

		// fmt.Println(line, " ", first, " ", second)
		part_one_sum += round1(first, second)

		// fmt.Println()

		if s1 == "" {
			s1 = line
		} else if s2 == "" {
			s2 = line
		} else {
			part_two_sum += round2(s1, s2, line)
			s1 = ""
			s2 = ""
		}

	}

	fmt.Println("Part one sum: ", part_one_sum)
	fmt.Println("Part two sum: ", part_two_sum)
	dat.Close()
}
