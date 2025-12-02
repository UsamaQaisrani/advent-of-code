package main
import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"math"
)

func main() {
	path := fmt.Sprintf("input.txt")

	rawInput, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error while reading file:", err)
		return
	}

	removedNewLine := strings.Split(string(rawInput), "\n")
	refinedInput := strings.Split(removedNewLine[0], ",")
	var sliceInput [][]string
	for _, input := range refinedInput {
		currRange := strings.Split(input, "-")
		if len(currRange) < 2 {
			//Skip end of file or invalid range
			continue
		}
		sliceInput = append(sliceInput, currRange)
	}

	partOneSol, err := partOne(sliceInput)
	if err != nil {
		fmt.Println("Error Occured:", err)
		return
	}
	fmt.Println("PartOne:", partOneSol)
}

func partOne(input [][]string) (int, error) {
	totalSum := 0
	for _, currRange := range input {
		low, err := strconv.Atoi(currRange[0])
		if err != nil {
			return -1, err
		}
		high, err := strconv.Atoi(currRange[1])
		if err != nil {
			return -1, err
		}

		for i := 2; i <= len(currRange[1]); i += 2 {
			currSeedLen := i / 2
			seed := int(math.Pow(10, float64(currSeedLen-1)))
			seedLimit := int(math.Pow(10, float64(currSeedLen))) - 1

			for {
				if seed > seedLimit {
					break
				}

				multiplier := int(math.Pow(10, float64(currSeedLen))) + 1
				target := seed * multiplier

				if target > high {
					break
				}
				if target >= low {
					totalSum += target
					fmt.Println("Found:", target)
				}
				seed++
			}
		}
	}

	return totalSum, nil
}
