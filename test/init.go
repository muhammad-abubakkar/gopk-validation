package test

import "reflect"

type (
	Tester interface {
		Test(value string, args []string, dataValues reflect.Value) error
	}

	ErrorBag map[string][]string

	List map[string] Tester
)