package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dadams2/go-shorter/internal/routes"
)



func main() {

	router := routes.SetupRouter()
	server := &http.Server{
		Addr: ":8080",
		Handler: router,
	}
	fmt.Println("Server started")

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal();
	}
	
}
