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

	userH := &userHandler{
		store: &datastore{
			m: map[string]contact{
				"1": contact{ID: "1", Name: "tonnytg", Phone: "+551199999999"},
			},
			RWMutex: &sync.RWMutex{},
		},
	}
	mux.Handle("/contacts", userH)
	mux.Handle("/contact/", userH)

	go receiver()

	fmt.Println("Conf connection: 0.0.0.0:3000")
	fmt.Println("Try access: http://localhost:3000/contacts")
	err := http.ListenAndServe("0.0.0.0:3000", mux)
	if err != nil {
		log.Println(err)
	}
}
