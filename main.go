package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"text/template"

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
	// seeder.Execute()
}

func CreateFactory(factoryName, args string) {
	tmpl, err := template.ParseFiles("./stubs/factory.stub")
	if err != nil {
		panic(err)
	}

	fields := FormatFields(args)

	path := "./database/factories"
	filepath := fmt.Sprintf("%s/%s.go", path, factoryName)

	err = os.MkdirAll(path, os.FileMode(0770))
	if err != nil {
		panic(err)
	}

	file, err := os.Create(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	type TemplateData struct {
		Name   string
		Fields string
	}
	err = tmpl.Execute(file, TemplateData{Name: factoryName, Fields: fields})
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
