package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Temporary variable for reading input
	var inputString string
	var rows int

	// Input
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("How many rows do you want? (Enter a number): ")
	inputString, _ = reader.ReadString('\n')
	inputString = strings.TrimSpace(inputString)

	rows, _ = strconv.Atoi(inputString)

	// Initialize the grid.
	var grid string = ""

	for row := 1; row <= rows; row += 1 {
		for column := 1; column <= row; column += 1 {
			grid = grid + strconv.Itoa(column)
		}
		// For every row that is complete, add a newline to create a new row.
		grid = grid + "\n"
	}

	// Output the completed grid.
	fmt.Print(grid)

	fmt.Println("\nDone.")
}
