package day5

// GetRowNumber returns the row in a plane
func GetRowNumber(rowData string, min, max int) int {
	for _, c := range rowData {
		switch c {
		case 'F':
			return GetRowNumber(rowData[1:], min, max/2+min)
		case 'B':
			return GetRowNumber(rowData[1:], (max-min)/2, max)
		default:
			panic("unknown character")
		}

	}
	return 0
}
