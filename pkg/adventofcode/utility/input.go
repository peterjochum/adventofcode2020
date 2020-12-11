package utility

import (
	"bufio"
	"os"
	"strconv"
)

// ReadNumbersFromFile reads all integers from a text file
func ReadNumbersFromFile(fileName string) ([]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	var numbers []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// skip empty lines
		if line == "" {
			continue
		}
		num, err := strconv.Atoi(line)
		if err != nil {
			return numbers, err
		}
		numbers = append(numbers, num)
	}
	return numbers, nil
}
