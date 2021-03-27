/*
 * novellia-api
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package novellia_api

// Token - Generalizes the data required to describe a token in a Cardano wallet.
type Token struct {

	// Token policy ID registered on Cardano
	PolicyId string `json:"policy_id"`

	// Amount of token held in Cardano wallet
	Amount float32 `json:"amount"`

	// Ticker as interpreted by Novellia (e.g. NVLA, ADA)
	Ticker string `json:"ticker,omitempty"`

	// Short description of token as interpreted by Novellia
	Description string `json:"description,omitempty"`
}
