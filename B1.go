/*
 * @author Marco Cantusci
 * @verion 1.0.0
 * @date 2025-11-19
 * @fileoverview Check if the user should buy the colour of sweater.
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// declare variables
	var choice string

	reader := bufio.NewReader(os.Stdin)

	// input
	fmt.Print("Please choose a sweater colour from the available choices: blue, black, red, white. \n")
	choice, _ = reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	// processing
	// if choice is black or white
	if strings.ToLower(choice) == "black" || strings.ToLower(choice) == "white" {
		fmt.Println("You have enough sweaters in this colour.")
		// if choice is red
	} else if strings.ToLower(choice) == "red" {
		fmt.Println(
			"This colour will look good on you \n",
			"How about a pair of jeans to go with the sweater?",
		)
		// if choice is blue
	} else if strings.ToLower(choice) == "blue" {
		fmt.Println("This colour doesn't go well with your eyes.")
	} else {
		fmt.Println("Your colour choice is invalid.")
	}

	fmt.Println("\nDone.")
}
