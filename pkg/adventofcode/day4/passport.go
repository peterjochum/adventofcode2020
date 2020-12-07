package day4

import (
	"io/ioutil"
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

// TestNewPassportDataFromFile returns a []string of Passports
func NewPassportDataFromFile(fileName string) []string {
	data, _ := ioutil.ReadFile(fileName)
	return strings.Split(string(data), passPortSep)
}

// CountValidPassports tests how many passports are valid in a given array
func CountValidPassports(fileName string) int {
	ppData := NewPassportDataFromFile(fileName)
	validCnt := 0
	for _, pp := range ppData {
		if PassportValid(pp) {
			validCnt++
		}
	}
	return validCnt
}
