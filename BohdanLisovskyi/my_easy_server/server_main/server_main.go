package main

import (
	"net/http"
	"github.com/LisovskyiB/softserve-academy/server"
)

func main() {

	http.HandleFunc("/name", server.Handler)

	http.ListenAndServe(":2005", nil)

	select {}
}
