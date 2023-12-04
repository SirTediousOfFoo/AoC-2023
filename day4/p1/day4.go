package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"slices"
	"strings"
)

func main() {
	sum := 0
	// i love this little guy
	re := regexp.MustCompile("([0-9]+)")
	lines := returnLines("inputs.txt")
	for _, line := range lines {
		split := strings.Split(line, "|")[:]
		a := re.FindAllString(split[0], -1)
		a = a[1:]
		fmt.Println("a", a)
		b := strings.Split(split[1], " ")
		power := -1.0
		for _, element := range a {
			if slices.Contains(b, element) {
				power = power + 1
			}
		}
		if power >= 0.0 {
			sum = sum + int(math.Pow(2, float64(power)))
		}
	}
	fmt.Println(sum)
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
