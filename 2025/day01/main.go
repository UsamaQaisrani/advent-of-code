package main
import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

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

func abs(a int) int {
    if a >= 0 {
        return a
    }
    return -a
}

func partTwo() (int, error) {
	answer := 0
	return answer, nil
}
