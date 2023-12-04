package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	sum := 0
	// i love this little guy
	re := regexp.MustCompile("([0-9])")
	numberOrDot := regexp.MustCompile("([0-9]|\\.)")
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
				gotIt := false
				// we check eveything that isn't totally on edge left or right, top or bottom
				for k := numRange[0] - 1; k <= numRange[1]+1; k++ {
					// check above and below in line
					if !numberOrDot.Match([]byte(string(lines[i+1][k]))) || !numberOrDot.Match([]byte(string(lines[i-1][k]))) {
						gotIt = true
					}
				}
				// check the right side and the left
				if !numberOrDot.Match([]byte(string(line[numRange[0]-1]))) || !numberOrDot.Match([]byte(string(line[numRange[1]+1]))) {
					gotIt = true
				}
				if gotIt {
					fmt.Println(makeNumber(line, numRange[:]))
					sum = sum + makeNumber(line, numRange[:])
				}
				numRange = [2]int{-1, -1}
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
