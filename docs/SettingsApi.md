# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSettings**](SettingsApi.md#CreateSettings) | **Post** /v1/settings/{providerId}/create | create a Settings
[**FindAllSettingss**](SettingsApi.md#FindAllSettingss) | **Get** /v1/settings/all | get all settings
[**FindSettings**](SettingsApi.md#FindSettings) | **Get** /v1/settings/{id} | get a Settings
[**GetPublicKey**](SettingsApi.md#GetPublicKey) | **Get** /v1/settings/public-key | Get account public key
[**GetTestNumber**](SettingsApi.md#GetTestNumber) | **Get** /v1/settings/test-numbers | Get Test numbers
[**RemoveSettings**](SettingsApi.md#RemoveSettings) | **Delete** /v1/settings/{id} | delete a Settings
[**UpdateSettings**](SettingsApi.md#UpdateSettings) | **Put** /v1/settings/update | update a Settings

# **CreateSettings**
> CreateSettings(ctx, body, providerId)
create a Settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SettingsSettings**](SettingsSettings.md)|  | 
  **providerId** | **string**| Provider id | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindAllSettingss**
> []SettingsSettings FindAllSettingss(ctx, optional)
get all settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SettingsApiFindAllSettingssOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SettingsApiFindAllSettingssOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **optional.String**| Filter. e.g. col1:v1,col2:v2 | 
 **order** | **optional.String**| Order. e.g. col1 desc,col2 | 
 **offset** | **optional.String**| Start position of result set. Must be an integer | 
 **limit** | **optional.String**| Limit the size of result set. Must be an integer | 

### Return type

[**[]SettingsSettings**](settings.Settings.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindSettings**
> SettingsSettings FindSettings(ctx, id)
get a Settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| identifier of the Settings | 

### Return type

[**SettingsSettings**](settings.Settings.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPublicKey**
> SharedPublicKey GetPublicKey(ctx, )
Get account public key

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SharedPublicKey**](shared.PublicKey.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTestNumber**
> []SharedTestNumber GetTestNumber(ctx, )
Get Test numbers

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]SharedTestNumber**](shared.TestNumber.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveSettings**
> RemoveSettings(ctx, id)
delete a Settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| identifier of the Settings | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSettings**
> UpdateSettings(ctx, body)
update a Settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SettingsSettings**](SettingsSettings.md)|  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

