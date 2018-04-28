package main

import "fmt"
import "os"
import "net/http"
//import "github.com/gorilla/mux"
import "flag"
import "strconv"
import "errors"

var isUsingFastCGI = flag.Bool("fcgi", false, "Should run in FastCGI mode?")
var nonFastCGIPort = flag.Int("no-fcgi", 80, "TCP port to use in non FastCGI mode")

type Route struct {
    Path string
    Handler http.HandlerFunc
}

var routes = []Route {
    Route{"/", defaultHandler},
}

func main() {
    flag.Parse()

    if !*isUsingFastCGI && len(os.Args) > 1 {
        port, err := strconv.Atoi(os.Args[1])
        if err == nil {
            *nonFastCGIPort = port
        }
    }

    if !*isUsingFastCGI && *nonFastCGIPort < 1 {
        checkError(errors.New("Invalid port number"))
    }

    if *isUsingFastCGI {
        startFastCGI()
    } else {
        start(*nonFastCGIPort)
    }
}

func start(port int) {
    fmt.Printf("Starting normal mode @ %d\n", port)

    /*router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes {
        router.Path(route.Path).Handler(route.Handler)
    }

    err := http.ListenAndServe(":8080", router)
    checkError(err)*/
}

func startFastCGI() {
    fmt.Println("Starting FastCGI mode")
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
