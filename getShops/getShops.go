package getShops

import (
	"database/sql"
	"encoding/json"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
)

type Shop struct {
	Id   int
	Name string
	Login string
	Password string
}
var db *sql.DB
var err error

func GetShops(w http.ResponseWriter, r *http.Request) {
	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/backend")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	w.Header().Set("Content-Type", "application/json")
	var shops []Shop
	result, err := db.Query("SELECT id, name, login, password from shops")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	for result.Next() {
		var shop Shop
		er := result.Scan(&shop.Id, &shop.Name, &shop.Login, &shop.Password)
		if er != nil {
			panic(er.Error())
		}
		shops = append(shops, shop)
	}
	json.NewEncoder(w).Encode(shops)
}
