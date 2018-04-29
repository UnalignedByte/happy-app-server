package main

import "fmt"
import "os"
import "net/http"
import "net/http/fcgi"
import "flag"
import "strconv"
import "errors"

var isUsingFastCGI = flag.Bool("fcgi", false, "Should run in FastCGI mode?")
var nonFastCGIPort = flag.Int("no-fcgi", 80, "TCP port to use in non FastCGI mode")


func main() {
    parseArguments()

    if *isUsingFastCGI {
        startFastCGI()
    } else {
        start(*nonFastCGIPort)
    }
}

func parseArguments() {
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
}

func startFastCGI() {
    fmt.Println("Starting in FastCGI")
    err := fcgi.Serve(nil, router)
    checkError(err)
}

func start(port int) {
    fmt.Printf("Starting in normal mode using port %d\n", port)

    portString := fmt.Sprintf(":%d", port)
    err := http.ListenAndServe(portString, router)
    checkError(err)
}

func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "%s\n", err.Error())
        os.Exit(1)
    }
}
