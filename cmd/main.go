package main

import (
	"log"
	"net/http"
	"time"

	"github.com/jgrecu/hello-api/config"
	"github.com/jgrecu/hello-api/handlers"
	"github.com/jgrecu/hello-api/handlers/rest"
	"github.com/jgrecu/hello-api/translation"
)

func main() {

	cfg := config.LoadConfiguration()
	addr := cfg.Port

	mux := http.NewServeMux()

	var translationService rest.Translator
	translationService = translation.NewStaticService()
	if cfg.LegacyEndpoint != "" {
		log.Printf("creating external translation client: %s", cfg.LegacyEndpoint)
		client := translation.NewHelloClient(cfg.LegacyEndpoint)
		translationService = translation.NewRemoteService(client)
	}

	translateHandler := rest.NewTranslateHandler(translationService)

	mux.HandleFunc("GET /translate/hello", translateHandler.TranslateHandler)
	mux.HandleFunc("GET /health", handlers.HealthCheck)
	mux.HandleFunc("GET /info", handlers.Info)

	server := &http.Server{Addr: addr, ReadHeaderTimeout: 3 * time.Second, Handler: mux}

	log.Printf("listening on %s\n", addr)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
