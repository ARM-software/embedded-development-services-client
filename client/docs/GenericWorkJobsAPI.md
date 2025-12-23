<!--
Copyright (C) 2020-2025 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# \GenericWorkJobsAPI

All URIs are relative to *https://all.api.keil.arm.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelGenericWorkJob**](GenericWorkJobsAPI.md#CancelGenericWorkJob) | **Post** /generic-work-jobs/{jobName}/cancel | Cancel a generic job.
[**CompleteGenericWorkJob**](GenericWorkJobsAPI.md#CompleteGenericWorkJob) | **Post** /generic-work-jobs/{jobName}/complete | Stop a generic job gracefully.
[**ControlGenericWorkJob**](GenericWorkJobsAPI.md#ControlGenericWorkJob) | **Post** /generic-work-jobs/{jobName}/control | Send a control command for a generic job.
[**DeleteGenericWorkJob**](GenericWorkJobsAPI.md#DeleteGenericWorkJob) | **Delete** /generic-work-jobs/{jobName} | Delete a generic job.
[**GetGenericWorkJob**](GenericWorkJobsAPI.md#GetGenericWorkJob) | **Get** /generic-work-jobs/{jobName} | Return the status of a Generic Work Job.
[**GetGenericWorkJobArtefactManager**](GenericWorkJobsAPI.md#GetGenericWorkJobArtefactManager) | **Get** /generic-work-jobs/{jobName}/artefacts/{artefactName} | Get the job&#39;s artefact manager for the artefact named &#x60;artefactName&#x60;.
[**GetGenericWorkJobMessages**](GenericWorkJobsAPI.md#GetGenericWorkJobMessages) | **Get** /generic-work-jobs/{jobName}/messages | GenericWorkJob Message Feed.
[**GetGenericWorkJobOutputArtefact**](GenericWorkJobsAPI.md#GetGenericWorkJobOutputArtefact) | **Get** /generic-work-jobs/{jobName}/artefacts/{artefactName}/artefact | Download the artefact for the corresponding generic work job.
[**ListGenericWorkJobArtefactManagers**](GenericWorkJobsAPI.md#ListGenericWorkJobArtefactManagers) | **Get** /generic-work-jobs/{jobName}/artefacts/ | Get the list of artefact managers for the given job.
[**ListGenericWorkJobs**](GenericWorkJobsAPI.md#ListGenericWorkJobs) | **Get** /generic-work-jobs/ | List all jobs performed by generic workers.
[**RetainGenericWorkJob**](GenericWorkJobsAPI.md#RetainGenericWorkJob) | **Post** /generic-work-jobs/{jobName}/retain | Update how long a generic work job will be retained before automatic deletion.



## CancelGenericWorkJob

> GenericWorkJobItem CancelGenericWorkJob(ctx, jobName).AcceptVersion(acceptVersion).Execute()

Cancel a generic job.



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
	jobName := "jobName_example" // string | Unique ID of the generic work Job.
	acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GenericWorkJobsAPI.CancelGenericWorkJob(context.Background(), jobName).AcceptVersion(acceptVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GenericWorkJobsAPI.CancelGenericWorkJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelGenericWorkJob`: GenericWorkJobItem
	fmt.Fprintf(os.Stdout, "Response from `GenericWorkJobsAPI.CancelGenericWorkJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobName** | **string** | Unique ID of the generic work Job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelGenericWorkJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 

### Return type

[**GenericWorkJobItem**](GenericWorkJobItem.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompleteGenericWorkJob

> GenericWorkJobItem CompleteGenericWorkJob(ctx, jobName).AcceptVersion(acceptVersion).Execute()

Stop a generic job gracefully.



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
	jobName := "jobName_example" // string | Unique ID of the generic work Job.
	acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GenericWorkJobsAPI.CompleteGenericWorkJob(context.Background(), jobName).AcceptVersion(acceptVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GenericWorkJobsAPI.CompleteGenericWorkJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompleteGenericWorkJob`: GenericWorkJobItem
	fmt.Fprintf(os.Stdout, "Response from `GenericWorkJobsAPI.CompleteGenericWorkJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobName** | **string** | Unique ID of the generic work Job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompleteGenericWorkJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 

### Return type

[**GenericWorkJobItem**](GenericWorkJobItem.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ControlGenericWorkJob

> GenericWorkJobItem ControlGenericWorkJob(ctx, jobName).GenericWorkJobControlCommandItem(genericWorkJobControlCommandItem).AcceptVersion(acceptVersion).Execute()

Send a control command for a generic job.



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
	jobName := "jobName_example" // string | Unique ID of the generic work Job.
	genericWorkJobControlCommandItem := *openapiclient.NewGenericWorkJobControlCommandItem("end-session", "arm") // GenericWorkJobControlCommandItem | A definition of the control command to send for a job.
	acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GenericWorkJobsAPI.ControlGenericWorkJob(context.Background(), jobName).GenericWorkJobControlCommandItem(genericWorkJobControlCommandItem).AcceptVersion(acceptVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GenericWorkJobsAPI.ControlGenericWorkJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ControlGenericWorkJob`: GenericWorkJobItem
	fmt.Fprintf(os.Stdout, "Response from `GenericWorkJobsAPI.ControlGenericWorkJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobName** | **string** | Unique ID of the generic work Job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiControlGenericWorkJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **genericWorkJobControlCommandItem** | [**GenericWorkJobControlCommandItem**](GenericWorkJobControlCommandItem.md) | A definition of the control command to send for a job. | 
 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 

### Return type

[**GenericWorkJobItem**](GenericWorkJobItem.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGenericWorkJob

> DeleteGenericWorkJob(ctx, jobName).AcceptVersion(acceptVersion).Execute()

Delete a generic job.



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
	jobName := "jobName_example" // string | Unique ID of the generic work Job.
	acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GenericWorkJobsAPI.DeleteGenericWorkJob(context.Background(), jobName).AcceptVersion(acceptVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GenericWorkJobsAPI.DeleteGenericWorkJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobName** | **string** | Unique ID of the generic work Job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGenericWorkJobRequest struct via the builder pattern


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


## GetGenericWorkJob

> GenericWorkJobItem GetGenericWorkJob(ctx, jobName).AcceptVersion(acceptVersion).IfNoneMatch(ifNoneMatch).Execute()

Return the status of a Generic Work Job.



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
	jobName := "jobName_example" // string | Unique ID of the generic work Job.
	acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)
	ifNoneMatch := "ifNoneMatch_example" // string | Caching: Optional header to improve performance. The value of this header should be the `ETag` of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GenericWorkJobsAPI.GetGenericWorkJob(context.Background(), jobName).AcceptVersion(acceptVersion).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GenericWorkJobsAPI.GetGenericWorkJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGenericWorkJob`: GenericWorkJobItem
	fmt.Fprintf(os.Stdout, "Response from `GenericWorkJobsAPI.GetGenericWorkJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobName** | **string** | Unique ID of the generic work Job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGenericWorkJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 
 **ifNoneMatch** | **string** | Caching: Optional header to improve performance. The value of this header should be the &#x60;ETag&#x60; of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content. | 

### Return type

[**GenericWorkJobItem**](GenericWorkJobItem.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGenericWorkJobArtefactManager

> ArtefactManagerItem GetGenericWorkJobArtefactManager(ctx, artefactName, jobName).AcceptVersion(acceptVersion).IfNoneMatch(ifNoneMatch).Execute()

Get the job's artefact manager for the artefact named `artefactName`.



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
	jobName := "jobName_example" // string | Unique ID of the generic work Job.
	acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)
	ifNoneMatch := "ifNoneMatch_example" // string | Caching: Optional header to improve performance. The value of this header should be the `ETag` of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GenericWorkJobsAPI.GetGenericWorkJobArtefactManager(context.Background(), artefactName, jobName).AcceptVersion(acceptVersion).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GenericWorkJobsAPI.GetGenericWorkJobArtefactManager``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGenericWorkJobArtefactManager`: ArtefactManagerItem
	fmt.Fprintf(os.Stdout, "Response from `GenericWorkJobsAPI.GetGenericWorkJobArtefactManager`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**artefactName** | **string** | Unique ID of a Job Artefact. | 
**jobName** | **string** | Unique ID of the generic work Job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGenericWorkJobArtefactManagerRequest struct via the builder pattern


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


## GetGenericWorkJobMessages

> NotificationFeed GetGenericWorkJobMessages(ctx, jobName).AcceptVersion(acceptVersion).IfNoneMatch(ifNoneMatch).Limit(limit).Offset(offset).Execute()

GenericWorkJob Message Feed.



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
	jobName := "jobName_example" // string | Unique ID of the generic work Job.
	acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)
	ifNoneMatch := "ifNoneMatch_example" // string | Caching: Optional header to improve performance. The value of this header should be the `ETag` of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content. (optional)
	limit := int32(20) // int32 | Paging: The maximum number of items to return in a resource. (optional) (default to 20)
	offset := int32(0) // int32 | Paging:  The index of the first item to return in the resource. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GenericWorkJobsAPI.GetGenericWorkJobMessages(context.Background(), jobName).AcceptVersion(acceptVersion).IfNoneMatch(ifNoneMatch).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GenericWorkJobsAPI.GetGenericWorkJobMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGenericWorkJobMessages`: NotificationFeed
	fmt.Fprintf(os.Stdout, "Response from `GenericWorkJobsAPI.GetGenericWorkJobMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobName** | **string** | Unique ID of the generic work Job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGenericWorkJobMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 
 **ifNoneMatch** | **string** | Caching: Optional header to improve performance. The value of this header should be the &#x60;ETag&#x60; of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content. | 
 **limit** | **int32** | Paging: The maximum number of items to return in a resource. | [default to 20]
 **offset** | **int32** | Paging:  The index of the first item to return in the resource. | [default to 0]

### Return type

[**NotificationFeed**](NotificationFeed.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGenericWorkJobOutputArtefact

> *os.File GetGenericWorkJobOutputArtefact(ctx, artefactName, jobName).AcceptVersion(acceptVersion).IfNoneMatch(ifNoneMatch).Execute()

Download the artefact for the corresponding generic work job.



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
	jobName := "jobName_example" // string | Unique ID of the generic work Job.
	acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)
	ifNoneMatch := "ifNoneMatch_example" // string | Caching: Optional header to improve performance. The value of this header should be the `ETag` of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GenericWorkJobsAPI.GetGenericWorkJobOutputArtefact(context.Background(), artefactName, jobName).AcceptVersion(acceptVersion).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GenericWorkJobsAPI.GetGenericWorkJobOutputArtefact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGenericWorkJobOutputArtefact`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `GenericWorkJobsAPI.GetGenericWorkJobOutputArtefact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**artefactName** | **string** | Unique ID of a Job Artefact. | 
**jobName** | **string** | Unique ID of the generic work Job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGenericWorkJobOutputArtefactRequest struct via the builder pattern


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


## ListGenericWorkJobArtefactManagers

> ArtefactManagerCollection ListGenericWorkJobArtefactManagers(ctx, jobName).AcceptVersion(acceptVersion).Embed(embed).IfNoneMatch(ifNoneMatch).Limit(limit).Offset(offset).Execute()

Get the list of artefact managers for the given job.



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
	jobName := "jobName_example" // string | Unique ID of the generic work Job.
	acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)
	embed := false // bool | Embedding: The whether or not to embed resources into the collection (rather than return links). (optional) (default to false)
	ifNoneMatch := "ifNoneMatch_example" // string | Caching: Optional header to improve performance. The value of this header should be the `ETag` of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content. (optional)
	limit := int32(20) // int32 | Paging: The maximum number of items to return in a resource. (optional) (default to 20)
	offset := int32(0) // int32 | Paging:  The index of the first item to return in the resource. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GenericWorkJobsAPI.ListGenericWorkJobArtefactManagers(context.Background(), jobName).AcceptVersion(acceptVersion).Embed(embed).IfNoneMatch(ifNoneMatch).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GenericWorkJobsAPI.ListGenericWorkJobArtefactManagers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGenericWorkJobArtefactManagers`: ArtefactManagerCollection
	fmt.Fprintf(os.Stdout, "Response from `GenericWorkJobsAPI.ListGenericWorkJobArtefactManagers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobName** | **string** | Unique ID of the generic work Job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListGenericWorkJobArtefactManagersRequest struct via the builder pattern


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


## ListGenericWorkJobs

> GenericWorkJobCollection ListGenericWorkJobs(ctx).AcceptVersion(acceptVersion).Embed(embed).IfNoneMatch(ifNoneMatch).Limit(limit).Offset(offset).Execute()

List all jobs performed by generic workers.



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
	resp, r, err := apiClient.GenericWorkJobsAPI.ListGenericWorkJobs(context.Background()).AcceptVersion(acceptVersion).Embed(embed).IfNoneMatch(ifNoneMatch).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GenericWorkJobsAPI.ListGenericWorkJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGenericWorkJobs`: GenericWorkJobCollection
	fmt.Fprintf(os.Stdout, "Response from `GenericWorkJobsAPI.ListGenericWorkJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListGenericWorkJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 
 **embed** | **bool** | Embedding: The whether or not to embed resources into the collection (rather than return links). | [default to false]
 **ifNoneMatch** | **string** | Caching: Optional header to improve performance. The value of this header should be the &#x60;ETag&#x60; of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content. | 
 **limit** | **int32** | Paging: The maximum number of items to return in a resource. | [default to 20]
 **offset** | **int32** | Paging:  The index of the first item to return in the resource. | [default to 0]

### Return type

[**GenericWorkJobCollection**](GenericWorkJobCollection.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetainGenericWorkJob

> GenericWorkJobItem RetainGenericWorkJob(ctx, jobName).AcceptVersion(acceptVersion).RetainBuildJobRequest(retainBuildJobRequest).Execute()

Update how long a generic work job will be retained before automatic deletion.



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
	jobName := "jobName_example" // string | Unique ID of the generic work Job.
	acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)
	retainBuildJobRequest := *openapiclient.NewRetainBuildJobRequest() // RetainBuildJobRequest | TTL configuration. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GenericWorkJobsAPI.RetainGenericWorkJob(context.Background(), jobName).AcceptVersion(acceptVersion).RetainBuildJobRequest(retainBuildJobRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GenericWorkJobsAPI.RetainGenericWorkJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetainGenericWorkJob`: GenericWorkJobItem
	fmt.Fprintf(os.Stdout, "Response from `GenericWorkJobsAPI.RetainGenericWorkJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobName** | **string** | Unique ID of the generic work Job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetainGenericWorkJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 
 **retainBuildJobRequest** | [**RetainBuildJobRequest**](RetainBuildJobRequest.md) | TTL configuration. | 

### Return type

[**GenericWorkJobItem**](GenericWorkJobItem.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

