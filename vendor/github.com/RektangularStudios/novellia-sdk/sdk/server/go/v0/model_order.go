/*
 * novellia-api
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Contact: contact@rektangularstudios.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package novellia_sdk

// Order - Describes a single order / transaction
type Order struct {

	// Unique identifier of product
	ProductId int32 `json:"product_id"`

	// Number of product purchased in order
	OrderSize int32 `json:"order_size"`

	// Currency product is priced with
	CurrencyPolicyId string `json:"currency_policy_id"`

	// Price for a single item of product type
	UnitPrice int32 `json:"unit_price"`
}