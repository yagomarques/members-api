package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Starting...")

	routes := mux.NewRouter().StrictSlash(true)

	routes.HandleFunc("/members", getMembers).Methods("GET")
	routes.HandleFunc("/members/{idNumber}", getOneMember).Methods("GET")
	routes.HandleFunc("/members", postMembers).Methods("POST")

	log.Fatal(http.ListenAndServe(":3000", routes))
}

type Member struct {
	ID        int    `json:"id"`
	NAME      string `json:"name"`
	EMAIL     string `json:"email"`
	CELLPHONE string `json:"cellphone"`
}

var members = []Member{} //simule database

type Response struct {
	CODE    string `json:"code"`
	MESSAGE string `json:"message"`
}

var responses = []Response{
	Response{CODE: "0", MESSAGE: "member registed with sucess"},
	Response{CODE: "1", MESSAGE: "quebrou"},
}

func getMembers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(members)
	fmt.Println("ping get")
}

func getOneMember(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var vars = mux.Vars(r)
	idNumber, _ := strconv.Atoi(vars["idNumber"])

	for _, values := range members {
		if values.ID == idNumber {
			json.NewEncoder(w).Encode(values)
		}
	}
}

func postMembers(w http.ResponseWriter, r *http.Request) {
	var newmember Member
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &newmember)

	members = append(members, newmember)
	json.NewEncoder(w).Encode(responses[0])
	fmt.Println("ping post")
}
