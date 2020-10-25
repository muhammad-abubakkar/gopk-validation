package test

import (
	"errors"
	"fmt"
	"strconv"
)

type Len struct {}

func (an Len) Test(value string, args []string) error {
	if len(args) == 0 {
		return errors.New("len arg not provided")
	}

	extLen, err := strconv.Atoi(args[0])

	if err != nil {
		return err
	}

	if len(value) != extLen {
		return fmt.Errorf("value should be exactly %d characters long", extLen)
	}

	return nil
}
