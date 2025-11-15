package faas

import (
	"net/http"

	"github.com/jgrecu/hello-api/handlers/rest"
	"github.com/jgrecu/hello-api/translation"
)

func Translate(w http.ResponseWriter, r *http.Request) {
	translationService := translation.NewStaticService()
	translateHandler := rest.NewTranslateHandler(translationService)
	translateHandler.TranslateHandler(w, r)
}
