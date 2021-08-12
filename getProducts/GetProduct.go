package getProducts

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
)

func GetProduct (w http.ResponseWriter, r *http.Request){

	db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/backend")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)  //{id}
	result, err := db.Query("SELECT id, category_id, name, description, likes, shop_id from products WHERE id = ?", params["id"])
	if err != nil {
	  panic(err.Error())
	}  
	defer result.Close()  
	var product Product  
	for result.Next() {
	  err := result.Scan(&product.Id, &product.CategoryId, &product.Name, &product.Description, 
		&product.Likes, &product.ShopId)
	  if err != nil {
		panic(err.Error())
	  }
	}  
	json.NewEncoder(w).Encode(product)
  }