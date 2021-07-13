package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
)

var (
	listContactRegex   = regexp.MustCompile(`^\/contacts[\/]*$`)
	getContactRegex    = regexp.MustCompile(`^\/contact\/(\d+)$`)
	createContactRegex = regexp.MustCompile(`^\/contact[\/]*$`)
)

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
