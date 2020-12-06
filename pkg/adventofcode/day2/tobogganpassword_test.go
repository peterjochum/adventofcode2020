package day2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTobogganPassword_IsValid(t *testing.T) {
	tpwd := NewTobboganPassword(getTestPwd1())
	assert.True(t, tpwd.IsValid())

	tpwd = NewTobboganPassword(getTestPwd2())
	assert.False(t, tpwd.IsValid())

	tpwd = NewTobboganPassword(getTestPwd3())
	assert.False(t, tpwd.IsValid())
}

func TestPasswordFile(t *testing.T) {
	pwdList, err := NewPasswordsFromFile("passwordData.txt")
	if err != nil {
		t.Error(err)
	}

	toboggan := NewTobboganFromPassword(pwdList)

	validatables := make([]Validatable, len(pwdList))
	for i := range toboggan {
		validatables[i] = toboggan[i]
	}
	actualValid := CountValid(validatables)
	expectedValid := 745
	assert.Equal(t, expectedValid, actualValid)
}
