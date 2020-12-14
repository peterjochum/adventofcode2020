package day11

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
