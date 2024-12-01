package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func parseInput(filePath string) ([]int, []int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var data []int
	fileReader := bufio.NewScanner(file)
	for fileReader.Scan() {
		line := fileReader.Text()
		numbers := strings.Fields(line)
		for _, n := range numbers {
			value, err := strconv.Atoi(n)
			if err == nil {
				data = append(data, value)
			}
		}
	}

	if err := fileReader.Err(); err != nil {
		return nil, nil, fmt.Errorf("syntax error: %w", err)
	}

	var left, right []int
	for index, value := range data {
		if index%2 == 0 {
			left = append(left, value)
		} else {
			right = append(right, value)
		}
	}

	return left, right, nil
}

func totalDistance(left, right []int) int {
	sort.Ints(left)
	sort.Ints(right)

	totalDifference := 0
	for i := 0; i < len(left); i++ {
		totalDifference += abs(left[i] - right[i])
	}

	return totalDifference
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func countOccurrences(slice []int, value int) int {
	count := 0
	for _, v := range slice {
		if v == value {
			count++
		}
	}
	return count
}

func similarityScore(left, right []int) int {
	sort.Ints(left)
	sort.Ints(right)

	weightedSum := 0
	for _, lVal := range left {
		count := countOccurrences(right, lVal)
		weightedSum += lVal * count
	}

	return weightedSum
}

func main() {
	filePath := "day1/input"
	left, right, err := parseInput(filePath)
	if err != nil {
		log.Fatalf("Error parsing input: %v", err)
	}

	distance := totalDistance(left, right)
	fmt.Println(distance)

	similarity := similarityScore(left, right)
	fmt.Println(similarity)
}
