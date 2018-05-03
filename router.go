package main

import "fmt"
import "net/http"
import "github.com/gorilla/mux"
import "encoding/json"

type Route struct {
    Path string
    Handler http.HandlerFunc
}

var routes = []Route {
    Route{"/", DefaultHandler},
    Route{"/api/happiness", HappinessHandler},
}

var router = mux.NewRouter().StrictSlash(true)

func init() {
    for _, route := range routes {
        router.Path(route.Path).Handler(route.Handler)
    }
}

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusNotFound)
    fmt.Fprintf(w, "Invalid Request")
}

func HappinessHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)

    json.NewEncoder(w).Encode(HappinessStatus{0})
}
