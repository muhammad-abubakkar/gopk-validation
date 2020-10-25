package test

import (
	"errors"
	"fmt"
	"strconv"
)

type Max struct {}

func (an Max) Test(value string, args []string) error {
	if len(args) == 0 {
		return errors.New("max arg not provided")
	}

	maxLen, err := strconv.Atoi(args[0])

	if err != nil {
		return err
	}

	if len(value) > maxLen {
		return fmt.Errorf("value should not be more than %d characters", maxLen)
	}

	return nil
}
