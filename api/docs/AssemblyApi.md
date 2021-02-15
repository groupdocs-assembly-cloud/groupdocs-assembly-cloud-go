# \AssemblyApi

All URIs are relative to *https://localhost/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssembleDocument**](AssemblyApi.md#AssembleDocument) | **Post** /assembly/assemble | Builds a document using document template and XML or JSON data passed in request.
[**CopyFile**](AssemblyApi.md#CopyFile) | **Put** /assembly/storage/file/copy/{srcPath} | Copy file
[**CopyFolder**](AssemblyApi.md#CopyFolder) | **Put** /assembly/storage/folder/copy/{srcPath} | Copy folder
[**CreateFolder**](AssemblyApi.md#CreateFolder) | **Put** /assembly/storage/folder/{path} | Create the folder
[**DeleteFile**](AssemblyApi.md#DeleteFile) | **Delete** /assembly/storage/file/{path} | Delete file
[**DeleteFolder**](AssemblyApi.md#DeleteFolder) | **Delete** /assembly/storage/folder/{path} | Delete folder
[**DownloadFile**](AssemblyApi.md#DownloadFile) | **Get** /assembly/storage/file/{path} | Download file
[**GetFilesList**](AssemblyApi.md#GetFilesList) | **Get** /assembly/storage/folder/{path} | Get all files and folders within a folder
[**GetSupportedFileFormats**](AssemblyApi.md#GetSupportedFileFormats) | **Get** /assembly/formats | Retrieves list of supported file formats.
[**MoveFile**](AssemblyApi.md#MoveFile) | **Put** /assembly/storage/file/move/{srcPath} | Move file
[**MoveFolder**](AssemblyApi.md#MoveFolder) | **Put** /assembly/storage/folder/move/{srcPath} | Move folder
[**UploadFile**](AssemblyApi.md#UploadFile) | **Put** /assembly/storage/file/{path} | Upload file


# **AssembleDocument**
> *os.File AssembleDocument(ctx, assembleOptions)
Builds a document using document template and XML or JSON data passed in request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **assembleOptions** | [**AssembleOptions**](AssembleOptions.md)| Assemble Options. It should be JSON or XML with TemplateFileInfo, SaveFormat, ReportData and etc.              | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

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

# **GetSupportedFileFormats**
> FileFormatsResponse GetSupportedFileFormats(ctx, )
Retrieves list of supported file formats.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**FileFormatsResponse**](FileFormatsResponse.md)

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

# **UploadFile**
> FilesUploadResult UploadFile(ctx, fileContent, path, optional)
Upload file

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fileContent** | ***os.File**| File to upload | 
  **path** | **string**| Path where to upload including filename and extension e.g. /file.ext or /Folder 1/file.ext              If the content is multipart and path does not contains the file name it tries to get them from filename parameter              from Content-Disposition header. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fileContent** | ***os.File**| File to upload | 
 **path** | **string**| Path where to upload including filename and extension e.g. /file.ext or /Folder 1/file.ext              If the content is multipart and path does not contains the file name it tries to get them from filename parameter              from Content-Disposition header. | 
 **storageName** | **string**| Storage name | 

### Return type

[**FilesUploadResult**](FilesUploadResult.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

