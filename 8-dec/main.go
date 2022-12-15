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

func main() {
	// dat, err := os.Open("test-input.txt")
	dat, err := os.Open("input.txt")
	check(err)

	fileScanner := bufio.NewScanner(dat)

	fileScanner.Split(bufio.ScanLines)

	// debug := true

	part_one_val := 0
	part_two_val := 0
	treeH := make([][]int, 0)
	treeV := make([][]int, 0)
	i := 0
	g := 0

	for fileScanner.Scan() {

		line := fileScanner.Text()
		g = len(line)

		treeH = append(treeH, make([]int, g))
		tmpArr := make([]int, g)
		treeV = append(treeV, make([]int, 0))
		j := 0
		for _, char := range line {
			treeH[i][j] = int(char) - 48
			tmpArr[j] = int(char) - 48
			j++
		}
		treeV = append(treeV, tmpArr)
		i++

	}

	for _, l := range treeH {
		fmt.Printf("%v", l)
	}
	fmt.Println()
	for _, l := range treeV {
		fmt.Printf("%v", l)
	}

	// this is the unefficient way of doing it :(
	for j := 1; j < i-1; j++ {
	INNER:
		for k := 1; k < g-1; k++ {
			testVal := treeH[j][k]
			fmt.Println("Test val: ", testVal)
			visH := true
			visV := true

			//vert
			for t := 0; t < i; t++ {
				// fmt.Println("Testing: [", t, "][", k, "]")
				if t == j {
					//same as test tree
					if visV {
						fmt.Println("Vis from Top")
						part_one_val++
						continue INNER
					}
					fmt.Println("Not vis from Top")
					visV = true
					continue
				}
				if treeH[t][k] >= testVal {
					visV = false
				}

			}
			if visV {
				fmt.Println("Vis from Bot")
				part_one_val++
				continue INNER
			}
			fmt.Println("Not vis from Bot")
			// fmt.Println()
			//hor
			for y := 0; y < g; y++ {
				// fmt.Println("Testing: [", j, "][", y, "]")
				if y == k {
					//same as test tree
					if visH {
						fmt.Println("Vis from Left")
						part_one_val++
						continue INNER
					}
					fmt.Println("Not vis from Left")
					visH = true
					continue
				}
				if treeH[j][y] >= testVal {
					visH = false
				}

			}
			if visH {
				fmt.Println("Vis from Right")
				part_one_val++
				continue INNER
			}
			fmt.Println("Not vis from Right")
			fmt.Println("-------------")
		}
	}
	fmt.Println("============")
	// part 2
	for j := 1; j < i-1; j++ {

		for k := 1; k < g-1; k++ {
			testVal := treeH[j][k]
			fmt.Println("Test val: ", testVal)
			left := 0
			right := 0
			top := 0
			bot := 0
			//to top
			for t := j - 1; t >= 0; t-- {
				fmt.Println("Testing: [", t, "][", k, "]")
				if treeH[t][k] >= testVal {
					top++
					break
				}
				top++
			}
			//to bottom
			fmt.Println("To Bottom")
			for t := j + 1; t < i; t++ {
				fmt.Println("Testing: [", t, "][", k, "]")
				if treeH[t][k] >= testVal {
					bot++
					break
				}
				bot++
			}

			//to left
			fmt.Println("To left")
			for y := k - 1; y >= 0; y-- {
				fmt.Println("Testing: [", j, "][", y, "]")
				if treeH[j][y] >= testVal {
					left++
					break
				}
				left++

			}

			//to right
			fmt.Println("To right")
			for y := k + 1; y < g; y++ {
				fmt.Println("Testing: [", j, "][", y, "]")
				if treeH[j][y] >= testVal {
					right++
					break
				}
				right++

			}

			score := left * top * right * bot
			if score >= part_two_val {
				part_two_val = score
			}

		}
	}

	fmt.Println(i + i + (g - 2) + (g - 2))
	fmt.Println(part_one_val)

	part_one_val += i + i + (g - 2) + (g - 2)

	fmt.Println("Part one: ", part_one_val)

	fmt.Println("Part two: ", part_two_val)

	dat.Close()
}
