package main

import "net/http"
import "log"
import "github.com/fizzwu/bmwg/chap_04/handlers"
import "github.com/fizzwu/bmwg/chap_04/data"

func main() {
	store := &data.MemoryStore{}

	err := http.ListenAndServe(":2323", &handlers.SearchHandler{DataStore: store})
	if err != nil {
		log.Fatal(err)
	}
}
