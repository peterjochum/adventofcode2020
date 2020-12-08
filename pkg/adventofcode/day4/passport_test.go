package day4

import (
	"fmt"
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
	validCnt := CountValidPassports("passportDataTest.txt", false)
	assert.Equal(t, 2, validCnt)
}

func TestCountValidPassportsReal(t *testing.T) {
	validCnt := CountValidPassports("passportData.txt", false)
	assert.Equal(t, 254, validCnt)
}

func TestGetColonField(t *testing.T) {
	ppExamples := []struct {
		ppStr     string
		fieldName string
		expected  string
		error     error
	}{
		{"foo:23", "foo", "23", nil},
		{"\nfoo:45 bar:baz\n", "bar", "baz", nil},
		{"", "baz", "", errFieldNotFound},
		{"abaz:hi notexact:23", "exact", "", errFieldNotFound},
	}

	for _, pe := range ppExamples {
		t.Run(fmt.Sprintf("Retrieving field %s", pe.fieldName), func(t *testing.T) {
			result, err := GetColonField(pe.ppStr, pe.fieldName)
			if err != pe.error {
				if err == nil {
					t.Errorf("Expected an error")
				} else {
					t.Errorf("Expected error %v but got %v",
						pe.error, err)
				}

			}
			if result != pe.expected {
				t.Errorf("Expected field %s to return \"%s\" but got %s",
					pe.fieldName, pe.expected, result)
			}
		})
	}
}

func TestCheckByr(t *testing.T) {
	assert.False(t, checkByr("byr:2003"))
	assert.True(t, checkByr("byr:2002"))
	assert.True(t, checkByr("byr:1920"))
	assert.False(t, checkByr("byr:1919"))
}

func TestCheckIyr(t *testing.T) {
	assert.False(t, checkIyr("iyr:2009"))
	assert.True(t, checkIyr("iyr:2010"))
	assert.True(t, checkIyr("iyr:2020"))
	assert.False(t, checkIyr("iyr:2021"))
}

func TestCheckEyr(t *testing.T) {
	assert.False(t, checkEyr("eyr:2019"))
	assert.True(t, checkEyr("eyr:2020"))
	assert.True(t, checkEyr("eyr:2030"))
	assert.False(t, checkEyr("eyr:2031"))
	assert.False(t, checkEyr("eyr:20305"))
}

func TestCheckHeight(t *testing.T) {
	heights := []struct {
		pwDdata string
		result  bool
	}{
		{"hgt:60in", true},
		{"hgt:190cm", true},
		{"hgt:190in", false},
		{"hgt:190", false},
	}
	for _, h := range heights {
		testName := fmt.Sprintf("Testing height %s", h.pwDdata)
		t.Run(testName, func(t *testing.T) {
			if checkHeight(h.pwDdata) != h.result {
				t.Errorf("Expected %s to be %v", h.pwDdata, h.result)
			}
		})
	}
}

func TestCheckHairColor(t *testing.T) {
	assert.True(t, checkHairColor("hcl:#123abc"))
	assert.False(t, checkHairColor("hcl:#123abz"))
	assert.False(t, checkHairColor("hcl:123abc"))
}

func TestCheckEyeColor(t *testing.T) {
	assert.True(t, checkEyeColor("ecl:blu"))
	assert.True(t, checkEyeColor("ecl:oth"))
	assert.False(t, checkEyeColor("ecl:grr"))
	assert.False(t, checkEyeColor("ecl:bff"))
}

func TestCheckPid(t *testing.T) {
	assert.True(t, checkPassportID("pid:000000001"))
	assert.False(t, checkPassportID("pid:0123456789"))
}

func TestPassportFieldsInvalid(t *testing.T) {
	validPP := CountValidPassports("passPortDataInvalidFields.txt", true)
	assert.Equal(t, 0, validPP)
}

func TestPassportFieldsValid(t *testing.T) {
	validPP := CountValidPassports("passPortDataValidFields.txt", true)
	assert.Equal(t, 4, validPP)
}

func TestPasswordFieldsReal(t *testing.T) {
	validPP := CountValidPassports("passportData.txt", true)
	assert.Equal(t, 184, validPP)
}
