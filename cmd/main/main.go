package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/pooulad/go-book-management/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBoostoreRoutes(r)
	r.Handle("/", r)

	err := http.ListenAndServe(os.Getenv("PROJECT_HOST"), r)
	if err != nil {
		log.Fatal(err)
	}
}
