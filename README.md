### Used this library to validate structs in golang
Inspired by PHP laravel framework
```go
package main

import (
    "fmt"
	gopk_validation "github.com/muhammad-abubakkar/gopk-validation"
)

type Person struct {
    Name    string `name:"first_name" tests:"required|max:50"`
    Email   string `name:"email" tests:"max:50"`
}

func main() {
    person := Person{
        Name: "Abu Bakkar",
        Email: "abubakkar@example.com",
    }
    bag, err := gopk_validation.Validate(person)

    fmt.Println(err)

    for field, errs := range bag {
        fmt.Printf("%v: \n", field)
        for i, e := range errs {
            fmt.Printf("\t%d: %v\n", i+1, e)
        }
    }
}
```