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

func init() {
	fmt.Printf("Accessing %s ... ", DbName)

	db, err = sql.Open(DatabaseDriver, DataSourceName)

	if err != nil {
		panic(err.Error())
		return
	}

	fmt.Println("Connected!")
	defer db.Close()
}

func sqlSelect() {

	rows, err := db.Query("SELECT id, name, phone FROM " + TableName)
	if err != nil {
		log.Fatal("Build Query:", err)
	}

	for rows.Next() {

		var contact contact

		err = rows.Scan(&contact.ID, &contact.Name, &contact.Phone)
		if err != nil {
			log.Fatal("Scan copy:",err)
		}
		fmt.Printf("%d\t%s\t%s \n", contact.ID, contact.Name, contact.Phone)
	}
	rows.Close()
}

func sqlInsert(id int, name, phone string) {

	rows := fmt.Sprintf("INSERT INTO %s VALUES ($1, $2, $3)", TableName)

	insert, err := db.Prepare(rows)
	if err != nil {
		log.Fatal("Prepare SQL:", err)
	}

	result, err := insert.Exec(id, name, phone)
	if err != nil {
		log.Fatalln("Insert SQL:", err)
	}

	affect, err := result.RowsAffected()
	if err != nil {
		log.Fatalln("Rows Affect SQL:", err)
	}
	fmt.Println(affect)
}
