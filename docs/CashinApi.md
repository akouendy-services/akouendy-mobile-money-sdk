# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCashin**](CashinApi.md#CreateCashin) | **Post** /v1/cashin/create | create a Cashin
[**FindAllCashins**](CashinApi.md#FindAllCashins) | **Get** /v1/cashin/all | get all cashins
[**FindCashin**](CashinApi.md#FindCashin) | **Get** /v1/cashin/{id} | get a Cashin
[**RemoveCashin**](CashinApi.md#RemoveCashin) | **Delete** /v1/cashin/{id} | delete a Cashin
[**UpdateCashin**](CashinApi.md#UpdateCashin) | **Put** /v1/cashin/update | update a Cashin

# **CreateCashin**
> CreateCashin(ctx, body)
create a Cashin

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CashinCashin**](CashinCashin.md)|  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindAllCashins**
> []CashinCashin FindAllCashins(ctx, optional)
get all cashins

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CashinApiFindAllCashinsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CashinApiFindAllCashinsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **optional.String**| Filter. e.g. col1:v1,col2:v2 | 
 **order** | **optional.String**| Order. e.g. col1 desc,col2 | 
 **offset** | **optional.String**| Start position of result set. Must be an integer | 
 **limit** | **optional.String**| Limit the size of result set. Must be an integer | 

### Return type

[**[]CashinCashin**](cashin.Cashin.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindCashin**
> CashinCashin FindCashin(ctx, id)
get a Cashin

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| identifier of the Cashin | 

### Return type

[**CashinCashin**](cashin.Cashin.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveCashin**
> RemoveCashin(ctx, id)
delete a Cashin

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| identifier of the Cashin | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCashin**
> UpdateCashin(ctx, body)
update a Cashin

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CashinCashin**](CashinCashin.md)|  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

