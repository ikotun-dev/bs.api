package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ikotun-dev/bookstore/pkg/routes"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))

}
