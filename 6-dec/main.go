package main

import (
	"bufio"
	"fmt"
	"os"
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

func startOfPacket(l string) bool {

	if len(map[interface{}]int{l[0]: 0, l[1]: 0, l[2]: 0, l[3]: 0}) < 4 {
		return false
	}
	return true
}

func startOfMessage(l string) bool {

	if len(map[interface{}]int{l[0]: 0, l[1]: 0, l[2]: 0, l[3]: 0, l[4]: 0, l[5]: 0, l[6]: 0, l[7]: 0, l[8]: 0, l[9]: 0, l[10]: 0, l[11]: 0, l[12]: 0, l[13]: 0}) < 14 {
		return false
	}
	return true
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

	for fileScanner.Scan() {
		line := fileScanner.Text()

		// start of packet
		for i := 4; i < len(line); i++ {
			if startOfPacket(line[i-4 : i]) {
				part_one_val = i
				break
			}
		}

		//start of message
		for i := 14; i < len(line); i++ {
			if startOfMessage(line[i-14 : i]) {
				part_two_val = i
				break
			}
		}

	}

	fmt.Println("Part one: ", part_one_val)

	fmt.Println("Part two: ", part_two_val)

	dat.Close()
}
