package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"sync"

	_ "github.com/lib/pq"
)

func init () {
	fmt.Printf("Accessing %s ... ", DbName)

	db, err = sql.Open(PostgresDriver, DataSourceName)

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Connected!")
	}

	defer db.Close()

	sqlSelect()
	sqlInsert(2,"test","+551188889999")
}

func main() {

	mux := http.NewServeMux()
	userH := &userHandler{
		store: &datastore{
			m: map[string]contact{
				"1": contact{ID: "1", Name: "tonnytg", Phone: "+551199999999"},
			},
			RWMutex: &sync.RWMutex{},
		},
	}
	mux.Handle("/list", userH)
	mux.Handle("/list/", userH)

	fmt.Println("Conf connection: 0.0.0.0:3000")
	fmt.Println("Try access: http://localhost:3000/list")
	err := http.ListenAndServe("0.0.0.0:3000", mux)
	if err != nil {
		fmt.Println(err)
	}
}
