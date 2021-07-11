package main

import (
	"database/sql"
	"fmt"
	"log"
)

var (
	db  *sql.DB
	err error
)

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func sqlSelect() {

	sqlStatement, err := db.Query("SELECT id, name, phone FROM " + TableName)
	checkErr(err)

	for sqlStatement.Next() {

		var contact Contact

		err = sqlStatement.Scan(&contact.ID, &contact.Name, &contact.Phone)
		checkErr(err)

		fmt.Printf("%d\t%s\t%s \n", contact.ID, contact.Name, contact.Phone)
	}
}

func sqlInsert(id int, name, phone string) {

	sqlStatement := fmt.Sprintf("INSERT INTO %s VALUES ($1, $2, $3)", TableName)

	insert, err := db.Prepare(sqlStatement)
	checkErr(err)

	result, err := insert.Exec(id, name, phone)
	if err != nil {
		log.Fatalln("Error inserting:", err)
	}

	affect, err := result.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
}
