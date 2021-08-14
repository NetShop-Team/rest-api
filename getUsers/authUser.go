package getUsers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)


func AuthUser(w http.ResponseWriter, r *http.Request) {

	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/backend")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)  //{id}
	var user User
	err := db.QueryRow("SELECT id, username, password FROM users where username = ? and password = ?", params["username"], params["password"]).Scan(&user.Id, &user.UserName, &user.Password)
	
	if err != nil {
		if string(err.Error())=="sql: no rows in result set" {
			fmt.Fprintf(w, "no users")
		}	else{
			panic(err.Error())
		}
	}	else {
		json.NewEncoder(w).Encode(user)
	}

  }
