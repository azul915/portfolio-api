package main

import (
  "fmt"
  "log"
  "net/http"
)

func homePage(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "Test")
  fmt.Println("Testlog")
}

func handleRequests() {
  http.HandleFunc("/", homePage)
  log.Fatal(http.ListenAndServe(":1999", nil))
}

func main() {
    handleRequests()
}