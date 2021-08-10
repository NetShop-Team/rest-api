package main

import (
	"github.com/gorilla/mux"
	"github.com/NetShop-Team/rest-api/getUsers"
	"github.com/NetShop-Team/rest-api/getShops"
	"net/http"
	"os"
	"fmt"
)

func main(){
	r := mux.NewRouter() 
	r.HandleFunc("/users", getUsers.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", getUsers.GetUser).Methods("GET")

	r.HandleFunc("/shops", getShops.GetShops).Methods("GET")
	r.HandleFunc("/shops/{id}", getShops.GetShop).Methods("GET")
	
	APP_IP := os.Getenv("APP_IP")
    APP_PORT := os.Getenv("APP_PORT")

    fmt.Println(APP_IP+":"+APP_PORT)
    http.ListenAndServe(APP_IP+":"+APP_PORT, nil)
}