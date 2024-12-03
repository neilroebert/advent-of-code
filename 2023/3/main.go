package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func slice_forward(x int, line []string) int {
	
}

func build_num(x int, y int, last_col int, lines [][]string) int {
	var num string = ""
	if x == 0 {
		// 		if first col - step forward
		tmp_x := x
		tmp_y := y
		for lines[y][tmp_x] != "." {

			tmp_x = tmp_x + 1
			if tmp_x > last_col {
				break
			}
		}
		ifunicode.IsDigit(rune(lines[y][x+1][0])) {
			return true
		}
		if !(unicode.IsDigit(rune(lines[y+1][x+1][0])) || lines[y+1][x+1] == ".") {
			return true
		}
		if !(unicode.IsDigit(rune(lines[y+1][x][0])) || lines[y+1][x] == ".") {
			return true
		}

	} else if x == last_col {
		//		if last col - step backward
		if !(unicode.IsDigit(rune(lines[y+1][x][0])) || lines[y+1][x] == ".") {
			//b
			return true
		}
		if !(unicode.IsDigit(rune(lines[y+1][x-1][0])) || lines[y+1][x-1] == ".") {
			//b d l
			return true
		}
		if !(unicode.IsDigit(rune(lines[y][x-1][0])) || lines[y][x-1] == ".") {
			//l
			return true
		}
	} else {
		// 		else middle col - check x-1, x, x+1 - check left and right and forward and backwards
		if !(unicode.IsDigit(rune(lines[y][x+1][0])) || lines[y][x+1] == ".") {
			return true
		}
		if !(unicode.IsDigit(rune(lines[y+1][x+1][0])) || lines[y+1][x+1] == ".") {
			return true
		}
		if !(unicode.IsDigit(rune(lines[y+1][x][0])) || lines[y+1][x] == ".") {
			return true
		}
		if !(unicode.IsDigit(rune(lines[y+1][x-1][0])) || lines[y+1][x-1] == ".") {
			return true
		}
		if !(unicode.IsDigit(rune(lines[y][x-1][0])) || lines[y][x-1] == ".") {
			return true
		}
	}
}

func find_symbol(x int, y int, last_col int, lines [][]string) bool {
	last_row := len(lines) - 1
	// fmt.Println(x, y, last_col, last_row)
	if y == 0 {
		// if first row - check y, y+1
		if x == 0 {
			// 		if first col - check x, x+1
			if !(unicode.IsDigit(rune(lines[y][x+1][0])) || lines[y][x+1] == ".") {
				return true
			}
			if !(unicode.IsDigit(rune(lines[y+1][x+1][0])) || lines[y+1][x+1] == ".") {
				return true
			}
			if !(unicode.IsDigit(rune(lines[y+1][x][0])) || lines[y+1][x] == ".") {
				return true
			}

		} else if x == last_col {
			// 		if last col - check x, x-1
			if !(unicode.IsDigit(rune(lines[y+1][x][0])) || lines[y+1][x] == ".") {
				//b
				return true
			}
			if !(unicode.IsDigit(rune(lines[y+1][x-1][0])) || lines[y+1][x-1] == ".") {
				//b d l
				return true
			}
			if !(unicode.IsDigit(rune(lines[y][x-1][0])) || lines[y][x-1] == ".") {
				//l
				return true
			}
		} else {
			// 		else middle col - check x-1, x, x+1
			if !(unicode.IsDigit(rune(lines[y][x+1][0])) || lines[y][x+1] == ".") {
				return true
			}
			if !(unicode.IsDigit(rune(lines[y+1][x+1][0])) || lines[y+1][x+1] == ".") {
				return true
			}
			if !(unicode.IsDigit(rune(lines[y+1][x][0])) || lines[y+1][x] == ".") {
				return true
			}
			if !(unicode.IsDigit(rune(lines[y+1][x-1][0])) || lines[y+1][x-1] == ".") {
				return true
			}
			if !(unicode.IsDigit(rune(lines[y][x-1][0])) || lines[y][x-1] == ".") {
				return true
			}
		}
	} else if y == last_row {
		// if last row - check y, y-1
		if x == 0 {
			// 		if first col - check x, x+1
			if !(unicode.IsDigit(rune(lines[y-1][x][0])) || lines[y-1][x] == ".") {
				//a
				return true
			}
			if !(unicode.IsDigit(rune(lines[y-1][x+1][0])) || lines[y-1][x+1] == ".") {
				//a d r
				return true
			}
			if !(unicode.IsDigit(rune(lines[y][x+1][0])) || lines[y][x+1] == ".") {
				//r
				return true
			}
		} else if x == last_col {
			// 		if last col - check x, x-1
			if !(unicode.IsDigit(rune(lines[y][x-1][0])) || lines[y][x-1] == ".") {
				//l
				return true
			}
			if !(unicode.IsDigit(rune(lines[y-1][x-1][0])) || lines[y-1][x-1] == ".") {
				//a d l
				return true
			}
			if !(unicode.IsDigit(rune(lines[y-1][x][0])) || lines[y-1][x] == ".") {
				//a
				return true
			}
		} else {
			// 		else middle col - check x-1, x, x+1
			if !(unicode.IsDigit(rune(lines[y][x-1][0])) || lines[y][x-1] == ".") {
				return true
			}
			if !(unicode.IsDigit(rune(lines[y-1][x-1][0])) || lines[y-1][x-1] == ".") {
				return true
			}
			if !(unicode.IsDigit(rune(lines[y-1][x][0])) || lines[y-1][x] == ".") {
				return true
			}
			if !(unicode.IsDigit(rune(lines[y-1][x+1][0])) || lines[y-1][x+1] == ".") {
				return true
			}
			if !(unicode.IsDigit(rune(lines[y][x+1][0])) || lines[y][x+1] == ".") {
				return true
			}
		}

	} else {
		// else middle rows check y-1, y, y+1
		if x == 0 {
			// 		if first col
			if !(unicode.IsDigit(rune(lines[y-1][x][0])) || lines[y-1][x] == ".") {
				//a
				return true
			}
			if !(unicode.IsDigit(rune(lines[y-1][x+1][0])) || lines[y-1][x+1] == ".") {
				//a d r
				return true
			}
			if !(unicode.IsDigit(rune(lines[y][x+1][0])) || lines[y][x+1] == ".") {
				//r
				return true
			}
			if !(unicode.IsDigit(rune(lines[y+1][x+1][0])) || lines[y+1][x+1] == ".") {
				//b d r
				return true
			}
			if !(unicode.IsDigit(rune(lines[y+1][x][0])) || lines[y+1][x] == ".") {
				//b
				return true
			}
		} else if x == last_col {
			// 		if last col
			if !(unicode.IsDigit(rune(lines[y+1][x][0])) || lines[y+1][x] == ".") {
				//b
				return true
			}
			if !(unicode.IsDigit(rune(lines[y+1][x-1][0])) || lines[y+1][x-1] == ".") {
				//b d l
				return true
			}
			if !(unicode.IsDigit(rune(lines[y][x-1][0])) || lines[y][x-1] == ".") {
				//l
				return true
			}
			if !(unicode.IsDigit(rune(lines[y-1][x-1][0])) || lines[y-1][x-1] == ".") {
				//a d l
				return true
			}
			if !(unicode.IsDigit(rune(lines[y-1][x][0])) || lines[y-1][x] == ".") {
				//a
				return true
			}
		} else {
			// 		else middle col
			if !(unicode.IsDigit(rune(lines[y-1][x][0])) || lines[y-1][x] == ".") {
				//a
				return true
			}
			if !(unicode.IsDigit(rune(lines[y-1][x+1][0])) || lines[y-1][x+1] == ".") {
				//a d r
				return true
			}
			if !(unicode.IsDigit(rune(lines[y][x+1][0])) || lines[y][x+1] == ".") {
				//r
				return true
			}
			if !(unicode.IsDigit(rune(lines[y+1][x+1][0])) || lines[y+1][x+1] == ".") {
				//b d r
				return true
			}
			if !(unicode.IsDigit(rune(lines[y+1][x][0])) || lines[y+1][x] == ".") {
				//b
				return true
			}
			if !(unicode.IsDigit(rune(lines[y+1][x-1][0])) || lines[y+1][x-1] == ".") {
				//b d l
				return true
			}
			if !(unicode.IsDigit(rune(lines[y][x-1][0])) || lines[y][x-1] == ".") {
				//l
				return true
			}
			if !(unicode.IsDigit(rune(lines[y-1][x-1][0])) || lines[y-1][x-1] == ".") {
				//a d l
				return true
			}
		}
	}

	return false
}

func part_one(lines [][]string) int {
	var num_list []int
	for i, line := range lines {
		fmt.Println(i, line)
		last_col := len(line) - 1
		// fmt.Println(lines[0])
		var touches_symbol = false
		var building_num bool = false
		var num string = ""
		for j, char := range line {

			if unicode.IsDigit(rune(char[0])) {
				fmt.Println("IsDigit", char)
				building_num = true
				tmp_num := num + char
				num = tmp_num
				fmt.Println("Num:", num)
				touches_symbol = touches_symbol || find_symbol(j, i, last_col, lines)
				fmt.Println(num, touches_symbol)
			} else {
				fmt.Println("NotDigit", char)
				if building_num && touches_symbol {
					fmt.Println("Building num")
					tmp_val, err := strconv.Atoi(num)
					check(err)
					num_list = append(num_list, tmp_val)
					touches_symbol = false
				}
				num = ""
				building_num = false

			}
			if last_col == j {
				fmt.Println("Last Col")
				if building_num && touches_symbol {
					tmp_val, err := strconv.Atoi(num)
					check(err)
					num_list = append(num_list, tmp_val)
				}
				touches_symbol = false
				num = ""
				building_num = false
			}
		}
	}
	var sum = 0
	for _, num := range num_list {
		sum = sum + num
	}

	fmt.Println(num_list)

	return sum
}

func main() {
	dat, err := os.Open("input.txt")
	check(err)

	fileScanner := bufio.NewScanner(dat)

	fileScanner.Split(bufio.ScanLines)

	var lines [][]string

	var part_one_total int = 0
	var part_two_total int = 0
	// var current int = 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line != "" {
			// fmt.Println(line)
			lines = append(lines, strings.Split(line, ""))
			lines_part_two = append(lines, line)

		} else {
			continue
		}

	}
	part_one_total = part_one(lines)
	part_two_total = part_two(lines_part_two)
	// fmt.Println(lines)

	// for j, line := range lines {
	// 	fmt.Println(j, line)
	// }

	fmt.Println("Part one total: ", part_one_total)
	fmt.Println("Part two total: ", part_two_total)
	dat.Close()
}
