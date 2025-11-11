package main

import (
    "fmt"
    "log"
    "net/http"
    "os"

    "github.com/jgrecu/hello-api/handlers"
    "github.com/jgrecu/hello-api/handlers/rest"
)

func main() {

    addr := fmt.Sprintf(":%s", os.Getenv("PORT"))
    if addr == ":" {
        addr = ":8080"
    }

    mux := http.NewServeMux()

    mux.HandleFunc("GET /hello", rest.TranslateHandler)
    mux.HandleFunc("GET /health", handlers.HealthCheck)

    log.Printf("listening on %s\n", addr)

    log.Fatal(http.ListenAndServe(addr, mux))
}
