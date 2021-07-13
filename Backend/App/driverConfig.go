package main

import "fmt"

const DatabaseDriver = "postgres"
const User = "postgres"
const Host = "database"
const Password = "Postgres2021!"
const DbName = "contacts"
const TableName = "contact"

var Port = "5432"
var DataSourceName = ""

func init() {
	if DatabaseDriver == "postgres" {
		var DataSourceName = fmt.Sprintf("host=%s port=%s user=%s "+
			"password=%s dbname=%s sslmode=disable", Host, Port, User, Password, DbName)
		fmt.Println(DataSourceName)
	} else {
		Port = "3306"
		var DataSourceName = fmt.Sprintf("%s:%s@%s(:%s)/%s",
			User, Password, Host, Port, DbName)
		fmt.Println(DataSourceName)
	}
}
