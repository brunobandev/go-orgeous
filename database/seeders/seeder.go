package main

import (
	"fmt"
	"github.com/brunobandev/go-orgeous/database/factories"
)

func main() {
	factory := factories.NewFactory()

	data := factory.PropertyFactory.Generate(15)
	users := factory.UserFactory.Generate(10)

	fmt.Println(data, users)
}
