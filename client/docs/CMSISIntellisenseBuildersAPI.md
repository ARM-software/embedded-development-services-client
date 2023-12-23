<!--
Copyright (C) 2020-2023 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# \CMSISIntellisenseBuildersAPI

All URIs are relative to *https://all.api.keil.arm.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCmsisIntellisense**](CMSISIntellisenseBuildersAPI.md#GetCmsisIntellisense) | **Get** /cmsis-intellisense/{builderName} | Return details of specific CMSIS Intellisense Builders.
[**ListCmsisIntellisense**](CMSISIntellisenseBuildersAPI.md#ListCmsisIntellisense) | **Get** /cmsis-intellisense/ | List available CMSIS Intellisense Builders.
[**StartCmsisIntellisense**](CMSISIntellisenseBuildersAPI.md#StartCmsisIntellisense) | **Post** /cmsis-intellisense/{builderName} | Initiate a compilation database generation using the specified CMSIS Intellisense Builder.



## GetCmsisIntellisense

> CmsisIntellisenseItem GetCmsisIntellisense(ctx, builderName).AcceptVersion(acceptVersion).IfNoneMatch(ifNoneMatch).Execute()

Return details of specific CMSIS Intellisense Builders.



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
	builderName := "builderName_example" // string | The ID of the CMSIS Intellisense Builder.
	acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)
	ifNoneMatch := "ifNoneMatch_example" // string | Caching: Optional header to improve performance. The value of this header should be the `ETag` of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CMSISIntellisenseBuildersAPI.GetCmsisIntellisense(context.Background(), builderName).AcceptVersion(acceptVersion).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CMSISIntellisenseBuildersAPI.GetCmsisIntellisense``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCmsisIntellisense`: CmsisIntellisenseItem
	fmt.Fprintf(os.Stdout, "Response from `CMSISIntellisenseBuildersAPI.GetCmsisIntellisense`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**builderName** | **string** | The ID of the CMSIS Intellisense Builder. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsisIntellisenseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 
 **ifNoneMatch** | **string** | Caching: Optional header to improve performance. The value of this header should be the &#x60;ETag&#x60; of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content. | 

### Return type

[**CmsisIntellisenseItem**](CmsisIntellisenseItem.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCmsisIntellisense

> CmsisIntellisenseCollection ListCmsisIntellisense(ctx).AcceptVersion(acceptVersion).Embed(embed).IfNoneMatch(ifNoneMatch).Limit(limit).Offset(offset).Execute()

List available CMSIS Intellisense Builders.



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
	resp, r, err := apiClient.CMSISIntellisenseBuildersAPI.ListCmsisIntellisense(context.Background()).AcceptVersion(acceptVersion).Embed(embed).IfNoneMatch(ifNoneMatch).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CMSISIntellisenseBuildersAPI.ListCmsisIntellisense``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCmsisIntellisense`: CmsisIntellisenseCollection
	fmt.Fprintf(os.Stdout, "Response from `CMSISIntellisenseBuildersAPI.ListCmsisIntellisense`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCmsisIntellisenseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 
 **embed** | **bool** | Embedding: The whether or not to embed resources into the collection (rather than return links). | [default to false]
 **ifNoneMatch** | **string** | Caching: Optional header to improve performance. The value of this header should be the &#x60;ETag&#x60; of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content. | 
 **limit** | **int32** | Paging: The maximum number of items to return in a resource. | [default to 20]
 **offset** | **int32** | Paging:  The index of the first item to return in the resource. | [default to 0]

### Return type

[**CmsisIntellisenseCollection**](CmsisIntellisenseCollection.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartCmsisIntellisense

> IntellisenseJobItem StartCmsisIntellisense(ctx, builderName).IntellisenseJobItem(intellisenseJobItem).AcceptVersion(acceptVersion).Execute()

Initiate a compilation database generation using the specified CMSIS Intellisense Builder.



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
	builderName := "builderName_example" // string | The ID of the CMSIS Intellisense Builder.
	intellisenseJobItem := *openapiclient.NewIntellisenseJobItem("TODO", "TODO", NullableInt32(50), NullableInt32(70), true, false, true, "332129b3-f14d-49d2-b9be-acd2abd80c6b", "/root/packs", "workspace/debug-build.csolution.yaml", "INITIALISING", false, "/root/toolchain/bin", "/root/toolchain/headers", "/root/workspace") // IntellisenseJobItem | A name of the CMSIS project to generate compilation database.
	acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CMSISIntellisenseBuildersAPI.StartCmsisIntellisense(context.Background(), builderName).IntellisenseJobItem(intellisenseJobItem).AcceptVersion(acceptVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CMSISIntellisenseBuildersAPI.StartCmsisIntellisense``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartCmsisIntellisense`: IntellisenseJobItem
	fmt.Fprintf(os.Stdout, "Response from `CMSISIntellisenseBuildersAPI.StartCmsisIntellisense`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**builderName** | **string** | The ID of the CMSIS Intellisense Builder. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartCmsisIntellisenseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **intellisenseJobItem** | [**IntellisenseJobItem**](IntellisenseJobItem.md) | A name of the CMSIS project to generate compilation database. | 
 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 

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

