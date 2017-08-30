// http://codekata.com/kata/kata04-data-munging/
// In weather.dat you’ll find daily weather data for Morristown, NJ for June 2002.
// Download this text file, then write a program to output the day number (column one)
// with the smallest temperature spread (the maximum temperature is the second column,
// the minimum the third column).

package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func spread(values []int) float64 {
	spread := float64(values[1]) - float64(values[2])
	return spread
}

func main() {

	file, err := os.Open("./weather.dat")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var day int
	var diff float64
	var curSpread float64
	var spreadDay int
	for scanner.Scan() {

		line := strings.Fields(scanner.Text())
		if len(line) > 0 {
			zero, _ := strconv.Atoi(line[0])
			one, _ := strconv.Atoi(line[1])
			two, _ := strconv.Atoi(line[2])
			diff = spread([]int{zero, one, two})
			day = zero
			if day == 0 {
				continue
			} else if day == 1 {
				curSpread = diff
				spreadDay = day
			} else if math.Abs(diff) < math.Abs(curSpread) {
				curSpread = diff
				spreadDay = day

			}
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Current Min spread is", math.Abs(curSpread))
	fmt.Println("On Day", spreadDay, "\n")

}
