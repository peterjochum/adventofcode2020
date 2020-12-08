package day4

import (
	"errors"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var mandatoryFields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

// passPortSep is the separating string between passwords
const passPortSep = "\n\n"

// PassportValid checks if all mandatory fields are in the passport
func PassportValid(pwdData string) bool {
	for _, field := range mandatoryFields {
		if !strings.Contains(pwdData, field) {
			return false
		}
	}
	// all mandatory fields contained
	return true
}

// PassportFieldsValid checks all mandatory fields for correct values
func PassportFieldsValid(ppData string) bool {
	if !PassportValid(ppData) {
		return false
	}
	return checkByr(ppData) && checkIyr(ppData) && checkEyr(ppData) &&
		checkHeight(ppData) && checkHairColor(ppData) &&
		checkEyeColor(ppData) && checkPassportID(ppData)
}

func checkByr(pwData string) bool {
	return checkRange(pwData, "byr", "1920", "2002", 4)
}

func checkIyr(pwData string) bool {
	return checkRange(pwData, "iyr", "2010", "2020", 4)
}

func checkEyr(pwData string) bool {
	return checkRange(pwData, "eyr", "2020", "2030", 4)
}

func checkHeight(pwData string) bool {
	height, err := GetColonField(pwData, "hgt")
	if err != nil {
		panic(err)
	}
	re := regexp.MustCompile(`(\d+)(in|cm)`)
	res := re.FindStringSubmatch(height)
	if res == nil || len(res) != 3 {
		return false
	}
	size, _ := strconv.Atoi(res[1])

	if res[2] == "cm" && size >= 150 && size <= 193 {
		return true
	} else if res[2] == "in" && size >= 59 && size <= 76 {
		return true
	}
	return false
}

func checkHairColor(passportData string) bool {
	hcl, _ := GetColonField(passportData, "hcl")
	re := regexp.MustCompile("^#[0-9a-f]{6}$")
	return re.MatchString(hcl)
}

func checkEyeColor(passportData string) bool {
	eyeClr, err := GetColonField(passportData, "ecl")
	if err != nil {
		panic(err)
	}
	switch eyeClr {
	case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
		return true
	}
	return false
}

func checkPassportID(passportData string) bool {
	pid, err := GetColonField(passportData, "pid")
	if err != nil {
		panic(err)
	}
	re := regexp.MustCompile("^[0-9]{9}$")
	return re.MatchString(pid)
}

func checkRange(pwData, field, lower, upper string, numberlen int) bool {
	val, err := GetColonField(pwData, field)
	if err != nil {
		panic("field not found")
	}
	if len(val) == numberlen && val >= lower && val <= upper {
		return true
	}
	return false
}

// NewPassportDataFromFile returns a []string of Passports
func NewPassportDataFromFile(fileName string) []string {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(data), passPortSep)
}

// CountValidPassports tests how many passports are valid in a given array
func CountValidPassports(fileName string, validateFields bool) int {
	ppData := NewPassportDataFromFile(fileName)
	validCnt := 0
	for _, pp := range ppData {
		if validateFields {
			if PassportFieldsValid(pp) {
				validCnt++
			}
		} else if PassportValid(pp) {
			validCnt++
		}
	}
	return validCnt
}

var errFieldNotFound = errors.New("field was not found")

// GetColonField returns a value from a fieldname "name:value" in a string
func GetColonField(passportStr, fieldName string) (string, error) {
	regexStr := fmt.Sprintf("(?:^|\\s+)%s:(\\S+)", fieldName)
	re := regexp.MustCompile(regexStr)
	res := re.FindStringSubmatch(passportStr)
	if res == nil {
		return "", errFieldNotFound
	}
	return res[1], nil
}
