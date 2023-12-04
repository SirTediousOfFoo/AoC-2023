package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type touching struct {
	up    []int
	down  []int
	left  []int
	right []int
}

func main() {
	sum := 0
	// i love this little guy
	re := regexp.MustCompile("([0-9])")
	gear := regexp.MustCompile("(\\*)")
	lines := initEmpty(returnLines("inputs.txt"))
	for _, line := range lines {
		fmt.Println(line)
	}
	for i, line := range lines {
		for j, char := range line {
			var touching touching
			// why use anything smart when you can just match everything to my pet regex
			if gear.Match([]byte(string(char))) {
				// we check eveything that isn't totally on edge left or right, top or bottom
				for k := 0; k <= 2; k++ {
					// check above and below in line
					if re.Match([]byte(string(lines[i+1][k]))) {
						touching.up = append(touching.up, k)
						fmt.Println("up", string(lines[i+1][k]))
					}
					if re.Match([]byte(string(lines[i-1][k]))) {
						touching.down = append(touching.down, k)
						fmt.Println("down", string(lines[i-1][k]))
					}
				}
				// check the right side and the left
				if re.Match([]byte(string(line[j-1]))) {
					touching.left = append(touching.left, j)
					fmt.Println("left", string(line[j-1]))
				}
				if re.Match([]byte(string(line[j+1]))) {
					fmt.Println("right", string(line[j-1]))

					touching.right = append(touching.right, j)
				}
				if len(touching.down) > 0 || len(touching.up) > 0 || len(touching.left) > 0 || len(touching.right) > 0 {
					fmt.Println(touching)
					touching = touching
				}
			}
		}
	}
	fmt.Print(sum)
}

func returnLines(inputFile string) []string {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

func makeNumber(line string, numRange []int) int {
	strNum := ""
	for i := numRange[0]; i <= numRange[1]; i++ {
		strNum = strNum + string(line[i])
	}
	numNum, _ := strconv.Atoi(strNum)

	return numNum
}

func initEmpty(lines []string) []string {
	for i := range lines {
		lines[i] = "." + lines[i] + "."
	}
	emptyRow := ""
	for i := 0; i < len(lines[0]); i++ {
		emptyRow = emptyRow + "."
	}
	lines = append([]string{emptyRow}, lines...)
	lines = append(lines, []string{emptyRow}...)
	return lines
}
