package main

import (
	"log"
	"net/http"

	"github.com/Gravitalia/AirBack/router"
)

func main() {
	http.HandleFunc("/create/post", router.Post)

	if err := http.ListenAndServe(":8040", nil); err != nil {
		log.Fatal(err)
	}
}
