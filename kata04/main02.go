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
	"math"
	"os"
	"strconv"
	"strings"
)

func spread(values []int) float64 {
	spread := float64(values[0]) - float64(values[1])
	return spread
}

func main() {

	file, err := os.Open("./football.dat")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var diff float64
	var curSpread float64
	var team string
	var myTeam string
	i := 0
	for scanner.Scan() {
		if i == 0 {
			i++
			continue
		}
		// fmt.Println(strings.Fields(scanner.Text()))
		line := strings.Fields(scanner.Text())
		// fmt.Println(line)
		if len(line) > 8 {
			pos1, _ := strconv.Atoi(line[6])
			pos2, _ := strconv.Atoi(line[8])

			if pos1 == 0 && pos2 == 0 {
				continue
			}
			team = line[1]

			diff = spread([]int{pos1, pos2})
			if i == 1 {
				curSpread = diff
				myTeam = team
			} else if math.Abs(diff) < math.Abs(curSpread) {
				curSpread = diff
				myTeam = team

			}
		}
		fmt.Println("curspread", curSpread)
		fmt.Println("team", team)
		i++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Smallest Diff is:", math.Abs(curSpread))
	fmt.Println("Team:", myTeam, "\n")
}
