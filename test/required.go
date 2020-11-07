package test

import (
	"errors"
	"reflect"
)

type Required struct {}

func (r Required) Test(value string, args []string, values reflect.Value) error {
	if len(value) == 0 {
		return errors.New("value is required")
	}
	return nil
}



