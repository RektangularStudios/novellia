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

// NovelliaProduct - Product listed on Novelia without a token
type NovelliaProduct struct {

	// Attribution of rights to product.
	Copyright string `json:"copyright"`

	// List of publishers or entities involved in token creation. Useful for onlookers to determine token origin.
	Publisher []string `json:"publisher"`

	// Iteration in update sequence for product.
	Version int32 `json:"version"`

	// Display name for token.
	Name string `json:"name"`

	// Tags for sorting and filtering. \"nsfw\" indicates NSFW content
	Tags []string `json:"tags"`

	Commission []Commission `json:"commission,omitempty"`

	Description DescriptionSet `json:"description"`

	Resource []OffChainResource `json:"resource"`

	// Token number in a set. Redundant field which makes no sense for tokens without a total-order.
	Id int32 `json:"id,omitempty"`
}
