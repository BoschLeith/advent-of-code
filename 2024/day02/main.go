package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func convertInput(fileName string) ([][]int, error) {
	var reports [][]int

	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		stringParts := strings.Fields(line)

		var report []int
		for _, part := range stringParts {
			level, err := strconv.Atoi(part)
			if err != nil {
				return nil, fmt.Errorf("error converting level number: %w", err)
			}
			report = append(report, level)
		}

		reports = append(reports, report)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return reports, nil
}

func main() {
	reports, err := convertInput("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	safeReports := 0
	for _, report := range reports {
		isIncreasing := true
		isDecreasing := true
		for i := 0; i < len(report)-1; i++ {
			diff := report[i+1] - report[i]

			if diff > 3 {
				isDecreasing = false
				isIncreasing = false
				break
			} else if diff < -3 {
				isDecreasing = false
				isIncreasing = false
				break
			} else if diff < 0 {
				isIncreasing = false
			} else if diff > 0 {
				isDecreasing = false
			} else {
				isIncreasing = false
				isDecreasing = false
				break
			}
		}

		if isIncreasing || isDecreasing {
			safeReports++
		}
	}

	fmt.Printf("Safe Reports: %d\n", safeReports)
}
