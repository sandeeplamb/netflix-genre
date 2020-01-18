package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func main() {
	var lastLine int
	rand.Seed(time.Now().UnixNano())
	randomNumber := randomInt(1, 18)

	fileName, err := os.Open("genre")
	check(err)

	scanner := bufio.NewScanner(fileName)
	for scanner.Scan() {
		lastLine++
		if lastLine == randomNumber {
			line := scanner.Text()
			fmt.Println("We will watch NetFlix Movie from Genre :", line)
		}
	}
}
