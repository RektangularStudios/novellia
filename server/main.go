package main

import (
	"fmt"
	"net/http"
	"os"
	"context"
	
	"github.com/shurcooL/graphql"

	nvla "github.com/RektangularStudios/novellia-sdk/sdk/server/go/v0"
	"github.com/RektangularStudios/novellia/internal/api"
	"github.com/RektangularStudios/novellia/internal/novellia_database"
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

	var apiService nvla.DefaultApiServicer
	if config.Mocked {
		apiService = api.NewMockedApiService()
	} else {
		cardanoGraphQLHostString := fmt.Sprintf("%s:%s", config.CardanoGraphQL.Host, config.CardanoGraphQL.Port)
		cardanoGraphQLClient := graphql.NewClient(cardanoGraphQLHostString, nil)
	
		cardanoGraphQLService := cardano_graphql.New(cardanoGraphQLClient)
	
		novelliaDatabaseService, err := novellia_database.New(
			ctx,
			config.Postgres.Username,
			config.Postgres.Password,
			config.Postgres.Host,
			config.Postgres.Database,
		)
		if err != nil {
			fmt.Printf("Failed to make Novellia database service: %+v\n", err)
			os.Exit(novelliaDatabaseErr)
		}
		defer novelliaDatabaseService.Close(ctx)

		err = novelliaDatabaseService.ListProductAttribution(ctx)
		if err != nil {
			fmt.Printf("Failed to get rows: %+v\n", err)
		}

		apiService = api.NewApiService(cardanoGraphQLService)	
	}

	apiController := nvla.NewDefaultApiController(apiService)

	router := nvla.NewRouter(apiController)

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
