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

import (
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"

	common "github.com/adyen/adyen-go-api-library/src/common"
)

// Checkout Checkout service
type Checkout common.Service

/*
PaymentLinksPost Creates a payment link.
Creates a payment link to our hosted payment form where shoppers can pay. The list of payment methods presented to the shopper depends on the &#x60;currency&#x60; and &#x60;country&#x60; parameters sent in the request.  For more information, refer to [Pay by Link documentation](https://docs.adyen.com/checkout/pay-by-link#create-payment-links-through-api).
 * @param request CreatePaymentLinkRequest - reference of CreatePaymentLinkRequest).
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return CreatePaymentLinkResponse
*/

func (a *Checkout) PaymentLinksPost(request *CreatePaymentLinkRequest, ctxs ..._context.Context) (CreatePaymentLinkResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue CreatePaymentLinkResponse
	)

	// create path and map variables
	localVarPath := a.BasePath() + "/paymentLinks"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := common.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := common.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	if request != nil {
		localVarPostBody = request
	}

	var ctx _context.Context
	if len(ctxs) == 1 {
		ctx = ctxs[0]
	}

	r, err := a.Client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := common.NewApiError(localVarBody, localVarHTTPResponse.Status)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := common.NewApiError(localVarBody, err.Error())
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

/*
PaymentMethodsPost Returns available payment methods.
Queries the available payment methods for a transaction based on the transaction context (like amount, country, and currency). Besides giving back a list of the available payment methods, the response also returns which input details you need to collect from the shopper (to be submitted to &#x60;/payments&#x60;).  Although we highly recommend using this endpoint to ensure you are always offering the most up-to-date list of payment methods, its usage is optional. You can, for example, also cache the &#x60;/paymentMethods&#x60; response and update it once a week.
 * @param request PaymentMethodsRequest - reference of PaymentMethodsRequest).
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return PaymentMethodsResponse
*/

func (a *Checkout) PaymentMethodsPost(request *PaymentMethodsRequest, ctxs ..._context.Context) (PaymentMethodsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue PaymentMethodsResponse
	)

	// create path and map variables
	localVarPath := a.BasePath() + "/paymentMethods"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := common.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := common.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	if request != nil {
		localVarPostBody = request
	}

	var ctx _context.Context
	if len(ctxs) == 1 {
		ctx = ctxs[0]
	}

	r, err := a.Client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := common.NewApiError(localVarBody, localVarHTTPResponse.Status)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := common.NewApiError(localVarBody, err.Error())
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

/*
PaymentSessionPost Creates a payment session.
Provides the data object that can be used to start the Checkout SDK. To set up the payment, pass its amount, currency, and other required parameters. We use this to optimise the payment flow and perform better risk assessment of the transaction.  For more information, refer to [How it works](https://docs.adyen.com/checkout#howitworks).
 * @param request PaymentSetupRequest - reference of PaymentSetupRequest).
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return PaymentSetupResponse
*/

func (a *Checkout) PaymentSessionPost(request *PaymentSetupRequest, ctxs ..._context.Context) (PaymentSetupResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue PaymentSetupResponse
	)

	// create path and map variables
	localVarPath := a.BasePath() + "/paymentSession"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := common.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := common.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	if request != nil {
		localVarPostBody = request
	}

	var ctx _context.Context
	if len(ctxs) == 1 {
		ctx = ctxs[0]
	}

	r, err := a.Client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := common.NewApiError(localVarBody, localVarHTTPResponse.Status)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := common.NewApiError(localVarBody, err.Error())
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

/*
PaymentsDetailsPost Submits details for a payment.
Submits details for a payment created using &#x60;/payments&#x60;. This step is only needed when no final state has been reached on the &#x60;/payments&#x60; request (for example for 3D Secure, or when getting redirected back directly from a payment method using an app switch).  The exact details, which need to be sent to this endpoint, are always specified in the response of the associated &#x60;/payments&#x60; request.  In addition, the endpoint can be used to verify a &#x60;payload&#x60;, which is returned after coming back from the Checkout SDK or any of the redirect based methods on the Checkout API.
 * @param request DetailsRequest - reference of DetailsRequest).
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return PaymentResponse
*/

func (a *Checkout) PaymentsDetailsPost(request *DetailsRequest, ctxs ..._context.Context) (PaymentResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue PaymentResponse
	)

	// create path and map variables
	localVarPath := a.BasePath() + "/payments/details"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := common.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := common.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	if request != nil {
		localVarPostBody = request
	}

	var ctx _context.Context
	if len(ctxs) == 1 {
		ctx = ctxs[0]
	}

	r, err := a.Client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := common.NewApiError(localVarBody, localVarHTTPResponse.Status)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := common.NewApiError(localVarBody, err.Error())
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

/*
PaymentsPost Starts a transaction.
Sends payment parameters (like amount, country, and currency) together with the input details collected from the shopper. The response returns the result of the payment request: * For some payment methods (e.g. Visa, Mastercard, and SEPA Direct Debits) you&#39;ll get a final state in the &#x60;resultCode&#x60; (e.g. &#x60;Authorised&#x60; or &#x60;Refused&#x60;). * For other payment methods, you&#39;ll receive &#x60;redirectShopper&#x60; as &#x60;resultCode&#x60; together with a &#x60;redirectUrl&#x60;. In this case, the shopper must finalize the payment on the page behind the &#x60;redirectUrl&#x60;.
 * @param request PaymentRequest - reference of PaymentRequest).
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return PaymentResponse
*/

func (a *Checkout) PaymentsPost(request *PaymentRequest, ctxs ..._context.Context) (PaymentResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue PaymentResponse
	)

	// create path and map variables
	localVarPath := a.BasePath() + "/payments"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := common.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := common.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	if request != nil {
		localVarPostBody = request
	}

	var ctx _context.Context
	if len(ctxs) == 1 {
		ctx = ctxs[0]
	}

	r, err := a.Client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := common.NewApiError(localVarBody, localVarHTTPResponse.Status)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := common.NewApiError(localVarBody, err.Error())
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

/*
PaymentsResultPost Verifies payment result.
Verifies the payment result using the payload returned from the Checkout SDK.  For more information, refer to [How it works](https://docs.adyen.com/checkout#howitworks).
 * @param request PaymentVerificationRequest - reference of PaymentVerificationRequest).
 * @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return PaymentVerificationResponse
*/

func (a *Checkout) PaymentsResultPost(request *PaymentVerificationRequest, ctxs ..._context.Context) (PaymentVerificationResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue PaymentVerificationResponse
	)

	// create path and map variables
	localVarPath := a.BasePath() + "/payments/result"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := common.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := common.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	if request != nil {
		localVarPostBody = request
	}

	var ctx _context.Context
	if len(ctxs) == 1 {
		ctx = ctxs[0]
	}

	r, err := a.Client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := common.NewApiError(localVarBody, localVarHTTPResponse.Status)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := common.NewApiError(localVarBody, err.Error())
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
