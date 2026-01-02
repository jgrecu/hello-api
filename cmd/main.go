package main

import (
	"log"
	"net/http"

	"github.com/jgrecu/hello-api/config"
	"github.com/jgrecu/hello-api/handlers"
	"github.com/jgrecu/hello-api/handlers/rest"
	"github.com/jgrecu/hello-api/translation"
)

func main() {

	cfg := config.LoadConfiguration()
	addr := cfg.Port

	mux := http.NewServeMux()

	translationService := translation.NewStaticService()
	translateHandler := rest.NewTranslateHandler(translationService)

	mux.HandleFunc("GET /translate/hello", translateHandler.TranslateHandler)
	mux.HandleFunc("GET /health", handlers.HealthCheck)

	server := &http.Server{Addr: addr, ReadHeaderTimeout: 3, Handler: mux}

	log.Printf("listening on %s\n", addr)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
