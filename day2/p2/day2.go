package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	sum := 0
	file, err := os.Open("inputs.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := []byte(scanner.Text())
		games := strings.Split(strings.Split(string(line), ":")[1], ";")
		blue := 0
		green := 0
		red := 0
		for _, drawings := range games {
			cubes := strings.Split(drawings, ",")
			for _, combination := range cubes {
				pair := strings.Split(strings.Trim(combination, " "), " ")
				amount, _ := strconv.Atoi(pair[0])
				if pair[1] == "blue" && amount > blue {
					blue = amount
				}
				if pair[1] == "green" && amount > green {
					green = amount
				}
				if pair[1] == "red" && amount > red {
					red = amount
				}
			}
		}

		sum = sum + red*green*blue
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(sum)
}
