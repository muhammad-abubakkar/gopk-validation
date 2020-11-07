package test

import (
	"errors"
	"reflect"
)

type Match struct {}

func (an Match) Test(value string, args []string, values reflect.Value) error {
	compare := values.FieldByName(args[0]).Interface().(string)
	if value != compare {
		return errors.New("value should match")
	}
	return nil
}
