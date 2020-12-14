package day10

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAdapter_IsCompatibleWith(t *testing.T) {
	testAdapter := Adapter{0, false}

	adapterCompatibility := []struct {
		adapter Adapter
		result  Compatibility
	}{
		{Adapter{0, false}, Incompatible},
		{Adapter{1, false}, Compatible},
		{Adapter{2, false}, Compatible},
		{Adapter{3, false}, Compatible},
		{Adapter{4, false}, Incompatible},
	}
	for _, ac := range adapterCompatibility {
		testMsg := fmt.Sprintf("Adapter %s compatible to %s?",
			testAdapter, ac.adapter)
		t.Run(testMsg, func(t *testing.T) {
			if testAdapter.IsCompatibleWith(ac.adapter) != ac.result {
				t.Errorf("Expected %s to be %s with %s, but was %s.",
					testAdapter, !ac.result, ac.adapter, ac.result)
			}
		})
	}
}

func TestGetAdapterListFromFile(t *testing.T) {
	adapters := GetAdapterListFromFile("example_adapters.txt")
	expectedJoltages := []int{1, 4, 5, 6, 7, 10, 11, 12, 15, 16, 19}
	expectedAdapters := GetAdaptersFromJoltages(expectedJoltages)

	if !reflect.DeepEqual(adapters, expectedAdapters) {
		t.Errorf("Adapters %v is not equal to expected %v",
			adapters, expectedAdapters)
	}
}

func TestGetDifferences(t *testing.T) {

	testAdapters := GetAdaptersFromJoltages([]int{1, 4, 7})
	exampleAdapters := GetAdapterListFromFile("example_adapters.txt")
	example2Adapters := GetAdapterListFromFile("example_adapters_2.txt")
	puzzleInputAdapters := GetAdapterListFromFile("adapters.txt")

	differentCases := []struct {
		name        string
		adapters    []Adapter
		differences [4]int
	}{
		{"Test", testAdapters, [...]int{0, 1, 0, 3}},
		{"Example 1", exampleAdapters, [...]int{0, 7, 0, 5}},
		{"Example 2", example2Adapters, [...]int{0, 22, 0, 10}},
		{"Puzzle", puzzleInputAdapters, [...]int{0, 69, 0, 34}},
	}

	for _, testCase := range differentCases {
		t.Run(fmt.Sprintf("Differences %s", testCase.name), func(t *testing.T) {
			differences, err := GetDifferences(testCase.adapters)
			if err != nil {
				t.Error(err)
			}
			if !reflect.DeepEqual(differences, testCase.differences) {
				t.Errorf("Expected %v, but go %v", testCase.differences, differences)
			}
		})

	}
}
