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
	currentDir := "/"
	tree := make(map[string]int)
	tree["/"] = 0

	for fileScanner.Scan() {
		line := fileScanner.Text()

		if strings.HasPrefix(line, "$") {
			// command
			// fmt.Println("Old: ", currentDir)
			// fmt.Println("cmd: ", line)
			if strings.Contains(line, "cd") {
				//cd commands
				if strings.Contains(line, "..") {
					// remove last dir from currentdir
					if currentDir != "/" {
						// fmt.Println(strings.LastIndex(currentDir, "/"))
						// fmt.Println(currentDir)
						currentDir = currentDir[0:strings.LastIndex(currentDir, "/")]
						currentDir = currentDir[0 : strings.LastIndex(currentDir, "/")+1]
						// fmt.Println(currentDir)
						// fmt.Println()
					}
				} else if strings.Contains(line, "/") {
					// cd to base
					currentDir = "/"
				} else {
					//
					currentDir += strings.Split(line, " ")[2] + "/"
				}
			}
			// fmt.Println("New: ", currentDir)
			// fmt.Println()

		} else if strings.HasPrefix(line, "dir") {
			// dir listing
			f := strings.Split(line, " ")
			tree[currentDir+f[1]+"/"] = 0
		} else {
			// file listing
			f := strings.Split(line, " ")
			size, err := strconv.Atoi(f[0])
			check(err)
			tree[currentDir+f[1]] = size

		}

	}

	for k, v := range tree {

		// fmt.Println(k, "size: ", v)
		if strings.HasSuffix(k, "/") {
			//in a dir
		} else {
			// file

			k = k[0 : strings.LastIndex(k, "/")+1]
			// fmt.Println("Folder access: ", k)
			tree[k] += v

			for strings.LastIndex(k, "/") > 0 {
				// fmt.Println(k)
				k = k[0:strings.LastIndex(k, "/")]
				k = k[0 : strings.LastIndex(k, "/")+1]
				// fmt.Println("Folder access: ", k)
				tree[k] += v

			}

		}
		// fmt.Println()
	}

	requiredSize := 30000000 - (70000000 - tree["/"])
	// fmt.Println(requiredSize)

	part_two_val = 70000000

	for k, v := range tree {

		if strings.HasSuffix(k, "/") {
			//in a dir
			if v <= 100000 {
				part_one_val += v
			}
			if v >= requiredSize {
				if v < part_two_val {
					part_two_val = v
				}
				// fmt.Println(k, "size: ", v)
			}

		}
	}

	fmt.Println("Part one: ", part_one_val)

	fmt.Println("Part two: ", part_two_val)

	dat.Close()
}
