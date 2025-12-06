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

	numbers := strSplit[:len(strSplit)-1]
	rawLastLine := strSplit[len(strSplit)-1]

	operationsPart1 := strings.Split(strings.ReplaceAll(rawLastLine, " ", ""), "")

	partOneSol, err := partOne(numbers, operationsPart1)
	if err != nil {
		log.Fatal("Error in part one:", err)
		os.Exit(1)
	}
	fmt.Println("PartOne:", partOneSol)

	operationsPart2 := []string{rawLastLine}
	partTwoSol, err := partTwo(numbers, operationsPart2)
	if err != nil {
		log.Fatal("Error in part two:", err)
		os.Exit(1)
	}
	fmt.Println("PartTwo:", partTwoSol)
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

			if j < len(operations) {
				if operations[j] == "+" {
					answersList[j] += currNum
				} else {
					answersList[j] *= currNum
				}
			}
		}
	}
	answer := 0
	for _, num := range answersList {
		answer += num
	}
	return answer, nil
}

func partTwo(numbers []string, operations []string) (int, error) {
	opLine := operations[0]
	var opPos []int
	for i, c := range opLine {
		if c == '+' || c == '*' {
			opPos = append(opPos, i)
		}
	}
	result := int64(0)
	for i := range opPos {
		start := opPos[i]
		end := len(opLine)
		if i+1 < len(opPos) {
			end = opPos[i+1]
		}
		var colText []string
		for _, line := range numbers {
			if start < len(line) {
				currEnd := end
				if currEnd > len(line) {
					currEnd = len(line)
				}
				colText = append(colText, line[start:currEnd])
			} else {
				colText = append(colText, "")
			}
		}

		nums := extractVertical(colText)
		op := rune(opLine[start])
		result += applyOp(nums, op)
	}

	return int(result), nil
}

func extractVertical(lines []string) []int {
	maxH := len(lines)
	var out []int

	if len(lines) == 0 {
		return out
	}

	width := len(lines[0])

	for col := 0; col < width; col++ {
		var b strings.Builder
		for row := 0; row < maxH; row++ {
			if col < len(lines[row]) && lines[row][col] != ' ' {
				b.WriteByte(lines[row][col])
			}
		}
		if b.Len() > 0 {
			n, _ := strconv.Atoi(b.String())
			out = append(out, n)
		}
	}
	return out
}

func applyOp(nums []int, op rune) int64 {
	if len(nums) == 0 {
		return 0
	}
	acc := int64(nums[0])
	for _, n := range nums[1:] {
		if op == '+' {
			acc += int64(n)
		} else {
			acc *= int64(n)
		}
	}
	return acc
}
