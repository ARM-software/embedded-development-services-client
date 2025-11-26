<!--
Copyright (C) 2020-2025 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# \FPGAPayloadsAPI

All URIs are relative to *https://all.api.keil.arm.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckFPGAPayloadStatus**](FPGAPayloadsAPI.md#CheckFPGAPayloadStatus) | **Post** /repositories/{repositoryName}/payloads/{fpgaPayloadName}/check | Check and update the payload status.
[**CreateFPGAPayload**](FPGAPayloadsAPI.md#CreateFPGAPayload) | **Post** /repositories/{repositoryName}/payloads | Create an FPGA payload.
[**CreateFPGAPayloadUploadSession**](FPGAPayloadsAPI.md#CreateFPGAPayloadUploadSession) | **Post** /payloads/upload-session | Create upload session for FPGA payload.
[**DeleteFpgaPayload**](FPGAPayloadsAPI.md#DeleteFpgaPayload) | **Delete** /repositories/{repositoryName}/payloads/{fpgaPayloadName} | Delete an FPGA payload.
[**GetAlternativeFpgaPayloadUploadProgress**](FPGAPayloadsAPI.md#GetAlternativeFpgaPayloadUploadProgress) | **Get** /payloads/upload-session/{uploadSessionName} | Return FPGA payload upload progress.
[**GetFpgaPayload**](FPGAPayloadsAPI.md#GetFpgaPayload) | **Get** /repositories/{repositoryName}/payloads/{fpgaPayloadName} | Return details of specific FPGA payload.
[**GetFpgaPayloadUploadJobMessages**](FPGAPayloadsAPI.md#GetFpgaPayloadUploadJobMessages) | **Get** /upload-jobs/{uploadJobName}/messages | FPGA payload upload job Message Feed.
[**GetFpgaPayloadUploadOptions**](FPGAPayloadsAPI.md#GetFpgaPayloadUploadOptions) | **Options** /payloads/upload-session | Return service TUS protocol support.
[**GetFpgaPayloadUploadProgress**](FPGAPayloadsAPI.md#GetFpgaPayloadUploadProgress) | **Head** /payloads/upload-session/{uploadSessionName} | Return FPGA payload upload progress.
[**ListFPGAPayloads**](FPGAPayloadsAPI.md#ListFPGAPayloads) | **Get** /repositories/{repositoryName}/payloads | List payloads in a repository.
[**ListPayloads**](FPGAPayloadsAPI.md#ListPayloads) | **Get** /payloads | List payloads.
[**UploadPayload**](FPGAPayloadsAPI.md#UploadPayload) | **Patch** /payloads/upload-session/{uploadSessionName} | Upload part of a payload.



## CheckFPGAPayloadStatus

> FPGAPayloadItem CheckFPGAPayloadStatus(ctx, fpgaPayloadName, repositoryName).AcceptVersion(acceptVersion).Execute()

Check and update the payload status.



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
	fpgaPayloadName := "fpgaPayloadName_example" // string | Unique ID of the FPGA payload.
	repositoryName := "repositoryName_example" // string | Unique ID of a repository.
	acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FPGAPayloadsAPI.CheckFPGAPayloadStatus(context.Background(), fpgaPayloadName, repositoryName).AcceptVersion(acceptVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FPGAPayloadsAPI.CheckFPGAPayloadStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckFPGAPayloadStatus`: FPGAPayloadItem
	fmt.Fprintf(os.Stdout, "Response from `FPGAPayloadsAPI.CheckFPGAPayloadStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fpgaPayloadName** | **string** | Unique ID of the FPGA payload. | 
**repositoryName** | **string** | Unique ID of a repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckFPGAPayloadStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 

### Return type

[**FPGAPayloadItem**](FPGAPayloadItem.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFPGAPayload

> FPGAPayloadItem CreateFPGAPayload(ctx, repositoryName).FPGAPayloadItem(fPGAPayloadItem).AcceptVersion(acceptVersion).Execute()

Create an FPGA payload.



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
	repositoryName := "repositoryName_example" // string | Unique ID of a repository.
	fPGAPayloadItem := *openapiclient.NewFPGAPayloadItem("TODO", "TODO", true, "4545aaf", *openapiclient.NewFPGAPayloadOwner("user123", "Jane Doe"), "READY", true, true, "Demo payload for FPGA 12", "http://some-location/4545aaf") // FPGAPayloadItem | An FPGA Payload to be created.
	acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FPGAPayloadsAPI.CreateFPGAPayload(context.Background(), repositoryName).FPGAPayloadItem(fPGAPayloadItem).AcceptVersion(acceptVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FPGAPayloadsAPI.CreateFPGAPayload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFPGAPayload`: FPGAPayloadItem
	fmt.Fprintf(os.Stdout, "Response from `FPGAPayloadsAPI.CreateFPGAPayload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Unique ID of a repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFPGAPayloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fPGAPayloadItem** | [**FPGAPayloadItem**](FPGAPayloadItem.md) | An FPGA Payload to be created. | 
 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 

### Return type

[**FPGAPayloadItem**](FPGAPayloadItem.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFPGAPayloadUploadSession

> CreateFPGAPayloadUploadSession(ctx).AcceptVersion(acceptVersion).TusResumable(tusResumable).UploadDeferLength(uploadDeferLength).UploadLength(uploadLength).UploadMetadata(uploadMetadata).Execute()

Create upload session for FPGA payload.



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
	tusResumable := "1.0.0" // string | Version of the Tus protocol being used. (optional)
	uploadDeferLength := int32(56) // int32 | Set to 1 if upload size is not known at the time. Any other value results in a 400 Bad Request. (optional)
	uploadLength := int64(789) // int64 | The size of the entire upload in bytes. (optional)
	uploadMetadata := "uploadMetadata_example" // string | Additional metadata for the upload request. The header consists of comma-separated key-value pairs. The key MUST NOT contain spaces or commas and MUST be ASCII encoded. The value MUST be Base64 encoded. The workspace key should be provided with the unique identifier for the payload to be uploaded to. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FPGAPayloadsAPI.CreateFPGAPayloadUploadSession(context.Background()).AcceptVersion(acceptVersion).TusResumable(tusResumable).UploadDeferLength(uploadDeferLength).UploadLength(uploadLength).UploadMetadata(uploadMetadata).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FPGAPayloadsAPI.CreateFPGAPayloadUploadSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFPGAPayloadUploadSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 
 **tusResumable** | **string** | Version of the Tus protocol being used. | 
 **uploadDeferLength** | **int32** | Set to 1 if upload size is not known at the time. Any other value results in a 400 Bad Request. | 
 **uploadLength** | **int64** | The size of the entire upload in bytes. | 
 **uploadMetadata** | **string** | Additional metadata for the upload request. The header consists of comma-separated key-value pairs. The key MUST NOT contain spaces or commas and MUST be ASCII encoded. The value MUST be Base64 encoded. The workspace key should be provided with the unique identifier for the payload to be uploaded to. | 

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


## DeleteFpgaPayload

> DeleteFpgaPayload(ctx, fpgaPayloadName, repositoryName).AcceptVersion(acceptVersion).Execute()

Delete an FPGA payload.



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
	fpgaPayloadName := "fpgaPayloadName_example" // string | Unique ID of the FPGA payload.
	repositoryName := "repositoryName_example" // string | Unique ID of a repository.
	acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FPGAPayloadsAPI.DeleteFpgaPayload(context.Background(), fpgaPayloadName, repositoryName).AcceptVersion(acceptVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FPGAPayloadsAPI.DeleteFpgaPayload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fpgaPayloadName** | **string** | Unique ID of the FPGA payload. | 
**repositoryName** | **string** | Unique ID of a repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFpgaPayloadRequest struct via the builder pattern


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


## GetAlternativeFpgaPayloadUploadProgress

> GetAlternativeFpgaPayloadUploadProgress(ctx, uploadSessionName).XHTTPMethodOverride(xHTTPMethodOverride).AcceptVersion(acceptVersion).Execute()

Return FPGA payload upload progress.



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
	uploadSessionName := "uploadSessionName_example" // string | Unique ID of an upload session.
	xHTTPMethodOverride := "HEAD" // string | Verb tunnelling when some methods are not supported.
	acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FPGAPayloadsAPI.GetAlternativeFpgaPayloadUploadProgress(context.Background(), uploadSessionName).XHTTPMethodOverride(xHTTPMethodOverride).AcceptVersion(acceptVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FPGAPayloadsAPI.GetAlternativeFpgaPayloadUploadProgress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uploadSessionName** | **string** | Unique ID of an upload session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlternativeFpgaPayloadUploadProgressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xHTTPMethodOverride** | **string** | Verb tunnelling when some methods are not supported. | 
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


## GetFpgaPayload

> FPGAPayloadItem GetFpgaPayload(ctx, fpgaPayloadName, repositoryName).AcceptVersion(acceptVersion).IfNoneMatch(ifNoneMatch).Execute()

Return details of specific FPGA payload.



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
	fpgaPayloadName := "fpgaPayloadName_example" // string | Unique ID of the FPGA payload.
	repositoryName := "repositoryName_example" // string | Unique ID of a repository.
	acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)
	ifNoneMatch := "ifNoneMatch_example" // string | Caching: Optional header to improve performance. The value of this header should be the `ETag` of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FPGAPayloadsAPI.GetFpgaPayload(context.Background(), fpgaPayloadName, repositoryName).AcceptVersion(acceptVersion).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FPGAPayloadsAPI.GetFpgaPayload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFpgaPayload`: FPGAPayloadItem
	fmt.Fprintf(os.Stdout, "Response from `FPGAPayloadsAPI.GetFpgaPayload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fpgaPayloadName** | **string** | Unique ID of the FPGA payload. | 
**repositoryName** | **string** | Unique ID of a repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFpgaPayloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 
 **ifNoneMatch** | **string** | Caching: Optional header to improve performance. The value of this header should be the &#x60;ETag&#x60; of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content. | 

### Return type

[**FPGAPayloadItem**](FPGAPayloadItem.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFpgaPayloadUploadJobMessages

> NotificationFeed GetFpgaPayloadUploadJobMessages(ctx, uploadJobName).AcceptVersion(acceptVersion).IfNoneMatch(ifNoneMatch).Limit(limit).Offset(offset).Execute()

FPGA payload upload job Message Feed.



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
	uploadJobName := "uploadJobName_example" // string | Unique ID of the FPGA payload upload job.
	acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)
	ifNoneMatch := "ifNoneMatch_example" // string | Caching: Optional header to improve performance. The value of this header should be the `ETag` of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content. (optional)
	limit := int32(20) // int32 | Paging: The maximum number of items to return in a resource. (optional) (default to 20)
	offset := int32(0) // int32 | Paging:  The index of the first item to return in the resource. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FPGAPayloadsAPI.GetFpgaPayloadUploadJobMessages(context.Background(), uploadJobName).AcceptVersion(acceptVersion).IfNoneMatch(ifNoneMatch).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FPGAPayloadsAPI.GetFpgaPayloadUploadJobMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFpgaPayloadUploadJobMessages`: NotificationFeed
	fmt.Fprintf(os.Stdout, "Response from `FPGAPayloadsAPI.GetFpgaPayloadUploadJobMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uploadJobName** | **string** | Unique ID of the FPGA payload upload job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFpgaPayloadUploadJobMessagesRequest struct via the builder pattern


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


## GetFpgaPayloadUploadOptions

> GetFpgaPayloadUploadOptions(ctx).AcceptVersion(acceptVersion).Execute()

Return service TUS protocol support.



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FPGAPayloadsAPI.GetFpgaPayloadUploadOptions(context.Background()).AcceptVersion(acceptVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FPGAPayloadsAPI.GetFpgaPayloadUploadOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFpgaPayloadUploadOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 

### Return type

 (empty response body)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFpgaPayloadUploadProgress

> GetFpgaPayloadUploadProgress(ctx, uploadSessionName).AcceptVersion(acceptVersion).Execute()

Return FPGA payload upload progress.



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
	uploadSessionName := "uploadSessionName_example" // string | Unique ID of an upload session.
	acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FPGAPayloadsAPI.GetFpgaPayloadUploadProgress(context.Background(), uploadSessionName).AcceptVersion(acceptVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FPGAPayloadsAPI.GetFpgaPayloadUploadProgress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uploadSessionName** | **string** | Unique ID of an upload session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFpgaPayloadUploadProgressRequest struct via the builder pattern


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


## ListFPGAPayloads

> FPGAPayloadCollection ListFPGAPayloads(ctx, repositoryName).AcceptVersion(acceptVersion).Embed(embed).IfNoneMatch(ifNoneMatch).Limit(limit).Offset(offset).Execute()

List payloads in a repository.



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
	repositoryName := "repositoryName_example" // string | Unique ID of a repository.
	acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)
	embed := false // bool | Embedding: The whether or not to embed resources into the collection (rather than return links). (optional) (default to false)
	ifNoneMatch := "ifNoneMatch_example" // string | Caching: Optional header to improve performance. The value of this header should be the `ETag` of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content. (optional)
	limit := int32(20) // int32 | Paging: The maximum number of items to return in a resource. (optional) (default to 20)
	offset := int32(0) // int32 | Paging:  The index of the first item to return in the resource. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FPGAPayloadsAPI.ListFPGAPayloads(context.Background(), repositoryName).AcceptVersion(acceptVersion).Embed(embed).IfNoneMatch(ifNoneMatch).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FPGAPayloadsAPI.ListFPGAPayloads``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFPGAPayloads`: FPGAPayloadCollection
	fmt.Fprintf(os.Stdout, "Response from `FPGAPayloadsAPI.ListFPGAPayloads`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Unique ID of a repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListFPGAPayloadsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 
 **embed** | **bool** | Embedding: The whether or not to embed resources into the collection (rather than return links). | [default to false]
 **ifNoneMatch** | **string** | Caching: Optional header to improve performance. The value of this header should be the &#x60;ETag&#x60; of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content. | 
 **limit** | **int32** | Paging: The maximum number of items to return in a resource. | [default to 20]
 **offset** | **int32** | Paging:  The index of the first item to return in the resource. | [default to 0]

### Return type

[**FPGAPayloadCollection**](FPGAPayloadCollection.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPayloads

> FPGAPayloadCollection ListPayloads(ctx).AcceptVersion(acceptVersion).Embed(embed).IfNoneMatch(ifNoneMatch).Limit(limit).Offset(offset).Execute()

List payloads.



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
	resp, r, err := apiClient.FPGAPayloadsAPI.ListPayloads(context.Background()).AcceptVersion(acceptVersion).Embed(embed).IfNoneMatch(ifNoneMatch).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FPGAPayloadsAPI.ListPayloads``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPayloads`: FPGAPayloadCollection
	fmt.Fprintf(os.Stdout, "Response from `FPGAPayloadsAPI.ListPayloads`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPayloadsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 
 **embed** | **bool** | Embedding: The whether or not to embed resources into the collection (rather than return links). | [default to false]
 **ifNoneMatch** | **string** | Caching: Optional header to improve performance. The value of this header should be the &#x60;ETag&#x60; of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content. | 
 **limit** | **int32** | Paging: The maximum number of items to return in a resource. | [default to 20]
 **offset** | **int32** | Paging:  The index of the first item to return in the resource. | [default to 0]

### Return type

[**FPGAPayloadCollection**](FPGAPayloadCollection.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadPayload

> UploadPayload(ctx, uploadSessionName).TusResumable(tusResumable).AcceptVersion(acceptVersion).UploadChecksum(uploadChecksum).UploadOffset(uploadOffset).Execute()

Upload part of a payload.



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
	uploadSessionName := "uploadSessionName_example" // string | Unique ID of an upload session.
	tusResumable := "1.0.0" // string | Version of the Tus protocol being used.
	acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)
	uploadChecksum := "uploadChecksum_example" // string | Information about the checksum of the current body payload. (optional)
	uploadOffset := int64(789) // int64 | The byte offset within a resource. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FPGAPayloadsAPI.UploadPayload(context.Background(), uploadSessionName).TusResumable(tusResumable).AcceptVersion(acceptVersion).UploadChecksum(uploadChecksum).UploadOffset(uploadOffset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FPGAPayloadsAPI.UploadPayload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uploadSessionName** | **string** | Unique ID of an upload session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadPayloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tusResumable** | **string** | Version of the Tus protocol being used. | 
 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 
 **uploadChecksum** | **string** | Information about the checksum of the current body payload. | 
 **uploadOffset** | **int64** | The byte offset within a resource. | 

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

