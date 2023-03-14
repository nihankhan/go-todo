package cmd

import (
	"log"
	"net/http"

	"github.com/nihankhan/go-todo/internal"
)

func main() {
	r := internal.Routers()

	log.Println("Server Running on 127.0.0.1:8080")

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: r,
	}

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
