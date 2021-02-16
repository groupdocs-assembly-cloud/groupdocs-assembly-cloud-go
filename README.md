# GroupDocs.Assembly Cloud SDK for Go
This repository contains GroupDocs.Assembly Cloud SDK for Go source code. This SDK allows you to work with GroupDocs.Assembly Cloud REST APIs in your C++ applications quickly and easily, with zero initial cost.
 
See [API Reference](https://apireference.groupdocs.cloud/assembly) for full API specification.

# Key Features
 * API to Define Templates, Fetch Data Source, Insert Data in Template & Generate on the fly Reports.


### Prerequisites

To use GroupDocs.Assembly for Cloud Go SDK you need to register an account with [GroupDocs Cloud](https://www.groupdocs.cloud/) and lookup/create App Key and SID at [Cloud Dashboard](https://dashboard.groupdocs.cloud/applications). There is free quota available. For more details, see [GroupDocs Cloud Pricing](https://purchase.groupdocs.cloud/pricing).

### Installation

#### Install GroupDocs.Assembly Cloud

From Visual Stuio Code:

	Add "github.com/groupdocs-assembly-cloud/groupdocs-assembly-cloud-go/api" in the import section of your code

From the command line:

	go get -v github.com/groupdocs-assembly-cloud/groupdocs-assembly-cloud-go/api

### Sample usage

The examples below show how your application have to initiate and get a result after document is assembled:

Config.json file:
```csharp
{
	"AppKey": "your app key",
	"AppSid": "your app sid",
	"BaseUrl": "https://api.groupdocs.cloud"
} 
```
Go code:

```csharp
import (
	"fmt"
	"github.com/groupdocs-assembly-cloud/groupdocs-assembly-cloud-go/api"
	"os"
)

// init groupdocs assembly cloud api
config, _ := api.NewConfiguration("config.json")
client, _ := api.NewAPIClient(config)
ctx, _ := client.NewContextWithToken(nil)

// upload localFilePath to a cloud
// remoteName is a name in the cloud
client, ctx := UploadFileToStorage(localFilePath, remoteName)
    data, fileErr := ioutil.ReadFile(GetLocalPath("TableData.json"))
    if fileErr != nil {
		fmt.Println(fileErr)
	}
    templateInfo := api.TemplateFileInfo{
        FilePath: remoteName,
    }
    
    assembleData := api.AssembleOptions{
        TemplateFileInfo: &templateInfo,
        SaveFormat: format,
        ReportData: string(data),
    }

	output, err := client.AssemblyApi.AssembleDocument(ctx, assembleData)
```

## Dependencies
The libraray doesn't uses any non-Google Golang packages.

[Product Page](https://products.groupdocs.cloud/assembly/go) | [Documentation](https://docs.groupdocs.cloud/display/assemblycloud/Home) | [API Reference](https://apireference.groupdocs.cloud/assembly/) | [Code Samples](https://github.com/groupdocs-assembly-cloud/groupdocs-assembly-cloud-go) | [Blog](https://blog.groupdocs.cloud/category/assembly/) | [Free Support](https://forum.groupdocs.cloud/c/assembly) | [Free Trial](https://dashboard.groupdocs.cloud/applications)
