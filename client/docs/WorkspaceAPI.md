<!--
Copyright (C) 2020-2025 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# \WorkspaceAPI

All URIs are relative to *https://all.api.keil.arm.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClearWorkspaceArchiveContent**](WorkspaceAPI.md#ClearWorkspaceArchiveContent) | **Delete** /workspaces/{workspaceName}/archive-content | Clear the content of this workspace.
[**ClearWorkspaceIncrementalChunk**](WorkspaceAPI.md#ClearWorkspaceIncrementalChunk) | **Delete** /workspaces/{workspaceName}/chunk | Clear the content of this chunked workspace.
[**ClearWorkspaceRepositoryContentManager**](WorkspaceAPI.md#ClearWorkspaceRepositoryContentManager) | **Delete** /workspaces/{workspaceName}/repository-content | Clear the content of this workspace.
[**CreateCompositeWorkspaceChunk**](WorkspaceAPI.md#CreateCompositeWorkspaceChunk) | **Post** /workspaces/{workspaceName} | Create workspace chunks of a composite workspace
[**CreateWorkspace**](WorkspaceAPI.md#CreateWorkspace) | **Post** /workspace-sources/{workspaceSourceName} | Creates a workspace based on the source.
[**DeleteWorkspace**](WorkspaceAPI.md#DeleteWorkspace) | **Delete** /workspaces/{workspaceName} | Delete a Workspace
[**EditWorkspaceRepositoryContentManager**](WorkspaceAPI.md#EditWorkspaceRepositoryContentManager) | **Put** /workspaces/{workspaceName}/repository-content | Edit the source for the content of the workspace.
[**GetChunkedWorkspaceUploadProgress**](WorkspaceAPI.md#GetChunkedWorkspaceUploadProgress) | **Head** /workspaces/{workspaceName}/chunk | Return workspace content upload progress.
[**GetCompositeWorkspaceUploadOptions**](WorkspaceAPI.md#GetCompositeWorkspaceUploadOptions) | **Options** /workspaces/{workspaceName}/details | Return workspace TUS protocol support.
[**GetCompositeWorkspaceUploadSupport**](WorkspaceAPI.md#GetCompositeWorkspaceUploadSupport) | **Options** /workspaces/{workspaceName} | Return workspace TUS protocol support.
[**GetWorkspace**](WorkspaceAPI.md#GetWorkspace) | **Get** /workspaces/{workspaceName} | Return the state of a workspace.
[**GetWorkspaceArchiveContent**](WorkspaceAPI.md#GetWorkspaceArchiveContent) | **Get** /workspaces/{workspaceName}/archive-content | Get the manager of the archive file containing the workspace content.
[**GetWorkspaceArchiveContentUploadProgress**](WorkspaceAPI.md#GetWorkspaceArchiveContentUploadProgress) | **Head** /workspaces/{workspaceName}/archive-content | Return progress about the upload.
[**GetWorkspaceDetails**](WorkspaceAPI.md#GetWorkspaceDetails) | **Get** /workspaces/{workspaceName}/details | Details about the workspace.
[**GetWorkspaceIncrementalChunk**](WorkspaceAPI.md#GetWorkspaceIncrementalChunk) | **Get** /workspaces/{workspaceName}/chunk | Get the manager of the file containing the chunked workspace content.
[**GetWorkspaceRepositoryContentManager**](WorkspaceAPI.md#GetWorkspaceRepositoryContentManager) | **Get** /workspaces/{workspaceName}/repository-content | Get the manager of the workspace content defined in a repository.
[**ListWorkspaces**](WorkspaceAPI.md#ListWorkspaces) | **Get** /workspaces/ | List all workspaces available.
[**RetainWorkspace**](WorkspaceAPI.md#RetainWorkspace) | **Post** /workspaces/{workspaceName}/retain | Update how long a workspace will be retained before automatic deletion..
[**UploadWorkspaceArchiveContent**](WorkspaceAPI.md#UploadWorkspaceArchiveContent) | **Put** /workspaces/{workspaceName}/archive-content | Upload and replace the content of the named workspace.
[**UploadWorkspaceIncrementalChunk**](WorkspaceAPI.md#UploadWorkspaceIncrementalChunk) | **Patch** /workspaces/{workspaceName}/chunk | Upload part of a workspace.



## ClearWorkspaceArchiveContent

> ClearWorkspaceArchiveContent(ctx, workspaceName).AcceptVersion(acceptVersion).TusResumable(tusResumable).Execute()

Clear the content of this workspace.



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
	workspaceName := "workspaceName_example" // string | Unique ID of the Workspace.
	acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)
	tusResumable := "1.0.0" // string | Version of the Tus protocol being used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkspaceAPI.ClearWorkspaceArchiveContent(context.Background(), workspaceName).AcceptVersion(acceptVersion).TusResumable(tusResumable).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceAPI.ClearWorkspaceArchiveContent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceName** | **string** | Unique ID of the Workspace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiClearWorkspaceArchiveContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 
 **tusResumable** | **string** | Version of the Tus protocol being used. | 

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


## ClearWorkspaceIncrementalChunk

> ClearWorkspaceIncrementalChunk(ctx, workspaceName).TusResumable(tusResumable).AcceptVersion(acceptVersion).Execute()

Clear the content of this chunked workspace.



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
	workspaceName := "workspaceName_example" // string | Unique ID of the Workspace.
	tusResumable := "1.0.0" // string | Version of the Tus protocol being used.
	acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkspaceAPI.ClearWorkspaceIncrementalChunk(context.Background(), workspaceName).TusResumable(tusResumable).AcceptVersion(acceptVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceAPI.ClearWorkspaceIncrementalChunk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceName** | **string** | Unique ID of the Workspace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiClearWorkspaceIncrementalChunkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tusResumable** | **string** | Version of the Tus protocol being used. | 
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


## ClearWorkspaceRepositoryContentManager

> ClearWorkspaceRepositoryContentManager(ctx, workspaceName).AcceptVersion(acceptVersion).Execute()

Clear the content of this workspace.



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
	workspaceName := "workspaceName_example" // string | Unique ID of the Workspace.
	acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkspaceAPI.ClearWorkspaceRepositoryContentManager(context.Background(), workspaceName).AcceptVersion(acceptVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceAPI.ClearWorkspaceRepositoryContentManager``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceName** | **string** | Unique ID of the Workspace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiClearWorkspaceRepositoryContentManagerRequest struct via the builder pattern


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


## CreateCompositeWorkspaceChunk

> CreateCompositeWorkspaceChunk(ctx, workspaceName).TusResumable(tusResumable).UploadConcat(uploadConcat).AcceptVersion(acceptVersion).UploadDeferLength(uploadDeferLength).UploadLength(uploadLength).UploadMetadata(uploadMetadata).Execute()

Create workspace chunks of a composite workspace



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
	workspaceName := "workspaceName_example" // string | Unique ID of the Workspace.
	tusResumable := "1.0.0" // string | Version of the Tus protocol being used.
	uploadConcat := "uploadConcat_example" // string | whether it is a partial upload or a final upload.  https://tus.io/protocols/resumable-upload#upload-concat
	acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)
	uploadDeferLength := int32(56) // int32 | Set to 1 if upload size is not known at the time. Any other value results in a 400 Bad Request. (optional)
	uploadLength := int64(789) // int64 | The size of the entire upload in bytes. (optional)
	uploadMetadata := "uploadMetadata_example" // string | Additional metadata for the upload request. The header consists of comma-separated key-value pairs. The key MUST NOT contain spaces or commas and MUST be ASCII encoded. The value MUST be Base64 encoded. The workspace key should be provided with the unique identifier for the payload to be uploaded to. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkspaceAPI.CreateCompositeWorkspaceChunk(context.Background(), workspaceName).TusResumable(tusResumable).UploadConcat(uploadConcat).AcceptVersion(acceptVersion).UploadDeferLength(uploadDeferLength).UploadLength(uploadLength).UploadMetadata(uploadMetadata).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceAPI.CreateCompositeWorkspaceChunk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceName** | **string** | Unique ID of the Workspace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCompositeWorkspaceChunkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tusResumable** | **string** | Version of the Tus protocol being used. | 
 **uploadConcat** | **string** | whether it is a partial upload or a final upload.  https://tus.io/protocols/resumable-upload#upload-concat | 
 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 
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


## CreateWorkspace

> WorkspaceItem CreateWorkspace(ctx, workspaceSourceName).WorkspaceItem(workspaceItem).AcceptVersion(acceptVersion).TusResumable(tusResumable).UploadDeferLength(uploadDeferLength).UploadLength(uploadLength).UploadMetadata(uploadMetadata).Execute()

Creates a workspace based on the source.



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
	workspaceSourceName := "workspaceSourceName_example" // string | The ID of the workspace source.
	workspaceItem := *openapiclient.NewWorkspaceItem("TODO", "TODO", "332129b3-f14d-49d2-b9be-acd2abd80c6b") // WorkspaceItem | A name of the source type to create the workspace from.
	acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)
	tusResumable := "1.0.0" // string | Version of the Tus protocol being used. (optional)
	uploadDeferLength := int32(56) // int32 | Set to 1 if upload size is not known at the time. Any other value results in a 400 Bad Request. (optional)
	uploadLength := int64(789) // int64 | The size of the entire upload in bytes. (optional)
	uploadMetadata := "uploadMetadata_example" // string | Additional metadata for the upload request. The header consists of comma-separated key-value pairs. The key MUST NOT contain spaces or commas and MUST be ASCII encoded. The value MUST be Base64 encoded. The workspace key should be provided with the unique identifier for the payload to be uploaded to. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspaceAPI.CreateWorkspace(context.Background(), workspaceSourceName).WorkspaceItem(workspaceItem).AcceptVersion(acceptVersion).TusResumable(tusResumable).UploadDeferLength(uploadDeferLength).UploadLength(uploadLength).UploadMetadata(uploadMetadata).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceAPI.CreateWorkspace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWorkspace`: WorkspaceItem
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceAPI.CreateWorkspace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceSourceName** | **string** | The ID of the workspace source. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workspaceItem** | [**WorkspaceItem**](WorkspaceItem.md) | A name of the source type to create the workspace from. | 
 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 
 **tusResumable** | **string** | Version of the Tus protocol being used. | 
 **uploadDeferLength** | **int32** | Set to 1 if upload size is not known at the time. Any other value results in a 400 Bad Request. | 
 **uploadLength** | **int64** | The size of the entire upload in bytes. | 
 **uploadMetadata** | **string** | Additional metadata for the upload request. The header consists of comma-separated key-value pairs. The key MUST NOT contain spaces or commas and MUST be ASCII encoded. The value MUST be Base64 encoded. The workspace key should be provided with the unique identifier for the payload to be uploaded to. | 

### Return type

[**WorkspaceItem**](WorkspaceItem.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWorkspace

> DeleteWorkspace(ctx, workspaceName).AcceptVersion(acceptVersion).TusResumable(tusResumable).Execute()

Delete a Workspace



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
	workspaceName := "workspaceName_example" // string | Unique ID of the Workspace.
	acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)
	tusResumable := "1.0.0" // string | Version of the Tus protocol being used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkspaceAPI.DeleteWorkspace(context.Background(), workspaceName).AcceptVersion(acceptVersion).TusResumable(tusResumable).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceAPI.DeleteWorkspace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceName** | **string** | Unique ID of the Workspace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 
 **tusResumable** | **string** | Version of the Tus protocol being used. | 

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


## EditWorkspaceRepositoryContentManager

> WorkspaceRepositoryContentManager EditWorkspaceRepositoryContentManager(ctx, workspaceName).IfMatch(ifMatch).WorkspaceRepositoryContentManager(workspaceRepositoryContentManager).AcceptVersion(acceptVersion).Execute()

Edit the source for the content of the workspace.



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
	workspaceName := "workspaceName_example" // string | Unique ID of the Workspace.
	ifMatch := "ifMatch_example" // string | Conditional Requests: This is required in order to perform an update of a resource. The value of this header should be the `ETag` of the resource when read (before being subsequently modified by the client).
	workspaceRepositoryContentManager := *openapiclient.NewWorkspaceRepositoryContentManager("TODO", "TODO", "example001", "23db6982caef9e9152f1a5b2589e6ca3", " https://github.com/cli/cli", "git") // WorkspaceRepositoryContentManager | Definition of the repository where the workspace content should be checked out from
	acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspaceAPI.EditWorkspaceRepositoryContentManager(context.Background(), workspaceName).IfMatch(ifMatch).WorkspaceRepositoryContentManager(workspaceRepositoryContentManager).AcceptVersion(acceptVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceAPI.EditWorkspaceRepositoryContentManager``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditWorkspaceRepositoryContentManager`: WorkspaceRepositoryContentManager
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceAPI.EditWorkspaceRepositoryContentManager`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceName** | **string** | Unique ID of the Workspace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditWorkspaceRepositoryContentManagerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **string** | Conditional Requests: This is required in order to perform an update of a resource. The value of this header should be the &#x60;ETag&#x60; of the resource when read (before being subsequently modified by the client). | 
 **workspaceRepositoryContentManager** | [**WorkspaceRepositoryContentManager**](WorkspaceRepositoryContentManager.md) | Definition of the repository where the workspace content should be checked out from | 
 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 

### Return type

[**WorkspaceRepositoryContentManager**](WorkspaceRepositoryContentManager.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChunkedWorkspaceUploadProgress

> GetChunkedWorkspaceUploadProgress(ctx, workspaceName).AcceptVersion(acceptVersion).TusResumable(tusResumable).Execute()

Return workspace content upload progress.



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
	workspaceName := "workspaceName_example" // string | Unique ID of the Workspace.
	acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)
	tusResumable := "1.0.0" // string | Version of the Tus protocol being used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkspaceAPI.GetChunkedWorkspaceUploadProgress(context.Background(), workspaceName).AcceptVersion(acceptVersion).TusResumable(tusResumable).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceAPI.GetChunkedWorkspaceUploadProgress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceName** | **string** | Unique ID of the Workspace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChunkedWorkspaceUploadProgressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 
 **tusResumable** | **string** | Version of the Tus protocol being used. | 

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


## GetCompositeWorkspaceUploadOptions

> GetCompositeWorkspaceUploadOptions(ctx, workspaceName).AcceptVersion(acceptVersion).Execute()

Return workspace TUS protocol support.



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
	workspaceName := "workspaceName_example" // string | Unique ID of the Workspace.
	acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkspaceAPI.GetCompositeWorkspaceUploadOptions(context.Background(), workspaceName).AcceptVersion(acceptVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceAPI.GetCompositeWorkspaceUploadOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceName** | **string** | Unique ID of the Workspace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompositeWorkspaceUploadOptionsRequest struct via the builder pattern


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


## GetCompositeWorkspaceUploadSupport

> GetCompositeWorkspaceUploadSupport(ctx, workspaceName).AcceptVersion(acceptVersion).Execute()

Return workspace TUS protocol support.



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
	workspaceName := "workspaceName_example" // string | Unique ID of the Workspace.
	acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkspaceAPI.GetCompositeWorkspaceUploadSupport(context.Background(), workspaceName).AcceptVersion(acceptVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceAPI.GetCompositeWorkspaceUploadSupport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceName** | **string** | Unique ID of the Workspace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompositeWorkspaceUploadSupportRequest struct via the builder pattern


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


## GetWorkspace

> WorkspaceItem GetWorkspace(ctx, workspaceName).AcceptVersion(acceptVersion).TusResumable(tusResumable).IfNoneMatch(ifNoneMatch).Execute()

Return the state of a workspace.



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
	workspaceName := "workspaceName_example" // string | Unique ID of the Workspace.
	acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)
	tusResumable := "1.0.0" // string | Version of the Tus protocol being used. (optional)
	ifNoneMatch := "ifNoneMatch_example" // string | Caching: Optional header to improve performance. The value of this header should be the `ETag` of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspaceAPI.GetWorkspace(context.Background(), workspaceName).AcceptVersion(acceptVersion).TusResumable(tusResumable).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceAPI.GetWorkspace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkspace`: WorkspaceItem
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceAPI.GetWorkspace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceName** | **string** | Unique ID of the Workspace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 
 **tusResumable** | **string** | Version of the Tus protocol being used. | 
 **ifNoneMatch** | **string** | Caching: Optional header to improve performance. The value of this header should be the &#x60;ETag&#x60; of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content. | 

### Return type

[**WorkspaceItem**](WorkspaceItem.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkspaceArchiveContent

> ArtefactManagerItem GetWorkspaceArchiveContent(ctx, workspaceName).AcceptVersion(acceptVersion).TusResumable(tusResumable).IfNoneMatch(ifNoneMatch).Execute()

Get the manager of the archive file containing the workspace content.



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
	workspaceName := "workspaceName_example" // string | Unique ID of the Workspace.
	acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)
	tusResumable := "1.0.0" // string | Version of the Tus protocol being used. (optional)
	ifNoneMatch := "ifNoneMatch_example" // string | Caching: Optional header to improve performance. The value of this header should be the `ETag` of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspaceAPI.GetWorkspaceArchiveContent(context.Background(), workspaceName).AcceptVersion(acceptVersion).TusResumable(tusResumable).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceAPI.GetWorkspaceArchiveContent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkspaceArchiveContent`: ArtefactManagerItem
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceAPI.GetWorkspaceArchiveContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceName** | **string** | Unique ID of the Workspace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkspaceArchiveContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 
 **tusResumable** | **string** | Version of the Tus protocol being used. | 
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


## GetWorkspaceArchiveContentUploadProgress

> GetWorkspaceArchiveContentUploadProgress(ctx, workspaceName).AcceptVersion(acceptVersion).TusResumable(tusResumable).Execute()

Return progress about the upload.



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
	workspaceName := "workspaceName_example" // string | Unique ID of the Workspace.
	acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)
	tusResumable := "1.0.0" // string | Version of the Tus protocol being used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkspaceAPI.GetWorkspaceArchiveContentUploadProgress(context.Background(), workspaceName).AcceptVersion(acceptVersion).TusResumable(tusResumable).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceAPI.GetWorkspaceArchiveContentUploadProgress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceName** | **string** | Unique ID of the Workspace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkspaceArchiveContentUploadProgressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 
 **tusResumable** | **string** | Version of the Tus protocol being used. | 

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


## GetWorkspaceDetails

> WorkspaceDetailsItem GetWorkspaceDetails(ctx, workspaceName).AcceptVersion(acceptVersion).TusResumable(tusResumable).IfNoneMatch(ifNoneMatch).Execute()

Details about the workspace.



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
	workspaceName := "workspaceName_example" // string | Unique ID of the Workspace.
	acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)
	tusResumable := "1.0.0" // string | Version of the Tus protocol being used. (optional)
	ifNoneMatch := "ifNoneMatch_example" // string | Caching: Optional header to improve performance. The value of this header should be the `ETag` of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspaceAPI.GetWorkspaceDetails(context.Background(), workspaceName).AcceptVersion(acceptVersion).TusResumable(tusResumable).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceAPI.GetWorkspaceDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkspaceDetails`: WorkspaceDetailsItem
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceAPI.GetWorkspaceDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceName** | **string** | Unique ID of the Workspace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkspaceDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 
 **tusResumable** | **string** | Version of the Tus protocol being used. | 
 **ifNoneMatch** | **string** | Caching: Optional header to improve performance. The value of this header should be the &#x60;ETag&#x60; of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content. | 

### Return type

[**WorkspaceDetailsItem**](WorkspaceDetailsItem.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkspaceIncrementalChunk

> ArtefactManagerItem GetWorkspaceIncrementalChunk(ctx, workspaceName).AcceptVersion(acceptVersion).TusResumable(tusResumable).IfNoneMatch(ifNoneMatch).Execute()

Get the manager of the file containing the chunked workspace content.



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
	workspaceName := "workspaceName_example" // string | Unique ID of the Workspace.
	acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)
	tusResumable := "1.0.0" // string | Version of the Tus protocol being used. (optional)
	ifNoneMatch := "ifNoneMatch_example" // string | Caching: Optional header to improve performance. The value of this header should be the `ETag` of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspaceAPI.GetWorkspaceIncrementalChunk(context.Background(), workspaceName).AcceptVersion(acceptVersion).TusResumable(tusResumable).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceAPI.GetWorkspaceIncrementalChunk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkspaceIncrementalChunk`: ArtefactManagerItem
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceAPI.GetWorkspaceIncrementalChunk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceName** | **string** | Unique ID of the Workspace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkspaceIncrementalChunkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 
 **tusResumable** | **string** | Version of the Tus protocol being used. | 
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


## GetWorkspaceRepositoryContentManager

> WorkspaceRepositoryContentManager GetWorkspaceRepositoryContentManager(ctx, workspaceName).AcceptVersion(acceptVersion).IfNoneMatch(ifNoneMatch).Execute()

Get the manager of the workspace content defined in a repository.



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
	workspaceName := "workspaceName_example" // string | Unique ID of the Workspace.
	acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)
	ifNoneMatch := "ifNoneMatch_example" // string | Caching: Optional header to improve performance. The value of this header should be the `ETag` of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspaceAPI.GetWorkspaceRepositoryContentManager(context.Background(), workspaceName).AcceptVersion(acceptVersion).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceAPI.GetWorkspaceRepositoryContentManager``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkspaceRepositoryContentManager`: WorkspaceRepositoryContentManager
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceAPI.GetWorkspaceRepositoryContentManager`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceName** | **string** | Unique ID of the Workspace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkspaceRepositoryContentManagerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 
 **ifNoneMatch** | **string** | Caching: Optional header to improve performance. The value of this header should be the &#x60;ETag&#x60; of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content. | 

### Return type

[**WorkspaceRepositoryContentManager**](WorkspaceRepositoryContentManager.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkspaces

> WorkspaceCollection ListWorkspaces(ctx).AcceptVersion(acceptVersion).Embed(embed).IfNoneMatch(ifNoneMatch).Limit(limit).Offset(offset).Execute()

List all workspaces available.



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
	resp, r, err := apiClient.WorkspaceAPI.ListWorkspaces(context.Background()).AcceptVersion(acceptVersion).Embed(embed).IfNoneMatch(ifNoneMatch).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceAPI.ListWorkspaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWorkspaces`: WorkspaceCollection
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceAPI.ListWorkspaces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWorkspacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 
 **embed** | **bool** | Embedding: The whether or not to embed resources into the collection (rather than return links). | [default to false]
 **ifNoneMatch** | **string** | Caching: Optional header to improve performance. The value of this header should be the &#x60;ETag&#x60; of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content. | 
 **limit** | **int32** | Paging: The maximum number of items to return in a resource. | [default to 20]
 **offset** | **int32** | Paging:  The index of the first item to return in the resource. | [default to 0]

### Return type

[**WorkspaceCollection**](WorkspaceCollection.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetainWorkspace

> WorkspaceItem RetainWorkspace(ctx, workspaceName).AcceptVersion(acceptVersion).RetainWorkspaceRequest(retainWorkspaceRequest).Execute()

Update how long a workspace will be retained before automatic deletion..



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
	workspaceName := "workspaceName_example" // string | Unique ID of the Workspace.
	acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)
	retainWorkspaceRequest := *openapiclient.NewRetainWorkspaceRequest() // RetainWorkspaceRequest | TTL configuration. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspaceAPI.RetainWorkspace(context.Background(), workspaceName).AcceptVersion(acceptVersion).RetainWorkspaceRequest(retainWorkspaceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceAPI.RetainWorkspace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetainWorkspace`: WorkspaceItem
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceAPI.RetainWorkspace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceName** | **string** | Unique ID of the Workspace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetainWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 
 **retainWorkspaceRequest** | [**RetainWorkspaceRequest**](RetainWorkspaceRequest.md) | TTL configuration. | 

### Return type

[**WorkspaceItem**](WorkspaceItem.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadWorkspaceArchiveContent

> ArtefactManagerItem UploadWorkspaceArchiveContent(ctx, workspaceName).IfMatch(ifMatch).Content(content).Hash(hash).AcceptVersion(acceptVersion).TusResumable(tusResumable).ContentMediaType(contentMediaType).Size(size).Title(title).Execute()

Upload and replace the content of the named workspace.



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
	workspaceName := "workspaceName_example" // string | Unique ID of the Workspace.
	ifMatch := "ifMatch_example" // string | Conditional Requests: This is required in order to perform an update of a resource. The value of this header should be the `ETag` of the resource when read (before being subsequently modified by the client).
	content := os.NewFile(1234, "some_file") // *os.File | artefact content
	hash := "hash_example" // string | Hash of the artefact (sha256) for network resilience
	acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)
	tusResumable := "1.0.0" // string | Version of the Tus protocol being used. (optional)
	contentMediaType := "contentMediaType_example" // string | Media type of the artefact according to https://www.iana.org/assignments/media-types/media-types.xhtml Technically redundant, but can be used by JSON Schema tools that may not be aware of the OpenAPI context. (optional)
	size := int64(789) // int64 | size in bytes of this artefact. Technically redundant, but can be used by JSON Schema tools that may not be aware of the OpenAPI context. (optional)
	title := "title_example" // string | Optional human readable name of the artefact. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspaceAPI.UploadWorkspaceArchiveContent(context.Background(), workspaceName).IfMatch(ifMatch).Content(content).Hash(hash).AcceptVersion(acceptVersion).TusResumable(tusResumable).ContentMediaType(contentMediaType).Size(size).Title(title).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceAPI.UploadWorkspaceArchiveContent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadWorkspaceArchiveContent`: ArtefactManagerItem
	fmt.Fprintf(os.Stdout, "Response from `WorkspaceAPI.UploadWorkspaceArchiveContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceName** | **string** | Unique ID of the Workspace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadWorkspaceArchiveContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **string** | Conditional Requests: This is required in order to perform an update of a resource. The value of this header should be the &#x60;ETag&#x60; of the resource when read (before being subsequently modified by the client). | 
 **content** | ***os.File** | artefact content | 
 **hash** | **string** | Hash of the artefact (sha256) for network resilience | 
 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 
 **tusResumable** | **string** | Version of the Tus protocol being used. | 
 **contentMediaType** | **string** | Media type of the artefact according to https://www.iana.org/assignments/media-types/media-types.xhtml Technically redundant, but can be used by JSON Schema tools that may not be aware of the OpenAPI context. | 
 **size** | **int64** | size in bytes of this artefact. Technically redundant, but can be used by JSON Schema tools that may not be aware of the OpenAPI context. | 
 **title** | **string** | Optional human readable name of the artefact. | 

### Return type

[**ArtefactManagerItem**](ArtefactManagerItem.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadWorkspaceIncrementalChunk

> UploadWorkspaceIncrementalChunk(ctx, workspaceName).TusResumable(tusResumable).AcceptVersion(acceptVersion).UploadChecksum(uploadChecksum).UploadOffset(uploadOffset).Execute()

Upload part of a workspace.



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
	workspaceName := "workspaceName_example" // string | Unique ID of the Workspace.
	tusResumable := "1.0.0" // string | Version of the Tus protocol being used.
	acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)
	uploadChecksum := "uploadChecksum_example" // string | Information about the checksum of the current body payload. (optional)
	uploadOffset := int64(789) // int64 | The byte offset within a resource. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkspaceAPI.UploadWorkspaceIncrementalChunk(context.Background(), workspaceName).TusResumable(tusResumable).AcceptVersion(acceptVersion).UploadChecksum(uploadChecksum).UploadOffset(uploadOffset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceAPI.UploadWorkspaceIncrementalChunk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceName** | **string** | Unique ID of the Workspace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadWorkspaceIncrementalChunkRequest struct via the builder pattern


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

