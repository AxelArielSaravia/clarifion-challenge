package main

import (
    "fmt"
    "net/http"
)

const PORT string = ":6969"


func logger(next http.Handler) http.Handler {
    return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
        fmt.Println(r.Method, r.URL)
        next.ServeHTTP(w, r)
    })
}

func main() {
    mux := http.NewServeMux();

    fileServer := http.FileServer(http.Dir("./src"))
    mux.Handle("/", fileServer)

    fmt.Printf("Starting server on http://127.0.0.1%s\n", PORT)
    err := http.ListenAndServe(PORT, logger(mux))
    panic(err);
}
