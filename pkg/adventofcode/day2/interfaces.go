package day2

// Validatable structs can tell if they are in a valid state
type Validatable interface {
	IsValid() bool
}

// CountValid counts the number of valid passwords in a ValidatableList
func CountValid(vList []Validatable) int {
	count := 0
	for _, v := range vList {
		if v.IsValid() {
			count++
		}
	}
	return count
}
