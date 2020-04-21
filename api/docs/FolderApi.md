# \FolderApi

All URIs are relative to *https://localhost/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CopyFolder**](FolderApi.md#CopyFolder) | **Put** /assembly/storage/folder/copy/{srcPath} | Copy folder
[**CreateFolder**](FolderApi.md#CreateFolder) | **Put** /assembly/storage/folder/{path} | Create the folder
[**DeleteFolder**](FolderApi.md#DeleteFolder) | **Delete** /assembly/storage/folder/{path} | Delete folder
[**GetFilesList**](FolderApi.md#GetFilesList) | **Get** /assembly/storage/folder/{path} | Get all files and folders within a folder
[**MoveFolder**](FolderApi.md#MoveFolder) | **Put** /assembly/storage/folder/move/{srcPath} | Move folder


# **CopyFolder**
> CopyFolder(ctx, destPath, srcPath, optional)
Copy folder

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **destPath** | **string**| Destination folder path e.g. &#39;/dst&#39; | 
  **srcPath** | **string**| Source folder path e.g. /Folder1 | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **destPath** | **string**| Destination folder path e.g. &#39;/dst&#39; | 
 **srcPath** | **string**| Source folder path e.g. /Folder1 | 
 **srcStorageName** | **string**| Source storage name | 
 **destStorageName** | **string**| Destination storage name | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateFolder**
> CreateFolder(ctx, path, optional)
Create the folder

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **path** | **string**| Target folder&#39;s path e.g. Folder1/Folder2/. The folders will be created recursively | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string**| Target folder&#39;s path e.g. Folder1/Folder2/. The folders will be created recursively | 
 **storageName** | **string**| Storage name | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFolder**
> DeleteFolder(ctx, path, optional)
Delete folder

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **path** | **string**| Folder path e.g. /Folder1s | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string**| Folder path e.g. /Folder1s | 
 **storageName** | **string**| Storage name | 
 **recursive** | **bool**| Enable to delete folders, subfolders and files | [default to false]

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFilesList**
> FilesList GetFilesList(ctx, path, optional)
Get all files and folders within a folder

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **path** | **string**| Folder path e.g. /Folder1 | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string**| Folder path e.g. /Folder1 | 
 **storageName** | **string**| Storage name | 

### Return type

[**FilesList**](FilesList.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MoveFolder**
> MoveFolder(ctx, destPath, srcPath, optional)
Move folder

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **destPath** | **string**| Destination folder path to move to e.g &#39;/dst&#39; | 
  **srcPath** | **string**| Source folder path e.g. /Folder1 | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **destPath** | **string**| Destination folder path to move to e.g &#39;/dst&#39; | 
 **srcPath** | **string**| Source folder path e.g. /Folder1 | 
 **srcStorageName** | **string**| Source storage name | 
 **destStorageName** | **string**| Destination storage name | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

