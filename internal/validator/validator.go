package validator

import (
	"regexp"
)

var (
	EmailRX = regexp.MustCompile("/^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$/")
)


type Validator struct {
	Errors map[string]string
}


func New() *Validator {
	return &Validator{Errors: make(map[string]string)}
}


func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

func (v *Validator) AddError(field, message string) {
	if _, ok := v.Errors[field]; !ok {
		v.Errors[field] = message
	}
}


func (v *Validator) Check(ok bool, key, message string) {
	if !ok {
		v.AddError(key, message)
	}
}


func In(value string, values ...string) bool {
	for _, v := range values {
		if v == value {
			return true
		}
	}
	return false
}


func Matches(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}


func Unique(values []string) bool {
	uniqueValues := map[string]struct{}{}
	for _, value := range values {
		uniqueValues[value] = struct{}{}
	}
	return len(uniqueValues) == len(values)
}