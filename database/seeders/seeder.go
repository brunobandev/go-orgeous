package seeder

import (
	"database/sql"
	"fmt"
	"reflect"
	"strconv"

	"github.com/brunobandev/go-orgeous/database/factories"
)

func Execute() {
	factory := factories.NewFactory()

	pf := factory.PropertyFactory
	rows := pf.Generate(15)
	fmt.Println(rows)

	//conn, err := GetConnection()
	//if err != nil {
	//	panic(err)
	//}
	//
	for _, row := range rows {
	fmt.Println(row.Email)
	fields := reflect.TypeOf(*row)
	values := reflect.ValueOf(*row)
	num := fields.NumField()
	for i := 0; i < num; i++ {
		field := fields.Field(i)
		value := values.Field(i)
		fmt.Print("Type ", field.Type, ",", field.Name, "=", value, "\n")

		switch value.Kind() {
		case reflect.String:
			v := value.String()
			fmt.Print(v, "\n")
		case reflect.Int:
			v := strconv.FormatInt(value.Int(), 10)
			fmt.Print(v, "\n")
		case reflect.Int32:
			v := strconv.FormatInt(value.Int(), 10)
			fmt.Print(v, "\n")
		case reflect.Int64:
			v := strconv.FormatInt(value.Int(), 10)
			fmt.Print(v, "\n")
		default:
			fmt.Print("Not support type of struct")
		}
	}
	}

	fmt.Println(rows)
}

func GetConnection() (*sql.DB, error) {
	db, err := sql.Open("postgres", "")
	if err != nil {
		return nil, err
	}

	return db, nil
}
