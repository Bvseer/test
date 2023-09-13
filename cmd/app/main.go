package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

const host = "localhost"
const port = "8080"

func getUsers(w http.ResponseWriter, r *http.Request) {
	hi := "hello world"
	w.WriteHeader(http.StatusOK)
	result, _ := json.Marshal(hi)
	w.Write(result)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users", getUsers)
	http.ListenAndServe(host+":"+port, r)
}

func doReq(url string) (content string) {

	resp, err := http.Get(url)

	if err != nil {

		log.Println(err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {

		log.Println(err)
		return
	}

	return string(body)
}

func getTitle(content string) (title string) {

	re := regexp.MustCompile("<title>(.*)</title>")

	parts := re.FindStringSubmatch(content)

	if len(parts) > 0 {
		return parts[1]
	} else {
		return "no title"
	}
}
