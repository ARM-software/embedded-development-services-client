<!--
Copyright (C) 2020-2023 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# \WorkspaceApi

All URIs are relative to *https://all.api.keil.arm.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClearWorkspaceArchiveContent**](WorkspaceApi.md#ClearWorkspaceArchiveContent) | **Delete** /workspaces/{workspaceName}/archive-content | Clear the content of this workspace.
[**ClearWorkspaceRepositoryContentManager**](WorkspaceApi.md#ClearWorkspaceRepositoryContentManager) | **Delete** /workspaces/{workspaceName}/repository-content | Clear the content of this workspace.
[**CreateWorkspace**](WorkspaceApi.md#CreateWorkspace) | **Post** /workspace-sources/{workspaceSourceName} | Creates a workspace based on the source.
[**DeleteWorkspace**](WorkspaceApi.md#DeleteWorkspace) | **Delete** /workspaces/{workspaceName} | Delete a Workspace
[**EditWorkspaceRepositoryContentManager**](WorkspaceApi.md#EditWorkspaceRepositoryContentManager) | **Put** /workspaces/{workspaceName}/repository-content | Edit the source for the content of the workspace.
[**GetWorkspace**](WorkspaceApi.md#GetWorkspace) | **Get** /workspaces/{workspaceName} | Return the state of a workspace.
[**GetWorkspaceArchiveContent**](WorkspaceApi.md#GetWorkspaceArchiveContent) | **Get** /workspaces/{workspaceName}/archive-content | Get the manager of the archive file containing the workspace content.
[**GetWorkspaceDetails**](WorkspaceApi.md#GetWorkspaceDetails) | **Get** /workspaces/{workspaceName}/details | Details about the workspace.
[**GetWorkspaceRepositoryContentManager**](WorkspaceApi.md#GetWorkspaceRepositoryContentManager) | **Get** /workspaces/{workspaceName}/repository-content | Get the manager of the workspace content defined in a repository.
[**ListWorkspaces**](WorkspaceApi.md#ListWorkspaces) | **Get** /workspaces/ | List all workspaces available.
[**RetainWorkspace**](WorkspaceApi.md#RetainWorkspace) | **Post** /workspaces/{workspaceName}/retain | Update how long a workspace will be retained before automatic deletion..
[**UploadWorkspaceArchiveContent**](WorkspaceApi.md#UploadWorkspaceArchiveContent) | **Put** /workspaces/{workspaceName}/archive-content | Upload and replace the content of the named workspace.



## ClearWorkspaceArchiveContent

> ClearWorkspaceArchiveContent(ctx, workspaceName).AcceptVersion(acceptVersion).Execute()

Clear the content of this workspace.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    workspaceName := "workspaceName_example" // string | Unique ID of the Workspace.
    acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkspaceApi.ClearWorkspaceArchiveContent(context.Background(), workspaceName).AcceptVersion(acceptVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceApi.ClearWorkspaceArchiveContent``: %v\n", err)
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
    openapiclient "./openapi"
)

func main() {
    workspaceName := "workspaceName_example" // string | Unique ID of the Workspace.
    acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkspaceApi.ClearWorkspaceRepositoryContentManager(context.Background(), workspaceName).AcceptVersion(acceptVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceApi.ClearWorkspaceRepositoryContentManager``: %v\n", err)
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


## CreateWorkspace

> WorkspaceItem CreateWorkspace(ctx, workspaceSourceName).WorkspaceItem(workspaceItem).AcceptVersion(acceptVersion).Execute()

Creates a workspace based on the source.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    workspaceSourceName := "workspaceSourceName_example" // string | The ID of the workspace source.
    workspaceItem := *openapiclient.NewWorkspaceItem("TODO", "TODO", "332129b3-f14d-49d2-b9be-acd2abd80c6b") // WorkspaceItem | A name of the source type to create the workspace from.
    acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkspaceApi.CreateWorkspace(context.Background(), workspaceSourceName).WorkspaceItem(workspaceItem).AcceptVersion(acceptVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceApi.CreateWorkspace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorkspace`: WorkspaceItem
    fmt.Fprintf(os.Stdout, "Response from `WorkspaceApi.CreateWorkspace`: %v\n", resp)
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

> DeleteWorkspace(ctx, workspaceName).AcceptVersion(acceptVersion).Execute()

Delete a Workspace



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    workspaceName := "workspaceName_example" // string | Unique ID of the Workspace.
    acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkspaceApi.DeleteWorkspace(context.Background(), workspaceName).AcceptVersion(acceptVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceApi.DeleteWorkspace``: %v\n", err)
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
    openapiclient "./openapi"
)

func main() {
    workspaceName := "workspaceName_example" // string | Unique ID of the Workspace.
    ifMatch := "ifMatch_example" // string | Conditional Requests: This is required in order to perform an update of a resource. The value of this header should be the `ETag` of the resource when read (before being subsequently modified by the client).
    workspaceRepositoryContentManager := *openapiclient.NewWorkspaceRepositoryContentManager("TODO", "TODO", "example001", "23db6982caef9e9152f1a5b2589e6ca3", " https://github.com/cli/cli", "git") // WorkspaceRepositoryContentManager | Definition of the repository where the workspace content should be checked out from
    acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkspaceApi.EditWorkspaceRepositoryContentManager(context.Background(), workspaceName).IfMatch(ifMatch).WorkspaceRepositoryContentManager(workspaceRepositoryContentManager).AcceptVersion(acceptVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceApi.EditWorkspaceRepositoryContentManager``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditWorkspaceRepositoryContentManager`: WorkspaceRepositoryContentManager
    fmt.Fprintf(os.Stdout, "Response from `WorkspaceApi.EditWorkspaceRepositoryContentManager`: %v\n", resp)
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


## GetWorkspace

> WorkspaceItem GetWorkspace(ctx, workspaceName).AcceptVersion(acceptVersion).IfNoneMatch(ifNoneMatch).Execute()

Return the state of a workspace.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    workspaceName := "workspaceName_example" // string | Unique ID of the Workspace.
    acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)
    ifNoneMatch := "ifNoneMatch_example" // string | Caching: Optional header to improve performance. The value of this header should be the `ETag` of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkspaceApi.GetWorkspace(context.Background(), workspaceName).AcceptVersion(acceptVersion).IfNoneMatch(ifNoneMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceApi.GetWorkspace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkspace`: WorkspaceItem
    fmt.Fprintf(os.Stdout, "Response from `WorkspaceApi.GetWorkspace`: %v\n", resp)
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

> ArtefactManagerItem GetWorkspaceArchiveContent(ctx, workspaceName).AcceptVersion(acceptVersion).IfNoneMatch(ifNoneMatch).Execute()

Get the manager of the archive file containing the workspace content.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    workspaceName := "workspaceName_example" // string | Unique ID of the Workspace.
    acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)
    ifNoneMatch := "ifNoneMatch_example" // string | Caching: Optional header to improve performance. The value of this header should be the `ETag` of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkspaceApi.GetWorkspaceArchiveContent(context.Background(), workspaceName).AcceptVersion(acceptVersion).IfNoneMatch(ifNoneMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceApi.GetWorkspaceArchiveContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkspaceArchiveContent`: ArtefactManagerItem
    fmt.Fprintf(os.Stdout, "Response from `WorkspaceApi.GetWorkspaceArchiveContent`: %v\n", resp)
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


## GetWorkspaceDetails

> WorkspaceDetailsItem GetWorkspaceDetails(ctx, workspaceName).AcceptVersion(acceptVersion).IfNoneMatch(ifNoneMatch).Execute()

Details about the workspace.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    workspaceName := "workspaceName_example" // string | Unique ID of the Workspace.
    acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)
    ifNoneMatch := "ifNoneMatch_example" // string | Caching: Optional header to improve performance. The value of this header should be the `ETag` of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkspaceApi.GetWorkspaceDetails(context.Background(), workspaceName).AcceptVersion(acceptVersion).IfNoneMatch(ifNoneMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceApi.GetWorkspaceDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkspaceDetails`: WorkspaceDetailsItem
    fmt.Fprintf(os.Stdout, "Response from `WorkspaceApi.GetWorkspaceDetails`: %v\n", resp)
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
    openapiclient "./openapi"
)

func main() {
    workspaceName := "workspaceName_example" // string | Unique ID of the Workspace.
    acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)
    ifNoneMatch := "ifNoneMatch_example" // string | Caching: Optional header to improve performance. The value of this header should be the `ETag` of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkspaceApi.GetWorkspaceRepositoryContentManager(context.Background(), workspaceName).AcceptVersion(acceptVersion).IfNoneMatch(ifNoneMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceApi.GetWorkspaceRepositoryContentManager``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkspaceRepositoryContentManager`: WorkspaceRepositoryContentManager
    fmt.Fprintf(os.Stdout, "Response from `WorkspaceApi.GetWorkspaceRepositoryContentManager`: %v\n", resp)
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
    openapiclient "./openapi"
)

func main() {
    acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)
    embed := false // bool | Embedding: The whether or not to embed resources into the collection (rather than return links). (optional) (default to false)
    ifNoneMatch := "ifNoneMatch_example" // string | Caching: Optional header to improve performance. The value of this header should be the `ETag` of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content. (optional)
    limit := int32(20) // int32 | Paging: The maximum number of items to return in a resource. (optional) (default to 20)
    offset := int32(0) // int32 | Paging:  The index of the first item to return in the resource. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkspaceApi.ListWorkspaces(context.Background()).AcceptVersion(acceptVersion).Embed(embed).IfNoneMatch(ifNoneMatch).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceApi.ListWorkspaces``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorkspaces`: WorkspaceCollection
    fmt.Fprintf(os.Stdout, "Response from `WorkspaceApi.ListWorkspaces`: %v\n", resp)
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
    openapiclient "./openapi"
)

func main() {
    workspaceName := "workspaceName_example" // string | Unique ID of the Workspace.
    acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)
    retainWorkspaceRequest := *openapiclient.NewRetainWorkspaceRequest() // RetainWorkspaceRequest | TTL configuration. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkspaceApi.RetainWorkspace(context.Background(), workspaceName).AcceptVersion(acceptVersion).RetainWorkspaceRequest(retainWorkspaceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceApi.RetainWorkspace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetainWorkspace`: WorkspaceItem
    fmt.Fprintf(os.Stdout, "Response from `WorkspaceApi.RetainWorkspace`: %v\n", resp)
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

> ArtefactManagerItem UploadWorkspaceArchiveContent(ctx, workspaceName).IfMatch(ifMatch).Content(content).Hash(hash).AcceptVersion(acceptVersion).ContentMediaType(contentMediaType).Size(size).Title(title).Execute()

Upload and replace the content of the named workspace.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    workspaceName := "workspaceName_example" // string | Unique ID of the Workspace.
    ifMatch := "ifMatch_example" // string | Conditional Requests: This is required in order to perform an update of a resource. The value of this header should be the `ETag` of the resource when read (before being subsequently modified by the client).
    content := os.NewFile(1234, "some_file") // *os.File | artefact content
    hash := "hash_example" // string | Hash of the artefact (sha256) for network resilience
    acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)
    contentMediaType := "contentMediaType_example" // string | Media type of the artefact according to https://www.iana.org/assignments/media-types/media-types.xhtml Technically redundant, but can be used by JSON Schema tools that may not be aware of the OpenAPI context. (optional)
    size := int64(789) // int64 | size in bytes of this artefact. Technically redundant, but can be used by JSON Schema tools that may not be aware of the OpenAPI context. (optional)
    title := "title_example" // string | Optional human readable name of the artefact. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkspaceApi.UploadWorkspaceArchiveContent(context.Background(), workspaceName).IfMatch(ifMatch).Content(content).Hash(hash).AcceptVersion(acceptVersion).ContentMediaType(contentMediaType).Size(size).Title(title).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceApi.UploadWorkspaceArchiveContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadWorkspaceArchiveContent`: ArtefactManagerItem
    fmt.Fprintf(os.Stdout, "Response from `WorkspaceApi.UploadWorkspaceArchiveContent`: %v\n", resp)
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

