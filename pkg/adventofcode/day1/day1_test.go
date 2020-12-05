package day1

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinearSumFinder_FindSum(t *testing.T) {
	expenses := []int{1721,
		979,
		366,
		299,
		675,
		1456}
	finder := LinearSumFinder{}

	expectedSum := 2020
	t.Run("Find 2020", func(t *testing.T) {
		a, b, err := finder.FindTwoElementSum(expenses, expectedSum)
		if err != nil {
			t.Error(err)
		}

		if a > b {
			a, b = b, a
		}

		assert.Equal(t, 299, a)
		assert.Equal(t, 1721, b)
	})

	t.Run("No sum available", func(t *testing.T) {
		// sum of 2021 not in array
		_, _, err := finder.FindTwoElementSum(expenses, 2021)
		if err != nil {
			if err != ErrNoSumFound {
				t.Errorf("\"%s\" error expected", msgNoSumFound)
			}
		} else {
			t.Errorf("Error expected")
		}
	})

	t.Run("Two numbers whose sum is 2020", func(t *testing.T) {
		a, b, err := finder.FindTwoElementSum(GetDay1Input(), expectedSum)
		assert.Equal(t, nil, err)

		if a+b != expectedSum {
			t.Errorf("Sum of %d could not be found.", expectedSum)
		}

		fmt.Printf("%d + %d = %d, multiplication result %d\n", a, b, a+b, a*b)
	})

	t.Run("Failing to find three numbers", func(t *testing.T) {
		n, err := finder.FindThreeElementSum(expenses, 2021)
		if err == nil {
			t.Errorf("Expected no sum error, but got %v", err)
		}
		for _, retVal := range n {
			if retVal != 0 {
				t.Errorf("Expected all results to be 0, but got %d in %v", retVal, n)
			}
		}
	})

	t.Run("Three numbers summing up to 2020", func(t *testing.T) {
		n, err := finder.FindThreeElementSum(GetDay1Input(), expectedSum)
		if err != nil {
			t.Error(err)
		}

		sum := n[0] + n[1] + n[2]
		mult := n[0] * n[1] * n[2]
		fmt.Printf("%d + %d + %d = %d, multiplication result %d\n",
			n[0], n[1], n[2], sum, mult)
	})

	t.Run("Exclude finding sum with self", func(t *testing.T) {
		troubleSum := []int{1000, 20, 25, 25}
		_, err := finder.FindThreeElementSum(troubleSum, 2020)
		if err == nil {
			t.Errorf("Expected no sum error, but got %v", err)
		}
	})

}
