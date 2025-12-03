package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	path := "input.txt"
	rawInput, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error while reading input:", err)
	}
	refinedInput := strings.Split(string(rawInput), "\n")

	partOneSol, err := partOne(refinedInput)
	if err != nil {
		fmt.Println("Error in PartOne:", err)
		return
	}
	fmt.Println("PartOne:", partOneSol)

	partTwoSol, err := partTwo(refinedInput)
	if err != nil {
		fmt.Println("Error in PartTwo:", err)
		return
	}
	fmt.Println("PartTwo:", partTwoSol)
}

func partOne(input []string) (int, error) {
	totalMax := 0	
	for _, bank := range input {
		left := 0
		right := 1
		currMax := 0
		for ; right<len(bank); {
			strNum := string(bank[left]) + string(bank[right])
			num, err := strconv.Atoi(strNum)
			if err != nil {
				return -1, err
			}
			if num > currMax {
				currMax = num
			}
			if bank[left] < bank[right] {
				left = right
			}
			right++
		}
		totalMax += currMax
	}
	return totalMax, nil
}

func partTwo(input []string) (int, error) {
	totalJolts := 0
	for _, bank := range input {
		requiredNums := 12
		var stack []int
		for index, numStr := range bank {
			remainingNums := len(bank) - index 
			num, err := strconv.Atoi(string(numStr))
			if err != nil {
				return -1, err
			}
			//If remaining numbers less than reqruired, just keep appending.
			if remainingNums <= requiredNums {
				stack = append(stack, num)
				requiredNums--
				continue
			}
			// If stack is empty append the number
			if len(stack) < 1 {
				stack = append(stack, num)
				requiredNums--
				continue
			}
			
			//Pop until a bigger number is found
			for len(stack) > 0 && num > stack[len(stack)-1] { 
				if len(stack)+remainingNums > 12 { 
					stack = stack[:len(stack)-1]
					requiredNums++
				} else {
					break 
				}
			}

			if len(stack) < 12 {
				stack = append(stack, num)
				requiredNums--
			}
		}
		joltStr := ""
		for _, num := range stack {
			joltStr += fmt.Sprintf("%d", num)
		}
		fmt.Println(joltStr)
		if joltStr == "" {
			continue
		}
		currJolts, err := strconv.Atoi(joltStr)
		if err != nil {
			return -1, err
		}
		totalJolts += currJolts
	}
	return totalJolts, nil
}
