package main

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
)

var (
	listContactRegex   = regexp.MustCompile(`^\/contacts[\/]*$`)
	getContactRegex    = regexp.MustCompile(`^\/contact\/(\d+)$`)
	createContactRegex = regexp.MustCompile(`^\/contact[\/]*$`)
)

func (h *userHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("content-type", "application/json")
	switch {
	case r.Method == http.MethodGet && listContactRegex.MatchString(r.URL.Path):
		h.List(w, r)
		return
	case r.Method == http.MethodGet && getContactRegex.MatchString(r.URL.Path):
		h.Get(w, r)
		return
	case r.Method == http.MethodPost && createContactRegex.MatchString(r.URL.Path):
		h.Create(w, r)
		return
	default:
		notFound(w, r)
		log.Fatal("not found:", w, r)
		return
	}
}

func (h *userHandler) List(w http.ResponseWriter, r *http.Request) {

	mc := sqlSelect()
	list := make([]contact, 0, len(mc))
	for _, v := range mc {
		list = append(list, v)
	}

	jsonBytes, err := json.Marshal(list)
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
	mc := sqlSelectWhere(matches[1])
	jsonBytes, err := json.Marshal(mc)
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

	err := sqlInsert(u.ID, u.Name, u.Phone)
	if err != nil {
		log.Fatal("Houston we have a problem!\nErr:", err)
	}

	jsonBytes, err := json.Marshal(u)
	if err != nil {
		internalServerError(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
	sender(string(jsonBytes[:]))
}

func internalServerError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("internal server error"))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("not found"))
}
