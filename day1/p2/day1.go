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
		ready := []byte(replaceNumbers(string(expr)))
		log.Print(string(ready))
		first := string(re.Find(ready))
		slices.Reverse(ready)
		number, _ := strconv.Atoi(first + string(re.Find(ready)))
		sum = sum + number
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	log.Print(sum)

}

func replaceNumbers(input string) string {
	replacements := map[string]string{
		"one":   "o1e",
		"two":   "t2o",
		"three": "t3e",
		"four":  "f4r",
		"five":  "f5e",
		"six":   "s6x",
		"seven": "s7n",
		"eight": "e8t",
		"nine":  "n9e",
	}

	regexPattern := `(one|two|three|four|five|six|seven|eight|nine)`

	re := regexp.MustCompile(regexPattern)
	result := re.ReplaceAllStringFunc(input, func(match string) string {
		return replacements[match]
	})
	result = re.ReplaceAllStringFunc(result, func(match string) string {
		return replacements[match]
	})
	return result
}
