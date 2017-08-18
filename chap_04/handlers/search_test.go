package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/fizzwu/bmwg/chap_04/data"
)

var mockStore data.MockStore

func TestSearchHandlerReturnsBadRequestWhenNoSearchCriterialIsSent(t *testing.T) {
	req, resp, h := setupTest(nil)

	h.ServeHTTP(resp, req)

	if resp.Code != http.StatusBadRequest {
		t.Errorf("Expected Bad Request, got %v", resp.Code)
	}
}

func TestSearchHandlerReturnsBadRequestWhenBlankSearchCriterialIsSent(t *testing.T) {
	req, resp, h := setupTest(&searchRequest{})

	h.ServeHTTP(resp, req)

	if resp.Code != http.StatusBadRequest {
		t.Errorf("Expected Bad Request, got %v", resp.Code)
	}
}

// func TestSearchHandlerCallsDataStoreWithValidQuery(t *testing.T) {
// 	req, resp, h := setupTest(&searchRequest{Query: "Fat Freddy's Cat"})
// 	mockStore.On("Search", "Fat Freddy's Cat").Return(make([]data.Kitten, 0))

// 	h.ServeHTTP(resp, req)
// 	mockStore.AssertExpectations(t)
// }

func TestSearchHandlerReturnKittensWithValidQuery(t *testing.T) {
	req, resp, h := setupTest(&searchRequest{Query: "Fat Freddy's Cat"})
	mockStore.On("Search", "Fat Freddy's Cat").Return(make([]data.Kitten, 1))

	h.ServeHTTP(resp, req)

	sr := searchResponse{}
	json.Unmarshal(resp.Body.Bytes(), &sr)

	assert.Equal(t, 1, len(sr.Kittens))
	assert.Equal(t, http.StatusOK, resp.Code)
}

func setupTest(d interface{}) (*http.Request, *httptest.ResponseRecorder, SearchHandler) {
	h := SearchHandler{DataStore: &data.MemoryStore{}}
	resp := httptest.NewRecorder()

	if d == nil {
		return httptest.NewRequest("POST", "/search", nil), resp, h
	}

	body, _ := json.Marshal(d)
	return httptest.NewRequest("POST", "/search", bytes.NewReader(body)), resp, h
}
