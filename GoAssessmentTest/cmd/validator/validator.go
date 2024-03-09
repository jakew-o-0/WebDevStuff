package validator

import (
	"unicode"
	"unicode/utf8"
)

type Validator struct {
    Errors map[string]string
}
func New() Validator {
    return Validator{
	Errors: make(map[string]string),
    }
}

func (v *Validator) IsValid() bool {
    return len(v.Errors) == 0
}

func (v *Validator) Validate(valid bool, error, message string) {
    if !valid {
	v.Errors[error] = message
    }
}




func (v *Validator) ValidateMinSize(input string, minSize int) bool {
    return utf8.RuneCount([]byte(input)) > minSize
}

func (v *Validator) ValidateMaxSize(input string, maxSize int) bool {
    return utf8.RuneCount([]byte(input)) < maxSize
}

func (v *Validator) ValidateMinMax(input string, minSize, maxSize int) bool {
    return v.ValidateMinSize(input, minSize) && v.ValidateMaxSize(input, maxSize)
}

func (v *Validator) ValidateStrings(input1, input2 string) bool {
    return input1 == input2
}


func (v *Validator) ValidateUsername(input string) bool {
    for _,c := range input {
	if unicode.IsPunct(c) || unicode.IsSymbol(c) {
	    return false
	}
    }
    return true
}

func (v *Validator) ValidatePassword(input string) bool {
    var (
	hasLetterUpper bool
	hasLetterLower bool
	hasNum bool
	hasSymbol bool
    )

    for _,ch := range input {
	switch{
	case unicode.IsUpper(ch):
	    hasLetterUpper = true
	case unicode.IsLower(ch):
	    hasLetterLower = true
	case unicode.IsNumber(ch):
	    hasNum = true
	case unicode.IsSymbol(ch) || unicode.IsPunct(ch) || unicode.IsSpace(ch):
	    hasSymbol = true
	}
    }

    if hasLetterLower && hasLetterUpper && hasNum && hasSymbol {
	return true
    }
    return false
}
