/*
 * novellia
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.9.2
 * Contact: contact@rektangularstudios.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package novellia

// CardanoTransaction - A Cardano transaction to be signed and submitted.
type CardanoTransaction struct {

	// text of transaction file
	Transaction string `json:"transaction"`

	// Cost to submit TX in lovelace (1 ADA = 1,000,000 lovelace)
	FeeLovelace int32 `json:"fee_lovelace"`

	// Indicates if the transaction is signed or raw
	Signed bool `json:"signed"`
}
