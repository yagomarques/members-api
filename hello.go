package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	routes := mux.NewRouter().StrictSlash(true)

	routes.HandleFunc("/members", getMembers).Methods("GET")
	routes.HandleFunc("/members", postMembers).Methods("POST")

	log.Fatal(http.ListenAndServe(":3000", routes))
}

type Member struct {
	ID        int    `json:"id"`
	NAME      string `json:"name"`
	EMAIL     string `json:"email"`
	CELLPHONE string `json:"cellphone"`
}

type Response struct {
	CODE    string `json:"code"`
	MESSAGE string `json:"message"`
}

var members = []Member{}

var responses = []Response{
	Response{CODE: "0", MESSAGE: "member registred with sucess"},
	Response{CODE: "1", MESSAGE: "omg"},
}

func getMembers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(members)
	fmt.Println("TEST NO GET")
}

func postMembers(w http.ResponseWriter, r *http.Request) {
	var newmember Member
	body, _ := ioutil.ReadAll(r.Body)

	json.Unmarshal(body, &newmember)
	members = append(members, newmember)
	//fmt.Println(members)
	fmt.Println("TEST NO POST")
	json.NewEncoder(w).Encode(responses[0])
}
