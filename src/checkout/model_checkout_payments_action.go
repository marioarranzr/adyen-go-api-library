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
// CheckoutPaymentsAction struct for CheckoutPaymentsAction
type CheckoutPaymentsAction struct {
	// The voucher alternative reference code.
	AlternativeReference string `json:"alternativeReference,omitempty"`
	// When the redirect URL must be accessed via POST, use this data to post to the redirect URL.
	Data *map[string]interface{} `json:"data,omitempty"`
	// The URL to download the voucher.
	DownloadUrl string `json:"downloadUrl,omitempty"`
	// An entity number of Multibanco.
	Entity string `json:"entity,omitempty"`
	// The date time of the voucher expiry.
	ExpiresAt string `json:"expiresAt,omitempty"`
	InitialAmount *Amount `json:"initialAmount,omitempty"`
	// The URL to the detailed instructions to make payment using the voucher.
	InstructionsUrl string `json:"instructionsUrl,omitempty"`
	// The issuer of the voucher.
	Issuer string `json:"issuer,omitempty"`
	// The shopper telephone number (partially masked).
	MaskedTelephoneNumber string `json:"maskedTelephoneNumber,omitempty"`
	// The merchant name.
	MerchantName string `json:"merchantName,omitempty"`
	// The merchant reference.
	MerchantReference string `json:"merchantReference,omitempty"`
	// Specifies the HTTP method, for example GET or POST.
	Method string `json:"method,omitempty"`
	// When non-empty, contains a value that you must submit to the `/payments/details` endpoint. In some cases, required for polling.
	PaymentData string `json:"paymentData,omitempty"`
	// Specifies the payment method.
	PaymentMethodType string `json:"paymentMethodType,omitempty"`
	// The contents of the QR code as a UTF8 string.
	QrCodeData string `json:"qrCodeData,omitempty"`
	// The voucher reference code.
	Reference string `json:"reference,omitempty"`
	// The shopper email.
	ShopperEmail string `json:"shopperEmail,omitempty"`
	// The shopper name.
	ShopperName string `json:"shopperName,omitempty"`
	Surcharge *Amount `json:"surcharge,omitempty"`
	// A token to pass to the 3DS2 Component to get the fingerprint/challenge.
	Token string `json:"token,omitempty"`
	TotalAmount *Amount `json:"totalAmount,omitempty"`
	// Enum that specifies the action that needs to be taken by the client.
	Type string `json:"type,omitempty"`
	// Specifies the URL to redirect to.
	Url string `json:"url,omitempty"`
}
