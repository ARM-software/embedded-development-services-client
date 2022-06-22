# \VirtualHardwareTargetInstanceApi

All URIs are relative to *https://all.api.keil.arm.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClearVhtInstanceArtefact**](VirtualHardwareTargetInstanceApi.md#ClearVhtInstanceArtefact) | **Delete** /vht-instances/{instanceName}/artefacts/{artefactName}/artefact | Clear the VHT artefact from the VHT instance.
[**DeleteVhtInstance**](VirtualHardwareTargetInstanceApi.md#DeleteVhtInstance) | **Delete** /vht-instances/{instanceName} | Delete a VHT instance
[**DownloadVhtInstanceArtefact**](VirtualHardwareTargetInstanceApi.md#DownloadVhtInstanceArtefact) | **Get** /vht-instances/{instanceName}/artefacts/{artefactName}/artefact | Download the artefact named &#x60;artefactName&#x60; present on this VHT instance.
[**GetVhtInstance**](VirtualHardwareTargetInstanceApi.md#GetVhtInstance) | **Get** /vht-instances/{instanceName} | Return status of a VHT instance.
[**GetVhtInstanceArtefactManager**](VirtualHardwareTargetInstanceApi.md#GetVhtInstanceArtefactManager) | **Get** /vht-instances/{instanceName}/artefacts/{artefactName} | Get the VHT artefact manager for the artefact named &#x60;artefactName&#x60; present of this VHT instance.
[**GetVhtInstanceMessages**](VirtualHardwareTargetInstanceApi.md#GetVhtInstanceMessages) | **Get** /vht-instances/{instanceName}/messages | Instance Message Feed.
[**ListVhtInstanceArtefactManagers**](VirtualHardwareTargetInstanceApi.md#ListVhtInstanceArtefactManagers) | **Get** /vht-instances/{instanceName}/artefacts/ | List all the managers of the artefacts (e.g. binary, test input) available on a specific VHT instance.
[**ListVhtInstances**](VirtualHardwareTargetInstanceApi.md#ListVhtInstances) | **Get** /vht-instances/ | List all VHT instances requested.
[**StartVhtRunJob**](VirtualHardwareTargetInstanceApi.md#StartVhtRunJob) | **Post** /vht-instances/{instanceName} | Starts a VHT Run job.
[**StopVhtInstance**](VirtualHardwareTargetInstanceApi.md#StopVhtInstance) | **Post** /vht-instances/{instanceName}/cancel | Stop this VHT instance.
[**UploadVhtInstanceArtefact**](VirtualHardwareTargetInstanceApi.md#UploadVhtInstanceArtefact) | **Put** /vht-instances/{instanceName}/artefacts/{artefactName}/artefact | Upload and replace the named VHT artefact on the given VHT instance.



## ClearVhtInstanceArtefact

> ClearVhtInstanceArtefact(ctx, artefactName, instanceName).AcceptVersion(acceptVersion).Execute()

Clear the VHT artefact from the VHT instance.



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
    artefactName := "artefactName_example" // string | The URL safe name of the artefact.
    instanceName := "instanceName_example" // string | Unique ID of the VHT instance.
    acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirtualHardwareTargetInstanceApi.ClearVhtInstanceArtefact(context.Background(), artefactName, instanceName).AcceptVersion(acceptVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualHardwareTargetInstanceApi.ClearVhtInstanceArtefact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**artefactName** | **string** | The URL safe name of the artefact. | 
**instanceName** | **string** | Unique ID of the VHT instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiClearVhtInstanceArtefactRequest struct via the builder pattern


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


## DeleteVhtInstance

> DeleteVhtInstance(ctx, instanceName).AcceptVersion(acceptVersion).Execute()

Delete a VHT instance



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
    instanceName := "instanceName_example" // string | Unique ID of the VHT instance.
    acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirtualHardwareTargetInstanceApi.DeleteVhtInstance(context.Background(), instanceName).AcceptVersion(acceptVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualHardwareTargetInstanceApi.DeleteVhtInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceName** | **string** | Unique ID of the VHT instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVhtInstanceRequest struct via the builder pattern


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


## DownloadVhtInstanceArtefact

> *os.File DownloadVhtInstanceArtefact(ctx, artefactName, instanceName).AcceptVersion(acceptVersion).IfNoneMatch(ifNoneMatch).Execute()

Download the artefact named `artefactName` present on this VHT instance.



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
    artefactName := "artefactName_example" // string | The URL safe name of the artefact.
    instanceName := "instanceName_example" // string | Unique ID of the VHT instance.
    acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)
    ifNoneMatch := "ifNoneMatch_example" // string | Caching: Optional header to improve performance. The value of this header should be the `ETag` of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirtualHardwareTargetInstanceApi.DownloadVhtInstanceArtefact(context.Background(), artefactName, instanceName).AcceptVersion(acceptVersion).IfNoneMatch(ifNoneMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualHardwareTargetInstanceApi.DownloadVhtInstanceArtefact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadVhtInstanceArtefact`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `VirtualHardwareTargetInstanceApi.DownloadVhtInstanceArtefact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**artefactName** | **string** | The URL safe name of the artefact. | 
**instanceName** | **string** | Unique ID of the VHT instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadVhtInstanceArtefactRequest struct via the builder pattern


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
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVhtInstance

> VhtInstanceItem GetVhtInstance(ctx, instanceName).AcceptVersion(acceptVersion).IfNoneMatch(ifNoneMatch).Execute()

Return status of a VHT instance.



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
    instanceName := "instanceName_example" // string | Unique ID of the VHT instance.
    acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)
    ifNoneMatch := "ifNoneMatch_example" // string | Caching: Optional header to improve performance. The value of this header should be the `ETag` of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirtualHardwareTargetInstanceApi.GetVhtInstance(context.Background(), instanceName).AcceptVersion(acceptVersion).IfNoneMatch(ifNoneMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualHardwareTargetInstanceApi.GetVhtInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVhtInstance`: VhtInstanceItem
    fmt.Fprintf(os.Stdout, "Response from `VirtualHardwareTargetInstanceApi.GetVhtInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceName** | **string** | Unique ID of the VHT instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVhtInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 
 **ifNoneMatch** | **string** | Caching: Optional header to improve performance. The value of this header should be the &#x60;ETag&#x60; of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content. | 

### Return type

[**VhtInstanceItem**](VhtInstanceItem.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVhtInstanceArtefactManager

> ArtefactManagerItem GetVhtInstanceArtefactManager(ctx, artefactName, instanceName).AcceptVersion(acceptVersion).IfNoneMatch(ifNoneMatch).Execute()

Get the VHT artefact manager for the artefact named `artefactName` present of this VHT instance.



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
    artefactName := "artefactName_example" // string | The URL safe name of the artefact.
    instanceName := "instanceName_example" // string | Unique ID of the VHT instance.
    acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)
    ifNoneMatch := "ifNoneMatch_example" // string | Caching: Optional header to improve performance. The value of this header should be the `ETag` of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirtualHardwareTargetInstanceApi.GetVhtInstanceArtefactManager(context.Background(), artefactName, instanceName).AcceptVersion(acceptVersion).IfNoneMatch(ifNoneMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualHardwareTargetInstanceApi.GetVhtInstanceArtefactManager``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVhtInstanceArtefactManager`: ArtefactManagerItem
    fmt.Fprintf(os.Stdout, "Response from `VirtualHardwareTargetInstanceApi.GetVhtInstanceArtefactManager`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**artefactName** | **string** | The URL safe name of the artefact. | 
**instanceName** | **string** | Unique ID of the VHT instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVhtInstanceArtefactManagerRequest struct via the builder pattern


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


## GetVhtInstanceMessages

> NotificationFeed GetVhtInstanceMessages(ctx, instanceName).AcceptVersion(acceptVersion).IfNoneMatch(ifNoneMatch).Limit(limit).Offset(offset).Execute()

Instance Message Feed.



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
    instanceName := "instanceName_example" // string | Unique ID of the VHT instance.
    acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)
    ifNoneMatch := "ifNoneMatch_example" // string | Caching: Optional header to improve performance. The value of this header should be the `ETag` of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content. (optional)
    limit := int32(20) // int32 | Paging: The maximum number of items to return in a resource. (optional) (default to 20)
    offset := int32(0) // int32 | Paging:  The index of the first item to return in the resource. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirtualHardwareTargetInstanceApi.GetVhtInstanceMessages(context.Background(), instanceName).AcceptVersion(acceptVersion).IfNoneMatch(ifNoneMatch).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualHardwareTargetInstanceApi.GetVhtInstanceMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVhtInstanceMessages`: NotificationFeed
    fmt.Fprintf(os.Stdout, "Response from `VirtualHardwareTargetInstanceApi.GetVhtInstanceMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceName** | **string** | Unique ID of the VHT instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVhtInstanceMessagesRequest struct via the builder pattern


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


## ListVhtInstanceArtefactManagers

> ArtefactManagerCollection ListVhtInstanceArtefactManagers(ctx, instanceName).AcceptVersion(acceptVersion).Embed(embed).IfNoneMatch(ifNoneMatch).Limit(limit).Offset(offset).Execute()

List all the managers of the artefacts (e.g. binary, test input) available on a specific VHT instance.



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
    instanceName := "instanceName_example" // string | Unique ID of the VHT instance.
    acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)
    embed := false // bool | Embedding: The whether or not to embed resources into the collection (rather than return links). (optional) (default to false)
    ifNoneMatch := "ifNoneMatch_example" // string | Caching: Optional header to improve performance. The value of this header should be the `ETag` of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content. (optional)
    limit := int32(20) // int32 | Paging: The maximum number of items to return in a resource. (optional) (default to 20)
    offset := int32(0) // int32 | Paging:  The index of the first item to return in the resource. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirtualHardwareTargetInstanceApi.ListVhtInstanceArtefactManagers(context.Background(), instanceName).AcceptVersion(acceptVersion).Embed(embed).IfNoneMatch(ifNoneMatch).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualHardwareTargetInstanceApi.ListVhtInstanceArtefactManagers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVhtInstanceArtefactManagers`: ArtefactManagerCollection
    fmt.Fprintf(os.Stdout, "Response from `VirtualHardwareTargetInstanceApi.ListVhtInstanceArtefactManagers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceName** | **string** | Unique ID of the VHT instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListVhtInstanceArtefactManagersRequest struct via the builder pattern


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


## ListVhtInstances

> VhtInstanceCollection ListVhtInstances(ctx).AcceptVersion(acceptVersion).Embed(embed).IfNoneMatch(ifNoneMatch).Limit(limit).Offset(offset).Execute()

List all VHT instances requested.



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
    resp, r, err := apiClient.VirtualHardwareTargetInstanceApi.ListVhtInstances(context.Background()).AcceptVersion(acceptVersion).Embed(embed).IfNoneMatch(ifNoneMatch).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualHardwareTargetInstanceApi.ListVhtInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVhtInstances`: VhtInstanceCollection
    fmt.Fprintf(os.Stdout, "Response from `VirtualHardwareTargetInstanceApi.ListVhtInstances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVhtInstancesRequest struct via the builder pattern


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


## StartVhtRunJob

> VhtRunJobItem StartVhtRunJob(ctx, instanceName).VhtRunJobItem(vhtRunJobItem).AcceptVersion(acceptVersion).Execute()

Starts a VHT Run job.



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
    instanceName := "instanceName_example" // string | Unique ID of the VHT instance.
    vhtRunJobItem := *openapiclient.NewVhtRunJobItem("TODO", "TODO", true, true, "332129b3-f14d-49d2-b9be-acd2abd80c6b", false, "PENDING", false, false) // VhtRunJobItem | A VHT run job.
    acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirtualHardwareTargetInstanceApi.StartVhtRunJob(context.Background(), instanceName).VhtRunJobItem(vhtRunJobItem).AcceptVersion(acceptVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualHardwareTargetInstanceApi.StartVhtRunJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartVhtRunJob`: VhtRunJobItem
    fmt.Fprintf(os.Stdout, "Response from `VirtualHardwareTargetInstanceApi.StartVhtRunJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceName** | **string** | Unique ID of the VHT instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartVhtRunJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vhtRunJobItem** | [**VhtRunJobItem**](VhtRunJobItem.md) | A VHT run job. | 
 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 

### Return type

[**VhtRunJobItem**](VhtRunJobItem.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopVhtInstance

> VhtInstanceItem StopVhtInstance(ctx, instanceName).AcceptVersion(acceptVersion).Execute()

Stop this VHT instance.



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
    instanceName := "instanceName_example" // string | Unique ID of the VHT instance.
    acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirtualHardwareTargetInstanceApi.StopVhtInstance(context.Background(), instanceName).AcceptVersion(acceptVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualHardwareTargetInstanceApi.StopVhtInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StopVhtInstance`: VhtInstanceItem
    fmt.Fprintf(os.Stdout, "Response from `VirtualHardwareTargetInstanceApi.StopVhtInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceName** | **string** | Unique ID of the VHT instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopVhtInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 

### Return type

[**VhtInstanceItem**](VhtInstanceItem.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadVhtInstanceArtefact

> ArtefactManagerItem UploadVhtInstanceArtefact(ctx, artefactName, instanceName).IfMatch(ifMatch).Content(content).Hash(hash).AcceptVersion(acceptVersion).ContentMediaType(contentMediaType).Size(size).Title(title).Execute()

Upload and replace the named VHT artefact on the given VHT instance.



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
    artefactName := "artefactName_example" // string | The URL safe name of the artefact.
    instanceName := "instanceName_example" // string | Unique ID of the VHT instance.
    ifMatch := "ifMatch_example" // string | Conditional Requests: This is required in order to perform an update of a resource. The value of this header should be the `ETag` of the resource when read (before being subsequently modified by the client).
    content := os.NewFile(1234, "some_file") // *os.File | artefact content
    hash := "hash_example" // string | Hash of the artefact (sha256) for network resilience
    acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)
    contentMediaType := "contentMediaType_example" // string | Media type of the artefact according to https://www.iana.org/assignments/media-types/media-types.xhtml Technically redundant, but can be used by JSON Schema tools that may not be aware of the OpenAPI context. (optional)
    size := int64(789) // int64 | size in bytes of this artefact. Technically redundant, but can be used by JSON Schema tools that may not be aware of the OpenAPI context. (optional)
    title := "title_example" // string | Optional human readable name of the artefact. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirtualHardwareTargetInstanceApi.UploadVhtInstanceArtefact(context.Background(), artefactName, instanceName).IfMatch(ifMatch).Content(content).Hash(hash).AcceptVersion(acceptVersion).ContentMediaType(contentMediaType).Size(size).Title(title).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualHardwareTargetInstanceApi.UploadVhtInstanceArtefact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadVhtInstanceArtefact`: ArtefactManagerItem
    fmt.Fprintf(os.Stdout, "Response from `VirtualHardwareTargetInstanceApi.UploadVhtInstanceArtefact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**artefactName** | **string** | The URL safe name of the artefact. | 
**instanceName** | **string** | Unique ID of the VHT instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadVhtInstanceArtefactRequest struct via the builder pattern


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

