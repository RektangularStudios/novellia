package main

import (
	"fmt"
	"net/http"

	nvla "github.com/RektangularStudios/novellia-sdk/sdk/server/go/v1"
	"github.com/RektangularStudios/novellia/internal/api"
)

const (
	server_port = 3555
)

func main() {
	fmt.Printf("Starting server on port %d", server_port)

	DefaultApiService := api.NewApiService()
	DefaultApiController := nvla.NewDefaultApiController(DefaultApiService)

	router := nvla.NewRouter(DefaultApiController)

	err := http.ListenAndServe(fmt.Sprintf(":%d", server_port), router)
	if err != nil {
		fmt.Printf("Error starting server")
	}
}
