package main
import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	partOneOutput, err := solution()
	if err != nil {
		fmt.Println("Error while reading file content:", err)
		return
	}
	fmt.Println("PartOne: ", partOneOutput)
}

func solution() (int, error) {
	path := "input.txt"
	answer := 0
	currPosition := 50

	rawInstructions, err := os.ReadFile(path)
	if err != nil {
		return -1, err
	}

	refinedInstructions := strings.Split(string(rawInstructions), "\n")

	for _, instruction := range refinedInstructions {
		if len(instruction) < 2 {
			// Not a valid instruction, skip
			continue
		}
		direction := instruction[:1]
		distance, err := strconv.Atoi(instruction[1:])
		if err != nil {
			return -1, err
		}

		if direction == "L" {
			distance *= -1
		} 

		newPosition := (currPosition + distance) % 100
		
		if newPosition < 0 {
			currPosition = 100 + newPosition
		} else if newPosition >= 100 {
			currPosition = newPosition - 100
		} else {
			currPosition = newPosition
		}

		if currPosition == 0 {
			answer += 1
		}
	}
	return answer, nil
}

func abs(a int) int {
    if a >= 0 {
        return a
    }
    return -a
}
