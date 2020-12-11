package day9

import (
	"errors"
	"fmt"

	"github.com/peterjochum/adventofcode2020/pkg/adventofcode/utility"
)

// Xmas interface to add preambles and numbers
type Xmas interface {
	// AddPreamble adds the given number to the preamble
	AddPreamble(n int)

	// AddNumber inserts the next number is a sum any 2 in the preamble
	AddNumber(n int) error
}

// InvalidXmasNumber shows that a number has no sum in the preamble
type InvalidXmasNumber struct {
	number int
}

func (i *InvalidXmasNumber) Error() string {
	return fmt.Sprintf("number %d is not a sum in preamble", i.number)
}

// InitXmasFromFile feeds all numbers from fileName into an Xmas implementation
func InitXmasFromFile(xmas Xmas, fileName string, preambleLength int) error {
	numbers, err := utility.ReadNumbersFromFile(fileName)
	if err != nil {
		panic(err)
	}

	for _, n := range numbers[:preambleLength] {
		xmas.AddPreamble(n)
	}

	for _, n := range numbers[preambleLength:] {
		if err := xmas.AddNumber(n); err != nil {
			return err
		}
	}
	return nil
}

// BruteXmas saves the day by summing up all numbers it can find
type BruteXmas struct {
	numbers []int
}

// FindWeakness searches for a range of integers which, when summed up result
// in the weak number. The sum of smallest and largest number in the range
// is returned.
func FindWeakness(numbers []int, weak int) (int, error) {
	for i, n := range numbers {
		if n > weak {
			break
		}
		sum := n
		for j, c := range numbers[i+1:] {
			sum += c
			if sum == weak {
				min := utility.Min(numbers[i : i+j+1])
				max := utility.Max(numbers[i : i+j+1])
				return min + max, nil
			}
			if sum > weak {
				continue
			}
		}
	}
	return 0, errors.New("no sum for weak number found")
}

// AddPreamble adds a number to the preamble
func (b *BruteXmas) AddPreamble(n int) {
	b.numbers = append(b.numbers, n)
}

// AddNumber checks and adds a number to the cipher
func (b *BruteXmas) AddNumber(n int) error {
	if b.isValid(n) {
		b.numbers = append(b.numbers[1:], n)
		return nil
	}
	return &InvalidXmasNumber{n}
}

// isValid checks if the number is a sum of any 2 in the preamble
func (b *BruteXmas) isValid(n int) bool {
	for i := 0; i < len(b.numbers)-1; i++ {
		for j := i + 1; j < len(b.numbers); j++ {
			if b.numbers[i]+b.numbers[j] == n {
				return true
			}
		}
	}
	return false
}
