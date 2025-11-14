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

	server := &http.Server{Addr: addr, ReadHeaderTimeout: 3, Handler: mux}

	log.Printf("listening on %s\n", addr)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
