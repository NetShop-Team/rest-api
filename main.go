package main

import (
	"github.com/gorilla/mux"
	"github.com/NetShop-Team/rest-api/getUsers"
	"net/http"
)

func main(){
	r := mux.NewRouter() 
	r.HandleFunc("/users", getUsers.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", getUsers.GetUser).Methods("GET")
	http.ListenAndServe(":8000", r)
}