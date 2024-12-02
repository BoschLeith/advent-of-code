package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/BoschLeith/advent-of-code/utils"
)

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
		subStrings := strings.Fields(line)

		if len(subStrings) < 2 {
			fmt.Println("Invalid line:", line)
			continue
		}

		leftNum, err := strconv.Atoi(subStrings[0])
		if err != nil {
			fmt.Println("Error converting left number:", err)
			return
		}

		rightNum, err := strconv.Atoi(subStrings[1])
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

	totalDistance := 0
	for i := range leftList {
		totalDistance += utils.Abs(leftList[i] - rightList[i])
	}

	similarityScore := 0
	for i := range leftList {
		count := 0
		for _, value := range rightList {
			if value == leftList[i] {
				count++
			}
		}
		similarityScore += leftList[i] * count
	}

	fmt.Printf("Total Distance: %d\n", totalDistance)
	fmt.Printf("Similarity Score: %d\n", similarityScore)
}
