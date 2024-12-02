package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/BoschLeith/advent-of-code/utils"
)

func main() {
	reports, err := convertInput("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	safeReports := 0
	tolerableReports := 0
	for _, report := range reports {
		if isReportSafe(report) {
			safeReports++
		} else if checkReportSafetyWithDeletion(report) {
			tolerableReports++
		}
	}

	fmt.Printf("Safe Reports: %d\n", safeReports)
	fmt.Printf("Tolerable Reports: %d\n", tolerableReports)
	fmt.Printf("Total Reports: %d\n", safeReports+tolerableReports)
}

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
		subStrings := strings.Fields(line)

		var report []int
		for _, str := range subStrings {
			level, err := strconv.Atoi(str)
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

func isReportSafe(report []int) bool {
	isIncreasing, isDecreasing := false, false
	for i := 1; i < len(report); i++ {
		diff := report[i] - report[i-1]

		if diff > 0 {
			isIncreasing = true
		} else if diff < 0 {
			isDecreasing = true
		} else {
			return false
		}

		if isIncreasing && isDecreasing || utils.Abs(diff) > 3 {
			return false
		}
	}
	return true
}

func checkReportSafetyWithDeletion(report []int) bool {

	for i := range report {
		isSafe := isReportSafeWithDeletion(report, i)
		if isSafe {
			return true
		}
	}

	return false
}

func isReportSafeWithDeletion(report []int, deleteIndex int) bool {
	copyReport := make([]int, len(report))
	copy(copyReport, report)

	if deleteIndex == len(copyReport)-1 {
		copyReport = copyReport[:deleteIndex]
	} else {
		copyReport = append(copyReport[:deleteIndex], copyReport[deleteIndex+1:]...)
	}
	return isReportSafe(copyReport)
}
