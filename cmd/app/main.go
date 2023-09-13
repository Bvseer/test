package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"net/http"
	"test/internal/database"
	"test/internal/models"
)

const host = "localhost"
const port = "8080"

func getUsers(w http.ResponseWriter, r *http.Request) {
	users := []models.User{}
	w.WriteHeader(http.StatusOK)
	db := sqlx.DB{}
	db.Select(&users, "SELECT * FROM users")

	fmt.Println(users)
}

func init() {
	database.ConnectToDB()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users", getUsers)
	http.ListenAndServe(host+":"+port, r)
}
