package cardano_test

import (
	"context"
	"testing"

	"github.com/RektangularStudios/novellia/internal/config"
	"github.com/RektangularStudios/novellia/internal/cardano"
	nvla "github.com/RektangularStudios/novellia-sdk/sdk/server/go/novellia/v0"
)

const (
	configPath = "/media/ninja/SSD_2/rektangular/novellia/config/local.yaml"
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

func TestDecodeStakeAddress(t *testing.T) {
	ctx := context.Background()

	cardanoService, err := setupTest(ctx)
	if err != nil {
		t.Errorf("failed to setup test: %+v", err)
	}
	defer cardanoService.Close(ctx)

	// addr1q86l6gs80s0a5dnj9q4nrf4g5yzzaxs59srgxp4s5r0x33d2gm8j6rd5lx9y9tuvfv6qm6mypfhc7p9qkjawyj7g3h2spum389
	paymentAddressBase16 := "01f5fd22077c1fda3672282b31a6a8a1042e9a142c068306b0a0de68c5aa46cf2d0db4f98a42af8c4b340deb640a6f8f04a0b4bae24bc88dd5"

	s, err := cardanoService.DecodeStakeAddressFromBase16(paymentAddressBase16)
	if err != nil {
		t.Errorf("failed to get decode stake address: %+v", err)
	}
	t.Errorf("got stake address: %s", s)
}
