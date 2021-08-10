package main

import (
    "os"
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Welcome to the Internet! Please follow me.\n")
    })

    // создадим переменные APP_IP и APP_PORT, а их значения возьмем из одноименных переменных окружения
    APP_IP := os.Getenv("APP_IP")
    APP_PORT := os.Getenv("APP_PORT")

    fmt.Println(APP_IP+":"+APP_PORT)
    http.ListenAndServe(APP_IP+":"+APP_PORT, nil)
}