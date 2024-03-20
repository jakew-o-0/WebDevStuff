package validator

import (
	"net/mail"
	"unicode"
	"unicode/utf8"

	"golang.org/x/crypto/bcrypt"
)

// checks number of utf8 runes are larger than min
func (v Validator) MinLen(input string, min int) bool {
    return utf8.RuneCount([]byte(input)) >= min
}


// checks number of utf8 runes are less than max
func (v Validator) MaxLen(input string, max int) bool {
    return utf8.RuneCount([]byte(input)) <= max
}


// does both min and max
func (v Validator) MinMaxLen(input string, min,max int) bool {
    return v.MinLen(input, min) && v.MaxLen(input, max)
}


// checks a string for valid unicode letters (no numbers or sybols)
func (v Validator) ValidName(input string) bool {
    for _,ch := range input {
        if !unicode.IsLetter(ch) {
            return false
        }
    }
    return true
}


// checks a string has atleast one valid unicode uppercase, lowercase, number, symbol
func (v Validator) ValidPassword(input string) bool {
    var (
        hasUpper bool
        hasLower bool
        hasNum bool
        hasSymbol bool
    )

    for _,ch := range input {
        switch{
        case unicode.IsUpper(ch):
            hasUpper = true
        case unicode.IsLower(ch):
            hasLower = true
        case unicode.IsNumber(ch):
            hasNum = true
        case unicode.IsSymbol(ch) || unicode.IsSpace(ch) || unicode.IsPunct(ch):
            hasSymbol = true
        }
    }

    if hasUpper && hasLower && hasNum && hasSymbol {
        return true
    }
    return false
}


// uses the go standard lib to check if string is in a valid email format
func (v Validator) ValidEmail(input string) bool {
    _,err := mail.ParseAddress(input)
    return err == nil
}


func (v Validator) ValidPasswordHash(hash, password []byte) bool {
    err := bcrypt.CompareHashAndPassword(hash, password)
    return err == nil
}


func (v Validator) StringsMatch(input1,input2 string) bool {
    return input1 == input2
}
