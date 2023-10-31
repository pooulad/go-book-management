package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/pooulad/go-book-management/pkg/routes"
	"github.com/joho/godotenv"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBoostoreRoutes(r)
	r.Handle("/", r)

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	host := os.Getenv("PROJECT_HOST")
	port := os.Getenv("PROJECT_PORT")

	hostAddress := fmt.Sprintf("%s:%s", host, port)
	fmt.Println(hostAddress)
	err = http.ListenAndServe(hostAddress, r)
	if err != nil {
		log.Fatal(err)
	}
}
