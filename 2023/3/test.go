package main

import (
	"fmt"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func test_line(x int, line string) ([]int, bool) {
	// make a zero size array
	ret_arr := make([]int, 0)

	if !('0' <= line[x] && line[x] <= '9') {
		left_val, is_num := test_left(x, line)
		if is_num {
			ret_arr = append(ret_arr, left_val)
		}
		right_val, is_num := test_right(x, line)
		if is_num {
			ret_arr = append(ret_arr, right_val)
		}

	} else {
		//step left through string until end of line is found or a non numeric char

	}

	return ret_arr, false

}

func test_left(x int, line string) (int, bool) {
	test_index := x - 1
	// fmt.Println(test_index)
	if test_index >= 0 {
		if line[test_index] == '.' {
			return 0, false
		} else if '0' <= line[test_index] && line[test_index] <= '9' {
			// fmt.Println(line[:test_index])
			val := strings.Split(line[:test_index+1], ".")
			tmp_val, err := strconv.Atoi(val[len(val)-1])
			check(err)
			return tmp_val, true
		}
	}
	return 0, false
}

func test_right(x int, line string) (int, bool) {
	test_index := x + 1
	// fmt.Println(test_index)
	if test_index <= len(line)-1 {
		if line[test_index] == '.' {
			return 0, false
		} else if '0' <= line[test_index] && line[test_index] <= '9' {
			// fmt.Println(line[test_index:])
			val := strings.Split(line[test_index:], ".")
			tmp_val, err := strconv.Atoi(val[0])
			check(err)
			return tmp_val, true
		}
	}
	return 0, false
}

func main() {
	var line = "467..114.."
	// val := strings.Split(line[:8], ".")
	// fmt.Println(val)
	// tmp_val, err := strconv.Atoi(val[len(val)-1])
	// check(err)
	fmt.Println(line)
	// for i, val := range val {
	// 	tmp_val, err := strconv.Atoi(val)
	// 	check(err)
	// 	fmt.Println(i, tmp_val)
	// }
	for i := 0; i < 10; i++ {
		// fmt.Println(line[:i], "<")
		left_val, isNum := test_left(i, line)
		fmt.Println(i, left_val, isNum)
		right_val, isNum := test_right(i, line)
		fmt.Println(i, right_val, isNum)
		fmt.Println("-----------")
	}

}
