package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"test/internal/database"
	"test/internal/models"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := []models.User{}
	err := database.DB.Select(&users, "SELECT * FROM users")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	result, _ := json.Marshal(users)
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(result); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	user := models.User{}
	fmt.Println(args["id"])
	if err := database.DB.Select(&user, fmt.Sprintf("SELECT * FROM users WHERE id = ?", args["id"])); err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	result, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(result)
}
