package rest

import (
    "encoding/json"
    "net/http"
    "strings"

    "github.com/jgrecu/hello-api/translation"
)

type Resp struct {
    Language    string `json:"language"`
    Translation string `json:"translation"`
}

const defaultLanguage = "english"

func TranslateHandler(w http.ResponseWriter, r *http.Request) {
    enc := json.NewEncoder(w)
    w.Header().Set("Content-Type", "application/json; charset=utf-8")

    language := r.URL.Query().Get("language")
    if language == "" {
        language = defaultLanguage
    }

    word := strings.ReplaceAll(r.URL.Path, "/", "")

    transl := translation.Translate(word, language)
    if transl == "" {
        w.WriteHeader(http.StatusNotFound)
        return
    }

    resp := Resp{
        Language:    language,
        Translation: transl,
    }
    if err := enc.Encode(resp); err != nil {
        panic("unable to encode respoonse")
    }
}
