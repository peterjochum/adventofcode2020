package day3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const fieldHeight = 11
const fieldWidth = 11
const testHillFile = "test_hill.txt"
const hillFile = "hill.txt"

func getHill(fileName string) *Hill {
	h, err := NewHillFromFile(fileName)
	if err != nil {
		panic(err)
	}
	return h
}

func TestHill_ImportFromFile(t *testing.T) {
	h := getHill(testHillFile)
	assert.Equal(t, fieldHeight, len(h.fields))
	assert.Equal(t, fieldWidth, len(h.fields[0]))
}

func TestHill_At(t *testing.T) {
	h := getHill(testHillFile)
	t.Run("Positions without overflow", func(t *testing.T) {
		assert.Equal(t, h.At(0, 0), Open)
		assert.Equal(t, h.At(4, 1), Tree)
		assert.Equal(t, h.At(6, 2), Tree)
		assert.Equal(t, h.At(10, 10), Tree)
	})

	t.Run("Positions with overflow in X dimension", func(t *testing.T) {
		assert.Equal(t, h.At(11, 0), Open)
		assert.Equal(t, h.At(27, 5), Tree)
	})

	t.Run("Overflow in Y direction (panic", func(t *testing.T) {
		assert.Panics(t, func() {
			h.At(0, 12)
		}, "Code did not panic")
	})
}

func TestHill_SlideDownAndCountTrees(t *testing.T) {

	h := getHill(testHillFile)
	trees := h.SlideDownAndCountTrees(3, 1)
	assert.Equal(t, 7, trees)
}

func TestHill_SlideDownAndCountTreesRealHill(t *testing.T) {
	h := getHill(hillFile)
	trees := h.SlideDownAndCountTrees(3, 1)
	assert.Equal(t, 145, trees)
}

func TestHill_SlideDownAndCountTreesVariousSlopes(t *testing.T) {
	h := getHill(hillFile)
	slopes := []struct {
		xSlope int
		ySlope int
		result int
	}{
		{1, 1, 0},
		{3, 1, 0},
		{5, 1, 0},
		{7, 1, 0},
		{1, 2, 0},
	}
	mult := 1
	for i, s := range slopes {
		slopes[i].result = h.SlideDownAndCountTrees(s.xSlope, s.ySlope)
		mult *= slopes[i].result
	}
	const expected = 3424528800
	if mult != expected {
		t.Errorf("Wrong multiplicationresult %d, expected %d",
			mult, expected)
	}
}
