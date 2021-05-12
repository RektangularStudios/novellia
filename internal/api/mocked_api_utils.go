package api

import (
	nvla "github.com/RektangularStudios/novellia-sdk/sdk/server/go/novellia/v0"
)

const (
	MockedNovelliaStandardTokenProductId = "PROD-01F45Q1NE112Y61FD1NSE9NZXN"
	MockedNovelliaProductProductId = "PROD-01F4A66RHHQ5NECHQSY7AFVV3G"
)

func (s *MockedApiService) GetMockNovelliaStandardTokenProduct() nvla.Product {
	// Iscara assets
	card_low_urls := []string{
		"sia://fAD4zjfcyDIU0vAsn6dIdmnykYPP6ftYp33pHH1LJK1aMw",
		"ipfs://QmXaLtSG1PPJAUm9PYv6KhRiG5tqfZGvRGvaFfdkaRfhp2",
		"https://api.rektangularstudios.com/static/xysn45hd2nlz/resource/card_low.jpg",
	}
	card_urls := []string{
		"sia://AAC5JIhKnJtKQSE6wLBqRNUkGQ010HG3JS0GkBSugt7YNw",
		"ipfs://QmbHWwXbBXzei684goupCagpTtRWjbMtqyFeHnX4sUXn5k",
		"https://api.rektangularstudios.com/static/xysn45hd2nlz/resource/card.png",
	}
	artwork_low_urls := []string{
		"sia://vAB4kSyh4FsF2u_XV0FXWtIVFvbhRxoHFSSUet9-ZhAUnA",
		"ipfs://QmQ6VqTjn8d4hsssx7fqE32xRFRbSa7StyGcKis6LZBJ2g",
		"https://api.rektangularstudios.com/static/xysn45hd2nlz/resource/artwork_low.jpg",
	}
	artwork_urls := []string{
		"sia://AACEvMtrih6fvBWRyJtVkTkO4dXvl1pNquob3fF5qDxyhQ",
		"ipfs://QmUxntRBXo5mWK5LzWFNpR7de1rwagJcuX92xbqRZdQfNW",
		"https://api.rektangularstudios.com/static/xysn45hd2nlz/resource/artwork.png",
	}
	character_urls := []string{
		"sia://AABMK9GQCfCFSUxjluxsQ0ameY1LQaHhbjH8NYYBlXsClQ",
		"ipfs://QmTf7ycZi6tDp7m9SFUMm9VFfgNFoVa5a1QWWvzoYceVzf",
		"https://api.rektangularstudios.com/static/xysn45hd2nlz/resource/character.json",
	}

	descriptionSet := nvla.DescriptionSet{
		Short: "Occulta Novellia Character",
		Long: "# Take back gaming!\nA character token for the surreal horror game Occulta Novellia. This token represents ownership of a playable character.",
	}

	// to be combined like "PolicyId.AssetId"
	nativeToken := nvla.NativeToken{
		PolicyId: "0xMyMultisigScriptHash",
		AssetId: "IscaraTheTenThousandGuns",
	}

	novelliaStandardToken := nvla.NovelliaStandardToken{
		NativeToken: nativeToken,
		Copyright: "Copyright Rektangular Studios Inc.; all rights reserved",
		Publisher: []string{"https://rektangularstudios.com"},
		NovelliaVersion: 1,
		Version: 1,
		Extension: []string{"novellia_1"},
		Id: 2,
		Name: "Iscara the Ten Thousand Guns",
		// TODO: verify this doesn't need to follow ipfs://ipfs/... stutter (will need to mint a token and see which clients break)
		Image: "ipfs://QmXaLtSG1PPJAUm9PYv6KhRiG5tqfZGvRGvaFfdkaRfhp2",
		Description: descriptionSet,
		Tags: []string{
			"Game Character",
			"Kinda Rare",
		},
		Commission: []nvla.Commission{
			nvla.Commission{
				Name: "Rektangular Studios Inc.",
				Address: "addr1q8chzck9gkzd8t2v3477klsw3hna0s0er5vwspxehelaryfw08lffp5n2kzt72ez93m5zev2v4fm9sawnrqnvllmyhmsdjfzg9",
				Percent: 0.03,
			},
		},
		Resource: []nvla.OffChainResource{
			nvla.OffChainResource{
				ResourceId: "Artwork",
				Description: "Low Resolution Illustration",
				Priority: 0,
				Multihash: "QmXaLtSG1PPJAUm9PYv6KhRiG5tqfZGvRGvaFfdkaRfhp2",
				HashSourceType: "ipfs",
				Url: artwork_low_urls,
				ContentType: "image/jpeg",
			},
			nvla.OffChainResource{
				ResourceId: "Artwork",
				Description: "High Resolution Illustration",
				Priority: 1,
				Multihash: "QmbHWwXbBXzei684goupCagpTtRWjbMtqyFeHnX4sUXn5k",
				HashSourceType: "ipfs",
				Url: artwork_urls,
				ContentType: "image/png",
			},
			nvla.OffChainResource{
				ResourceId: "Card",
				Description: "Low Resolution Character Card",
				Priority: 0,
				Multihash: "QmQ6VqTjn8d4hsssx7fqE32xRFRbSa7StyGcKis6LZBJ2g",
				HashSourceType: "ipfs",
				Url: card_low_urls,
				ContentType: "image/jpeg",
			},
			nvla.OffChainResource{
				ResourceId: "Card",
				Description: "High Resolution Character Card",
				Priority: 1,
				Multihash: "QmUxntRBXo5mWK5LzWFNpR7de1rwagJcuX92xbqRZdQfNW",
				HashSourceType: "ipfs",
				Url: card_urls,
				ContentType: "image/png",
			},
			nvla.OffChainResource{
				ResourceId: "OccultaNovelliaCharacter",
				Description: "Occulta Novellia character play information such as stats and moves",
				Priority: 0,
				Multihash: "QmTf7ycZi6tDp7m9SFUMm9VFfgNFoVa5a1QWWvzoYceVzf",
				HashSourceType: "ipfs",
				Url: character_urls,
				ContentType: "application/json",
			},
		},
	}

	product := nvla.Product{
		Pricing: nvla.ProductPricing{
			PriceCurrencyId: "ada",
			PriceUnitAmount: 20,
			MaxOrderSize: 5,
		},
		Organization: nvla.ProductOrganization{
			Name: "Rektangular Studios Inc.",
			OrganizationId: "ORG-01F45PHP58QWYSWJPFC0RYYGJ2",
			Description: "[Rektangular Studios](https://rektangularstudios.com) is a software development company bringing new products to the Cardano ecosystem.",
		},
		Market: nvla.ProductMarket{
			Name: "Occulta Novellia",
			MarketId: "MKT-01F45PJXRNEM8V7CP48NR39639",
			Description: "[Occulta Novellia](https://rektangularstudios.com/occulta-novellia/) is a surreal horror multiverse.\nEach new set of collectibles is minted on the Cardano blockchain to incorporate new stories, building up the many facets of it.",
		},
		Stock: nvla.ProductStock{
			Available: 2300,
			TotalSupply: 2500,
		},
		Metadata: nvla.ProductMetadata{
			// PST time (-08:00) at 2:00 PM
			DateListed: "2021-05-03T22:00:00-0000",
			DateAvailable: "2021-05-17T22:00:00-0000",
		},
		Product: nvla.ProductProduct{
			NovelliaStandardToken: &novelliaStandardToken,
			ProductId: MockedNovelliaStandardTokenProductId,
		},
		Attribution: []nvla.Attribution{
			nvla.Attribution{
				AuthorName: "John Doe",
				Url: []string{
					"https://www.artstation.com/",
					"https://www.deviantart.com/",
				},
				WorkAttributed: "Iscara the Ten Thousand Guns Illustration",
			},
		},
	}

	return product
}

func (s* MockedApiService) GetMockNovelliaProduct() nvla.Product {
	// Booster assets
	artwork_low_urls := []string{
		"sia://PAPVFe5nrds0v74t-5dnvEG1MYh5k2hxvtJ9BfqF2kvDNg",
		"ipfs://QmX3pkMjRMB32nyHdrfMiBhjPc5H7g595Hr242keBF2co1",
		"https://api.rektangularstudios.com/static/iwqw13wtrnqj/resource/artwork_low.jpg",
	}
	artwork_urls := []string{
		"sia://_A0DrUket_aS275WqhkWLUFdxBk5VCFgDfBIyM9ijqf-Xg",
		"ipfs://QmWFfcMeob1u15LWqNpuf8o7MzyW35bZjVxxUz3QJmVejs",
		"https://api.rektangularstudios.com/static/iwqw13wtrnqj/resource/artwork.png",
	}

	descriptionSet := nvla.DescriptionSet{
		Short: "Occulta Novellia Booster",
		Long: "# Take back gaming!\nOcculta Novellia booster containing 3 random character tokens according to the following distribution:\n- 10% chance of Rare\n- 20% chance of Kinda Rare\n- 70% chance of Not That Rare",
	}

	novelliaProduct := nvla.NovelliaProduct{
		Copyright: "Copyright Rektangular Studios Inc.; all rights reserved",
		Publisher: []string{"https://rektangularstudios.com"},
		Name: "Booster Pack",
		Description: descriptionSet,
		Tags: []string{
			"Game Item",
			"Game Character",
			"Bundle",
		},
		Commission: []nvla.Commission{
			nvla.Commission{
				Name: "Rektangular Studios Inc.",
				Address: "addr1q8chzck9gkzd8t2v3477klsw3hna0s0er5vwspxehelaryfw08lffp5n2kzt72ez93m5zev2v4fm9sawnrqnvllmyhmsdjfzg9",
				Percent: 0.03,
			},
		},
		Resource: []nvla.OffChainResource{
			nvla.OffChainResource{
				ResourceId: "Artwork",
				Description: "Low Resolution Product Image",
				Priority: 0,
				Multihash: "QmX3pkMjRMB32nyHdrfMiBhjPc5H7g595Hr242keBF2co1",
				HashSourceType: "ipfs",
				Url: artwork_low_urls,
				ContentType: "image/jpeg",
			},
			nvla.OffChainResource{
				ResourceId: "Artwork",
				Description: "High Resolution Product Image",
				Priority: 1,
				Multihash: "QmWFfcMeob1u15LWqNpuf8o7MzyW35bZjVxxUz3QJmVejs",
				HashSourceType: "ipfs",
				Url: artwork_urls,
				ContentType: "image/png",
			},
		},
	}

	product := nvla.Product{
		Pricing: nvla.ProductPricing{
			PriceCurrencyId: "ada",
			PriceUnitAmount: 30,
			MaxOrderSize: 5,
		},
		Organization: nvla.ProductOrganization{
			Name: "Rektangular Studios Inc.",
			OrganizationId: "ORG-01F45PHP58QWYSWJPFC0RYYGJ2",
		},
		Market: nvla.ProductMarket{
			Name: "Occulta Novellia",
			MarketId: "MKT-01F45PJXRNEM8V7CP48NR39639",
		},
		Stock: nvla.ProductStock{
			Available: 500,
		},
		Metadata: nvla.ProductMetadata{
			// PST time (-08:00) at 2:00 PM
			DateListed: "2021-05-03T14:00:00-08:00",
			DateAvailable: "2021-05-17T14:00:00-08:00",
		},
		Product: nvla.ProductProduct{
			NovelliaProduct: &novelliaProduct,
			ProductId: MockedNovelliaProductProductId,
		},
	}

	return product
}
