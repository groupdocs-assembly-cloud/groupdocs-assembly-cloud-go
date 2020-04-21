# \AssemblyApi

All URIs are relative to *https://localhost/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssembleDocument**](AssemblyApi.md#AssembleDocument) | **Post** /assembly/assemble | Builds a document using document template and XML or JSON data passed in request.


# **AssembleDocument**
> *os.File AssembleDocument(ctx, assembleData)
Builds a document using document template and XML or JSON data passed in request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **assembleData** | [**AssembleOptionsData**](AssembleOptionsData.md)| Assemble Options. It should be JSON with TemplateName, SaveFormat, ReportData and etc.              | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

