package getUsers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"io/ioutil"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)


func CreateUser(w http.ResponseWriter, r *http.Request) {

	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/backend")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO users(name) VALUES(?)")
	if err != nil {
	  panic(err.Error())
	}  
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
	  panic(err.Error())
	}  
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	name := keyVal["name"]  
	_, err = stmt.Exec(name)
	if err != nil {
	  panic(err.Error())
	}  
	fmt.Fprintf(w, "New post was created")
  }
