To generate a factory we should use the following command:

```go
go run main.go -factory Property -fields Name:string,CreatedAt:string
```

Then find a file called PorpetyFactory and add your faker package, we highly recommend you to use [faker](https://github.com/go-faker/faker).

To install the package:

```sh
go get -u github.com/go-faker/faker/v4
```

Fill the following code:

```go
func NewProperty() *Property {
	return &Property{
		Name: faker.Name(),
	}
}
```
