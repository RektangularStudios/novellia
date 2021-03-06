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

// ProductListElement - Item in an array of product IDs and metadata
type ProductListElement struct {

	// Product ID
	ProductId string `json:"product_id"`

	// Last modified date of the product
	Modified string `json:"modified,omitempty"`

	// Optional token ID returned to associate with product
	NativeTokenId string `json:"native_token_id,omitempty"`
}
