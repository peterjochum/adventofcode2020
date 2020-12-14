package day10

import (
	"fmt"
	"sort"

	"github.com/peterjochum/adventofcode2020/pkg/adventofcode/utility"
)

// Adapter charges your devices in no time
type Adapter struct {
	Joltage int
	Used    bool
}

// String returns a formatted String representation of an Adapter
func (a Adapter) String() string {
	return fmt.Sprintf("%d", a.Joltage)
}

// IsCompatibleWith checks if two adapters are compatible
func (a Adapter) IsCompatibleWith(thatAdapter Adapter) Compatibility {
	if thatAdapter.Joltage > a.Joltage && thatAdapter.Joltage <= a.Joltage+3 {
		return Compatible
	}
	return Incompatible
}

// ByJoltage compares adapters according to their Joltage
type ByJoltage []Adapter

func (a ByJoltage) Len() int           { return len(a) }
func (a ByJoltage) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByJoltage) Less(i, j int) bool { return a[i].Joltage < a[j].Joltage }

// Compatibility (bool) - just wanted to try a bool enum with String() :)
type Compatibility bool

const (
	// Compatible adapters fit well together to charge your devices
	Compatible Compatibility = true
	// Incompatible adapters shall not be used - fire hazard
	Incompatible Compatibility = false
)

// GetDifferences returns the summed differences of a complete joltage chain
func GetDifferences(adapters []Adapter) ([4]int, error) {
	differences := [4]int{}
	// Initialize a virtual adapter (outlet with 0 jolts)
	currentAdapter := Adapter{0, false}
	for _, adapter := range adapters {
		if currentAdapter.IsCompatibleWith(adapter) {
			joltDiff := adapter.Joltage - currentAdapter.Joltage
			differences[joltDiff]++
			currentAdapter.Joltage = adapter.Joltage
			adapter.Used = true
		}
	}
	// finally always add 1 +3 difference as the device is rated 3 jolts
	// higher than the highest adapter
	differences[3]++

	return differences, nil
}

func (c Compatibility) String() string {
	switch c {
	case Compatible:
		return "compatible"
	case Incompatible:
		return "incompatible"
	}
	panic("unknown compatible type")
}

// GetAdaptersFromJoltages produces a slice of adapters using []int
func GetAdaptersFromJoltages(joltages []int) []Adapter {
	var adapters []Adapter
	for _, j := range joltages {
		adapters = append(adapters, Adapter{j, false})
	}
	return adapters
}

// GetAdapterListFromFile returns a slice of adapters according to Joltage info
func GetAdapterListFromFile(fileName string) []Adapter {
	joltages, err := utility.ReadNumbersFromFile(fileName)
	if err != nil {
		panic(err)
	}
	adapters := GetAdaptersFromJoltages(joltages)
	sort.Sort(ByJoltage(adapters))
	return adapters
}

// CountOptionalAdapters counts how many adapters could be skipped (WIP:)
func CountOptionalAdapters(adapters []Adapter) int {
	optional := 0
	outlet := Adapter{0, false}
	deviceJoltage := adapters[len(adapters)-1].Joltage + 3
	device := Adapter{deviceJoltage, false}

	// Add the outlet and the device as start and end points
	var completeAdapters []Adapter
	completeAdapters = append(completeAdapters, outlet)
	completeAdapters = append(completeAdapters, adapters...)
	completeAdapters = append(completeAdapters, device)

	// Check for every adapter if the previous and next are compatible
	for i := 1; i < len(adapters)-1; i++ {
		if completeAdapters[i-1].IsCompatibleWith(completeAdapters[i+1]) {
			// I am optional
			optional++
		}
	}
	return optional
}
