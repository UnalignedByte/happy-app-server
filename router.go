package main

import "fmt"
import "net/http"
import "github.com/gorilla/mux"

type Route struct {
    Path string
    Method string
    Handler http.HandlerFunc
}

var routes = []Route {
    Route{"/", "GET", DefaultHandler},
    Route{"/api/happiness", "GET", HappinessGetHandler},
    Route{"/api/happiness", "POST", HappinessPostHandler},
}

var router = mux.NewRouter().StrictSlash(true)

func init() {
    for _, route := range routes {
        router.
               Path(route.Path).
               Methods(route.Method).
               Handler(route.Handler)
    }
}

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusNotFound)
    fmt.Fprintf(w, "Invalid Request")
}
