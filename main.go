package hello

import (
    "fmt"
    "net/http"
    "google.golang.org/appengine"

)

func init() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/create",createhandler)
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, world!")
}

func createhandler(w http.ResponseWriter, r *http.Request){
	result := f(appengine.NewContext(r))
	fmt.Fprint(w, result)
}
