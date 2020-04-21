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
	"os"
	"io/ioutil"
	"encoding/json"
)

// Linger please
var (
	_ context.Context
)

type AssemblyApiService service

func (w *AssemblyApiErrorResponse) Error() string {
	out, err := json.Marshal(w)
	if err != nil {
		return err.Error()
	}

	return string(out)
}


/* AssemblyApiService Builds a document using document template and XML or JSON data passed in request.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param assembleData Assemble Options. It should be JSON with TemplateName, SaveFormat, ReportData and etc.             
 @return *os.File*/
func (a *AssemblyApiService) AssembleDocument(ctx context.Context, assembleData IAssembleOptionsData) (  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localFiles map[string][]byte = make(map[string][]byte)
	)

	// create path and map variables
	localVarPath := a.client.cfg.BaseUrl + "/assembly/assemble"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


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
	// body params
	localVarPostBody = &assembleData
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
	return localVarHttpResponse, err
}

