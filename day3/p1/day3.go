package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	// i love this little guy
	re := regexp.MustCompile("([0-9])")
	lines := returnLines("inputs.txt")
	//var numIndex [][]int
	for i, line := range lines {
		if i != 0 {
			for j, char := range line {
				// why use anything smart when you can just match everything to my pet regex
				if re.Match([]byte(string(char))) {
					// okay so we find a number, record its index and then
					// keep going while the next char is also a number.
					// that gives us the start and end indices of our number
					// after that loop from start to end ID in line before and after
					// also check same line on id-1 and id+1 and if any of those are
					// a symbol add the number to a sum
					fmt.Println(string(char), j)
				}
			}
		}
	}
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

func lookAround(lineBefore, line, lineAfter string) int {

	return 0
}
