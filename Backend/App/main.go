package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	_ "github.com/lib/pq"
)

type contact struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type catalog struct {
	m map[int]contact
}

type datastore struct {
	catalog []catalog
	*sync.RWMutex
}

type userHandler struct {
	store *datastore
}

func main() {

	mux := http.NewServeMux()

	catalog := &userHandler{store: &datastore{}}

	mux.Handle("/contacts", catalog)
	mux.Handle("/contact/", catalog)

	go receiver()

	fmt.Println("Conf connection: 0.0.0.0:3001")
	fmt.Println("Try access: http://localhost:3001/contacts")
	err := http.ListenAndServe("0.0.0.0:3001", mux)
	if err != nil {
		log.Println(err)
	}
}
