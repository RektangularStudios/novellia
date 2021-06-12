package main

import (
	"fmt"
	"net/http"
	"os"
	"context"

	nvla "github.com/RektangularStudios/novellia-sdk/sdk/server/go/novellia/v0"
	"github.com/RektangularStudios/novellia/internal/config"
	prometheus_monitoring "github.com/RektangularStudios/novellia/internal/monitoring"
	"github.com/RektangularStudios/novellia/internal/api"
	"github.com/RektangularStudios/novellia/internal/novellia_database"
	"github.com/RektangularStudios/novellia/internal/cardano"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	ctx := context.Background()

	fmt.Printf("Novellia Server - Version %s\n", version)

	configPath, err := config.GetConfigPath()
	if err != nil {
		fmt.Printf("Failed to get config path: %v\n", err)
		os.Exit(configPathErr)
	}

	err = config.LoadConfig(configPath)
	if err != nil {
		fmt.Printf("Failed to load config: %v\n", err)
		os.Exit(configLoadErr)
	}
	config, err := config.GetConfig()
	if err != nil {
		fmt.Printf("Failed to get config from env: %v\n", err)
		os.Exit(configGetErr)
	}

	fmt.Printf("Starting server with configuration (%s):\n %+v\n", configPath, config)

	var apiService nvla.DefaultApiServicer
	if config.Mocked {
		apiService = api.NewMockedApiService()
	} else {
		cardanoService, err := cardano.New(ctx)
		if err != nil {
			fmt.Printf("failed to make Cardano service: %v\n", err)
			os.Exit(cardanoServiceErr)
		}
		defer cardanoService.Close(ctx)

		novelliaDatabaseService, err := novellia_database.New(ctx)
		if err != nil {
			fmt.Printf("Failed to make Novellia database service: %+v\n", err)
			os.Exit(novelliaDatabaseErr)
		}
		defer novelliaDatabaseService.Close(ctx)

		apiService = api.NewApiService(
			cardanoService,
			novelliaDatabaseService,
		)	
	}

	apiController := nvla.NewDefaultApiController(apiService)

	router := nvla.NewRouter(apiController)

	// add Prometheus metrics to router
	prometheus_monitoring.Init(config.Monitoring.Namespace)
	prometheus_monitoring.RecordMetrics()
	router.Handle("/metrics", promhttp.Handler())

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
