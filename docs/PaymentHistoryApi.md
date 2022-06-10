# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePaymentHistory**](PaymentHistoryApi.md#CreatePaymentHistory) | **Post** /v1/paymenthistory/create | create a PaymentHistory
[**FindAllPaymentHistorys**](PaymentHistoryApi.md#FindAllPaymentHistorys) | **Get** /v1/paymenthistory/all | get all paymenthistorys
[**FindPaymentHistory**](PaymentHistoryApi.md#FindPaymentHistory) | **Get** /v1/paymenthistory/{id} | get a PaymentHistory
[**RemovePaymentHistory**](PaymentHistoryApi.md#RemovePaymentHistory) | **Delete** /v1/paymenthistory/{id} | delete a PaymentHistory
[**UpdatePaymentHistory**](PaymentHistoryApi.md#UpdatePaymentHistory) | **Put** /v1/paymenthistory/update | update a PaymentHistory

# **CreatePaymentHistory**
> CreatePaymentHistory(ctx, body)
create a PaymentHistory

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PaymentPaymentHistory**](PaymentPaymentHistory.md)|  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindAllPaymentHistorys**
> []PaymentPaymentHistory FindAllPaymentHistorys(ctx, optional)
get all paymenthistorys

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PaymentHistoryApiFindAllPaymentHistorysOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PaymentHistoryApiFindAllPaymentHistorysOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **optional.String**| Filter. e.g. col1:v1,col2:v2 | 
 **order** | **optional.String**| Order. e.g. col1 desc,col2 | 
 **offset** | **optional.String**| Start position of result set. Must be an integer | 
 **limit** | **optional.String**| Limit the size of result set. Must be an integer | 

### Return type

[**[]PaymentPaymentHistory**](payment.PaymentHistory.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindPaymentHistory**
> PaymentPaymentHistory FindPaymentHistory(ctx, id)
get a PaymentHistory

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| identifier of the PaymentHistory | 

### Return type

[**PaymentPaymentHistory**](payment.PaymentHistory.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemovePaymentHistory**
> RemovePaymentHistory(ctx, id)
delete a PaymentHistory

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| identifier of the PaymentHistory | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePaymentHistory**
> UpdatePaymentHistory(ctx, body)
update a PaymentHistory

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PaymentPaymentHistory**](PaymentPaymentHistory.md)|  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

