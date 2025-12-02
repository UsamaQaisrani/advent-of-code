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

	partTwoSol, err := partTwo(sliceInput)
	if err != nil {
		fmt.Println("Error Occured:", err)
		return
	}
	fmt.Println("PartTwo:", partTwoSol)
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
				}
				seed++
			}
		}
	}

	return totalSum, nil
}

func partTwo(input [][]string) (int, error) {
	uniqueNums := make(map[int64]struct{})

	for _, row := range input {
		if len(row) < 2 {
			continue
		}

		L, err := strconv.ParseInt(row[0], 10, 64)
		if err != nil {
			return 0, fmt.Errorf("invalid L: %w", err)
		}
		R, err := strconv.ParseInt(row[1], 10, 64)
		if err != nil {
			return 0, fmt.Errorf("invalid R: %w", err)
		}

		for seedLen := 1; seedLen <= 9; seedLen++ {

			minSeedVal := int64(1)
			for i := 0; i < seedLen-1; i++ {
				minSeedVal *= 10
			}
			maxSeedVal := minSeedVal*10 - 1
			if seedLen == 1 {
				minSeedVal = 1 
			}

			shift := int64(1)
			for i := 0; i < seedLen; i++ {
				shift *= 10
			}

			M := int64(1) 
			
			for k := 2; ; k++ {
				if M > (math.MaxInt64-1)/shift {
					break 
				}
				M = M*shift + 1

				reqMin := (L + M - 1) / M  
				
				reqMax := R / M           

				start := max(reqMin, minSeedVal)
				end := min(reqMax, maxSeedVal)

					break 
				}
				if minSeedVal*M > R {
					break 
				}

				if start <= end {
					for s := start; s <= end; s++ {
						val := s * M
						uniqueNums[val] = struct{}{}
					}
				}
			}
		}
	}

	totalSum := 0
	for num := range uniqueNums {
		totalSum += int(num)
	}

	return totalSum, nil
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
