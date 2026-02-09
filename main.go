package main

import (
        "fmt"
        "net/http"
)

func main() {
        http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
                fmt.Fprintf(w, "Hello, I am updated Go launched with Docker-Compose! Now, updated one more time!!! AND AGAIN....")
        })

        http.HandleFunc("/greet/", func(w http.ResponseWriter, r *http.Request) {
                name := r.URL.Path[len("/greet/"):]
                fmt.Fprintf(w, "Hello %s\n", name)
        })

        http.ListenAndServe(":9990", nil)
}
