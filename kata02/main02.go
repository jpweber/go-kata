// http://codekata.com/kata/kata02-karate-chop/
// Write a binary chop method that takes an integer search target and a
// sorted array of integers. It should return the integer index of the
// target in the array, or -1 if the target is not in the array.
// The signature will logically be:
// chop(int, array_of_int)  -> int

package main

import (
	"fmt"
	"math"
)

func chop(myInt int, myArray []int) int {
	var idx = -1
	arrayLen := len(myArray)

	fmt.Println("Array Lengh:", arrayLen)
	if arrayLen == 1 {
		return 0
	}

	if arrayLen > 0 {
		for i, v := range myArray {
			if math.Mod(float64(i), 2.0) == 0 {
				fmt.Println("Value in loop:", v)
				fmt.Println("myInt value:", myInt)
				if v == myInt {
					fmt.Println("Found a match. Loop exiting\n")
					return i
				}
			}
		}

		for i, v := range myArray {
			if math.Mod(float64(i), 2.0) != 0 {
				fmt.Println("Value in loop:", v)
				fmt.Println("myInt value:", myInt)
				if v == myInt {
					fmt.Println("Found a match. Loop exiting\n")
					return i
				}
			}
		}

	}

	return idx
}

func main() {

	testArray := []int{1, 3, 5, 7}
	testValue := 3

	theIndex := chop(testValue, testArray)
	if theIndex < 0 {
		fmt.Println("Could not find integer in array")
	}
	if theIndex >= 0 {
		fmt.Println("The index:", theIndex)
		theValue := testArray[theIndex]
		fmt.Println("The value:", theValue)
	}

}
