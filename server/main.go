package main

import (
	"fmt"
	"context"
	"net/http"
	"os"
	
	"github.com/shurcooL/graphql"

	nvla "github.com/RektangularStudios/novellia-sdk/sdk/server/go/v0"
	"github.com/RektangularStudios/novellia/internal/api"
	cardano_graphql "github.com/RektangularStudios/novellia/internal/cardano/graphql"
)

func main() {
	ctx := context.Background()

	fmt.Printf("Novellia Server - Version %s\n", version)

	configPath, err := getConfigPath()
	if err != nil {
		fmt.Printf("Failed to get config path: %v\n", err)
		os.Exit(configPathErr)
	}

	config, err := loadConfig(configPath)
	if err != nil {
		fmt.Printf("Failed to load config: %v\n", err)
		os.Exit(configLoadErr)
	}

	fmt.Printf("Starting server with configuration (%s):\n %+v\n", configPath, config)

	cardanoGraphQLHostString := fmt.Sprintf("%s:%s", config.CardanoGraphQL.Host, config.CardanoGraphQL.Port)
	cardanoGraphQLClient := graphql.NewClient(cardanoGraphQLHostString, nil)

	cardanoGraphQLService := cardano_graphql.New(cardanoGraphQLClient)

	a,b,c := cardanoGraphQLService.Initialized(ctx)
	fmt.Printf("\n%v, %v, %v\n", a,b,c)

	DefaultApiService := api.NewApiService(cardanoGraphQLService)
	DefaultApiController := nvla.NewDefaultApiController(DefaultApiService)

	router := nvla.NewRouter(DefaultApiController)

	hostString := fmt.Sprintf("%s:%s", config.Server.Host, config.Server.Port)
	server := http.Server {
		Addr: hostString,
		Handler: router,
	}
	err = server.ListenAndServe()
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}

	os.Exit(successCode)
}
