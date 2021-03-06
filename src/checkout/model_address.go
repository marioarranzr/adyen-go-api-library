/*
 * Adyen Checkout API
 *
 * Adyen Checkout API provides a simple and flexible way to initiate and authorise online payments. You can use the same integration for payments made with cards (including 3D Secure), mobile wallets, and local payment methods (for example, iDEAL and Sofort).  This API reference provides information on available endpoints and how to interact with them. To learn more about the API, visit [Checkout documentation](https://docs.adyen.com/checkout).  ## Authentication Each request to the Checkout API must be signed with an API key. For this, obtain an API Key from your Customer Area, as described in [How to get the API key](https://docs.adyen.com/user-management/how-to-get-the-api-key). Then set this key to the `X-API-Key` header value, for example:  ``` curl -H \"Content-Type: application/json\" \\ -H \"X-API-Key: Your_Checkout_API_key\" \\ ... ``` Note that when going live, you need to generate a new API Key to access the [live endpoints](https://docs.adyen.com/development-resources/live-endpoints).  ## Versioning Checkout API supports versioning of its endpoints through a version suffix in the endpoint URL. This suffix has the following format: \"vXX\", where XX is the version number.  For example: ``` https://checkout-test.adyen.com/v52/payments ```
 *
 * API version: 52
 * Contact: support@adyen.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package checkout
// Address struct for Address
type Address struct {
	// The name of the city.
	City string `json:"city"`
	// The two-character country code as defined in ISO-3166-1 alpha-2. For example, **US**. > If you don't know the country or are not collecting the country from the shopper, provide `country` as `ZZ`.
	Country string `json:"country"`
	// The number or name of the house.
	HouseNumberOrName string `json:"houseNumberOrName"`
	// A maximum of five digits for an address in the US, or a maximum of ten characters for an address in all other countries.
	PostalCode string `json:"postalCode"`
	// State or province codes as defined in ISO 3166-2. For example, **CA** in the US or **ON** in Canada. > Required for the US and Canada.
	StateOrProvince string `json:"stateOrProvince,omitempty"`
	// The name of the street. > The house number should not be included in this field; it should be separately provided via `houseNumberOrName`.
	Street string `json:"street"`
}
