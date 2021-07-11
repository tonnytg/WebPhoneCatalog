package main

import (
	"fmt"
)

type Contact struct {
	ID    string
	Name  string
	Phone string
}

const PostgresDriver = "postgres"
const User = "postgres"
const Host = "database"
const Port = "5432"
const Password = "Postgres2021!"
const DbName = "contacts"
const TableName = "contact"

var DataSourceName = fmt.Sprintf("host=%s port=%s user=%s "+
	"password=%s dbname=%s sslmode=disable", Host, Port, User, Password, DbName)
