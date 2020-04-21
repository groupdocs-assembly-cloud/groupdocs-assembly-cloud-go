//
// MIT License

// Copyright (c) 2020 Aspose Pty Ltd

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
//

package api_test

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/groupdocs-assembly-cloud/groupdocs-assembly-cloud-go/api"
	"github.com/google/uuid"
)
var remoteBaseTestDataFolder string = "Temp/SdkTests/TestData"

func GetLocalPath(filename string) string {
	_, currentFilePath, _, _ := runtime.Caller(0)
	currentFolder := filepath.Dir(currentFilePath)
	return filepath.Join(currentFolder, "TestData", filename)
}

func ReadConfiguration(t *testing.T) (cfg *api.Configuration) {
	_, filename, _, _ := runtime.Caller(0)
	configFile := filepath.Join(filepath.Dir(filename), "../config.json")
	configuration, err := api.NewConfiguration(configFile)

	if err != nil {
		t.Error(err)
	}

	return configuration
}

func PrepareTest(t *testing.T, config *api.Configuration) (apiClient *api.APIClient, ctx context.Context) {
	client, err := api.NewAPIClient(config)

	if err != nil {
		t.Error(err)
	}

	context, err := client.NewContextWithToken(nil)

	if err != nil {
		t.Error(err)
	}

	return client, context
}

func UploadFileToStorage(t *testing.T, fileName string, path string) (*api.APIClient, context.Context) {
	config := ReadConfiguration(t)
	client, ctx := PrepareTest(t, config)

	UploadNextFileToStorage(t, ctx, client, fileName, path)

	return client, ctx
}

func UploadNextFileToStorage(t *testing.T, ctx context.Context, client *api.APIClient, fileName string, path string) {
	file, fileErr := os.Open(fileName)
	if fileErr != nil {
		t.Error(fileErr)
	}

	_, _, err := client.FileApi.UploadFile(ctx, file, path, nil)
	if err != nil {
		t.Error(err)
	}
}

func TestAssembleDocument(t *testing.T) {

	format := "pdf"
	localFilePath := GetLocalPath("TableFeatures.odt")
	document, fileErr := os.Open(localFilePath)
	if fileErr != nil {
		t.Error(fileErr)
	}

    remoteName := path.Join(remoteBaseTestDataFolder, "Go", "TableFeatures.odt")
	config := ReadConfiguration(t)
	client, ctx := PrepareTest(t, config)
    data, fileErr := ioutil.ReadFile(GetLocalPath("TableData.json"))
    if fileErr != nil {
		t.Error(fileErr)
	}
    templateInfo := api.TemplateFileInfo{
        FilePath: remoteName,    
    }
    
    assembleData := api.AssembleOptionsData{
        TemplateFileInfo: &templateInfo,
        SaveFormat: "pdf",
        ReportData: data,
    }

	output, err := client.AssemblyApi.AssembleDocument(ctx, assembleData)
	defer output.Body.Close()

	if err != nil {
		t.Error(err)
	}
}