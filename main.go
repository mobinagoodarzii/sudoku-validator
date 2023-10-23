package main

import (
	"fmt"
)

func checkingInSlice(allRows [][]int, input int) bool {
	isDuplicate := false
	for _, eachVal := range allRows[len(allRows)-1] {
		if input == eachVal {
			isDuplicate = true
			fmt.Printf("%d is not valid in slice.\n", input)
			break
		}
	}
	if isDuplicate == false {
		fmt.Printf("%d is valid in slice.\n", input)
	}
	return isDuplicate
}
func checkinInEqualIndex(allRows [][]int, input int) bool {
	isDuplicate := false
	inputIndex := len(allRows[len(allRows)-1])
	desiredRows := allRows[:len(allRows)-1]
	for _, echSlice := range desiredRows {
		if input == echSlice[inputIndex] {
			isDuplicate = true
			fmt.Printf("%d is not valid in equal index.\n", input)
			break
		}
	}
	if isDuplicate == false {
		fmt.Printf("%d is valid in equal index.\n", input)
	}
	return isDuplicate
}
func checkinInSquare(allRows [][]int, input int) bool {
	isDuplicate := false
	y := len(allRows)
	x := len(allRows[len(allRows)-1]) + 1
	startOfY := y - ((y - 3*((y-1)/3)) - 1)
	startOfX := x - ((x - 3*((x-1)/3)) - 1)
	for j := startOfY; j < startOfY+((y-3*((y-1)/3))-1); j++ {
		for i := startOfX; i < startOfX+3; i++ {
			//fmt.Println("y is:", y, "x is:", x)         //(to test)
			//fmt.Println(input, allRows[j-1][i-1], j, i) //(to test)
			if input == allRows[j-1][i-1] {
				isDuplicate = true
				fmt.Printf("%d is not valid in square.\n", input)
				break
			}
		}
		if isDuplicate == true {
			break
		}
	}
	if isDuplicate == false {
		fmt.Printf("%d is valid in square.\n", input)
	}
	return isDuplicate
}
func main() {
	var allRows [][]int
	var input int
	for i := 1; i <= 9; i++ {
		var rowValues []int
		fmt.Printf("Please enter the numbers in list %d:\nIf you don't want to enter another number, enter zero!\n", i)
		for j := 1; j <= 9; j++ {
			fmt.Scan(&input)
			if input == 0 {
				break
			} else {
				rowValues = append(rowValues, input)
			}
		}
		allRows = append(allRows, rowValues)
		if input == 0 {
			break
		}
	}
	fmt.Println(allRows)
	fmt.Println("Please enter the next number:")
	fmt.Scan(&input)
	if checkingInSlice(allRows, input) == false {
		if checkinInEqualIndex(allRows, input) == false {
			if checkinInSquare(allRows, input) == false {
				fmt.Printf("%d is valid.\n", input)
			}
		}
	}
}
