package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var fileName string
	var args string

	flag.StringVar(&fileName, "factory", "", "Specify a factory name")
	flag.StringVar(&args, "fields", "", "Specify the fields")
	flag.Parse()

	fmt.Println(fileName)
	fmt.Println(args)

	if fileName == "" {
		fmt.Println("Invalid")
	} else {
		file, err := os.Open("./stubs/factory.stub")
		if err != nil {
			log.Fatal(err)
		}

		data := make([]byte, 369)
		count, err := file.Read(data)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("read %d bytes: %q\n", count, data[:count])
		content := strings.ReplaceAll(string(data[:count]), "_{}_", fileName)

		newArgs := strings.ReplaceAll(args, ",", "\n")
		newArgs = strings.ReplaceAll(newArgs, ":", " ")
		content = strings.ReplaceAll(content, "_[]_", newArgs)
		fmt.Println(content)

		path := "./database/factories"
		filepath := fmt.Sprintf("%s/%s.go", path, fileName)

		err = os.MkdirAll(path, os.FileMode(0770))
		if err != nil {
			panic(err)
		}

		_, err = os.Create(filepath)
		if err != nil {
			panic(err)
		}

		err = os.WriteFile(filepath, []byte(content), 0644)
		if err != nil {
			panic(err)
		}

		fmt.Println("File created.")
	}
}
