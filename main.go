package main

import (
    "net/http"
    "database/sql"
    "github.com/gorilla/mux"
	"github.com/NetShop-Team/rest-api/getProducts"
    "github.com/NetShop-Team/rest-api/getShops"
    "github.com/NetShop-Team/rest-api/getUsers"
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

    r.HandleFunc("/products", getProducts.GetProducts).Methods("GET")
    r.HandleFunc("/products/{id}", getProducts.GetProduct).Methods("GET")

	r.HandleFunc("/shops", getShops.GetShops).Methods("GET")
    r.HandleFunc("/shops/{id}", getShops.GetShop).Methods("GET")
	
	r.HandleFunc("/users", getUsers.GetUsers).Methods("GET")
    r.HandleFunc("/users/{id}", getUsers.GetUser).Methods("GET")
    http.Handle("/", r)
    http.ListenAndServe("localhost:8000", nil)
}