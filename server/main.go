package main

import (
	"fmt"
	"net/http"

	novellia_api "github.com/RektangularStudios/novellia/generated/novellia-api"
	"github.com/RektangularStudios/novellia/internal/api"
)

const (
	server_port = 3555
)

func main() {
	fmt.Printf("Starting server on port %d", server_port)

	DefaultApiService := api.NewApiService()
	DefaultApiController := novellia_api.NewDefaultApiController(DefaultApiService)

	router := novellia_api.NewRouter(DefaultApiController)

	err := http.ListenAndServe(fmt.Sprintf(":%d", server_port), router)
	if err != nil {
		fmt.Printf("Error starting server")
	}
}
