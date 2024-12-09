package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func correct(pages map[string]int, rules [][]string) bool {
	for _, order := range rules {
		left, leftExists := pages[order[0]]
		right, rightExists := pages[order[1]]

		if !leftExists || !rightExists {
			continue
		}

		if left > right {
			return false
		}
	}
	return true
}

func sortPages(pagesList []string, rules [][]string) {
	sort.Slice(pagesList, func(i, j int) bool {
		p1, p2 := pagesList[i], pagesList[j]
		for _, order := range rules {
			if p1 == order[0] && p2 == order[1] {
				return true
			}
			if p1 == order[1] && p2 == order[0] {
				return false
			}
		}
		return false // Maintain original order if no order is defined
	})
}

func solve(text []string) (int, int) {
	r1, r2 := 0, 0
	rules := [][]string{}

	for _, line := range text {
		if strings.Contains(line, "|") {
			pages := strings.Split(line, "|")
			rules = append(rules, pages)
		}

		if strings.Contains(line, ",") {
			pages := strings.Split(line, ",")
			middlePageNumber := len(pages) / 2

			pagesMap := make(map[string]int)
			pagesList := make([]string, len(pages))

			for number, page := range pages {
				pagesMap[page] = number
				pagesList[number] = page
			}

			if correct(pagesMap, rules) {
				if value, err := strconv.Atoi(pagesList[middlePageNumber]); err == nil {
					r1 += value
				}
			} else {
				sortPages(pagesList, rules)
				if value, err := strconv.Atoi(pagesList[middlePageNumber]); err == nil {
					r2 += value
				}
			}
		}
	}

	return r1, r2
}

func main() {
	// Open the file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
	}
	defer file.Close()

	var text []string
	scanner := bufio.NewScanner(file)

	// Read lines from the file
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v", err)
	}

	resultOne, resultTwo := solve(text)
	fmt.Println("Result 1:", resultOne)
	fmt.Println("Result 2:", resultTwo)
}
