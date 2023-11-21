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

func round1(ec string, mc string) int {

	part_one_sum := 0
	elf := map[string]int{"A": 1, "B": 2, "C": 3}
	shapes := map[string]int{"X": 1, "Y": 2, "Z": 3}

	if elf[ec] < shapes[mc] {
		if elf[ec] == 1 && shapes[mc] == 3 {
			//Loss
			part_one_sum += 0 + shapes[mc]
			// fmt.Println("loss: ", line, 0 + shapes[mc])
		} else {
			//Win
			part_one_sum += 6 + shapes[mc]
			// fmt.Println("win:  ", line, 6 + shapes[mc])
		}

	} else if elf[ec] == shapes[mc] {
		// Draw Same
		part_one_sum += 3 + shapes[mc]
		// fmt.Println("draw: ", line, 3 + shapes[mc])
	} else {
		if elf[ec] == 3 && shapes[mc] == 1 {
			//Win
			part_one_sum += 6 + shapes[mc]
			// fmt.Println("win:  ", line, 6 + shapes[mc])
		} else {
			//loss
			part_one_sum += 0 + shapes[mc]
			// fmt.Println("loss: ", line, 0 + shapes[mc])
		}

	}
	return part_one_sum
}

func round2(ec string, mc string) int {

	part_two_sum := 0
	win := map[string]int{"A": 2, "B": 3, "C": 1}
	draw := map[string]int{"A": 1, "B": 2, "C": 3}
	loss := map[string]int{"A": 3, "B": 1, "C": 2}

	if mc == "X" {
		//loss
		part_two_sum += loss[ec]
		// fmt.Println("Loss: ",ec," ",mc," ", loss[ec])
	} else if mc == "Y" {
		//draw
		part_two_sum += draw[ec] + 3
		// fmt.Println("Draw: ",ec," ",mc," ", draw[ec])
	} else {
		//win
		part_two_sum += win[ec] + 6
		// fmt.Println("Win:  ",ec," ",mc," ", win[ec])
	}
	return part_two_sum

}

// Day 2 of learning go, so no judgement ;)
func main() {
	dat, err := os.Open("input.txt")
	check(err)

	fileScanner := bufio.NewScanner(dat)

	fileScanner.Split(bufio.ScanLines)

	var part_one_sum int = 0
	var part_two_sum int = 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		//elf choice
		ec := string(line[0])
		//my choice
		mc := string(line[2])

		part_one_sum += round1(ec, mc)
		part_two_sum += round2(ec, mc)

	}

	fmt.Println("Part one sum: ", part_one_sum)
	fmt.Println("Part two sum: ", part_two_sum)
	dat.Close()
}
