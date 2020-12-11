package utility

import (
	"reflect"
	"testing"
)

func TestReadNumbersFromFile(t *testing.T) {
	numbers, err := ReadNumbersFromFile("number_input.txt")
	if err != nil {
		t.Error(err)
	}
	expectedNumbers := []int{-1, 2, 25}
	if !reflect.DeepEqual(numbers, expectedNumbers) {
		t.Errorf("Expected %v but got %v",
			expectedNumbers, numbers)
	}
}

func TestReadNumbersFromFile_Error(t *testing.T) {
	numbers, err := ReadNumbersFromFile("number_input_illegal.txt")
	if err == nil {
		t.Errorf("Expected an error parsing the illegal file")
	}
	expectedNumCount := 1
	if len(numbers) != expectedNumCount {
		t.Fatalf("Expected %d numbers", expectedNumCount)
	}

	expectedNum := 22
	if numbers[0] != expectedNum {
		t.Errorf("Expected %d to be returned despite of the error",
			expectedNum)
	}

}
