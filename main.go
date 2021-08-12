package main

import (
    "net/http"
    "database/sql"
    "github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
)
var db *sql.DB
var err error

type Product struct{
    Id int
    CategoryId int
    Name string
    Description string
    Likes int
    ShopId int
}

    
func main() {

    r := mux.NewRouter()

    r.HandleFunc("/products", GetProducts).Methods("GET")
    r.HandleFunc("/products/{id}", GetProduct).Methods("GET")
    http.Handle("/", r)
    http.ListenAndServe("localhost:8000", nil)
}