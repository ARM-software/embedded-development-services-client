<!--
Copyright (C) 2020-2026 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# \IntellisenseJobsAPI

All URIs are relative to *https://all.api.keil.arm.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelIntellisenseJob**](IntellisenseJobsAPI.md#CancelIntellisenseJob) | **Post** /intellisense-jobs/{jobName}/cancel | Cancel an Intellisense Job.
[**DeleteIntellisenseJob**](IntellisenseJobsAPI.md#DeleteIntellisenseJob) | **Delete** /intellisense-jobs/{jobName} | Delete an Intellisense Job.
[**GetIntellisenseArtefact**](IntellisenseJobsAPI.md#GetIntellisenseArtefact) | **Get** /intellisense-jobs/{jobName}/artefacts/{artefactName} | Download the named Intellisense Artefact for the given Intellisense Job.
[**GetIntellisenseJob**](IntellisenseJobsAPI.md#GetIntellisenseJob) | **Get** /intellisense-jobs/{jobName} | Return status of an Intellisense Job.
[**GetIntellisenseJobArtefactManager**](IntellisenseJobsAPI.md#GetIntellisenseJobArtefactManager) | **Get** /intellisense-jobs/{jobName}/outputs/{artefactName} | Get the Intellisense Job artefact manager for the artefact named &#x60;artefactName&#x60; present of this Build Job.
[**GetIntellisenseMessages**](IntellisenseJobsAPI.md#GetIntellisenseMessages) | **Get** /intellisense-jobs/{jobName}/messages | Intellisense Message Feed.
[**GetIntellisenseOutputArtefact**](IntellisenseJobsAPI.md#GetIntellisenseOutputArtefact) | **Get** /intellisense-jobs/{jobName}/outputs/{artefactName}/artefact | Download the artefact for the corresponding Intellisense job.
[**ListIntellisenseArtefacts**](IntellisenseJobsAPI.md#ListIntellisenseArtefacts) | **Get** /intellisense-jobs/{jobName}/artefacts/ | List all the available Intellisense Artefacts for the given Intellisense Job.
[**ListIntellisenseJob**](IntellisenseJobsAPI.md#ListIntellisenseJob) | **Get** /intellisense-jobs/ | List all Intellisense Jobs.
[**ListIntellisenseOutputManagers**](IntellisenseJobsAPI.md#ListIntellisenseOutputManagers) | **Get** /intellisense-jobs/{jobName}/outputs/ | Get the list of artefact managers for the given Intellisense job.
[**RetainIntellisenseJob**](IntellisenseJobsAPI.md#RetainIntellisenseJob) | **Post** /intellisense-jobs/{jobName}/retain | Update how long an intellisense job will be retained before automatic deletion.



## CancelIntellisenseJob

> IntellisenseJobItem CancelIntellisenseJob(ctx, jobName).AcceptVersion(acceptVersion).Execute()

Cancel an Intellisense Job.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ARM-software/embedded-development-services-client/client"
)

func main() {
	jobName := "jobName_example" // string | Unique ID of the Intellisense Job.
	acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntellisenseJobsAPI.CancelIntellisenseJob(context.Background(), jobName).AcceptVersion(acceptVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntellisenseJobsAPI.CancelIntellisenseJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelIntellisenseJob`: IntellisenseJobItem
	fmt.Fprintf(os.Stdout, "Response from `IntellisenseJobsAPI.CancelIntellisenseJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobName** | **string** | Unique ID of the Intellisense Job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelIntellisenseJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 

### Return type

[**IntellisenseJobItem**](IntellisenseJobItem.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIntellisenseJob

> DeleteIntellisenseJob(ctx, jobName).AcceptVersion(acceptVersion).Execute()

Delete an Intellisense Job.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ARM-software/embedded-development-services-client/client"
)

func main() {
	jobName := "jobName_example" // string | Unique ID of the Intellisense Job.
	acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IntellisenseJobsAPI.DeleteIntellisenseJob(context.Background(), jobName).AcceptVersion(acceptVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntellisenseJobsAPI.DeleteIntellisenseJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobName** | **string** | Unique ID of the Intellisense Job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIntellisenseJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 

### Return type

 (empty response body)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIntellisenseArtefact

> *os.File GetIntellisenseArtefact(ctx, artefactName, jobName).AcceptVersion(acceptVersion).Execute()

Download the named Intellisense Artefact for the given Intellisense Job.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ARM-software/embedded-development-services-client/client"
)

func main() {
	artefactName := "artefactName_example" // string | Unique ID of a Job Artefact.
	jobName := "jobName_example" // string | Unique ID of the Intellisense Job.
	acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntellisenseJobsAPI.GetIntellisenseArtefact(context.Background(), artefactName, jobName).AcceptVersion(acceptVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntellisenseJobsAPI.GetIntellisenseArtefact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIntellisenseArtefact`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `IntellisenseJobsAPI.GetIntellisenseArtefact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**artefactName** | **string** | Unique ID of a Job Artefact. | 
**jobName** | **string** | Unique ID of the Intellisense Job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntellisenseArtefactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIntellisenseJob

> IntellisenseJobItem GetIntellisenseJob(ctx, jobName).AcceptVersion(acceptVersion).IfNoneMatch(ifNoneMatch).Execute()

Return status of an Intellisense Job.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ARM-software/embedded-development-services-client/client"
)

func main() {
	jobName := "jobName_example" // string | Unique ID of the Intellisense Job.
	acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)
	ifNoneMatch := "ifNoneMatch_example" // string | Caching: Optional header to improve performance. The value of this header should be the `ETag` of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntellisenseJobsAPI.GetIntellisenseJob(context.Background(), jobName).AcceptVersion(acceptVersion).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntellisenseJobsAPI.GetIntellisenseJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIntellisenseJob`: IntellisenseJobItem
	fmt.Fprintf(os.Stdout, "Response from `IntellisenseJobsAPI.GetIntellisenseJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobName** | **string** | Unique ID of the Intellisense Job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntellisenseJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 
 **ifNoneMatch** | **string** | Caching: Optional header to improve performance. The value of this header should be the &#x60;ETag&#x60; of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content. | 

### Return type

[**IntellisenseJobItem**](IntellisenseJobItem.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIntellisenseJobArtefactManager

> ArtefactManagerItem GetIntellisenseJobArtefactManager(ctx, artefactName, jobName).AcceptVersion(acceptVersion).IfNoneMatch(ifNoneMatch).Execute()

Get the Intellisense Job artefact manager for the artefact named `artefactName` present of this Build Job.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ARM-software/embedded-development-services-client/client"
)

func main() {
	artefactName := "artefactName_example" // string | Unique ID of a Job Artefact.
	jobName := "jobName_example" // string | Unique ID of the Intellisense Job.
	acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)
	ifNoneMatch := "ifNoneMatch_example" // string | Caching: Optional header to improve performance. The value of this header should be the `ETag` of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntellisenseJobsAPI.GetIntellisenseJobArtefactManager(context.Background(), artefactName, jobName).AcceptVersion(acceptVersion).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntellisenseJobsAPI.GetIntellisenseJobArtefactManager``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIntellisenseJobArtefactManager`: ArtefactManagerItem
	fmt.Fprintf(os.Stdout, "Response from `IntellisenseJobsAPI.GetIntellisenseJobArtefactManager`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**artefactName** | **string** | Unique ID of a Job Artefact. | 
**jobName** | **string** | Unique ID of the Intellisense Job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntellisenseJobArtefactManagerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 
 **ifNoneMatch** | **string** | Caching: Optional header to improve performance. The value of this header should be the &#x60;ETag&#x60; of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content. | 

### Return type

[**ArtefactManagerItem**](ArtefactManagerItem.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIntellisenseMessages

> IntellisenseMessageItem GetIntellisenseMessages(ctx, jobName).AcceptVersion(acceptVersion).IfNoneMatch(ifNoneMatch).Limit(limit).Offset(offset).Execute()

Intellisense Message Feed.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ARM-software/embedded-development-services-client/client"
)

func main() {
	jobName := "jobName_example" // string | Unique ID of the Intellisense Job.
	acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)
	ifNoneMatch := "ifNoneMatch_example" // string | Caching: Optional header to improve performance. The value of this header should be the `ETag` of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content. (optional)
	limit := int32(20) // int32 | Paging: The maximum number of items to return in a resource. (optional) (default to 20)
	offset := int32(0) // int32 | Paging:  The index of the first item to return in the resource. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntellisenseJobsAPI.GetIntellisenseMessages(context.Background(), jobName).AcceptVersion(acceptVersion).IfNoneMatch(ifNoneMatch).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntellisenseJobsAPI.GetIntellisenseMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIntellisenseMessages`: IntellisenseMessageItem
	fmt.Fprintf(os.Stdout, "Response from `IntellisenseJobsAPI.GetIntellisenseMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobName** | **string** | Unique ID of the Intellisense Job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntellisenseMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 
 **ifNoneMatch** | **string** | Caching: Optional header to improve performance. The value of this header should be the &#x60;ETag&#x60; of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content. | 
 **limit** | **int32** | Paging: The maximum number of items to return in a resource. | [default to 20]
 **offset** | **int32** | Paging:  The index of the first item to return in the resource. | [default to 0]

### Return type

[**IntellisenseMessageItem**](IntellisenseMessageItem.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIntellisenseOutputArtefact

> *os.File GetIntellisenseOutputArtefact(ctx, artefactName, jobName).AcceptVersion(acceptVersion).IfNoneMatch(ifNoneMatch).Execute()

Download the artefact for the corresponding Intellisense job.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ARM-software/embedded-development-services-client/client"
)

func main() {
	artefactName := "artefactName_example" // string | Unique ID of a Job Artefact.
	jobName := "jobName_example" // string | Unique ID of the Intellisense Job.
	acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)
	ifNoneMatch := "ifNoneMatch_example" // string | Caching: Optional header to improve performance. The value of this header should be the `ETag` of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntellisenseJobsAPI.GetIntellisenseOutputArtefact(context.Background(), artefactName, jobName).AcceptVersion(acceptVersion).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntellisenseJobsAPI.GetIntellisenseOutputArtefact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIntellisenseOutputArtefact`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `IntellisenseJobsAPI.GetIntellisenseOutputArtefact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**artefactName** | **string** | Unique ID of a Job Artefact. | 
**jobName** | **string** | Unique ID of the Intellisense Job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntellisenseOutputArtefactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 
 **ifNoneMatch** | **string** | Caching: Optional header to improve performance. The value of this header should be the &#x60;ETag&#x60; of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIntellisenseArtefacts

> SimpleCollection ListIntellisenseArtefacts(ctx, jobName).AcceptVersion(acceptVersion).IfNoneMatch(ifNoneMatch).Limit(limit).Offset(offset).Execute()

List all the available Intellisense Artefacts for the given Intellisense Job.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ARM-software/embedded-development-services-client/client"
)

func main() {
	jobName := "jobName_example" // string | Unique ID of the Intellisense Job.
	acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)
	ifNoneMatch := "ifNoneMatch_example" // string | Caching: Optional header to improve performance. The value of this header should be the `ETag` of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content. (optional)
	limit := int32(20) // int32 | Paging: The maximum number of items to return in a resource. (optional) (default to 20)
	offset := int32(0) // int32 | Paging:  The index of the first item to return in the resource. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntellisenseJobsAPI.ListIntellisenseArtefacts(context.Background(), jobName).AcceptVersion(acceptVersion).IfNoneMatch(ifNoneMatch).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntellisenseJobsAPI.ListIntellisenseArtefacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListIntellisenseArtefacts`: SimpleCollection
	fmt.Fprintf(os.Stdout, "Response from `IntellisenseJobsAPI.ListIntellisenseArtefacts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobName** | **string** | Unique ID of the Intellisense Job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListIntellisenseArtefactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 
 **ifNoneMatch** | **string** | Caching: Optional header to improve performance. The value of this header should be the &#x60;ETag&#x60; of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content. | 
 **limit** | **int32** | Paging: The maximum number of items to return in a resource. | [default to 20]
 **offset** | **int32** | Paging:  The index of the first item to return in the resource. | [default to 0]

### Return type

[**SimpleCollection**](SimpleCollection.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIntellisenseJob

> IntellisenseJobCollection ListIntellisenseJob(ctx).AcceptVersion(acceptVersion).Embed(embed).IfNoneMatch(ifNoneMatch).Limit(limit).Offset(offset).Execute()

List all Intellisense Jobs.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ARM-software/embedded-development-services-client/client"
)

func main() {
	acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)
	embed := false // bool | Embedding: The whether or not to embed resources into the collection (rather than return links). (optional) (default to false)
	ifNoneMatch := "ifNoneMatch_example" // string | Caching: Optional header to improve performance. The value of this header should be the `ETag` of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content. (optional)
	limit := int32(20) // int32 | Paging: The maximum number of items to return in a resource. (optional) (default to 20)
	offset := int32(0) // int32 | Paging:  The index of the first item to return in the resource. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntellisenseJobsAPI.ListIntellisenseJob(context.Background()).AcceptVersion(acceptVersion).Embed(embed).IfNoneMatch(ifNoneMatch).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntellisenseJobsAPI.ListIntellisenseJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListIntellisenseJob`: IntellisenseJobCollection
	fmt.Fprintf(os.Stdout, "Response from `IntellisenseJobsAPI.ListIntellisenseJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListIntellisenseJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 
 **embed** | **bool** | Embedding: The whether or not to embed resources into the collection (rather than return links). | [default to false]
 **ifNoneMatch** | **string** | Caching: Optional header to improve performance. The value of this header should be the &#x60;ETag&#x60; of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content. | 
 **limit** | **int32** | Paging: The maximum number of items to return in a resource. | [default to 20]
 **offset** | **int32** | Paging:  The index of the first item to return in the resource. | [default to 0]

### Return type

[**IntellisenseJobCollection**](IntellisenseJobCollection.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIntellisenseOutputManagers

> ArtefactManagerCollection ListIntellisenseOutputManagers(ctx, jobName).AcceptVersion(acceptVersion).Embed(embed).IfNoneMatch(ifNoneMatch).Limit(limit).Offset(offset).Execute()

Get the list of artefact managers for the given Intellisense job.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ARM-software/embedded-development-services-client/client"
)

func main() {
	jobName := "jobName_example" // string | Unique ID of the Intellisense Job.
	acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)
	embed := false // bool | Embedding: The whether or not to embed resources into the collection (rather than return links). (optional) (default to false)
	ifNoneMatch := "ifNoneMatch_example" // string | Caching: Optional header to improve performance. The value of this header should be the `ETag` of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content. (optional)
	limit := int32(20) // int32 | Paging: The maximum number of items to return in a resource. (optional) (default to 20)
	offset := int32(0) // int32 | Paging:  The index of the first item to return in the resource. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntellisenseJobsAPI.ListIntellisenseOutputManagers(context.Background(), jobName).AcceptVersion(acceptVersion).Embed(embed).IfNoneMatch(ifNoneMatch).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntellisenseJobsAPI.ListIntellisenseOutputManagers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListIntellisenseOutputManagers`: ArtefactManagerCollection
	fmt.Fprintf(os.Stdout, "Response from `IntellisenseJobsAPI.ListIntellisenseOutputManagers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobName** | **string** | Unique ID of the Intellisense Job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListIntellisenseOutputManagersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 
 **embed** | **bool** | Embedding: The whether or not to embed resources into the collection (rather than return links). | [default to false]
 **ifNoneMatch** | **string** | Caching: Optional header to improve performance. The value of this header should be the &#x60;ETag&#x60; of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content. | 
 **limit** | **int32** | Paging: The maximum number of items to return in a resource. | [default to 20]
 **offset** | **int32** | Paging:  The index of the first item to return in the resource. | [default to 0]

### Return type

[**ArtefactManagerCollection**](ArtefactManagerCollection.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetainIntellisenseJob

> IntellisenseJobItem RetainIntellisenseJob(ctx, jobName).AcceptVersion(acceptVersion).RetainBuildJobRequest(retainBuildJobRequest).Execute()

Update how long an intellisense job will be retained before automatic deletion.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/ARM-software/embedded-development-services-client/client"
)

func main() {
	jobName := "jobName_example" // string | Unique ID of the Intellisense Job.
	acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)
	retainBuildJobRequest := *openapiclient.NewRetainBuildJobRequest() // RetainBuildJobRequest | TTL configuration. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntellisenseJobsAPI.RetainIntellisenseJob(context.Background(), jobName).AcceptVersion(acceptVersion).RetainBuildJobRequest(retainBuildJobRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntellisenseJobsAPI.RetainIntellisenseJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetainIntellisenseJob`: IntellisenseJobItem
	fmt.Fprintf(os.Stdout, "Response from `IntellisenseJobsAPI.RetainIntellisenseJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobName** | **string** | Unique ID of the Intellisense Job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetainIntellisenseJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 
 **retainBuildJobRequest** | [**RetainBuildJobRequest**](RetainBuildJobRequest.md) | TTL configuration. | 

### Return type

[**IntellisenseJobItem**](IntellisenseJobItem.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

