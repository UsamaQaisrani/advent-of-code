package main
import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

func abs(a int) int {
    if a >= 0 {
        return a
    }
    return -a
}

func main() {
	path := "input.txt"

	rawInstructions, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error while reading file content:", err)
	}

	refinedInstructions := strings.Split(string(rawInstructions), "\n")

	partOneOutput, err := partOne(refinedInstructions)
	if err != nil {
		fmt.Println("Error while computing solution:", err)
		return
	}
	fmt.Println("PartOne: ", partOneOutput)

	partTwoOutput, err := partTwo(refinedInstructions)
	if err != nil {
		fmt.Println("Error while computing solution:", err)
		return
	}
	fmt.Println("PartTwo: ", partTwoOutput)
}

func partOne(input []string) (int, error) {
	answer := 0
	currPosition := 50

	for _, instruction := range input {
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

func partTwo(input []string) (int, error) {
	answer := 0
	currPos := 50

	for _, instruction := range input {
		if len(instruction) < 2 {
			continue
		}

		prevPos := currPos

		direction := instruction[:1]
		distance, err := strconv.Atoi(instruction[1:])
		if err != nil {
			return -1, err
		}

		dir := 1
		if direction == "L" {
			dir = -1
		}

		if distance >= 100 {
			answer += distance / 100
		}

		currPos += (distance % 100) * dir

		if currPos > 99 {
			currPos -= 100
			if currPos != 0 && prevPos != 0 {
				answer += 1
			}
		} else if currPos < 0 {
			currPos += 100
			if currPos != 0 && prevPos != 0 {
				answer += 1
			}
		}

		if currPos == 0 {
			answer++
		}
	}

	return answer, nil
}
