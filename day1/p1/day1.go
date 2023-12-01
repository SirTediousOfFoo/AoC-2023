package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
)

func main() {
	sum := 0
	re := regexp.MustCompilePOSIX("([0-9])")
	file, err := os.Open("inputs.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		expr := []byte(scanner.Text())
		first := string(re.Find(expr))
		slices.Reverse(expr)
		number, _ := strconv.Atoi(first + string(re.Find(expr)))
		sum = sum + number
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	log.Print(sum)
}