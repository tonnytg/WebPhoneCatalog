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

func sqlSelect() {

	fmt.Printf("Accessing %s ... ", DbName)

	db, err = sql.Open(DatabaseDriver, DataSourceName)

	if err != nil {
		panic(err.Error())
		return
	}

	fmt.Println("Connected!")

	rows, err := db.Query("SELECT id, name, phone FROM " + TableName)
	if err != nil {
		log.Fatal("Build Query:", err)
	}

	for rows.Next() {

		var contact contact

		err = rows.Scan(&contact.ID, &contact.Name, &contact.Phone)
		if err != nil {
			log.Fatal("Scan copy:", err)
		}
		fmt.Printf("%d\t%s\t%s \n", contact.ID, contact.Name, contact.Phone)
	}
	defer db.Close()
}

func sqlInsert(id, name, phone string) error {

	fmt.Printf("Accessing %s ... ", DbName)

	db, err = sql.Open(DatabaseDriver, DataSourceName)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Connected!")

	rows := fmt.Sprintf("INSERT INTO %s VALUES ($1, $2, $3)", TableName)

	insert, err := db.Prepare(rows)
	if err != nil {
		log.Fatal("Prepare SQL:", err)
		return err
	}

	result, err := insert.Exec(2, name, phone)
	if err != nil {
		log.Fatalln("Insert SQL:", err)
		return err
	}
	defer db.Close()
	affect, err := result.RowsAffected()
	if err != nil {
		log.Fatalln("Rows Affect SQL:", err)
		return err
	}
	fmt.Println(affect)
	return nil
}
