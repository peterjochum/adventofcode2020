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
	Policy      Policy
	StrPassword string
}

// IsValid checks if the Password fulfills its policy
func (p Password) IsValid() bool {
	charCount := strings.Count(p.StrPassword, p.Policy.Character)
	if charCount >= p.Policy.MinOccurrence && charCount <= p.Policy.MaxOccurrence {
		return true
	}
	return false
}

// NewPasswordFromString creates a new password and its policy from a string
//
// Format: <MinOccurrence>-<MaxOccurrence> <Character>: <StrPassword>
// Example: 8-9 d: dddddddndd
func NewPasswordFromString(str string) (*Password, error) {
	// Commented ?P-Version for brevity
	// pwdRegex := "(?P<MinOccurrence>[0-9]+)-(?P<MaxOccurrence>[0-9]+) (?P<Character>[a-z]): (?P<StrPassword>[a-zA-Z]+)"
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
	p.StrPassword = parsedStrings[4]
	return &p, nil
}

// NewPasswordsFromFile imports all Password strings from a file
//
// See NewPasswordFromString for the expected format
func NewPasswordsFromFile(fileName string) ([]Password, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	lines := string(data)

	var passwords []Password
	for _, line := range strings.Split(lines, "\n") {
		if line == "" {
			continue
		}
		p, err := NewPasswordFromString(line)
		if err != nil {
			return nil, err
		}
		passwords = append(passwords, *p)
	}
	return passwords, nil
}
