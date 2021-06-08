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

	paymentAddress := "0132488c6dc87ed3e2c51436fd46cdfba158f4d99ddc6e708ccc87e06ce6bfe1302efb8d5f5b9891e41f14fa1573a2038176a3c6bd9d4c14fb"
	paymentAddressBech32 := "addr1qyey3rrdepld8ck9zsm063kdlws43axenhwxuuyvejr7qm8xhlsnqthm3404hxy3us03f7s4ww3q8qtk50rtm82vznasv64fhk"
	stakeAddress := "e6bfe1302efb8d5f5b9891e41f14fa1573a2038176a3c6bd9d4c14fb"
	stakeAddressBech32 := "stake1u8ntlcfs9mac6h6mnzg7g8c5lg2h8gsrs9m2834an4xpf7cp8zaff"

	s, err := cardanoService.DecodeStakeAddress(paymentAddressBech32)
	if err != nil {
		t.Errorf("failed to get decode stake address: %+v", err)
	}
	t.Errorf("got stake address: %s", s)
}
