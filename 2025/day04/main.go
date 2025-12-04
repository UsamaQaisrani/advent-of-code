package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	path := "input.txt"
	rawInput, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error while reading the input file:", err)
		return
	}
	var refinedInput [][]rune
	inputRows := strings.Split(string(rawInput), "\n")
	for _, row := range inputRows {
		//Skip the eof newline
		if row == "" {
			continue
		}
		currRow := []rune(row)
		refinedInput = append(refinedInput,currRow)
	}
	partOneSol, err := partOne(refinedInput)
	if err != nil {
		fmt.Println("Error occured in PartOne:", err)
		return
	}
	fmt.Println("PartOne:",partOneSol)
}

func partOne(input [][]rune) (int, error) {
	count := 0
	for i, row := range input {
		for j, _ := range row {
			currCount := expand(i,j, input, 1)
			fmt.Println("Count:", currCount)
			if currCount < 4 && string(input[i][j]) != "." {
				count += 1
			}
		}
	}
	return count, nil
}

func validPosition(row, col, totalRows, totalCols int) bool {
	validRow := 0 <= row && row < totalRows
	validCols := 0 <= col && col < totalCols
	return validRow && validCols
}

func expand(r, c int, grid [][]rune, steps int) int {
//	directions := [][]int{{0,1},{0,-1},{1,1},{1,-1},{-1,1},{-1,-1},{1,0},{-1,0}}
	totalRows := len(grid)
	totalCols := len(grid[0])

	//base cases
	if !validPosition(r, c, totalRows, totalCols) {
		return 0
	}
	if string(grid[r][c]) != "@" {
		return 0
	}
	if steps == 0 {
		return 1
	}
	nextStep := steps - 1
	n1 := expand(r+1,c, grid, nextStep)
	n2 := expand(r,c+1, grid, nextStep)
	n3 := expand(r-1,c, grid, nextStep)
	n4 := expand(r,c-1, grid, nextStep)
	n5 := expand(r+1,c+1, grid, nextStep)
	n6 := expand(r-1,c-1, grid, nextStep)
	n7 := expand(r+1,c-1, grid, nextStep)
	n8 := expand(r-1,c+1, grid, nextStep)

	return n1 + n2 + n3 + n4 + n5 + n6 + n7 + n8
}
