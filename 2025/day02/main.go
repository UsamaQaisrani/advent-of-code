package main
import (
	"fmt"
	"os"
	"strings"
	"stconv"
)

func main() {
	path := fmt.Sprintf("input.txt")

	rawInput, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error while reading file:", err)
		return
	}

	refinedInput := strings.Split(string(rawInput), ",")
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

	for _, currRange := range input {
		fmt.Println(currRange)
		low := strconv.Atoi(currRange[0])
		high := strconv.Atoi(currRange[1])
		seedLength := len(high)

		for {

		}
	}

	return 0, nil
}
