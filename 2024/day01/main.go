package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	var leftList, rightList []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		stringParts := strings.Fields(line)

		if len(stringParts) < 2 {
			fmt.Println("Invalid line:", line)
			continue
		}

		leftNum, err := strconv.Atoi(stringParts[0])
		if err != nil {
			fmt.Println("Error converting left number:", err)
			return
		}

		rightNum, err := strconv.Atoi(stringParts[1])
		if err != nil {
			fmt.Println("Error converting right number:", err)
			return
		}

		leftList = append(leftList, leftNum)
		rightList = append(rightList, rightNum)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

	sort.Ints(leftList)
	sort.Ints(rightList)

	total := 0
	for i := range leftList {
		total += abs(leftList[i] - rightList[i])
	}

	fmt.Println(total)
}
