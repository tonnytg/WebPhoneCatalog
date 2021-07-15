package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	_ "github.com/lib/pq"
)

type contact struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type datastore struct {
	m map[string]contact
	*sync.RWMutex
}

type userHandler struct {
	store *datastore
}

func main() {

	mux := http.NewServeMux()

	catalog := &userHandler{ store: &datastore{}}

	mux.Handle("/contacts", catalog)
	mux.Handle("/contact/", catalog)

	go receiver()

	fmt.Println("Conf connection: 0.0.0.0:3000")
	fmt.Println("Try access: http://localhost:3000/contacts")
	err := http.ListenAndServe("0.0.0.0:3000", mux)
	if err != nil {
		log.Println(err)
	}
}
