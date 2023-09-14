package main

import (
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"net/http"
	"test/api"
	"test/internal/database"
)

const host = "localhost"
const port = "8080"

func init() {
	database.ConnectToDB()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users", api.GetUsers).Methods(http.MethodGet)
	r.HandleFunc("/users/{id:[0-9]+}", api.GetUser).Methods(http.MethodGet)
	r.HandleFunc("/health", api.Health)
	err := http.ListenAndServe(host+":"+port, r)
	if err != nil {
		return
	}
}
