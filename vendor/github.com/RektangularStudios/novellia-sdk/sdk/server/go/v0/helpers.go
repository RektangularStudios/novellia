/*
 * novellia-api
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.4.0
 * Contact: contact@rektangularstudios.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package novellia_sdk

//Response return a ImplResponse struct filled
func Response(code int, body interface{}) ImplResponse {
	return ImplResponse{Code: code, Body: body}
}

