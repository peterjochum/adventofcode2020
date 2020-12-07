package day3

import (
	"errors"
	"io/ioutil"
	"strings"
)

const (
	// Open is a free (open) field
	Open = '.'
	// Tree represents a tree on the hill. Count it when you hit it :)
	Tree = '#'
)

// Hill is an object one can slide down in a tobbogan
type Hill struct {
	fields [][]rune
}

var errInvalidCharacter = errors.New("Invalid character in input file")

// NewHillFromFile creates a new hill from a datafile
func NewHillFromFile(fileName string) (*Hill, error) {
	var h Hill
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(data), "\n")
	xdim := len(lines[0])
	for i, line := range lines {
		if line == "" {
			continue
		}
		h.fields = append(h.fields, make([]rune, xdim))
		for j, char := range line {
			if char == Open || char == Tree {
				h.fields[i][j] = char
			} else {
				return nil, errInvalidCharacter
			}
		}
	}
	return &h, nil
}

// At returns the object type at a specified position
func (h *Hill) At(xPos int, yPos int) rune {
	if yPos > len(h.fields) {
		panic("Hill not that long!")
	}
	xPos %= len(h.fields[0])
	return h.fields[yPos][xPos]
}

// SlideDownAndCountTrees starts from the top-left of the hill and counts
// the number of trees while sliding down
func (h *Hill) SlideDownAndCountTrees(xSlope, ySlope int) int {
	xPos, yPos, trees := 0, 0, 0
	for {
		field := h.At(xPos, yPos)
		if field == Tree {
			trees++
		}
		xPos += xSlope
		yPos += ySlope
		if yPos >= len(h.fields) {
			break
		}
	}
	return trees
}
