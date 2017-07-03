package main

import (
	"net/http"
	"os"

	"log"

	"github.com/dharnnie/tai/site/handlers"
	"github.com/gorilla/mux"
)

func main() {
	serve()
}

func serve() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/assets/", handlers.ServeResource)

	myMux := mux.NewRouter()
	myMux.HandleFunc("/", handlers.Home)

	http.Handle("/", myMux)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("Server error", err)
	}
}
