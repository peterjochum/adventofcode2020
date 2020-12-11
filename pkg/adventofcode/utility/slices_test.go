package utility

import "testing"

func TestMin(t *testing.T) {
	minimum := 1
	var a = []int{3, 2, minimum}
	min := Min(a)
	if min != minimum {
		t.Errorf("minimum of %v is not %d", a, min)
	}
}

func TestMax(t *testing.T) {
	eMax := 4534
	var a = []int{2, 4, eMax, 99}
	max := Max(a)
	if max != eMax {
		t.Errorf("minimum of %v is not %d", a, max)
	}
}
