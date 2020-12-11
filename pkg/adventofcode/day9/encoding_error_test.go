package day9

import (
	"fmt"
	"testing"

	"github.com/peterjochum/adventofcode2020/pkg/adventofcode/utility"
	"github.com/stretchr/testify/assert"
)

func TestBruteXmas_AddPreamble(t *testing.T) {
	b := BruteXmas{}
	n := 35
	b.AddPreamble(n)
	assert.Contains(t, b.numbers, n)
}

func TestBruteXmas_AddNumber(t *testing.T) {
	b := BruteXmas{}
	b.AddPreamble(1)
	secondVal := 2
	b.AddPreamble(secondVal)

	t.Run("Add a valid number", func(t *testing.T) {
		err := b.AddNumber(3)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, 2, len(b.numbers))
		assert.Contains(t, b.numbers, secondVal)
	})

	t.Run("Add an invalid number", func(t *testing.T) {
		invalidVal := 4
		err := b.AddNumber(invalidVal) // 2,3 -> 4 is not a sum
		if err == nil {
			t.Errorf("Expected an error when adding %d", invalidVal)
		}
	})
}

func TestInitXmasFromFile(t *testing.T) {
	inputs := []struct {
		fileName    string
		weakNum     int
		preambleLen int
		weakSum     int
	}{
		{"input_test.txt", 127, 5, 62},
		{"input_a.txt", 90433990, 25, 11691646},
	}

	for _, input := range inputs {
		testName := fmt.Sprintf("Xmas from file %s", input.fileName)
		b := BruteXmas{}
		t.Run(testName, func(t *testing.T) {
			err := InitXmasFromFile(&b, input.fileName, input.preambleLen)
			if err == nil {
				t.Errorf("expected an error at number %d", input.weakNum)
			}
			invXmasNum, ok := err.(*InvalidXmasNumber)
			if ok {
				if invXmasNum.number != input.weakNum {
					t.Errorf("expected number %d in error but got %d",
						input.weakNum, invXmasNum.number)
				}
			} else {
				t.Errorf("Expected InvalidXmasNumber error class")
			}
		})
		weakTestName := fmt.Sprintf("Find weakness in %s", input.fileName)
		t.Run(weakTestName, func(t *testing.T) {
			num, err := utility.ReadNumbersFromFile(input.fileName)
			if err != nil {
				t.Error(err)
			}

			weakSum, err := FindWeakness(num, input.weakNum)
			if err != nil {
				t.Error(err)
			}

			if weakSum != input.weakSum {
				t.Errorf("Expected %d weak sum, but got %d",
					input.weakSum, weakSum)
			}
		})
	}
}
