package factories

type Factory struct {
	PropertyFactory PropertyFactory
}

func NewFactory() *Factory {
	return &Factory{
		PropertyFactory: NewPropertyFactory(),
	}
}
