# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConfirmPayment**](MerchantApi.md#ConfirmPayment) | **Post** /v1/merchant/payment/{transactionId}/confirm | Init payment request
[**GenerateOtp**](MerchantApi.md#GenerateOtp) | **Post** /v1/merchant/payment/gen-otp | Generate Otp
[**InitPayment**](MerchantApi.md#InitPayment) | **Post** /v1/merchant/init-payment | Init payment request
[**InitPaymentQrcode**](MerchantApi.md#InitPaymentQrcode) | **Post** /v1/merchant/init-payment-qrcode | Init payment with qrcode
[**OneStepPayment**](MerchantApi.md#OneStepPayment) | **Post** /v1/merchant/payment/one-step | Generate Otp

# **ConfirmPayment**
> ConfirmPayment(ctx, body, transactionId, optional)
Init payment request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SharedIdentifier**](SharedIdentifier.md)|  | 
  **transactionId** | **string**| payment id | 
 **optional** | ***MerchantApiConfirmPaymentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MerchantApiConfirmPaymentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **env** | **optional.**| api environment (sandbox,prod) | [default to sandbox]

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GenerateOtp**
> GenerateOtp(ctx, body, optional)
Generate Otp

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SharedIdentifier**](SharedIdentifier.md)|  | 
 **optional** | ***MerchantApiGenerateOtpOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MerchantApiGenerateOtpOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **env** | **optional.**| api environment (sandbox,prod) | [default to sandbox]

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InitPayment**
> InitPayment(ctx, body, optional)
Init payment request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MerchantInitPaymentRequest**](MerchantInitPaymentRequest.md)|  | 
 **optional** | ***MerchantApiInitPaymentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MerchantApiInitPaymentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **env** | **optional.**| api environment (sandbox,prod) | [default to sandbox]

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InitPaymentQrcode**
> MerchantPaymentQrcode InitPaymentQrcode(ctx, body, optional)
Init payment with qrcode

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MerchantInitPaymentQrcode**](MerchantInitPaymentQrcode.md)|  | 
 **optional** | ***MerchantApiInitPaymentQrcodeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MerchantApiInitPaymentQrcodeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **env** | **optional.**| api environment (sandbox,prod) | [default to sandbox]

### Return type

[**MerchantPaymentQrcode**](merchant.PaymentQrcode.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OneStepPayment**
> OneStepPayment(ctx, body, transactionId, optional)
Generate Otp

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MerchantOneStepPayment**](MerchantOneStepPayment.md)|  | 
  **transactionId** | **string**| payment id | 
 **optional** | ***MerchantApiOneStepPaymentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MerchantApiOneStepPaymentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **env** | **optional.**| api environment (sandbox,prod) | [default to sandbox]

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

