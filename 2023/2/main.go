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

func is_game_possible(picks []string) bool {
	// fmt.Println(picks)
	for i := 0; i < len(picks); i = i + 2 {
		cube_num, err := strconv.Atoi(picks[i])
		check(err)
		if cube_num > 12 && strings.HasPrefix(picks[i+1], "red") {
			return false
		} else if cube_num > 13 && strings.HasPrefix(picks[i+1], "green") {
			return false
		} else if cube_num > 14 && strings.HasPrefix(picks[i+1], "blue") {
			return false
		}
	}
	return true
}

func power_of_set(picks []string) int {
	var blue, red, green = 0, 0, 0
	for i := 0; i < len(picks); i = i + 2 {
		cube_num, err := strconv.Atoi(picks[i])
		check(err)
		if strings.HasPrefix(picks[i+1], "red") {
			if red < cube_num {
				red = cube_num
			}
		} else if strings.HasPrefix(picks[i+1], "green") {
			if green < cube_num {
				green = cube_num
			}
		} else if strings.HasPrefix(picks[i+1], "blue") {
			if blue < cube_num {
				blue = cube_num
			}
		}
	}
	return blue * red * green

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
			line_split := strings.Split(line, " ")
			// fmt.Println(line_split[0])
			game_num_text := strings.TrimSuffix(line_split[1], ":")
			game_num, err := strconv.Atoi(string(game_num_text))
			check(err)
			if is_game_possible(line_split[2:]) {
				fmt.Println("Game is possible")
				part_one_total = part_one_total + game_num
			}

			part_two_total = part_two_total + power_of_set(line_split[2:])

		} else {
			continue
		}

	}

	fmt.Println("Part one total: ", part_one_total)
	fmt.Println("Part two total: ", part_two_total)
	dat.Close()
}
