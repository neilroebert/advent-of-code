package main

import (
    "bufio"
    "fmt"
    "os"
	"strconv"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func testArray(test [3]int, val int) [3]int {
	if val > test[2] {
		if val >= test[1]{
			if val >= test[0]{
				test[2] = test[1]
				test[1] = test[0]
				test[0] = val

			} else {
				test[2] = test[1]
				test[1] = val
			}
		} else {
			test[2] = val
		}
	}

	return test
}
// This is my journey learning go, so no judgement ;)
func main() {
	dat, err := os.Open("input.txt")
	check(err)

	fileScanner := bufio.NewScanner(dat)
 
    fileScanner.Split(bufio.ScanLines)
  
	var part_one_highest int = 0
	var part_two_highest [3]int
	var current int = 0
    for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == "" {
			if current > part_one_highest {
				part_one_highest = current
			}
			part_two_highest = testArray(part_two_highest, current)
			current = 0
			
		} else {
			tmp, err := strconv.Atoi(line)
			check(err)

			current += tmp
		}
        
    }
  
	fmt.Println("Part one highest: ", part_one_highest)
	fmt.Println("Part two sum of highest: ", (part_two_highest[0] + part_two_highest[1] + part_two_highest[2]))
    dat.Close()
}