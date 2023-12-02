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
	re := regexp.MustCompile("^Game [0-9]+: ((((1[0-4] blue|[1-9] blue)|(1[0-2] red|[1-9] red)|(1[0-3] green|[1-9] green))*)[,\\s;]*)+$")
	re2 := regexp.MustCompilePOSIX("[0-9]+")
	sum := 0

	file, err := os.Open("inputs.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := []byte(scanner.Text())
		match := re.Match(line)
		if match {
			gameNo, _ := strconv.Atoi(string(re2.Find(line)))
			sum = sum + gameNo
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(sum)
}
