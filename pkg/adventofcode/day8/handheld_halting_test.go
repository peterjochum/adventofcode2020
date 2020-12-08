package day8

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunBootCode_Test(t *testing.T) {
	accumulator := RunBootCode("boot_code_test.txt")
	assert.Equal(t, 5, accumulator)
}

func TestRunBootCode_Real(t *testing.T) {
	accumulator := RunBootCode("boot_code.txt")
	assert.Equal(t, 1810, accumulator)
}

func TestFixBootCode_Testdata(t *testing.T) {
	accumulator, err := FixBootCode("fix_boot_code_test.txt")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 8, accumulator)
}

func TestFixBootCode_Real(t *testing.T) {
	accumulator, err := FixBootCode("boot_code.txt")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 969, accumulator)
}

func TestExecuteInstructions(t *testing.T) {
	var iList InstructionList
	accInstruction1 := Instruction{ACC, 5, 0}
	accInstruction2 := Instruction{ACC, -2, 0}

	t.Run("Test single ACC", func(t *testing.T) {
		iList = append(iList, &accInstruction1)
		accumulator, termCause := ExecuteInstructions(iList)
		assert.Equal(t, 5, accumulator)
		assert.Equal(t, 1, accInstruction1.RunCount)
		assert.Equal(t, EndReached, termCause)
		iList.Reset()
	})

	t.Run("Test negative ACC", func(t *testing.T) {
		iList = append(iList, &accInstruction2)
		accumulator, termCause := ExecuteInstructions(iList)
		assert.Equal(t, 3, accumulator)
		assert.Equal(t, 1, accInstruction1.RunCount)
		assert.Equal(t, EndReached, termCause)
		iList.Reset()
	})

}

func TestNewInstructionListFromFile(t *testing.T) {
	iList := NewInstructionListFromFile("boot_code_test.txt")
	assert.Equal(t, 9, len(iList))
}

func TestNewInstructionFromString(t *testing.T) {
	instructions := []struct {
		str       string
		eTyp      IType
		eVal      int
		eRunCount int
	}{
		{"acc +1", ACC, 1, 0},
		{"nop +0", NOP, 0, 0},
		{"acc -99", ACC, -99, 0},
		{"jmp -4", JMP, -4, 0},
	}

	for _, iExample := range instructions {
		testTitle := fmt.Sprintf("Testing %s instruction", iExample.str)
		t.Run(testTitle, func(t *testing.T) {
			i, err := NewInstructionFromString(iExample.str)
			if err != nil {
				t.Error(err)
			}
			assert.Equal(t, iExample.eTyp, i.Typ)
			assert.Equal(t, iExample.eVal, i.Val)
			assert.Equal(t, iExample.eRunCount, i.RunCount)
		})
	}

}
