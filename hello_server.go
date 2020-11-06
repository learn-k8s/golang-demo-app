package main

import (
    "fmt"
    "net/http"
    "log"
)

func main() {
    http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello World!")
        query := r.URL.Query()
        name := query.Get("name")
        if name != "" {
            log.Printf("Received request for %s\n", name)
        }
    })
    log.Println("Starting Server")
    if err :=  http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
    log.Println("Server stopped")
}
