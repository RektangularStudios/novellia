package main

import (
	"fmt"
	"net/http"
	"os"

	yaml "gopkg.in/yaml.v3"
	"github.com/shurcooL/graphql"

	nvla "github.com/RektangularStudios/novellia-sdk/sdk/server/go/v0"
	"github.com/RektangularStudios/novellia/internal/api"
)

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"server"`
	CardanoGraphQL struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"cardano-graphql"`
}

func getConfigPath() (string, error) {
	if (len(os.Args) != 2) {
		return "", fmt.Errorf("expected config path as first argument")
	}
	
	configPath := os.Args[1]
	return configPath, nil
}

func loadConfig(configPath string) (*Config, error) {
	var config Config

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)

	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func main() {
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

	DefaultApiService := api.NewApiService(cardanoGraphQLClient)
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
