package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type validationHandler struct {
	next http.Handler
}

type helloWorldHandler struct {
}

type helloWorldRequest struct {
	Name string `json:"name"`
}

type helloWorldResponse struct {
	Message string `json:"message"`
}

func newValidationHandler(next http.Handler) http.Handler {
	return validationHandler{next: next}
}

func newHelloWorldHandler() http.Handler {
	return helloWorldHandler{}
}

func (h validationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var request helloWorldRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	h.next.ServeHTTP(w, r)
}

func (h helloWorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := helloWorldResponse{Message: "Hello"}
	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}

func main() {
	port := 8080

	handler := newValidationHandler(newHelloWorldHandler())
	http.Handle("/helloworld", handler)
	log.Printf("Server listening on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
