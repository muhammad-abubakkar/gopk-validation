package test

import "errors"

type Required struct {}

func (r Required) Test(value string, args []string) error {
	if len(value) == 0 {
		return errors.New("value is required")
	}
	return nil
}



