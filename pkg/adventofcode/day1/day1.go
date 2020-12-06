package day1

import (
	"errors"
	"log"
)

const msgNoSumFound = "no sum found"

// SumFinder has methods to search for values that match a given sum
type SumFinder interface {
	// Searches a list of numbers to find a matching sum
	FindSum(elements []int, sum int) (int, int)

	// Searches a list of numbers for three element that sum up to given sum
	FindTripleSum(elements []int, sum int) ([]int, error)
}

// ErrNoSumFound gets returned if no sum was found
var ErrNoSumFound = errors.New(msgNoSumFound)

// LinearSumFinder tries to find sums by trying all possible combinations
type LinearSumFinder struct {
}

// FindTwoElementSum searches the array using two loops for a given sum
func (*LinearSumFinder) FindTwoElementSum(elements []int, sum int) (int, int, error) {
	compCounter := 0
	for i := 0; i < len(elements)-1; i++ {
		for j := i; j < len(elements); j++ {
			if elements[i]+elements[j] == sum {
				log.Printf("%d comparisons \n", compCounter)
				return elements[i], elements[j], nil
			}
			compCounter++
		}
	}
	return 0, 0, ErrNoSumFound
}

// FindThreeElementSum searches a, b, c in a list of integers with a+b+c=sum
func (l *LinearSumFinder) FindThreeElementSum(elements []int, sum int) ([3]int, error) {
	for idx, element := range elements {
		remainingSum := sum - element
		// prevent element from being summed up with itself [..] [x] [..]
		remainingElements := append(elements[:idx], elements[idx+1:]...)
		a, b, err := l.FindTwoElementSum(remainingElements, remainingSum)
		if err != nil {
			continue
		}
		// Check if we got the sum
		if a+b+element == sum {
			return [3]int{a, b, element}, nil
		}
	}
	return [3]int{}, ErrNoSumFound
}
