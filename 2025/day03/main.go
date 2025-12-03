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
