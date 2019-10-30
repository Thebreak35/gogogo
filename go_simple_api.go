package main

import (
    "fmt"
    "net/http"
    "encoding/json"
)

type myData struct {
    FirstName   string
    LastName    string
    ID          int
    Phone       string
}

func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Welcome to HomePage")
}

func getData(w http.ResponseWriter, r *http.Request) {
    exampleData := myData {
                    FirstName:  "Mr.A",
                    LastName:   "Boiz",
                    ID:         1234,
                    Phone:      "0812345678",
    }
    json.NewEncoder(w).Encode(exampleData)
}

func handleRequest() {
    http.HandleFunc("/", homePage)
    http.HandleFunc("/getData", getData)
    http.ListenAndServe(":8080", nil)
}

func main() {
    handleRequest()
}