package test

import (
	"errors"
	"regexp"
)

type AlphaNum struct {}

func (an AlphaNum) Test(value string, args []string) error {
	regTest := regexp.MustCompile(`^[a-z][a-z0-9_]+$`)
	if !regTest.MatchString(value) {
		return errors.New("value contains invalid characters")
	}
	return nil
}
