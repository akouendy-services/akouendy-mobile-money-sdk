# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangePinCode**](AccountApi.md#ChangePinCode) | **Patch** /v1/account/change/pincode | Change Pin Code
[**GetAccount**](AccountApi.md#GetAccount) | **Get** /v1/account/{msisdn}/{type} | get a Settings
[**GetRetailerBalance**](AccountApi.md#GetRetailerBalance) | **Post** /v1/account/retailer/balance | Get retailer balance
[**GetTransactionStatus**](AccountApi.md#GetTransactionStatus) | **Get** /v1/account/{transactionId}/status | get a Settings
[**GetTransactions**](AccountApi.md#GetTransactions) | **Get** /v1/account/transactions | get a transactions

# **ChangePinCode**
> ChangePinCode(ctx, body, optional)
Change Pin Code

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SharedChangePinIdentifier**](SharedChangePinIdentifier.md)|  | 
 **optional** | ***AccountApiChangePinCodeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountApiChangePinCodeOpts struct
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

# **GetAccount**
> SharedAccountProfile GetAccount(ctx, msisdn, type_, optional)
get a Settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **msisdn** | **string**| account msisdn | 
  **type_** | **string**| account type (retailer,customer) | 
 **optional** | ***AccountApiGetAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountApiGetAccountOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **env** | **optional.String**| api environment (sandbox,prod) | [default to sandbox]

### Return type

[**SharedAccountProfile**](shared.AccountProfile.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRetailerBalance**
> GetRetailerBalance(ctx, body, optional)
Get retailer balance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SharedIdentifier**](SharedIdentifier.md)|  | 
 **optional** | ***AccountApiGetRetailerBalanceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountApiGetRetailerBalanceOpts struct
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

# **GetTransactionStatus**
> SharedTransactionStatus GetTransactionStatus(ctx, transactionId, optional)
get a Settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **transactionId** | **string**| transaction Id | 
 **optional** | ***AccountApiGetTransactionStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountApiGetTransactionStatusOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **env** | **optional.String**| api environment (sandbox,prod) | [default to sandbox]

### Return type

[**SharedTransactionStatus**](shared.TransactionStatus.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTransactions**
> []SharedTransaction GetTransactions(ctx, optional)
get a transactions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountApiGetTransactionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountApiGetTransactionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **env** | **optional.String**| api environment (sandbox,prod) | [default to sandbox]
 **status** | **optional.String**| Transactions status. e.g. SUCCESS, FAILED | 
 **fromDateTime** | **optional.String**| Sart date | 
 **toDateTime** | **optional.String**| End date | 

### Return type

[**[]SharedTransaction**](shared.Transaction.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

