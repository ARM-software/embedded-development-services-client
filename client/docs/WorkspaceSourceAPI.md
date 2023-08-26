<!--
Copyright (C) 2020-2023 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# \WorkspaceSourceAPI

All URIs are relative to *https://all.api.keil.arm.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWorkspaceSource**](WorkspaceSourceAPI.md#GetWorkspaceSource) | **Get** /workspace-sources/{workspaceSourceName} | Return details of the specific workspace source.
[**ListWorkspaceSources**](WorkspaceSourceAPI.md#ListWorkspaceSources) | **Get** /workspace-sources/ | List available workspace sources.



## GetWorkspaceSource

> WorkspaceSourceItem GetWorkspaceSource(ctx, workspaceSourceName).AcceptVersion(acceptVersion).IfNoneMatch(ifNoneMatch).Execute()

Return details of the specific workspace source.



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
    acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)
    ifNoneMatch := "ifNoneMatch_example" // string | Caching: Optional header to improve performance. The value of this header should be the `ETag` of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkspaceSourceAPI.GetWorkspaceSource(context.Background(), workspaceSourceName).AcceptVersion(acceptVersion).IfNoneMatch(ifNoneMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceSourceAPI.GetWorkspaceSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkspaceSource`: WorkspaceSourceItem
    fmt.Fprintf(os.Stdout, "Response from `WorkspaceSourceAPI.GetWorkspaceSource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceSourceName** | **string** | The ID of the workspace source. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkspaceSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 
 **ifNoneMatch** | **string** | Caching: Optional header to improve performance. The value of this header should be the &#x60;ETag&#x60; of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content. | 

### Return type

[**WorkspaceSourceItem**](WorkspaceSourceItem.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkspaceSources

> WorkspaceSourceCollection ListWorkspaceSources(ctx).AcceptVersion(acceptVersion).Embed(embed).IfNoneMatch(ifNoneMatch).Limit(limit).Offset(offset).Execute()

List available workspace sources.



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
    resp, r, err := apiClient.WorkspaceSourceAPI.ListWorkspaceSources(context.Background()).AcceptVersion(acceptVersion).Embed(embed).IfNoneMatch(ifNoneMatch).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceSourceAPI.ListWorkspaceSources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorkspaceSources`: WorkspaceSourceCollection
    fmt.Fprintf(os.Stdout, "Response from `WorkspaceSourceAPI.ListWorkspaceSources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWorkspaceSourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 
 **embed** | **bool** | Embedding: The whether or not to embed resources into the collection (rather than return links). | [default to false]
 **ifNoneMatch** | **string** | Caching: Optional header to improve performance. The value of this header should be the &#x60;ETag&#x60; of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content. | 
 **limit** | **int32** | Paging: The maximum number of items to return in a resource. | [default to 20]
 **offset** | **int32** | Paging:  The index of the first item to return in the resource. | [default to 0]

### Return type

[**WorkspaceSourceCollection**](WorkspaceSourceCollection.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

