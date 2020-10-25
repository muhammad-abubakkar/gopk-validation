package test

type (
	Tester interface {
		Test(value string, args []string) error
	}

	ErrorBag map[string][]string

	List map[string] Tester
)