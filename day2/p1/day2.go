package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	sum := 0
	re := regexp.MustCompile("[0-9]+")
	file, err := os.Open("inputs.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := []byte(scanner.Text())
		gameNo, _ := strconv.Atoi(string(re.Find(line)))
		games := strings.Split(strings.Split(string(line), ":")[1], ";")
		blue := true
		green := true
		red := true
		for _, drawings := range games {
			cubes := strings.Split(drawings, ",")
			for _, combination := range cubes {
				pair := strings.Split(strings.Trim(combination, " "), " ")
				amount, _ := strconv.Atoi(pair[0])
				if pair[1] == "blue" && amount > 14 {
					blue = false
				}
				if pair[1] == "green" && amount > 13 {
					green = false
				}
				if pair[1] == "red" && amount > 12 {
					red = false
				}
			}
		}
		fmt.Println("Game No: ", gameNo, "rgb", red, green, blue, "sum ", sum)
		if red && green && blue {
			sum = sum + gameNo
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(sum)
}
