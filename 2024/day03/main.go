package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// Compile regular expressions for matching multiplication and control functions
	regexPart1 := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	regexPart2 := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`)

	// Open the input file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Read lines from the file
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

	// Calculate and print results for both parts
	fmt.Printf("Multiplication Result Total: %d\n", calculateResults(lines, regexPart1))
	fmt.Printf("Enabled Multiplication Result Total: %d\n", calculateEnabledResults(lines, regexPart2))
}

// calculateResults computes the total multiplication result for the given regex
func calculateResults(lines []string, regex *regexp.Regexp) int {
	var total int
	for _, line := range lines {
		matches := regex.FindAllString(line, -1)
		total += sumMultiplications(matches)
	}
	return total
}

// calculateEnabledResults computes the total multiplication result considering enabled functions
func calculateEnabledResults(lines []string, regex *regexp.Regexp) int {
	var matches []string
	for _, line := range lines {
		matches = append(matches, regex.FindAllString(line, -1)...)
	}

	// Filter out the "do()" and "don't()" sections
	filteredMatches := filterSlice(matches)
	return sumMultiplications(filteredMatches)
}

// sumMultiplications sums the results of multiplication expressions
func sumMultiplications(expressions []string) int {
	var total int
	for _, expr := range expressions {
		if strings.HasPrefix(expr, "mul(") && strings.HasSuffix(expr, ")") {
			numbers := strings.TrimSuffix(strings.TrimPrefix(expr, "mul("), ")")
			numSlice := strings.Split(numbers, ",")
			if len(numSlice) == 2 {
				if n1, err1 := strconv.Atoi(numSlice[0]); err1 == nil {
					if n2, err2 := strconv.Atoi(numSlice[1]); err2 == nil {
						total += n1 * n2
					}
				}
			}
		}
	}
	return total
}

// filterSlice filters out expressions based on "do()" and "don't()" control functions
func filterSlice(slice []string) []string {
	var result []string
	inDontSection := false
	for _, item := range slice {
		if item == "don't()" {
			inDontSection = true
		} else if item == "do()" {
			inDontSection = false
		}
		if !inDontSection {
			result = append(result, item)
		}
	}
	return result
}
