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
	lines := returnLines("inputs.txt")

	numRange := [2]int{-1, -1}
	for i, line := range lines {
		for j, char := range line {
			// why use anything smart when you can just match everything to my pet regex
			if re.Match([]byte(string(char))) {
				if j == 0 {
					numRange[0] = 0
				}
				if j != 0 && !re.Match([]byte(string(line[j-1]))) {
					numRange[0] = j
				}
				if j == len(line)-1 {
					numRange[1] = j
				}
				if j != len(line)-1 && !re.Match([]byte(string(line[j+1]))) {
					numRange[1] = j
				}
				if numRange[0] != -1 && numRange[1] != -1 {
					gotIt := false
					// we check eveything that isn't totally on edge left or right, top or bottom
					if numRange[0] != 0 && numRange[1] != len(line)-1 && i != 0 && i != len(lines)-1 {
						for k := numRange[0]; k <= numRange[1]; k++ {
							// check above and below in line
							if !numberOrDot.Match([]byte(string(lines[i+1][k]))) || !numberOrDot.Match([]byte(string(lines[i-1][k]))) {
								gotIt = true
							}
						}
						// check the right side and the left
						if !numberOrDot.Match([]byte(string(line[numRange[0]-1]))) || !numberOrDot.Match([]byte(string(lines[i+1][numRange[0]-1]))) || !numberOrDot.Match([]byte(string(lines[i-1][numRange[0]-1]))) || !numberOrDot.Match([]byte(string(line[numRange[1]+1]))) || !numberOrDot.Match([]byte(string(lines[i+1][numRange[1]+1]))) || !numberOrDot.Match([]byte(string(lines[i-1][numRange[1]+1]))) {
							gotIt = true
						}
					}
					// if not top or bot row check left
					if i != 0 && i != len(lines)-1 {
						if numRange[1] == len(line)-1 {
							if !numberOrDot.Match([]byte(string(line[numRange[0]-1]))) || !numberOrDot.Match([]byte(string(lines[i+1][numRange[0]-1]))) {
								gotIt = true
							}
						}
					}
					// if not top or bot row check right
					if i != 0 && i != len(lines)-1 {
						if numRange[0] == 0 {
							if !numberOrDot.Match([]byte(string(line[numRange[1]+1]))) || !numberOrDot.Match([]byte(string(lines[i+1][numRange[1]+1]))) {
								gotIt = true
							}
						}
					}
					// check the top row
					if i == 0 {
						if numRange != [2]int{0, len(line) - 1} {
							for k := numRange[0]; k <= numRange[1]; k++ {
								// check middle below
								if !numberOrDot.Match([]byte(string(lines[i+1][k]))) {
									gotIt = true
								}
							}
							// check left right below
							if numRange[0] != 0 && numRange[1] != len(line)-1 {
								if !numberOrDot.Match([]byte(string(lines[i+1][numRange[0]-1]))) || !numberOrDot.Match([]byte(string(lines[i+1][numRange[1]+1]))) {
									gotIt = true
								}
							}
							// check the right side top row
							if numRange[0] == 0 {
								if !numberOrDot.Match([]byte(string(line[numRange[1]+1]))) || !numberOrDot.Match([]byte(string(lines[i+1][numRange[1]+1]))) {
									gotIt = true
								}
							}
							// check the left side top row
							if numRange[1] == len(line)-1 {
								if !numberOrDot.Match([]byte(string(line[numRange[0]-1]))) || !numberOrDot.Match([]byte(string(lines[i+1][numRange[0]-1]))) {
									gotIt = true
								}
							}
						}
					}
					// check the bottom row
					if i == len(lines)-1 {
						if numRange != [2]int{0, len(line) - 1} {
							for k := numRange[0]; k <= numRange[1]; k++ {
								// check middle below
								if !numberOrDot.Match([]byte(string(lines[i-1][k]))) {
									gotIt = true
								}
							}
							// check left right above
							if numRange[0] != 0 && numRange[1] != len(line)-1 {
								if !numberOrDot.Match([]byte(string(lines[i-1][numRange[0]-1]))) || !numberOrDot.Match([]byte(string(lines[i-1][numRange[1]+1]))) {
									gotIt = true
								}
							}
							// check the right side bot row
							if numRange[0] == 0 {
								if !numberOrDot.Match([]byte(string(line[numRange[1]+1]))) || !numberOrDot.Match([]byte(string(lines[i-1][numRange[1]+1]))) {
									gotIt = true
								}
							}
							// check the left side bot row
							if numRange[1] == len(line)-1 {
								if !numberOrDot.Match([]byte(string(line[numRange[0]-1]))) || !numberOrDot.Match([]byte(string(lines[i-1][numRange[0]-1]))) {
									gotIt = true
								}
							}
						}
					}
					if gotIt {
						fmt.Println(makeNumber(line, numRange[:]))
						sum = sum + makeNumber(line, numRange[:])
					}
					numRange = [2]int{-1, -1}
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
	emptyRow := []string{}
	for i := 0; i < len(lines[0]); i++ {
		emptyRow = append(emptyRow, ".")
	}
	lines = append(emptyRow, lines...)
	lines = append(lines, emptyRow...)
	return lines
}
