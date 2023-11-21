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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func round1(first string, second string) int {

	elfVal1 := strings.Split(first, "-")
	elfVal2 := strings.Split(second, "-")

	min1, err := strconv.Atoi(elfVal1[0])
	check(err)
	min2, err := strconv.Atoi(elfVal2[0])
	check(err)
	min := min(min1, min2)

	max1, err := strconv.Atoi(elfVal1[1])
	check(err)
	max2, err := strconv.Atoi(elfVal2[1])
	check(err)
	max := max(max1, max2)

	if min1 == min && max1 == max {
		return 1
	} else if min2 == min && max2 == max {
		return 1
	} else {
		return 0
	}

}

func round2(first string, second string) int {

	elfVal1 := strings.Split(first, "-")
	elfVal2 := strings.Split(second, "-")

	min1, err := strconv.Atoi(elfVal1[0])
	check(err)
	min2, err := strconv.Atoi(elfVal2[0])
	check(err)

	max1, err := strconv.Atoi(elfVal1[1])
	check(err)
	max2, err := strconv.Atoi(elfVal2[1])
	check(err)

	if min1 <= min2 && min2 <= max1 {
		return 1
	} else if min2 <= min1 && min1 <= max2 {
		return 1
	} else {
		return 0
	}

}

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

		s := strings.Split(line, ",")
		s1 = s[0]
		s2 = s[1]

		// fmt.Println(line, " ", first, " ", second)
		part_one_sum += round1(s1, s2)
		part_two_sum += round2(s1, s2)

	}

	fmt.Println("Part one sum: ", part_one_sum)
	fmt.Println("Part two sum: ", part_two_sum)
	dat.Close()
}
