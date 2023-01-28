package factories

type UserFactory struct {
}

func NewUserFactory() UserFactory {
	return UserFactory{}
}

type User struct {
	Name string
}

func NewUser() *User {
	return &User{
	}
}

func (f *UserFactory) Generate(quantity int) []User {
	var generated []User

	for i := 0; i < quantity; i++ {
		generated = append(generated, *NewUser())
	}

	return generated
}

