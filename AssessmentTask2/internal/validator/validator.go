package validator

type Validator struct {
    Errors map[string]string
}

func New() Validator {
    return Validator{
        Errors: make(map[string]string),
    }
}

func (v Validator) Validate(valid bool, error, message string) {
    if !valid {
        v.Errors[error] = message
    }
}

func (v *Validator) IsValid() bool {
    if len(v.Errors) > 0 {
        return false
    }
    return true
}
