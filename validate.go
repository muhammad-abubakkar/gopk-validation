package gopk_validation

import (
	"errors"
	"github.com/muhammad-abubakkar/gopk-validation/test"
	"reflect"
	"strings"
)

func runTest(testName, value string, args []string, values reflect.Value) error {
	if tester, ok := allTests[testName]; ok {
		return tester.Test(value, args, values)
	}
	return nil
}

func Validate(data interface{}) (test.ErrorBag, error) {
	dataErrors := test.ErrorBag{}
	values := reflect.ValueOf(data)
	dataType := reflect.TypeOf(data)

	if dataType.Kind() != reflect.Struct {
		return dataErrors, errors.New("value should be of type struct")
	}

	nFields := dataType.NumField()

	for i := 0; i < nFields; i++ {
		field := dataType.Field(i)
		name := field.Name
		value := values.Field(i)
		jName := field.Tag.Get("json")
		rules := field.Tag.Get("tests")

		if rules == "" {
			continue
		}

		if jName != "" {
			name = jName
		}

		segments := strings.Split(rules, "|")
		for _, segment := range segments {
			rule := strings.Split(segment, ":")
			ruleName := rule[0]
			var ruleArgs []string
			if len(rule) > 1 {
				ruleArgs = strings.Split(rule[1], ",")
			}
			err := runTest(ruleName, value.String(), ruleArgs, values)
			if err != nil {
				errs := dataErrors[name]
				dataErrors[name] = append(errs, err.Error())
			}
		}
	}

	if len(dataErrors) > 0 {
		return dataErrors, errors.New("fix validation errors")
	}

	return dataErrors, nil
}
