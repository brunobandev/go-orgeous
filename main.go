package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	seeder "github.com/brunobandev/go-orgeous/database/seeders"
)

func main() {
	var factory string
	var args string
	var seed string

	flag.StringVar(&factory, "factory", "", "Specify a factory name")
	flag.StringVar(&args, "fields", "", "Specify the fields")

	flag.StringVar(&seed, "seed", "", "Specify a seed name")

	flag.Parse()

	fmt.Println(args)

	if factory != "" {
		CreateFactory(factory, args)
	}

	if seed != "" {
		seeder.Execute()
	}
	//	seeder.Execute()

}

func CreateFactory(factoryName, args string) {
	content, err := os.ReadFile("./stubs/factory.stub")
	if err != nil {
		panic(err)
	}

	sContent := string(content)
	sContent = strings.ReplaceAll(sContent, "_{}_", factoryName)

	fields := FormatFields(args)
	sContent = strings.ReplaceAll(sContent, "_[]_", fields)

	path := "./database/factories"
	filepath := fmt.Sprintf("%s/%s.go", path, factoryName)

	err = os.MkdirAll(path, os.FileMode(0770))
	if err != nil {
		panic(err)
	}

	_, err = os.Create(filepath)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(filepath, []byte(sContent), 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("File created.")
}

func FormatFields(fields string) string {
	var properties []string
	args := strings.Split(fields, ",")
	for _, s := range args {
		arg := fmt.Sprintf("\t%s", s)
		properties = append(properties, strings.ReplaceAll(arg, ":", " "))
	}

	return strings.Join(properties, "\n")
}
