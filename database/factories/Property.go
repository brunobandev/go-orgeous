package factories

type PropertyFactory struct {
}

func NewPropertyFactory() PropertyFactory {
	return PropertyFactory{}
}

type Property struct {
	Name string
}

func NewProperty() *Property {
	return &Property{
	}
}

func (f *PropertyFactory) Generate(quantity int) []Property {
	var generated []Property

	for i := 0; i < quantity; i++ {
		generated = append(generated, *NewProperty())
	}

	return generated
}

