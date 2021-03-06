/*
 * Adyen BinLookup API
 *
 * The BIN Lookup API provides endpoints for retrieving information, such as cost estimates, and 3D Secure supported version based on a given BIN.
 *
 * API version: 50
 * Contact: support@adyen.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package binlookup
// DSPublicKeyDetail struct for DSPublicKeyDetail
type DSPublicKeyDetail struct {
	// Card brand.
	Brand string `json:"brand,omitempty"`
	// Directory Server (DS) identifier.
	DirectoryServerId string `json:"directoryServerId,omitempty"`
	// Public key. The 3D Secure 2 SDK encrypts the device information by using the DS public key.
	PublicKey string `json:"publicKey,omitempty"`
}
