package factories

import (
	"github.com/go-faker/faker/v4"
)

type PropertyFactory struct {
}

func NewPropertyFactory() PropertyFactory {
	return PropertyFactory{}
}

type Property struct {
	Name  string
	Email string
}

func NewProperty() *Property {
	return &Property{
		Name:  faker.Name(),
		Email: faker.Email(),
	}
}

func (f *PropertyFactory) Generate(quantity int) []*Property {
	var generated []*Property

	for i := 0; i < quantity; i++ {
		generated = append(generated, NewProperty())
	}

	return generated
}
