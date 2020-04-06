package gorillamuxtesting

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/gorilla/mux"
)

// RunHTTPEndpoint test a mux router based endpoint
func RunHTTPEndpoint(method string, url string, body string, route string, token string, handlerFunc func(http.ResponseWriter, *http.Request), middleware ...mux.MiddlewareFunc) (string, error) {
	req, _ := http.NewRequest(method, url, strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	if token != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	for _, mwf := range middleware {
		router.Use(mwf)
	}
	router.HandleFunc(route, handlerFunc)
	router.ServeHTTP(rr, req)

	output := rr.Body.String()

	if status := rr.Code; status != http.StatusOK {
		return output, fmt.Errorf("Invalid HTTP Response: Actual: %v vs. Expected: %v Body: %s", status, http.StatusOK, output)
	}

	return output, nil
}
