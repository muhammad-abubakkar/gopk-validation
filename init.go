package gopk_validation

import "github.com/muhammad-abubakkar/gopk-validation/test"

var allTests test.List

func init() {
	allTests = test.List{
		"required":  test.Required{},
		"max":       test.Max{},
		"min":       test.Min{},
		"len":       test.Len{},
		"alpha_num": test.AlphaNum{},
		"match":     test.Match{},
	}
}

func SetTest(name string, tester test.Tester) {
	allTests[name] = tester
}
