package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"sync"

	_ "github.com/lib/pq"
)

var (
	listContactRegex   = regexp.MustCompile(`^\/list[\/]*$`)
	getContactRegex    = regexp.MustCompile(`^\/list\/(\d+)$`)
	createContactRegex = regexp.MustCompile(`^\/list[\/]*$`)

	db  *sql.DB
	err error
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

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func (h *userHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	switch {
	case r.Method == http.MethodGet && listContactRegex.MatchString(r.URL.Path):
		h.List(w, r)
		fmt.Println("List All accessed!")
		return
	case r.Method == http.MethodGet && getContactRegex.MatchString(r.URL.Path):
		h.Get(w, r)
		fmt.Println("Get One Contact accessed!")
		return
	case r.Method == http.MethodPost && createContactRegex.MatchString(r.URL.Path):
		h.Create(w, r)
		fmt.Println("Creat Contact accessed!")
		return
	default:
		notFound(w, r)
		return
	}
}

func (h *userHandler) List(w http.ResponseWriter, r *http.Request) {
	h.store.RLock()
	contact := make([]contact, 0, len(h.store.m))
	for _, v := range h.store.m {
		contact = append(contact, v)
	}
	h.store.RUnlock()
	jsonBytes, err := json.Marshal(contact)
	if err != nil {
		internalServerError(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (h *userHandler) Get(w http.ResponseWriter, r *http.Request) {
	matches := getContactRegex.FindStringSubmatch(r.URL.Path)
	if len(matches) < 2 {
		notFound(w, r)
		return
	}
	h.store.RLock()
	u, ok := h.store.m[matches[1]]
	h.store.RUnlock()
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("contact not found"))
		return
	}
	jsonBytes, err := json.Marshal(u)
	if err != nil {
		internalServerError(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (h *userHandler) Create(w http.ResponseWriter, r *http.Request) {
	var u contact
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		internalServerError(w, r)
		return
	}
	h.store.Lock()
	h.store.m[u.ID] = u
	h.store.Unlock()
	jsonBytes, err := json.Marshal(u)
	if err != nil {
		internalServerError(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func internalServerError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("internal server error"))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("not found"))
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

	result, err := insert.Exec(  id, name, phone)
	if err != nil {
		log.Fatalln("Error inserting:", err)
	}

	affect, err := result.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
}

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
