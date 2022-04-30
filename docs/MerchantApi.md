# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConfirmPayment**](MerchantApi.md#ConfirmPayment) | **Post** /v1/merchant/payment/{transactionId}/confirm | Init payment request
[**CreateMerchant**](MerchantApi.md#CreateMerchant) | **Post** /v1/merchant/create | create a Merchant
[**FindAllMerchants**](MerchantApi.md#FindAllMerchants) | **Get** /v1/merchant/all | get all merchants
[**FindMerchant**](MerchantApi.md#FindMerchant) | **Get** /v1/merchant/{id} | get a Merchant
[**GenerateOtp**](MerchantApi.md#GenerateOtp) | **Post** /v1/merchant/payment/gen-otp | Generate Otp
[**InitPayment**](MerchantApi.md#InitPayment) | **Post** /v1/merchant/init-payment | Init payment request
[**OneStepPayment**](MerchantApi.md#OneStepPayment) | **Post** /v1/merchant/payment/one-step | Generate Otp
[**RemoveMerchant**](MerchantApi.md#RemoveMerchant) | **Delete** /v1/merchant/{id} | delete a Merchant
[**UpdateMerchant**](MerchantApi.md#UpdateMerchant) | **Put** /v1/merchant/update | update a Merchant

# **ConfirmPayment**
> ConfirmPayment(ctx, body, transactionId)
Init payment request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SharedIdentifier**](SharedIdentifier.md)|  | 
  **transactionId** | **string**| payment id | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMerchant**
> CreateMerchant(ctx, body)
create a Merchant

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MerchantMerchant**](MerchantMerchant.md)|  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindAllMerchants**
> []MerchantMerchant FindAllMerchants(ctx, optional)
get all merchants

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MerchantApiFindAllMerchantsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MerchantApiFindAllMerchantsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **optional.String**| Filter. e.g. col1:v1,col2:v2 | 
 **order** | **optional.String**| Order. e.g. col1 desc,col2 | 
 **offset** | **optional.String**| Start position of result set. Must be an integer | 
 **limit** | **optional.String**| Limit the size of result set. Must be an integer | 

### Return type

[**[]MerchantMerchant**](merchant.Merchant.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindMerchant**
> MerchantMerchant FindMerchant(ctx, id)
get a Merchant

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| identifier of the Merchant | 

### Return type

[**MerchantMerchant**](merchant.Merchant.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GenerateOtp**
> GenerateOtp(ctx, body, transactionId)
Generate Otp

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SharedIdentifier**](SharedIdentifier.md)|  | 
  **transactionId** | **string**| payment id | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InitPayment**
> InitPayment(ctx, body)
Init payment request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MerchantInitPaymentRequest**](MerchantInitPaymentRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OneStepPayment**
> OneStepPayment(ctx, body, transactionId)
Generate Otp

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MerchantOneStepPayment**](MerchantOneStepPayment.md)|  | 
  **transactionId** | **string**| payment id | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveMerchant**
> RemoveMerchant(ctx, id)
delete a Merchant

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| identifier of the Merchant | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMerchant**
> UpdateMerchant(ctx, body)
update a Merchant

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MerchantMerchant**](MerchantMerchant.md)|  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

