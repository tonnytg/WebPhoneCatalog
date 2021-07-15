package main

import "fmt"

const DatabaseDriver = "postgres"
const User = "postgres"
const Host = "database"
const Password = "Postgres2021!"
const DbName = "contacts"
const TableName = "contact"

var Port = "5432"

var DataSourceName = fmt.Sprintf("host=%s port=%s user=%s "+
	"password=%s dbname=%s sslmode=disable", Host, Port, User, Password, DbName)
