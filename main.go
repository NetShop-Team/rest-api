package main
   // "net/http"
    // "database/sql"
    // "github.com/gorilla/mux"
	// "github.com/NetShop-Team/rest-api/getProducts"
    // "github.com/NetShop-Team/rest-api/getShops"
    // "github.com/NetShop-Team/rest-api/getUsers"
	// _ "github.com/go-sql-driver/mysql"
import (
     "github.com/gorilla/mux"
      "net/http"
    "fmt"
)
// var db *sql.DB
// var err error

// type Product struct{
//     Id int
//     CategoryId int
//     Name string
//     Description string
//     Likes int
//     ShopId int
// }

func setDefaultHeaders(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "helloooooooo")
   }
func main() {
    fmt.Println("helloooooo")
    r := mux.NewRouter()

    // r.HandleFunc("/products", getProducts.GetProducts).Methods("GET")
    // r.HandleFunc("/products/{id}", getProducts.GetProduct).Methods("GET")

	// r.HandleFunc("/shops", getShops.GetShops).Methods("GET")
    // r.HandleFunc("/shops/{id}", getShops.GetShop).Methods("GET")
	
	// r.HandleFunc("/users", getUsers.GetUsers).Methods("GET")
	// r.HandleFunc("/users", getUsers.CreateUser).Methods("POST")
    // r.HandleFunc("/users/{id}", getUsers.GetUser).Methods("GET")
	// r.HandleFunc("/users/auth/{username}/{password}", getUsers.AuthUser).Methods("GET")

	 r.HandleFunc("/", setDefaultHeaders)
 	http.ListenAndServe("185.234.114.233:8000", r)
}