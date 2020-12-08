package day8

import (
	"errors"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

// IType enumerates the types of supported Instruction types
type IType string

const (
	// ACC causes the number to be added to the accumulator
	ACC IType = "acc"
	// JMP makes a relative jump to another line of code
	JMP IType = "jmp"
	// NOP does nothing (no operation)
	NOP IType = "nop"
)

// TermCause enumerates the states of InstructionList execution termination
type TermCause int

const (
	// EndReached means the program reached the last instruction
	EndReached TermCause = iota
	// LoopDetected means the program was terminated by loop detection
	LoopDetected
	// OutOfBounds somehow (probably by a wrong jump) the instruction pointer
	// jumped outside of the available instruction set
	OutOfBounds
)

// Instruction represents a single instruction that stores its RunCount
type Instruction struct {
	Typ      IType
	Val      int
	RunCount int
}

// InstructionList is a slice of Instruction pointers (a.k.a a program)
type InstructionList []*Instruction

// Reset sets all run counters to 0
func (iList InstructionList) Reset() {
	for _, i := range iList {
		i.RunCount = 0
	}
}

// RunBootCode runs a file containing bootcode and returns the accumulator
func RunBootCode(fileName string) int {
	iList := NewInstructionListFromFile(fileName)
	acc, _ := ExecuteInstructions(iList)
	return acc
}

// FixBootCode searches for single errorenous instructions causeing code loops
func FixBootCode(fileName string) (int, error) {
	iList := NewInstructionListFromFile(fileName)

	// Try from JMP->NOP and from NOP->JMP
	instrChange := []struct {
		fromType IType
		toType   IType
	}{
		{JMP, NOP},
		{NOP, NOP},
	}
	for _, change := range instrChange {
		// Search lines for culprit instruction type
		for fixIdx := 0; fixIdx < len(iList); fixIdx++ {
			currentInstruction := iList[fixIdx]
			if currentInstruction.Typ != change.fromType {
				continue
			}
			// Change type
			currentInstruction.Typ = change.toType
			acc, termCause := ExecuteInstructions(iList)
			if termCause == EndReached {
				fmt.Printf("Fixing line %d made the code terminate\n",
					fixIdx)
				return acc, nil
			}
			// Restore the original type
			currentInstruction.Typ = change.fromType
			// Reset the instruction counters
			iList.Reset()
		}
	}
	return 0, errors.New("code could not be fixed")
}

// ExecuteInstructions runs an InstructionList and returns the accumulator
func ExecuteInstructions(iList InstructionList) (int, TermCause) {
	iPtr := 0
	acc := 0
	var termCause TermCause
	for {
		if iPtr < 0 || iPtr > len(iList) {
			termCause = OutOfBounds
		}
		if iPtr == len(iList) {
			// Program reached the end
			termCause = EndReached
			break
		}
		i := iList[iPtr]
		// Check the runcount to avoid loops
		if i.RunCount == 1 {
			termCause = LoopDetected
			break
		}
		// Increase the runcount
		i.RunCount++
		switch i.Typ {
		case ACC:
			acc += i.Val
		case JMP:
			iPtr += i.Val
			continue
		}
		// Jump to next instruction
		iPtr++
	}
	return acc, termCause
}

// NewInstructionListFromFile parses the file under fileName and returns an
// InstructionList
func NewInstructionListFromFile(fileName string) InstructionList {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic("err")
	}
	lines := string(data)
	var iList InstructionList
	for _, line := range strings.Split(lines, "\n") {
		if line == "" {
			continue
		}
		instruction, err := NewInstructionFromString(line)
		if err != nil {
			panic(err)
		}
		iList = append(iList, instruction)
	}
	return iList
}

// NewInstructionFromString parses an Instruction from a string
func NewInstructionFromString(instruction string) (*Instruction, error) {
	re := regexp.MustCompile(`(acc|jmp|nop) (\+|-)(\d+)`)
	parts := re.FindStringSubmatch(instruction)
	if parts == nil {
		return nil, fmt.Errorf("could not parse instruction: %s", instruction)
	}
	val, _ := strconv.Atoi(parts[3])
	if parts[2] == "-" {
		val *= -1
	}
	return &Instruction{
		Typ:      IType(parts[1]),
		Val:      val,
		RunCount: 0,
	}, nil
}
