package handlers

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/fizzwu/bmwg/chap_04/data"
)

func BenchmarkSearchHandler(b *testing.B) {
	var mockStore data.MockStore
	mockStore.On("Search", "Fat Freddy's Cat").Return([]data.Kitten{
		data.Kitten{
			Name: "Fat Freddy's Cat",
		},
	})

	sh := SearchHandler{DataStore: &data.MemoryStore{}}

	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest("POST", "/search", bytes.NewReader([]byte(`{"query":"Fat Freddy's Cat"}`)))
		resp := httptest.NewRecorder()
		sh.ServeHTTP(resp, req)
	}
}
