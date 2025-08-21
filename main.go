package main

import (
  "fmt"
  "log"
  "net/http"
  "os"
)

// http server to display the following
// hello-world
// the machine hostname it is running on

func helloHandler(w http.ResponseWriter, r *http.Request) {
  hostname, err := os.Hostname()
  if err != nil {
    panic(err)
  }

  fmt.Fprintln(w, "hello world!\n")
  fmt.Fprintln(w, "hostname:", hostname)
}

func main() {
  http.HandleFunc("/", helloHandler)
  
  log.Println("listening at port 8080")
  http.ListenAndServe(":8080", nil)
}