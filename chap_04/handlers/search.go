package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/fizzwu/bmwg/chap_04/data"
)

type searchRequest struct {
	Query string `json:"query"`
}

type searchResponse struct {
	Kittens []data.Kitten `json:"kittens"`
}

type SearchHandler struct {
	DataStore data.Store
}

func (h *SearchHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	req := new(searchRequest)
	err := decoder.Decode(req)
	if err != nil || len(req.Query) < 1 {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	kittens := h.DataStore.Search(req.Query)
	encoder := json.NewEncoder(w)
	encoder.Encode(searchResponse{Kittens: kittens})

}
