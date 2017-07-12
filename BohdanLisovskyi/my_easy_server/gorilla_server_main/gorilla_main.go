package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/LisovskyiB/softserve-academy/server"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/name", server.Handler).Methods("GET")

	log.Fatal(http.ListenAndServe(":2006", r))
}
