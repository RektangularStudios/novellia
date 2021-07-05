package cardano

import (
	"fmt"
	"context"
	"strings"
	"database/sql"

	nvla "github.com/RektangularStudios/novellia-sdk/sdk/server/go/novellia/v0"
)

// TODO: use this function to support additional search type
// returns asset fingerprint identifiers such as asset1vs9ve7ztxx5lm6taatrhh9v4p0ntlwj7j2nj9m
/*
func (s *ServiceImpl) getFingerprintTokenIdentifiers(tokenSearch nvla.TokenSearch) []string {
	fingerprints := []string{}
	for _, id := range tokenSearch.CardanoIdentifiers {
		if strings.HasPrefix(id, "asset") {
			fingerprints = append(fingerprints, id)
		}
	}

	return fingerprints
}
*/

func (s *ServiceImpl) getSpecialTokenIdentifier(tokenSearch nvla.TokenSearch) string {
	for _, id := range tokenSearch.CardanoIdentifiers {
		switch id {
		case "random": 
		default:
			continue
		}

		// only get first special identifier
		return id
	}

	return ""
}

// returns identifiers of form policyID.assetID
func (s *ServiceImpl) getNativeTokenIDTokenIdentifiers(tokenSearch nvla.TokenSearch) []string {
	nativeTokenIDs := []string{}
	for _, id := range tokenSearch.CardanoIdentifiers {
		split := strings.SplitN(id, ".", 2)
		if len(split) == 2 {
			nativeTokenIDs = append(nativeTokenIDs, id)
		}
	}

	return nativeTokenIDs
}

// return policy and name identifiers such as d27dadf7c5f24bfe9e377927c2d811d63d19222e1a53bb50cbb51772 or Draculi
func (s *ServiceImpl) getPolicyOrNameTokenIdentifiers(tokenSearch nvla.TokenSearch) []string {
	policiesOrNames := []string{}
	for _, id := range tokenSearch.CardanoIdentifiers {
		// filter our fully declared identifiers
		if len(strings.SplitN(id, ".", 2)) == 2 {
			continue
		}
		// filter our fingerprint identifiers
		if strings.HasPrefix(id, "asset") {
			continue
		}

		policiesOrNames = append(policiesOrNames, id)
	}

	return policiesOrNames
}

func (s *ServiceImpl) QueryTokensRandom(ctx context.Context) ([]nvla.Token, error) {
	rows, err := s.pool.Query(ctx, s.queries[queryTokensRandom])
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tokens := []nvla.Token{}
	for rows.Next() {
		var policyID, assetID []byte
		err = rows.Scan(
			&policyID,
			&assetID,
		)
		if err != nil {
			return nil, fmt.Errorf("query random tokens failed: %v", err)
		}

		nativeTokenID := fmt.Sprintf("%x.%s", policyID, string(assetID))
		tokens = append(tokens, nvla.Token{
			NativeTokenId: nativeTokenID,
		})
	}
	return tokens, nil
}

func (s *ServiceImpl) QueryTokensBySpecialIdentifier(ctx context.Context, tokenSearch nvla.TokenSearch) ([]nvla.Token, error) {
	specialIdentifier := s.getSpecialTokenIdentifier(tokenSearch)
	tokens := []nvla.Token{}

	// only ever consume a single special identifier
	switch specialIdentifier {
	case "random":
		return s.QueryTokensRandom(ctx)
	}

	return tokens, nil
}

func (s *ServiceImpl) QueryTokensByNativeTokenID(ctx context.Context, tokenSearch nvla.TokenSearch) ([]nvla.Token, error) {
	nativeTokenIDs := s.getNativeTokenIDTokenIdentifiers(tokenSearch)
	tokens := []nvla.Token{}

	for _, id := range nativeTokenIDs {
		idParts := strings.SplitN(id, ".", 2)
		if len(idParts) != 2 {
			return nil, fmt.Errorf("query tokens by native token ID failed, invalid native token ID")
		}

		var policyID, assetID []byte
		err := s.pool.QueryRow(ctx, s.queries[queryTokenByNativeTokenID], idParts[0], idParts[1]).Scan(
			&policyID,
			&assetID,
		)
		if err == sql.ErrNoRows {
			continue
		}
		if err != nil {
			return nil, fmt.Errorf("query tokens by native token ID failed: %v", err)
		}
		
		nativeTokenID := fmt.Sprintf("%x.%s", policyID, string(assetID))
		tokens = append(tokens, nvla.Token{
			NativeTokenId: nativeTokenID,
		})
	}

	return tokens, nil
}

func (s *ServiceImpl) QueryTokensByPolicyOrName(ctx context.Context, tokenSearch nvla.TokenSearch) ([]nvla.Token, error) {
	policiesOrNames := s.getPolicyOrNameTokenIdentifiers(tokenSearch)

	rows, err := s.pool.Query(ctx, s.queries[queryTokensByPolicyOrName], policiesOrNames)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tokens := []nvla.Token{}
	for rows.Next() {
		var policyID, assetID []byte
		err = rows.Scan(
			&policyID,
			&assetID,
		)
		if err != nil {
			return nil, fmt.Errorf("query tokens failed: %v", err)
		}

		nativeTokenID := fmt.Sprintf("%x.%s", policyID, string(assetID))
		tokens = append(tokens, nvla.Token{
			NativeTokenId: nativeTokenID,
		})
	}

	return tokens, nil
}

// Query tokens on Cardano from search identifiers
func (s *ServiceImpl) QueryTokens(ctx context.Context, search nvla.TokenSearch) ([]nvla.Token, error) {
	tokens := []nvla.Token{}
	
	t, err := s.QueryTokensBySpecialIdentifier(ctx, search)
	if err != nil {
		return nil, err
	}

	if len(t) != 0 {
		// do not do additional querying after consuming special identifier
		return t, nil
	}
	if len(s.getSpecialTokenIdentifier(search)) != 0 {
		return tokens, nil
	}


	t, err = s.QueryTokensByPolicyOrName(ctx, search)
	if err != nil {
		return nil, err
	}
	tokens = append(tokens, t...)
	
	t, err = s.QueryTokensByNativeTokenID(ctx, search)
	if err != nil {
		return nil, err
	}
	tokens = append(tokens, t...)

	return tokens, nil
}
