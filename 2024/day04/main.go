package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Open the file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var text []string
	scanner := bufio.NewScanner(file)

	// Read lines from the file
	for scanner.Scan() {
		line := scanner.Text()
		text = append(text, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	substrings := []string{"XMAS", "SAMX"}
	counts := make(map[string]int)

	// Count occurrences in all directions
	countOccurrences(text, substrings, counts)

	// Print results
	for substring, count := range counts {
		fmt.Printf("The substring '%s' was found %d times.\n", substring, count)
	}
}

func countOccurrences(text []string, substrings []string, counts map[string]int) {
	countHorizontally(text, substrings, counts)
	countVertically(text, substrings, counts)
	countDiagonally(text, substrings, counts)
}

func countHorizontally(text []string, substrings []string, counts map[string]int) {
	for _, line := range text {
		for _, substring := range substrings {
			counts[substring] += strings.Count(line, substring)
		}
	}
}

func countVertically(text []string, substrings []string, counts map[string]int) {
	numRows := len(text)
	numCols := len(text[0])

	for col := 0; col < numCols; col++ {
		var verticalString strings.Builder
		for row := 0; row < numRows; row++ {
			verticalString.WriteByte(text[row][col])
		}
		for _, substring := range substrings {
			counts[substring] += strings.Count(verticalString.String(), substring)
		}
	}
}

func countDiagonally(text []string, substrings []string, counts map[string]int) {
	numRows := len(text)
	numCols := len(text[0])

	// Top-left to bottom-right
	for startRow := 0; startRow < numRows; startRow++ {
		countDiagonal(text, startRow, 0, 1, 1, substrings, counts)
	}
	for startCol := 1; startCol < numCols; startCol++ {
		countDiagonal(text, 0, startCol, 1, 1, substrings, counts)
	}

	// Top-right to bottom-left
	for startRow := 0; startRow < numRows; startRow++ {
		countDiagonal(text, startRow, numCols-1, 1, -1, substrings, counts)
	}
	for startCol := numCols - 2; startCol >= 0; startCol-- {
		countDiagonal(text, 0, startCol, 1, -1, substrings, counts)
	}
}

func countDiagonal(text []string, startRow, startCol, rowStep, colStep int, substrings []string, counts map[string]int) {
	numRows := len(text)
	numCols := len(text[0])
	var diagonalString strings.Builder

	for row, col := startRow, startCol; row < numRows && col >= 0 && col < numCols; row, col = row+rowStep, col+colStep {
		diagonalString.WriteByte(text[row][col])
	}

	for _, substring := range substrings {
		counts[substring] += strings.Count(diagonalString.String(), substring)
	}
}
