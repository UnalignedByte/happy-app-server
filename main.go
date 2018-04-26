package main

import "fmt"
import "os"
import "net/http"
import "github.com/gorilla/mux"

type Route struct {
    Path string
    Handler http.HandlerFunc
}

var routes = []Route {
    Route{"/", defaultHandler},
}

func main() {
    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes {
        router.Path(route.Path).Handler(route.Handler)
    }

    err := http.ListenAndServe(":8080", router)
    checkError(err)
}

func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "%s\n", err.Error())
        os.Exit(1)
    }
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Invalid Request")
}
