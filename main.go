package main

import (
	"github.com/gorilla/mux"
	"github.com/NetShop-Team/rest-api/getUsers"
	"github.com/NetShop-Team/rest-api/getShops"
	"net/http"
)

func main(){
	r := mux.NewRouter() 
	r.HandleFunc("/users", getUsers.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", getUsers.GetUser).Methods("GET")

	r.HandleFunc("/shops", getShops.GetShops).Methods("GET")
	r.HandleFunc("/shops/{id}", getShops.GetShop).Methods("GET")
	http.ListenAndServe(":8000", r)
}