package getShops

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
)

func GetShop(w http.ResponseWriter, r *http.Request) {

	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/backend")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)  //{id}
	result, err := db.Query("SELECT id, name, login, password FROM shops WHERE id = ?", params["id"])
	if err != nil {
	  panic(err.Error())
	}  
	defer result.Close()  
	var shop Shop  
	for result.Next() {
	  err := result.Scan(&shop.Id, &shop.Name, &shop.Login, &shop.Password)
	  if err != nil {
		panic(err.Error())
	  }
	}  
	json.NewEncoder(w).Encode(shop)
  }
