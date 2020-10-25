package test

import (
	"errors"
	"fmt"
	"strconv"
)

type Min struct {}

func (an Min) Test(value string, args []string) error {
	if len(args) == 0 {
		return errors.New("min arg not provided")
	}

	minLen, err := strconv.Atoi(args[0])

	if err != nil {
		return err
	}

	if len(value) < minLen {
		return fmt.Errorf("value should not be more than %d characters", minLen)
	}

	return nil
}
