<!--
Copyright (C) 2020-2023 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# \VirtualHardwareTargetApi

All URIs are relative to *https://all.api.keil.arm.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVht**](VirtualHardwareTargetApi.md#GetVht) | **Get** /vhts/{vhtName} | Return details of the specific VHT.
[**ListFilteredVhtInstances**](VirtualHardwareTargetApi.md#ListFilteredVhtInstances) | **Get** /vhts/{vhtName}/vht-instances/ | List all VHT instances related to this specific VHT.
[**ListVhts**](VirtualHardwareTargetApi.md#ListVhts) | **Get** /vhts/ | List available VHTs.
[**StartVhtInstance**](VirtualHardwareTargetApi.md#StartVhtInstance) | **Post** /vhts/{vhtName} | Spawns a VHT instance matching this VHT specification.



## GetVht

> VhtItem GetVht(ctx, vhtName).AcceptVersion(acceptVersion).IfNoneMatch(ifNoneMatch).Execute()

Return details of the specific VHT.



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
    vhtName := "vhtName_example" // string | The ID of the VHT.
    acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)
    ifNoneMatch := "ifNoneMatch_example" // string | Caching: Optional header to improve performance. The value of this header should be the `ETag` of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirtualHardwareTargetApi.GetVht(context.Background(), vhtName).AcceptVersion(acceptVersion).IfNoneMatch(ifNoneMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualHardwareTargetApi.GetVht``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVht`: VhtItem
    fmt.Fprintf(os.Stdout, "Response from `VirtualHardwareTargetApi.GetVht`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vhtName** | **string** | The ID of the VHT. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVhtRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 
 **ifNoneMatch** | **string** | Caching: Optional header to improve performance. The value of this header should be the &#x60;ETag&#x60; of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content. | 

### Return type

[**VhtItem**](VhtItem.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFilteredVhtInstances

> VhtInstanceCollection ListFilteredVhtInstances(ctx, vhtName).AcceptVersion(acceptVersion).Embed(embed).IfNoneMatch(ifNoneMatch).Limit(limit).Offset(offset).Execute()

List all VHT instances related to this specific VHT.



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
    vhtName := "vhtName_example" // string | The ID of the VHT.
    acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)
    embed := false // bool | Embedding: The whether or not to embed resources into the collection (rather than return links). (optional) (default to false)
    ifNoneMatch := "ifNoneMatch_example" // string | Caching: Optional header to improve performance. The value of this header should be the `ETag` of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content. (optional)
    limit := int32(20) // int32 | Paging: The maximum number of items to return in a resource. (optional) (default to 20)
    offset := int32(0) // int32 | Paging:  The index of the first item to return in the resource. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirtualHardwareTargetApi.ListFilteredVhtInstances(context.Background(), vhtName).AcceptVersion(acceptVersion).Embed(embed).IfNoneMatch(ifNoneMatch).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualHardwareTargetApi.ListFilteredVhtInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFilteredVhtInstances`: VhtInstanceCollection
    fmt.Fprintf(os.Stdout, "Response from `VirtualHardwareTargetApi.ListFilteredVhtInstances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vhtName** | **string** | The ID of the VHT. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListFilteredVhtInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 
 **embed** | **bool** | Embedding: The whether or not to embed resources into the collection (rather than return links). | [default to false]
 **ifNoneMatch** | **string** | Caching: Optional header to improve performance. The value of this header should be the &#x60;ETag&#x60; of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content. | 
 **limit** | **int32** | Paging: The maximum number of items to return in a resource. | [default to 20]
 **offset** | **int32** | Paging:  The index of the first item to return in the resource. | [default to 0]

### Return type

[**VhtInstanceCollection**](VhtInstanceCollection.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVhts

> VhtCollection ListVhts(ctx).AcceptVersion(acceptVersion).Embed(embed).IfNoneMatch(ifNoneMatch).Limit(limit).Offset(offset).Execute()

List available VHTs.



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
    resp, r, err := apiClient.VirtualHardwareTargetApi.ListVhts(context.Background()).AcceptVersion(acceptVersion).Embed(embed).IfNoneMatch(ifNoneMatch).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualHardwareTargetApi.ListVhts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVhts`: VhtCollection
    fmt.Fprintf(os.Stdout, "Response from `VirtualHardwareTargetApi.ListVhts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVhtsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 
 **embed** | **bool** | Embedding: The whether or not to embed resources into the collection (rather than return links). | [default to false]
 **ifNoneMatch** | **string** | Caching: Optional header to improve performance. The value of this header should be the &#x60;ETag&#x60; of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content. | 
 **limit** | **int32** | Paging: The maximum number of items to return in a resource. | [default to 20]
 **offset** | **int32** | Paging:  The index of the first item to return in the resource. | [default to 0]

### Return type

[**VhtCollection**](VhtCollection.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartVhtInstance

> VhtInstanceItem StartVhtInstance(ctx, vhtName).VhtInstanceItem(vhtInstanceItem).AcceptVersion(acceptVersion).Execute()

Spawns a VHT instance matching this VHT specification.



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
    vhtName := "vhtName_example" // string | The ID of the VHT.
    vhtInstanceItem := *openapiclient.NewVhtInstanceItem("TODO", "TODO", false, "332129b3-f14d-49d2-b9be-acd2abd80c6b", false, false, "INITIALISING", true) // VhtInstanceItem | A name of the VHT to request.
    acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirtualHardwareTargetApi.StartVhtInstance(context.Background(), vhtName).VhtInstanceItem(vhtInstanceItem).AcceptVersion(acceptVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualHardwareTargetApi.StartVhtInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartVhtInstance`: VhtInstanceItem
    fmt.Fprintf(os.Stdout, "Response from `VirtualHardwareTargetApi.StartVhtInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vhtName** | **string** | The ID of the VHT. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartVhtInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vhtInstanceItem** | [**VhtInstanceItem**](VhtInstanceItem.md) | A name of the VHT to request. | 
 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 

### Return type

[**VhtInstanceItem**](VhtInstanceItem.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

