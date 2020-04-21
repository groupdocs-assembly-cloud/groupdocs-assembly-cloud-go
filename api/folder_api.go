/*
* MIT License

* Copyright (c) 2020 Aspose Pty Ltd

* Permission is hereby granted, free of charge, to any person obtaining a copy
* of this software and associated documentation files (the "Software"), to deal
* in the Software without restriction, including without limitation the rights
* to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
* copies of the Software, and to permit persons to whom the Software is
* furnished to do so, subject to the following conditions:

* The above copyright notice and this permission notice shall be included in all
* copies or substantial portions of the Software.

* THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
* IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
* FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
* AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
* LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
* OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
* SOFTWARE.
*/


package api

import (
	"net/url"
	"net/http"
	"strings"
	"golang.org/x/net/context"
	"encoding/json"
	"fmt"
)

// Linger please
var (
	_ context.Context
)

type FolderApiService service

func (w *AssemblyApiErrorResponse) Error() string {
	out, err := json.Marshal(w)
	if err != nil {
		return err.Error()
	}

	return string(out)
}


/* FolderApiService Copy folder
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param destPath Destination folder path e.g. &#39;/dst&#39;
 @param srcPath Source folder path e.g. /Folder1
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "srcStorageName" (string) Source storage name
     @param "destStorageName" (string) Destination storage name
 @return */
func (a *FolderApiService) CopyFolder(ctx context.Context, destPath string, srcPath string, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/assembly/storage/folder/copy/{srcPath}"
	localVarPath = strings.Replace(localVarPath, "{"+"srcPath"+"}", fmt.Sprintf("%v", srcPath), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["srcStorageName"], "string", "srcStorageName"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(localVarOptionals["destStorageName"], "string", "destStorageName"); err != nil {
		return nil, err
	}

	localVarQueryParams.Add("DestPath", parameterToString(destPath, ""))
	if localVarTempParam, localVarOk := localVarOptionals["srcStorageName"].(string); localVarOk {
		localVarQueryParams.Add("SrcStorageName", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["destStorageName"].(string); localVarOk {
		localVarQueryParams.Add("DestStorageName", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json", "application/xml",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		"application/xml",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError AssemblyApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	return localVarHttpResponse, err
}

/* FolderApiService Create the folder
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param path Target folder&#39;s path e.g. Folder1/Folder2/. The folders will be created recursively
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storageName" (string) Storage name
 @return */
func (a *FolderApiService) CreateFolder(ctx context.Context, path string, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/assembly/storage/folder/{path}"
	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", fmt.Sprintf("%v", path), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["storageName"], "string", "storageName"); err != nil {
		return nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storageName"].(string); localVarOk {
		localVarQueryParams.Add("StorageName", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json", "application/xml",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		"application/xml",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError AssemblyApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	return localVarHttpResponse, err
}

/* FolderApiService Delete folder
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param path Folder path e.g. /Folder1s
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storageName" (string) Storage name
     @param "recursive" (bool) Enable to delete folders, subfolders and files
 @return */
func (a *FolderApiService) DeleteFolder(ctx context.Context, path string, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/assembly/storage/folder/{path}"
	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", fmt.Sprintf("%v", path), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["storageName"], "string", "storageName"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(localVarOptionals["recursive"], "bool", "recursive"); err != nil {
		return nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storageName"].(string); localVarOk {
		localVarQueryParams.Add("StorageName", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["recursive"].(bool); localVarOk {
		localVarQueryParams.Add("Recursive", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json", "application/xml",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		"application/xml",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError AssemblyApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	return localVarHttpResponse, err
}

/* FolderApiService Get all files and folders within a folder
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param path Folder path e.g. /Folder1
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storageName" (string) Storage name
 @return FilesList*/
func (a *FolderApiService) GetFilesList(ctx context.Context, path string, localVarOptionals map[string]interface{}) ( FilesList,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
		successPayload FilesList
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/assembly/storage/folder/{path}"
	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", fmt.Sprintf("%v", path), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["storageName"], "string", "storageName"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storageName"].(string); localVarOk {
		localVarQueryParams.Add("StorageName", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json", "application/xml",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		"application/xml",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError AssemblyApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return successPayload, localVarHttpResponse, err
		}
		
		return successPayload, localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}
	
	return successPayload, localVarHttpResponse, err
}

/* FolderApiService Move folder
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param destPath Destination folder path to move to e.g &#39;/dst&#39;
 @param srcPath Source folder path e.g. /Folder1
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "srcStorageName" (string) Source storage name
     @param "destStorageName" (string) Destination storage name
 @return */
func (a *FolderApiService) MoveFolder(ctx context.Context, destPath string, srcPath string, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/assembly/storage/folder/move/{srcPath}"
	localVarPath = strings.Replace(localVarPath, "{"+"srcPath"+"}", fmt.Sprintf("%v", srcPath), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["srcStorageName"], "string", "srcStorageName"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(localVarOptionals["destStorageName"], "string", "destStorageName"); err != nil {
		return nil, err
	}

	localVarQueryParams.Add("DestPath", parameterToString(destPath, ""))
	if localVarTempParam, localVarOk := localVarOptionals["srcStorageName"].(string); localVarOk {
		localVarQueryParams.Add("SrcStorageName", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["destStorageName"].(string); localVarOk {
		localVarQueryParams.Add("DestStorageName", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json", "application/xml",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		"application/xml",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localFiles)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	if localVarHttpResponse.StatusCode >= 300 {
		defer localVarHttpResponse.Body.Close()
		
		var apiError AssemblyApiErrorResponse;

		if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&apiError); err != nil {
			return localVarHttpResponse, err
		}
		
		return localVarHttpResponse, &apiError
	}
	defer localVarHttpResponse.Body.Close()
	return localVarHttpResponse, err
}

