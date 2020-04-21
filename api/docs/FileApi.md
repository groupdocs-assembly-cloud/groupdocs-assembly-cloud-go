# \FileApi

All URIs are relative to *https://localhost/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CopyFile**](FileApi.md#CopyFile) | **Put** /assembly/storage/file/copy/{srcPath} | Copy file
[**DeleteFile**](FileApi.md#DeleteFile) | **Delete** /assembly/storage/file/{path} | Delete file
[**DownloadFile**](FileApi.md#DownloadFile) | **Get** /assembly/storage/file/{path} | Download file
[**MoveFile**](FileApi.md#MoveFile) | **Put** /assembly/storage/file/move/{srcPath} | Move file
[**UploadFile**](FileApi.md#UploadFile) | **Put** /assembly/storage/file/{path} | Upload file


# **CopyFile**
> CopyFile(ctx, destPath, srcPath, optional)
Copy file

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **destPath** | **string**| Destination file path | 
  **srcPath** | **string**| Source file&#39;s path e.g. &#39;/Folder 1/file.ext&#39; or &#39;/Bucket/Folder 1/file.ext&#39; | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **destPath** | **string**| Destination file path | 
 **srcPath** | **string**| Source file&#39;s path e.g. &#39;/Folder 1/file.ext&#39; or &#39;/Bucket/Folder 1/file.ext&#39; | 
 **srcStorageName** | **string**| Source storage name | 
 **destStorageName** | **string**| Destination storage name | 
 **versionId** | **string**| File version ID to copy | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFile**
> DeleteFile(ctx, path, optional)
Delete file

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **path** | **string**| Path of the file including file name and extension e.g. /Folder1/file.ext | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string**| Path of the file including file name and extension e.g. /Folder1/file.ext | 
 **storageName** | **string**| Storage name | 
 **versionId** | **string**| File version ID to delete | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DownloadFile**
> *os.File DownloadFile(ctx, path, optional)
Download file

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **path** | **string**| Path of the file including the file name and extension e.g. /folder1/file.ext | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string**| Path of the file including the file name and extension e.g. /folder1/file.ext | 
 **storageName** | **string**| Storage name | 
 **versionId** | **string**| File version ID to download | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MoveFile**
> MoveFile(ctx, destPath, srcPath, optional)
Move file

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **destPath** | **string**| Destination file path e.g. &#39;/dest.ext&#39; | 
  **srcPath** | **string**| Source file&#39;s path e.g. &#39;/Folder 1/file.ext&#39; or &#39;/Bucket/Folder 1/file.ext&#39; | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **destPath** | **string**| Destination file path e.g. &#39;/dest.ext&#39; | 
 **srcPath** | **string**| Source file&#39;s path e.g. &#39;/Folder 1/file.ext&#39; or &#39;/Bucket/Folder 1/file.ext&#39; | 
 **srcStorageName** | **string**| Source storage name | 
 **destStorageName** | **string**| Destination storage name | 
 **versionId** | **string**| File version ID to move | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadFile**
> FilesUploadResult UploadFile(ctx, file, path, optional)
Upload file

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **file** | ***os.File**| File to upload | 
  **path** | **string**| Path where to upload including filename and extension e.g. /file.ext or /Folder 1/file.ext              If the content is multipart and path does not contains the file name it tries to get them from filename parameter              from Content-Disposition header. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File**| File to upload | 
 **path** | **string**| Path where to upload including filename and extension e.g. /file.ext or /Folder 1/file.ext              If the content is multipart and path does not contains the file name it tries to get them from filename parameter              from Content-Disposition header. | 
 **storageName** | **string**| Storage name | 

### Return type

[**FilesUploadResult**](FilesUploadResult.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/json, application/xml, multipart/form-data
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

