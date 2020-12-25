package main

import (
	"net/http"

	"github.com/nireo/obviate"
)

func exampleHandler1(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Example handler 1"))
}

func exampleHandler2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{message:"hello from json"}`))
}

func main() {
	router := obviate.NewRouter()

	router.GET("/normal", exampleHandler1)
	router.GET("/json", exampleHandler2)

	router.Listen(":8080")
}
