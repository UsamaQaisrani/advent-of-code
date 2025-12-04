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
}

func partOne(input [][]rune) (int, error) {
	totalRow := input[0]
	totalCols := input[0][0]
	var visited map[string]bool
	for i, row := range input {
		for j, col := range row {
			if string(input[i][j]) == "@" {
				bfs(i, j)
			}
		}
	}
	return 0, nil
}

func validPosition(row, col, totalRows, totalCols int) bool {
	validRow := 0 <= row && row < totalRows
	validCols := 0 <= col && col < totalCols
	return validRow && validCols
}

func bfs(r, c int) {
	directions := [][]int{ { 0,1 },{ 0,-1 },{ 1,1 },{ 1,-1 },{ -1,1 },{ -1,-1 },{ 1,0 },{ -1,0 } }
	fmt.Println(directions)
}
