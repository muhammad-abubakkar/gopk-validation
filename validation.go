package gopk_validation

import (
"errors"
"reflect"
"strings"
)

func Validate(value interface{}) (ErrorBag, error) {
	vErrors := NewErrorBag()
	vValue := reflect.ValueOf(value)
	vType := reflect.TypeOf(value)
	vKind := vType.Kind()

	if vKind != reflect.Struct {
		return vErrors, errors.New("value should be of type struct")
	}

	nFields := vType.NumField()

	for i := 0; i < nFields; i++ {
		field := vType.Field(i)
		name := field.Name
		value := vValue.Field(i)
		fName := field.Tag.Get("name")
		rules := field.Tag.Get("rules")

		if rules == "" {
			continue
		}

		if fName != "" {
			name = fName
		}

		segments := strings.Split(rules, "|")
		for _, segment := range segments {
			rule := strings.Split(segment, ":")
			ruleName := rule[0]
			var ruleArgs []string
			if len(rule) > 1 {
				ruleArgs = strings.Split(rule[1], ",")
			}
			err := Check(ruleName, value.String(), ruleArgs)
			if err != nil {
				errs := vErrors.errors[ruleName]
				vErrors.errors[name] = append(errs, err.Error())
			}
		}
	}

	if len(vErrors.errors) > 0 {
		return vErrors, errors.New("fix validation errors")
	}

	return vErrors, nil
}
