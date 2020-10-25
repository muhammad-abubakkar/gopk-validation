package gopk_validation

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

type Validator struct {}

func (v Validator) Max(value string, args ...string) error {
	if len(args) == 0 {
		return errors.New("max arg not provided")
	}

	maxLen, err := strconv.Atoi(args[0])

	if err != nil {
		return err
	}

	if len(value) > maxLen {
		return fmt.Errorf("value should be less then %d characters", maxLen)
	}

	return nil
}

func (v Validator) Min(value string, args ...string) error {
	if len(args) == 0 {
		return errors.New("min arg not provided")
	}

	minLen, err := strconv.Atoi(args[0])

	if err != nil {
		return err
	}

	if len(value) < minLen {
		return fmt.Errorf("value should be more then %d characters", minLen)
	}

	return nil
}

func (v Validator) Len(value string, args ...string) error {
	if len(args) == 0 {
		return errors.New("len arg not provided")
	}

	extLen, err := strconv.Atoi(args[0])

	if err != nil {
		return err
	}

	if len(value) != extLen {
		return fmt.Errorf("value should be %d characters long", extLen)
	}

	return nil
}


func Check(ruleName, value string, args []string) error {
	validator := reflect.ValueOf(Validator{})
	method := validator.MethodByName(ruleName)
	inputs := make([]reflect.Value, len(args) + 1)
	inputs[0] = reflect.ValueOf(value)
	for i, arg := range args {
		inputs[i+1] = reflect.ValueOf(arg)
	}
	if method.IsValid() {
		val := method.Call(inputs)[0]
		err, ok := val.Interface().(error)
		if ok {
			return err
		}
	}
	return nil
}
