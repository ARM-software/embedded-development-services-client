<!--
Copyright (C) 2020-2025 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# \PersonalAccessTokenAPI

All URIs are relative to *https://all.api.keil.arm.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePAT**](PersonalAccessTokenAPI.md#CreatePAT) | **Post** /personal-access-tokens/ | Create a new personal access token for a user
[**DeletePAT**](PersonalAccessTokenAPI.md#DeletePAT) | **Delete** /personal-access-tokens/{patName} | Delete a personal access token
[**GetPAT**](PersonalAccessTokenAPI.md#GetPAT) | **Get** /personal-access-tokens/{patName} | Get a personal access token
[**ListPATs**](PersonalAccessTokenAPI.md#ListPATs) | **Get** /personal-access-tokens/ | List all personal access tokens for the user



## CreatePAT

> PATCreationItem CreatePAT(ctx).PATCreationItem(pATCreationItem).AcceptVersion(acceptVersion).Execute()

Create a new personal access token for a user



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
	pATCreationItem := *openapiclient.NewPATCreationItem(int64(5), "TODO", "TODO", "88a6137e-1d99-4cde-8db8-015312f7d5e6", "arm_pat_abcdefgh...", "My token") // PATCreationItem | Data required to create a new token.
	acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PersonalAccessTokenAPI.CreatePAT(context.Background()).PATCreationItem(pATCreationItem).AcceptVersion(acceptVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PersonalAccessTokenAPI.CreatePAT``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePAT`: PATCreationItem
	fmt.Fprintf(os.Stdout, "Response from `PersonalAccessTokenAPI.CreatePAT`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePATRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pATCreationItem** | [**PATCreationItem**](PATCreationItem.md) | Data required to create a new token. | 
 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 

### Return type

[**PATCreationItem**](PATCreationItem.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePAT

> DeletePAT(ctx, patName).AcceptVersion(acceptVersion).Execute()

Delete a personal access token



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
	patName := "patName_example" // string | Unique ID of the personal access token.
	acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PersonalAccessTokenAPI.DeletePAT(context.Background(), patName).AcceptVersion(acceptVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PersonalAccessTokenAPI.DeletePAT``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**patName** | **string** | Unique ID of the personal access token. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePATRequest struct via the builder pattern


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


## GetPAT

> PATItem GetPAT(ctx, patName).AcceptVersion(acceptVersion).IfNoneMatch(ifNoneMatch).Execute()

Get a personal access token



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
	patName := "patName_example" // string | Unique ID of the personal access token.
	acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)
	ifNoneMatch := "ifNoneMatch_example" // string | Caching: Optional header to improve performance. The value of this header should be the `ETag` of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PersonalAccessTokenAPI.GetPAT(context.Background(), patName).AcceptVersion(acceptVersion).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PersonalAccessTokenAPI.GetPAT``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPAT`: PATItem
	fmt.Fprintf(os.Stdout, "Response from `PersonalAccessTokenAPI.GetPAT`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**patName** | **string** | Unique ID of the personal access token. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPATRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 
 **ifNoneMatch** | **string** | Caching: Optional header to improve performance. The value of this header should be the &#x60;ETag&#x60; of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content. | 

### Return type

[**PATItem**](PATItem.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPATs

> PATCollection ListPATs(ctx).AcceptVersion(acceptVersion).Embed(embed).IfNoneMatch(ifNoneMatch).Execute()

List all personal access tokens for the user



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PersonalAccessTokenAPI.ListPATs(context.Background()).AcceptVersion(acceptVersion).Embed(embed).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PersonalAccessTokenAPI.ListPATs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPATs`: PATCollection
	fmt.Fprintf(os.Stdout, "Response from `PersonalAccessTokenAPI.ListPATs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPATsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 
 **embed** | **bool** | Embedding: The whether or not to embed resources into the collection (rather than return links). | [default to false]
 **ifNoneMatch** | **string** | Caching: Optional header to improve performance. The value of this header should be the &#x60;ETag&#x60; of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content. | 

### Return type

[**PATCollection**](PATCollection.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

