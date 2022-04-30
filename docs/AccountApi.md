# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccount**](AccountApi.md#GetAccount) | **Get** /v1/account/{msisdn}/{type} | get a Settings
[**GetRetailerBalance**](AccountApi.md#GetRetailerBalance) | **Post** /v1/account/retailer/balance | Get retailer balance
[**GetTransactionStatus**](AccountApi.md#GetTransactionStatus) | **Get** /v1/account/{transactionId}/status | get a Settings

# **GetAccount**
> SharedAccountProfile GetAccount(ctx, msisdn, type_)
get a Settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msisdn** | **string**| account msisdn | 
  **type_** | **string**| account type (retailer,customer) | 

### Return type

[**SharedAccountProfile**](shared.AccountProfile.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRetailerBalance**
> GetRetailerBalance(ctx, body)
Get retailer balance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SharedIdentifier**](SharedIdentifier.md)|  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTransactionStatus**
> SharedTransactionStatus GetTransactionStatus(ctx, transactionId)
get a Settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **transactionId** | **string**| transaction Id | 

### Return type

[**SharedTransactionStatus**](shared.TransactionStatus.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

