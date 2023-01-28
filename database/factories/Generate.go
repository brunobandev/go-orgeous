package factories

type Factory struct {
	PropertyFactory PropertyFactory
	UserFactory     UserFactory
}

func NewFactory() *Factory {
	return &Factory{
		PropertyFactory: NewPropertyFactory(),
		UserFactory:     NewUserFactory(),
	}
}
