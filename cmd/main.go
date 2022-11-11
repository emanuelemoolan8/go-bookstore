package main

import (
	"github.com/emanuelemoolan8/go-store/pkg/routes"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.BookstoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9000", nil))
}
