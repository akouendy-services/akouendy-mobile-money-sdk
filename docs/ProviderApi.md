# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProvider**](ProviderApi.md#CreateProvider) | **Post** /v1/provider/create | create a Provider
[**FindAllProviders**](ProviderApi.md#FindAllProviders) | **Get** /v1/provider/all | get all providers
[**FindProvider**](ProviderApi.md#FindProvider) | **Get** /v1/provider/{id} | get a Provider
[**RemoveProvider**](ProviderApi.md#RemoveProvider) | **Delete** /v1/provider/{id} | delete a Provider
[**UpdateProvider**](ProviderApi.md#UpdateProvider) | **Put** /v1/provider/update | update a Provider

# **CreateProvider**
> CreateProvider(ctx, body)
create a Provider

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ProviderProvider**](ProviderProvider.md)|  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindAllProviders**
> []ProviderProvider FindAllProviders(ctx, optional)
get all providers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProviderApiFindAllProvidersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProviderApiFindAllProvidersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **optional.String**| Filter. e.g. col1:v1,col2:v2 | 
 **order** | **optional.String**| Order. e.g. col1 desc,col2 | 
 **offset** | **optional.String**| Start position of result set. Must be an integer | 
 **limit** | **optional.String**| Limit the size of result set. Must be an integer | 

### Return type

[**[]ProviderProvider**](provider.Provider.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindProvider**
> ProviderProvider FindProvider(ctx, id)
get a Provider

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| identifier of the Provider | 

### Return type

[**ProviderProvider**](provider.Provider.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveProvider**
> RemoveProvider(ctx, id)
delete a Provider

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| identifier of the Provider | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateProvider**
> UpdateProvider(ctx, body)
update a Provider

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ProviderProvider**](ProviderProvider.md)|  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

