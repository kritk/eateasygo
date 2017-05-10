package hello

import (
    "fmt"
    "net/http"
)

func init() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/create",createhandler)
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, world!")
}

func createhandler(w http.ResponseWriter, r *http.Request){
	f(r.Context())
	fmt.Fprint(w,"Save successful")
}
