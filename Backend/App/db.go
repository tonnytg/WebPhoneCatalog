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

func sqlSelect() map[string]contact {

	mc := make(map[string]contact)

	fmt.Printf("Accessing %s ... ", DbName)

	db, err = sql.Open(DatabaseDriver, DataSourceName)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Connected!")

	rows, err := db.Query("SELECT id, name, phone FROM " + TableName)
	if err != nil {
		log.Fatal("Build Query:", err)
	}

	for rows.Next() {

		var c contact

		err = rows.Scan(&c.ID, &c.Name, &c.Phone)
		if err != nil {
			log.Fatal("Scan copy:", err)
		}
		fmt.Printf("%d\t%s\t%s \n", c.ID, c.Name, c.Phone)

		mc[c.Name] = contact{c.ID, c.Name, c.Phone}
	}
	defer db.Close()
	return mc
}

func sqlInsert(id int, name, phone string) error {

	fmt.Printf("Accessing %s ... ", DbName)
	db, err = sql.Open(DatabaseDriver, DataSourceName)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Connected!")
	defer db.Close()
	rows := fmt.Sprintf("INSERT INTO %s VALUES ($1, $2, $3)", TableName)

	insert, err := db.Prepare(rows)
	if err != nil {
		log.Fatal("Prepare SQL:", err)
		return err
	}

	result, err := insert.Exec(id, name, phone)
	if err != nil {
		log.Fatalln("Insert SQL:", err)
		return err
	}

	affect, err := result.RowsAffected()
	if err != nil {
		log.Fatalln("Rows Affect SQL:", err)
		return err
	}
	fmt.Println(affect)
	return nil
}
