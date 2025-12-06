package main

import (
	"fmt"
	"log"
	"os"
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
	strSplit := strings.Split(strings.TrimRight(string(fileContent), "\n"), "\n")
	operations := strings.Split(strings.ReplaceAll(strSplit[len(strSplit)-1], " ", ""), "")
	numbers := strSplit[:len(strSplit)-1]
	partOneSol, err := partOne(numbers, operations)
	if err != nil {
		log.Fatal("Error in part one:", err)
		os.Exit(1)
	}
	fmt.Println("PartOne:", partOneSol)
}

func partOne(numbers []string, operations []string) (int, error) {
	var answersList []int
	for i, row := range numbers {
		numsList := strings.Fields(row)
		for j, col := range numsList {
			currNum, err := strconv.Atoi(string(col))
			if err != nil {
				return 0, err
			}
			if i == 0 {
				answersList = append(answersList, currNum)
				continue
			}

			if operations[j] == "+" {
				answersList[j] += currNum
			} else {
				answersList[j] *= currNum
			}
		}
	}
	fmt.Println(answersList)
	answer := 0
	for _, num := range answersList {
		answer += num
	}
	return answer, nil
}
