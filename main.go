package main

import (
  "fmt"
  "log"
  "net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
  if err := r.ParseForm(); err != nil {
    fmt.Fprintf(w, "ParseForm() err: $v\n", err)
    return
  }
  fmt.Fprintf(w, "POST request successful\n")
  name := r.FormValue("name")
  address := r.FormValue("address")
  fmt.Fprintf(w, "Name: %s\nAddress: %s\n", name, address)
}

func main() {
  fileServer := http.FileServer(http.Dir("./static"))
  http.Handle("/", fileServer)
  http.HandleFunc("/form", formHandler)

  port := ":8080"
  fmt.Printf("Listening on port %s\n", port)
  if err := http.ListenAndServe(":8080", nil); err != nil {
    log.Fatal(err)
  }
}
