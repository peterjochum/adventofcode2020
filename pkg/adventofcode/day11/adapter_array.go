package day11

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

// IsCompatibleWith checks if two adapters are compatible
func (a Adapter) IsCompatibleWith(thatAdapter Adapter) Compatibility {
	if thatAdapter.Joltage > a.Joltage && thatAdapter.Joltage <= a.Joltage+3 {
		return Compatible
	}
	return Incompatible
}

// String returns a formatted String representation of an Adapter
func (a Adapter) String() string {
	return fmt.Sprintf("Ad%dj", a.Joltage)
}
