package faas

import (
	"net/http"

	"github.com/jgrecu/hello-api/handlers/rest"
)

func Translate(w http.ResponseWriter, r *http.Request) {
	rest.TranslateHandler(w, r)
}
