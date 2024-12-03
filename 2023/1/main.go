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

func strip_part_one(str string) int {
	var left string = ""
	var right string = ""
	// fmt.Println(str)
	for _, c := range str {
		// fmt.Println(c)
		if '0' <= c && c <= '9' {
			// fmt.Println(c)
			left = string(c)
			break
		}
		continue
	}
	for i := len(str) - 1; i >= 0; i-- {
		if '0' <= str[i] && str[i] <= '9' {
			right = string(str[i])
			break
		}
		continue
	}
	// fmt.Println(left + right)
	tmp, err := strconv.Atoi(left + right)
	check(err)

	return tmp
}

func strip_part_two(str string) int {
	numeric_chars := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	fmt.Println("----")

	// var left int = 0
	// var right int = 0

	var strlen = len(str)

	var left_num = 0
	var left_num_ind = strlen
	var right_num = 0
	var right_num_ind = -1

	var left_char = 0
	var left_char_ind = strlen
	var right_char = 0
	var right_char_ind = -1

	// fmt.Println(str)
	//find numeric indeces
	for i, c := range str {
		// fmt.Println(c)
		if '0' <= c && c <= '9' {
			// fmt.Println(c)
			left_num_ind = i
			tmp, err := strconv.Atoi(string(c))
			check(err)
			left_num = tmp
			break
		}
		continue
	}
	for i := len(str) - 1; i >= 0; i-- {
		if '0' <= str[i] && str[i] <= '9' {
			right_num_ind = i
			tmp, err := strconv.Atoi(string(str[i]))
			check(err)
			right_num = tmp
			break
		}
		continue
	}

	//find char indeces
	for k, _ := range str {
		// fmt.Println(str[k:])
		var found = false
		for i := 0; i < len(numeric_chars); i++ {

			if strings.HasPrefix(str[k:], numeric_chars[i]) {
				found = true
				left_char = i
				left_char_ind = k
				break
			}
		}
		if found {
			break
		} else {
			continue
		}
	}
	for k := len(str) - 1; k >= 0; k-- {
		// fmt.Println(str[k:])
		var found = false
		for i := 0; i < len(numeric_chars); i++ {

			if strings.HasPrefix(str[k:], numeric_chars[i]) {
				found = true
				right_char = i
				right_char_ind = k
				break
			}
		}
		if found {
			break
		} else {
			continue
		}
	}

	if left_char_ind < left_num_ind {
		if right_char_ind < right_num_ind {
			var tmp = left_char*10 + right_num
			fmt.Println(tmp)
			return tmp
		} else {
			var tmp = left_char*10 + right_char
			fmt.Println(tmp)
			return tmp
		}

	} else {
		if right_char_ind < right_num_ind {
			var tmp = left_num*10 + right_num
			fmt.Println(tmp)
			return tmp
		} else {
			var tmp = left_num*10 + right_char
			fmt.Println(tmp)
			return tmp
		}
	}

	// tmp, err := strconv.Atoi(left + right)
	// check(err)

	// return tmp
}

func main() {
	dat, err := os.Open("input.txt")
	check(err)

	fileScanner := bufio.NewScanner(dat)

	fileScanner.Split(bufio.ScanLines)

	var part_one_total int = 0
	var part_two_total int = 0
	// var current int = 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line != "" {
			fmt.Println(line)
			// part_one_total = part_one_total + strip_part_one(line)
			part_two_total = part_two_total + strip_part_two(line)

		} else {
			continue
		}

	}

	fmt.Println("Part one total: ", part_one_total)
	fmt.Println("Part two total: ", part_two_total)
	dat.Close()
}
