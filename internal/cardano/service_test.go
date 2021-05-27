package cardano_test

import (
	"context"
	"testing"

	"github.com/RektangularStudios/novellia/internal/config"
	"github.com/RektangularStudios/novellia/internal/cardano"
	nvla "github.com/RektangularStudios/novellia-sdk/sdk/server/go/novellia/v0"
)

const (
	configPath = "null"
)

func setupTest(ctx context.Context) (cardano.Service, error) {
	err := config.LoadConfig(configPath)
	if err != nil {
		return nil, err
	}

	cardanoService, err := cardano.New(ctx)
	if err != nil {
		return nil, err
	}

	return cardanoService, nil
}

func TestGet721Metadata(t *testing.T) {
	ctx := context.Background()

	cardanoService, err := setupTest(ctx)
	if err != nil {
		t.Errorf("failed to setup test: %+v", err)
	}
	defer cardanoService.Close(ctx)

	token := nvla.Token{
		NativeTokenId: "d27dadf7c5f24bfe9e377927c2d811d63d19222e1a53bb50cbb51772.Draculi",
	}
	modifiedTokens, err := cardanoService.Add721Metadata(ctx, []nvla.Token{token})
	if err != nil {
		t.Errorf("failed to get 721 metadata: %+v", err)
	}
	t.Errorf("got metadata in tokens: %+v", modifiedTokens)
}
