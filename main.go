/*
	* Author: Zachary Fowler
	* Version: 1.0.0
	* Date: 2025-11-15
	* This file creates a program that will display and calculate the average of given numbers
	*/

package main

import "fmt"

func main() {
	// assign variables
	var number1 float64 = 56.9
	var number2 float64 = 89.7
	var number3 float64 = 90.2

	// INPUT - none

	// PROCESS
	// calculate average
	var average float64 = (number1 + number2 + number3) / 3

	// OUTPUT
	fmt.Println("The average of", number1, number2, "and", number3, "is", average)
	
	fmt.Println("\nDone.")
	}