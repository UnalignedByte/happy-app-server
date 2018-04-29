package main

import "fmt"
import "net/http"
import "github.com/gorilla/mux"

type Route struct {
    Path string
    Handler http.HandlerFunc
}

var routes = []Route {
    Route{"/", defaultHandler},
}

var router = mux.NewRouter().StrictSlash(true)

func init() {
    for _, route := range routes {
        router.Path(route.Path).Handler(route.Handler)
    }
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Invalid Request")
}
