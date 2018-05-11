package main

import "fmt"
import "os"
import "net/http"
import "net/http/fcgi"
import "flag"
import "strconv"
import "errors"

var nonFastCGIPort = flag.Int("port", 0, "TCP port to use in non FastCGI mode")


func main() {
    parseArguments()

    if *nonFastCGIPort > 0 {
        start(*nonFastCGIPort)
    } else {
        startFastCGI()
    }
}

func parseArguments() {
    flag.Parse()

    if len(os.Args) > 1 {
        port, err := strconv.Atoi(os.Args[1])
        if err == nil {
            *nonFastCGIPort = port
        }
    }

    if *nonFastCGIPort < 0 {
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
