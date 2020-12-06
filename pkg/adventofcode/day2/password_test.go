package day2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getTestPwd3() Password {
	return Password{
		Policy: Policy{
			MinOccurrence: 2,
			MaxOccurrence: 9,
			Character:     "c",
		},
		StrPassword: "ccccccccc",
	}
}

func getTestPwd2() Password {
	return Password{
		Policy: Policy{
			MinOccurrence: 1,
			MaxOccurrence: 3,
			Character:     "b",
		},
		StrPassword: "cdefg",
	}
}

func getTestPwd1() Password {
	return Password{
		Policy: Policy{
			MinOccurrence: 1,
			MaxOccurrence: 3,
			Character:     "a",
		},
		StrPassword: "abcde",
	}
}

func TestPasswords(t *testing.T) {
	pwd1 := getTestPwd1()
	pwd2 := getTestPwd2()
	pwd3 := getTestPwd3()
	t.Run("TestPassword_IsValid", func(t *testing.T) {
		assert.True(t, pwd1.IsValid())
		assert.False(t, pwd2.IsValid())
		assert.True(t, pwd3.IsValid())
	})
}

func TestNewPasswordFromString(t *testing.T) {
	t.Run("Test a valid string", func(t *testing.T) {
		p, err := NewPasswordFromString("2-4 p: vpkpp")
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, 2, p.Policy.MinOccurrence)
		assert.Equal(t, 4, p.Policy.MaxOccurrence)
		assert.Equal(t, "p", p.Policy.Character)
		assert.Equal(t, "vpkpp", p.StrPassword)
	})

	invalidPasswordStrings := []string{"", "abcde", "a-3 c sdfs"}

	for _, pwd := range invalidPasswordStrings {
		_, err := NewPasswordFromString(pwd)
		if err == nil {
			t.Errorf("Expected error on unparsable string")
		}
	}
}

func TestNewPasswordsFromFile(t *testing.T) {
	t.Run("Load main password file", func(t *testing.T) {
		pwList, err := NewPasswordsFromFile("passwordData.txt")
		if err != nil {
			t.Error(err)
		}
		expectedPwCount := 1000
		actualPwCount := len(pwList)
		if actualPwCount != expectedPwCount {
			t.Errorf("Expected %d passwords, but got %d",
				expectedPwCount, actualPwCount)
		}
	})

	t.Run("Load non-existing password file", func(t *testing.T) {
		_, err := NewPasswordsFromFile("nonexisting.txt")
		if err == nil {
			t.Errorf("Loaded non existing file - error expected")
		}
	})
}

func TestPasswordList_CountValid(t *testing.T) {
	t.Run("Valid count of test passwords", func(t *testing.T) {
		var pwdList = []Validatable{
			getTestPwd1(),
			getTestPwd2(),
			getTestPwd3(),
		}

		expectedValid := 2
		actualValid := CountValid(pwdList)
		if actualValid != expectedValid {
			t.Errorf("Expected %d valid passwords, got %d", expectedValid, actualValid)
		}
	})

	t.Run("Test password file", func(t *testing.T) {
		pwdList, err := NewPasswordsFromFile("passwordData.txt")
		if err != nil {
			t.Error(err)
		}
		validatables := make([]Validatable, len(pwdList))
		for i := range pwdList {
			validatables[i] = pwdList[i]
		}
		actualValid := CountValid(validatables)
		expectedValid := 474
		assert.Equal(t, expectedValid, actualValid)
	})
}
