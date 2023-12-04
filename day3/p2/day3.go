package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type gear struct {
	x int
	y int
}

func main() {
	sum := 0
	fuckle := make(map[gear][]int)
	// i love this little guy
	re := regexp.MustCompile("([0-9])")
	gearex := regexp.MustCompile("\\*")
	lines := initEmpty(returnLines("inputs.txt"))
	for _, line := range lines {
		fmt.Println(line)
	}
	numRange := [2]int{-1, -1}
	for i, line := range lines {
		for j, char := range line {
			// why use anything smart when you can just match everything to my pet regex
			if re.Match([]byte(string(char))) && !re.Match([]byte(string(line[j-1]))) {
				numRange[0] = j
			}
			if re.Match([]byte(string(char))) && !re.Match([]byte(string(line[j+1]))) {
				numRange[1] = j
			}
			if numRange[0] != -1 && numRange[1] != -1 {
				// we check eveything that isn't totally on edge left or right, top or bottom
				for k := numRange[0] - 1; k <= numRange[1]+1; k++ {
					// check above and below in line
					if gearex.Match([]byte(string(lines[i+1][k]))) {
						fuckle[gear{x: i + 1, y: k}] = append(fuckle[gear{x: i + 1, y: k}], makeNumber(line, numRange[:]))
					}
					if gearex.Match([]byte(string(lines[i-1][k]))) {
						fmt.Println("up", string(lines[i-1][k]), "number", makeNumber(line, numRange[:]), "x", i-1, "y", k)
						fuckle[gear{x: i - 1, y: k}] = append(fuckle[gear{x: i - 1, y: k}], makeNumber(line, numRange[:]))
					}
				}
				// check the right side and the left
				if gearex.Match([]byte(string(line[numRange[0]-1]))) {
					fuckle[gear{x: i, y: numRange[0] - 1}] = append(fuckle[gear{x: i, y: numRange[0] - 1}], makeNumber(line, numRange[:]))
				}
				if gearex.Match([]byte(string(line[numRange[1]+1]))) {
					fuckle[gear{x: i, y: numRange[1] + 1}] = append(fuckle[gear{x: i, y: numRange[1] + 1}], makeNumber(line, numRange[:]))
				}

				numRange = [2]int{-1, -1}
			}
		}
	}

	for _, value := range fuckle {
		if len(value) == 2 {
			sum = sum + value[0]*value[1]
		}
	}
	fmt.Println(fuckle)
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
