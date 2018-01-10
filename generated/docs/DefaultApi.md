# \DefaultApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAPIKey**](DefaultApi.md#CreateAPIKey) | **Post** /consumers/{consumer_id}/key-auth | 
[**CreateCertificate**](DefaultApi.md#CreateCertificate) | **Post** /certificates | 
[**CreateConsumer**](DefaultApi.md#CreateConsumer) | **Post** /consumers | 
[**CreateJWTCredential**](DefaultApi.md#CreateJWTCredential) | **Post** /consumers/{consumer_id}/jwt | 
[**CreatePlugin**](DefaultApi.md#CreatePlugin) | **Post** /apis/{api_id}/plugins | 
[**DeleteAPIKey**](DefaultApi.md#DeleteAPIKey) | **Delete** /consumers/{consumer_id}/key-auth/{apikey_id} | 
[**DeleteConsumer**](DefaultApi.md#DeleteConsumer) | **Delete** /consumers | 
[**GetCertificate**](DefaultApi.md#GetCertificate) | **Get** /certificates/{sni} | 
[**GetConsumer**](DefaultApi.md#GetConsumer) | **Get** /consumers/{consumer_id} | 
[**ListAPIKeys**](DefaultApi.md#ListAPIKeys) | **Get** /consumers/{consumer_id}/key-auth | 
[**ListCertificates**](DefaultApi.md#ListCertificates) | **Get** /certificates | 
[**ListJWTCredentials**](DefaultApi.md#ListJWTCredentials) | **Get** /consumers/{consumer_id}/jwt | 
[**ListPlugins**](DefaultApi.md#ListPlugins) | **Get** /apis/{api_id}/plugins | 
[**Oauth2Get**](DefaultApi.md#Oauth2Get) | **Get** /oauth2 | 
[**Oauth2TokensTokenGet**](DefaultApi.md#Oauth2TokensTokenGet) | **Get** /oauth2_tokens/{token} | 


# **CreateAPIKey**
> ApiKey CreateAPIKey(consumerId, optional)


Create an API Key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **consumerId** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **consumerId** | **string**|  | 
 **empty** | [**interface{}**](interface{}.md)| An empty body. | 

### Return type

[**ApiKey**](APIKey.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateCertificate**
> CreateCertificate(optional)


asdf

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **plugin** | [**Certificate**](Certificate.md)| The plugin name to activate | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateConsumer**
> CreateConsumer(consumerInput)


Create a consumer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **consumerInput** | [**Consumer**](Consumer.md)| ConsumerInput | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateJWTCredential**
> JwtCredential CreateJWTCredential(consumerId, jwtCredential)


create a jwt credential for a consumer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **consumerId** | **string**|  | 
  **jwtCredential** | [**JwtCredential**](JwtCredential.md)|  | 

### Return type

[**JwtCredential**](JWTCredential.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreatePlugin**
> Plugin CreatePlugin(apiId, optional)


asdf

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **apiId** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiId** | **string**|  | 
 **plugin** | [**Plugin**](Plugin.md)| The plugin name to activate | 

### Return type

[**Plugin**](Plugin.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAPIKey**
> DeleteAPIKey(consumerId, apikeyId)


Delete an API Key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **consumerId** | **string**|  | 
  **apikeyId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteConsumer**
> DeleteConsumer(consumerInput)


Delete a consumer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **consumerInput** | [**Consumer**](Consumer.md)| ConsumerInput | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCertificate**
> Certificate GetCertificate(sni)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **sni** | **string**|  | 

### Return type

[**Certificate**](Certificate.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConsumer**
> Consumer GetConsumer(consumerId)


Get details about a consumer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **consumerId** | **string**|  | 

### Return type

[**Consumer**](Consumer.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAPIKeys**
> InlineResponse200 ListAPIKeys(consumerId)


List API Keys

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **consumerId** | **string**|  | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCertificates**
> []Certificate ListCertificates()


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Certificate**](Certificate.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListJWTCredentials**
> InlineResponse2001 ListJWTCredentials(consumerId)


list the jwt credentials for a consumer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **consumerId** | **string**|  | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPlugins**
> InlineResponse2002 ListPlugins(apiId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **apiId** | **string**|  | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Oauth2Get**
> InlineResponse2003 Oauth2Get(clientId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **clientId** | **string**|  | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Oauth2TokensTokenGet**
> InlineResponse2004 Oauth2TokensTokenGet(token)


Lookup an oauth2 token

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
  **token** | **string**|  | 

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

