// http://codekata.com/kata/kata04-data-munging/
// The file football.dat contains the results from the English Premier League for 2001/2.
// The columns labeled ‘F’ and ‘A’ contain the total number of goals scored for and against
// each team in that season (so Arsenal scored 79 goals against opponents, and had 36 goals
// scored against them). Write a program to print the name of the team with the smallest
// difference in ‘for’ and ‘against’ goals.

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	file, err := os.Open("./football.dat")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
