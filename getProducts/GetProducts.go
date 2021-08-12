package getProducts

import (
	"database/sql"
	"encoding/json"
	"net/http"
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

func GetProducts (w http.ResponseWriter, r *http.Request) {
	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/backend")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	w.Header().Set("Content-Type", "application/json")
	var products []Product
	result, err := db.Query("SELECT id, category_id, name, description, likes, shop_id from products")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	for result.Next() {
		var product Product
		er := result.Scan(&product.Id, &product.CategoryId, &product.Name, &product.Description, 
			&product.Likes, &product.ShopId)
		if er != nil {
			panic(er.Error())
		}
		products = append(products, product)
	}
	json.NewEncoder(w).Encode(products)
}