package getUsers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
)


func GetUser(w http.ResponseWriter, r *http.Request) {

	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/backend")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)  //{id}
	result, err := db.Query("SELECT id, username, password FROM users WHERE id = ?", params["id"])
	if err != nil {
	  panic(err.Error())
	}  
	defer result.Close()  
	var user User  
	for result.Next() {
	  err := result.Scan(&user.Id, &user.UserName, &user.Password)
	  if err != nil {
		panic(err.Error())
	  }
	}  
	json.NewEncoder(w).Encode(user)
  }
