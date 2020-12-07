package day4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getTestPassportData() []string {
	return NewPassportDataFromFile("passportDataTest.txt")
}

func getTestPassport(index int) string {
	if index > 3 {
		panic("Only 4 test passwords available")
	}
	return getTestPassportData()[index]

}

func TestNewPassportDataFromFile(t *testing.T) {
	data := getTestPassportData()
	assert.Equal(t, 4, len(data))

	testPassPort3 := "hcl:#cfa07d eyr:2025 pid:166559648\niyr:2011 ecl:brn hgt:59in\n"
	assert.Equal(t, testPassPort3, getTestPassport(3))

}

func TestPassportValid(t *testing.T) {
	testppt := []struct {
		idx     int
		isValid bool
	}{
		{
			idx:     0,
			isValid: true,
		}, {
			idx:     1,
			isValid: false,
		}, {
			idx:     2,
			isValid: true,
		}, {
			idx:     3,
			isValid: false,
		},
	}

	for _, p := range testppt {
		passportStr := getTestPassport(p.idx)
		assert.Equal(t, p.isValid, PassportValid(passportStr))
	}
}

func TestCountValidPassports(t *testing.T) {
	validCnt := CountValidPassports("passportDataTest.txt")
	assert.Equal(t, 2, validCnt)
}

func TestCountValidPassportsReal(t *testing.T) {
	validCnt := CountValidPassports("passportData.txt")
	assert.Equal(t, 254, validCnt)
}
