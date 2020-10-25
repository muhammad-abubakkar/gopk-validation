package gopk_validation

type ErrorBag struct {
	errors map[string][]string
}

func NewErrorBag() ErrorBag {
	return ErrorBag{
		errors: make(map[string][]string),
	}
}