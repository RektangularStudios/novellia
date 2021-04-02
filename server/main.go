package main

import (
	"fmt"
	"net/http"

	"github.com/shurcooL/graphql"

	nvla "github.com/RektangularStudios/novellia-sdk/sdk/server/go/v0"
	"github.com/RektangularStudios/novellia/internal/api"
)

const (
	server_port = 3555
)

func main() {
	fmt.Printf("Starting server on port %d", server_port)

	graphqlClient := graphql.NewClient("https://relay1.rektangularstudios.com/graphql", nil)

	DefaultApiService := api.NewApiService()
	DefaultApiController := nvla.NewDefaultApiController(DefaultApiService)

	router := nvla.NewRouter(DefaultApiController)

	err := http.ListenAndServe(fmt.Sprintf(":%d", server_port), router)
	if err != nil {
		fmt.Printf("Error starting server")
	}
}
