<!--
Copyright (C) 2020-2026 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# \FPGAEntitlementsAPI

All URIs are relative to *https://all.api.keil.arm.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFPGAEntitlement**](FPGAEntitlementsAPI.md#GetFPGAEntitlement) | **Get** /fpga-entitlements/{fpgaEntitlementName} | Return details of specific FPGA entitlement.
[**ListFPGAEntitlements**](FPGAEntitlementsAPI.md#ListFPGAEntitlements) | **Get** /fpga-entitlements | List all the entitlements that are managable by the FPGA admin group.
[**RemoveFPGAEntitlement**](FPGAEntitlementsAPI.md#RemoveFPGAEntitlement) | **Delete** /fpga-entitlements/{fpgaEntitlementName} | Remove all the permissions associated with an FPGA entitlement
[**UpdateFPGAEntitlement**](FPGAEntitlementsAPI.md#UpdateFPGAEntitlement) | **Put** /fpga-entitlements | Update an FPGA entitlement that defines which users are permitted to interact with an FPGA.



## GetFPGAEntitlement

> FPGAEntitlementItem GetFPGAEntitlement(ctx, fpgaEntitlementName).AcceptVersion(acceptVersion).IfNoneMatch(ifNoneMatch).Execute()

Return details of specific FPGA entitlement.



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
	fpgaEntitlementName := "fpgaEntitlementName_example" // string | Unique ID of the FPGA entitlement.
	acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)
	ifNoneMatch := "ifNoneMatch_example" // string | Caching: Optional header to improve performance. The value of this header should be the `ETag` of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FPGAEntitlementsAPI.GetFPGAEntitlement(context.Background(), fpgaEntitlementName).AcceptVersion(acceptVersion).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FPGAEntitlementsAPI.GetFPGAEntitlement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFPGAEntitlement`: FPGAEntitlementItem
	fmt.Fprintf(os.Stdout, "Response from `FPGAEntitlementsAPI.GetFPGAEntitlement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fpgaEntitlementName** | **string** | Unique ID of the FPGA entitlement. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFPGAEntitlementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 
 **ifNoneMatch** | **string** | Caching: Optional header to improve performance. The value of this header should be the &#x60;ETag&#x60; of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content. | 

### Return type

[**FPGAEntitlementItem**](FPGAEntitlementItem.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFPGAEntitlements

> SimpleCollection ListFPGAEntitlements(ctx).AcceptVersion(acceptVersion).IfNoneMatch(ifNoneMatch).Limit(limit).Offset(offset).Execute()

List all the entitlements that are managable by the FPGA admin group.



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
	ifNoneMatch := "ifNoneMatch_example" // string | Caching: Optional header to improve performance. The value of this header should be the `ETag` of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content. (optional)
	limit := int32(20) // int32 | Paging: The maximum number of items to return in a resource. (optional) (default to 20)
	offset := int32(0) // int32 | Paging:  The index of the first item to return in the resource. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FPGAEntitlementsAPI.ListFPGAEntitlements(context.Background()).AcceptVersion(acceptVersion).IfNoneMatch(ifNoneMatch).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FPGAEntitlementsAPI.ListFPGAEntitlements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFPGAEntitlements`: SimpleCollection
	fmt.Fprintf(os.Stdout, "Response from `FPGAEntitlementsAPI.ListFPGAEntitlements`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFPGAEntitlementsRequest struct via the builder pattern


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


## RemoveFPGAEntitlement

> RemoveFPGAEntitlement(ctx, fpgaEntitlementName).AcceptVersion(acceptVersion).Execute()

Remove all the permissions associated with an FPGA entitlement



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
	fpgaEntitlementName := "fpgaEntitlementName_example" // string | Unique ID of the FPGA entitlement.
	acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FPGAEntitlementsAPI.RemoveFPGAEntitlement(context.Background(), fpgaEntitlementName).AcceptVersion(acceptVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FPGAEntitlementsAPI.RemoveFPGAEntitlement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fpgaEntitlementName** | **string** | Unique ID of the FPGA entitlement. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveFPGAEntitlementRequest struct via the builder pattern


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


## UpdateFPGAEntitlement

> FPGAEntitlementItem UpdateFPGAEntitlement(ctx).IfMatch(ifMatch).FPGAEntitlementItem(fPGAEntitlementItem).AcceptVersion(acceptVersion).Execute()

Update an FPGA entitlement that defines which users are permitted to interact with an FPGA.



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
	ifMatch := "ifMatch_example" // string | Conditional Requests: This is required in order to perform an update of a resource. The value of this header should be the `ETag` of the resource when read (before being subsequently modified by the client).
	fPGAEntitlementItem := *openapiclient.NewFPGAEntitlementItem("TODO", "TODO", []string{"Fpgas_example"}, "samsung", []string{"Repositories_example"}, []string{"Users_example"}) // FPGAEntitlementItem | 
	acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FPGAEntitlementsAPI.UpdateFPGAEntitlement(context.Background()).IfMatch(ifMatch).FPGAEntitlementItem(fPGAEntitlementItem).AcceptVersion(acceptVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FPGAEntitlementsAPI.UpdateFPGAEntitlement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFPGAEntitlement`: FPGAEntitlementItem
	fmt.Fprintf(os.Stdout, "Response from `FPGAEntitlementsAPI.UpdateFPGAEntitlement`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFPGAEntitlementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ifMatch** | **string** | Conditional Requests: This is required in order to perform an update of a resource. The value of this header should be the &#x60;ETag&#x60; of the resource when read (before being subsequently modified by the client). | 
 **fPGAEntitlementItem** | [**FPGAEntitlementItem**](FPGAEntitlementItem.md) |  | 
 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 

### Return type

[**FPGAEntitlementItem**](FPGAEntitlementItem.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

