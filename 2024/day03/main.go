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
	pattern := `mul\(\d{1,3},\d{1,3}\)`
	re := regexp.MustCompile(pattern)

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Read all lines at once
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

	// Process all lines
	var total int
	var result [][]int

	for _, line := range lines {
		matches := re.FindAllString(line, -1)

		for _, expr := range matches {
			numbers := strings.TrimSuffix(strings.TrimPrefix(expr, "mul("), ")")
			numSlice := strings.Split(numbers, ",")

			var sublist []int
			for _, num := range numSlice {
				if n, err := strconv.Atoi(num); err == nil {
					sublist = append(sublist, n)
				}
			}
			result = append(result, sublist)
		}
	}

	for i := range result {
		total += result[i][0] * result[i][1]
	}

	fmt.Printf("Multiplication Result Total: %d\n", total)

}
