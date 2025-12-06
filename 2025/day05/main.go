package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	path := "input.txt"
	fileContent, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("Error while reading file:", err)
		os.Exit(1)
	}
	contentStr := strings.TrimRight(string(fileContent), "\n")
	input := strings.Split(contentStr, "\n\n")
	intervals := strings.Split(input[0], "\n")
	ids := strings.Split(input[1], "\n")
	mergedIntervals, err := mergeIntervals(intervals)
	fmt.Println("Merged Intervals:", mergedIntervals)
	partOneSol, err := partOne(ids, mergedIntervals)
	if err != nil {
		log.Fatal("Error in part one:", err)
		os.Exit(1)
	}
	fmt.Println("PartOne:", partOneSol)
}

func mergeIntervals(intervals []string) ([][]int, error) {
	sort.Slice(intervals, func(i, j int) bool {
		startI, _ := strconv.Atoi(strings.Split(intervals[i], "-")[0])
		startJ, _ := strconv.Atoi(strings.Split(intervals[j], "-")[0])
		return startI < startJ
	})
	var uniqueIntervals [][]int
	prevStart, err := strconv.Atoi(strings.Split(intervals[0], "-")[0])
	if err != nil {
		return nil, err
	}
	prevEnd, err := strconv.Atoi(strings.Split(intervals[0], "-")[1])
	if err != nil {
		return nil, err
	}
	for _, interval := range intervals {
		splitString := strings.Split(interval, "-")
		currStart, err := strconv.Atoi(splitString[0])
		if err != nil {
			return nil, err
		}
		currEnd, err := strconv.Atoi(splitString[1])
		if err != nil {
			return nil, err
		}

		if currStart <= prevEnd {
			prevEnd = max(currEnd, prevEnd)
		} else {
			uniqueIntervals = append(uniqueIntervals, []int{prevStart, prevEnd})
			prevStart = currStart
			prevEnd = currEnd
		}
	}
	uniqueIntervals = append(uniqueIntervals, []int{prevStart, prevEnd})
	return uniqueIntervals, nil
}

func max(a, b int) int {
	if a <= b {
		return b
	} else {
		return a
	}
}

func partOne(ids []string, intervals [][]int) (int, error) {
	count := 0
	for _, id := range ids {
		intId, err := strconv.Atoi(id)
		if err != nil {
			return 0, err
		}

		for _, interval := range intervals {
			start := interval[0]
			end := interval[1]
			if intId >= start && intId <= end {
				count++
			}
		}
	}
	return count, nil
}
