/*
 * novellia
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.8.0
 * Contact: contact@rektangularstudios.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package novellia

// DescriptionSet - Set of descriptions
type DescriptionSet struct {

	// A short description makes a good header
	Short string `json:"short"`

	// A long description makes a good paragraph body. Supports Markdown.
	Long string `json:"long"`
}