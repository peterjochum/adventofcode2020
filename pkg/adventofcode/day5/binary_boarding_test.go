package day5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRowNumber(t *testing.T) {
	t.Skip()
	n := GetRowNumber("FBFBBFF", 0, 127)
	assert.Equal(t, 44, n)
}
