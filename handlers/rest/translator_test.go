package rest_test

import (
	"encoding/json"
	"hello-api/handlers/rest"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTranslateAPI(t *testing.T) {
	tt := []struct {
		Endpoint            string
		StatusCode          int
		ExpectedLanguage    string
		ExpectedTranslation string
	}{
		{
			Endpoint:            "/hello",
			StatusCode:          http.StatusOK,
			ExpectedLanguage:    "english",
			ExpectedTranslation: "hello",
		},
		{
			Endpoint:            "/hello?language=german",
			StatusCode:          http.StatusOK,
			ExpectedLanguage:    "german",
			ExpectedTranslation: "hallo",
		},
		{
			Endpoint: "/hello?language=dutch",
			StatusCode: http.StatusNotFound,
			ExpectedLanguage: "",
			ExpectedTranslation: "",
		},
	}

	handler := http.HandlerFunc(rest.TranslateHandler)

	for _, test := range tt {
		// Arrange
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", test.Endpoint, nil)

		// Act
		handler.ServeHTTP(rr, req)

		// Assert
		if rr.Code != test.StatusCode {
			t.Errorf(`expected status %d but received %d`, test.StatusCode, rr.Code)
		}

		var resp rest.Resp
		json.Unmarshal(rr.Body.Bytes(), &resp)

		if resp.Language != test.ExpectedLanguage {
			t.Errorf(`expected language %q but received %s`, test.ExpectedLanguage, resp.Language)
		}

		if resp.Translation != test.ExpectedTranslation {
			t.Errorf(`expected Translation %q but received %s`, test.ExpectedTranslation, resp.Translation)
		}

	}
}
