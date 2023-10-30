package main

import (
	"fmt"
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

	hostAddress := fmt.Sprintf("%s:%s", os.Getenv("PROJECT_HOST"), os.Getenv("PROJECT_PORT"))
	err := http.ListenAndServe(hostAddress, r)
	if err != nil {
		log.Fatal(err)
	}
}
