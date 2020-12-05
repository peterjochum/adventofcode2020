package day2

import (
	"errors"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

// Policy defines the requirements of a password
type Policy struct {
	MinOccurrence int
	MaxOccurrence int
	Character     string
}

// Password structs are strings with policy data attached.
//
// The Policy allows to verify the Password using the IsValid method.
type Password struct {
	Policy   Policy
	Password string
}

// PasswordList is an alias for a collection of passwords
type PasswordList []*Password

// NewPasswordFromString creates a new password and its policy from a string
//
// Format: <MinOccurrence>-<MaxOccurrence> <Character>: <Password>
// Example: 8-9 d: dddddddndd
func NewPasswordFromString(str string) (*Password, error) {
	// Commented ?P-Version for brevity
	// pwdRegex := "(?P<MinOccurrence>[0-9]+)-(?P<MaxOccurrence>[0-9]+) (?P<Character>[a-z]): (?P<Password>[a-zA-Z]+)"
	pwdRegex := "([0-9]+)-([0-9]+) ([a-z]): ([a-zA-Z]+)"
	re := regexp.MustCompile(pwdRegex)
	parsedStrings := re.FindStringSubmatch(str)
	if parsedStrings == nil {
		msg := fmt.Sprintf("Invalid input %s - expected p.E. 1-3 p: pretzel", str)
		return nil, errors.New(msg)
	}
	p := Password{}
	// Ignore parse errors as regex already checks strings
	p.Policy.MinOccurrence, _ = strconv.Atoi(parsedStrings[1])
	p.Policy.MaxOccurrence, _ = strconv.Atoi(parsedStrings[2])
	p.Policy.Character = parsedStrings[3]
	p.Password = parsedStrings[4]
	return &p, nil
}

// NewPasswordsFromFile imports all Password strings from a file
//
// See NewPasswordFromString for the expected format
func NewPasswordsFromFile(fileName string) (PasswordList, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	lines := string(data)

	var passwords []*Password
	for _, line := range strings.Split(lines, "\n") {
		if line == "" {
			continue
		}
		p, err := NewPasswordFromString(line)
		if err != nil {
			return nil, err
		}
		passwords = append(passwords, p)
	}
	return passwords, nil
}

// IsValid checks if the Password fulfills its policy
func (p *Password) IsValid() bool {
	charCount := strings.Count(p.Password, p.Policy.Character)
	if charCount >= p.Policy.MinOccurrence && charCount <= p.Policy.MaxOccurrence {
		return true
	}
	return false
}

// CountValid counts the number of valid passwords in a PasswordList
func (pwList *PasswordList) CountValid() int {
	count := 0
	for _, p := range *pwList {
		if p.IsValid() {
			count++
		}
	}
	return count
}
