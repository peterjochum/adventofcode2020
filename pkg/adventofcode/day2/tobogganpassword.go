package day2

// TobogganPassword have a different policy but the same data
type TobogganPassword struct {
	Password
}

// IsValid tests if a TobogganPassword is valid
func (p TobogganPassword) IsValid() bool {
	apos := p.Policy.MinOccurrence
	letterA := p.StrPassword[apos-1 : apos]
	bpos := p.Policy.MaxOccurrence
	letterB := p.StrPassword[bpos-1 : bpos]

	if letterA == letterB {
		return false
	}

	if letterA == p.Policy.Character || letterB == p.Policy.Character {
		return true
	}
	return false
}

// NewTobboganPassword creates a TobogganPassword from a Password
func NewTobboganPassword(password Password) TobogganPassword {
	var tbpwd TobogganPassword
	tbpwd.Password = password
	return tbpwd
}

// NewTobboganFromPassword converts the type of TobogganPassword to Validatable
func NewTobboganFromPassword(pwdList []Password) []TobogganPassword {
	tpwd := make([]TobogganPassword, len(pwdList))
	for i := range pwdList {
		tpwd[i] = NewTobboganPassword(pwdList[i])
	}
	return tpwd
}
