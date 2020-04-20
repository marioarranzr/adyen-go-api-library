/*
 * Adyen Payout API
 *
 * A set of API endpoints that allow you to store payout details, confirm, or decline a payout.  For more information, refer to [Online payouts](https://docs.adyen.com/checkout/online-payouts).
 *
 * API version: 52
 * Contact: support@adyen.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package payouts

// ThreeDSecureData struct for ThreeDSecureData
type ThreeDSecureData struct {
	// In 3D Secure 1, the authentication response if the shopper was redirected.  In 3D Secure 2, this is the `transStatus` from the challenge result. If the transaction was frictionless, omit this parameter.
	AuthenticationResponse string `json:"authenticationResponse,omitempty"`
	// The cardholder authentication value (base64 encoded, 20 bytes in a decoded form).
	Cavv string `json:"cavv,omitempty"`
	// The CAVV algorithm used. Include this only for 3D Secure 1.
	CavvAlgorithm string `json:"cavvAlgorithm,omitempty"`
	// In 3D Secure 1, this is the enrollment response from the 3D directory server.    In 3D Secure 2, this is the `transStatus` from the `ARes`.
	DirectoryResponse string `json:"directoryResponse,omitempty"`
	// Supported for 3D Secure 2. The unique transaction identifier assigned by the Directory Server (DS) to identify a single transaction.
	DsTransID string `json:"dsTransID,omitempty"`
	// The electronic commerce indicator.
	Eci string `json:"eci,omitempty"`
	// The version of the 3D Secure protocol.
	ThreeDSVersion string `json:"threeDSVersion,omitempty"`
	// Supported for 3D Secure 1. The transaction identifier (Base64-encoded, 20 bytes in a decoded form).
	Xid string `json:"xid,omitempty"`
}