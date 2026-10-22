package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dadams2/go-shorter/internal/app"
	"github.com/dadams2/go-shorter/internal/routes"
)

func main() {
	application, err := app.NewApplication()
	if err != nil {
		log.Fatal(err)
	}
	defer application.DB.Close()

	router := routes.SetupRouter(application)
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	fmt.Println("Server started")

	err = server.ListenAndServe()
	if err != nil {
		// Return normally so the deferred database close runs.
		log.Print(err)
	}
}
