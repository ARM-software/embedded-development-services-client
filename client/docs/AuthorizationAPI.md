<!--
Copyright (C) 2020-2025 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# \AuthorizationAPI

All URIs are relative to *https://all.api.keil.arm.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckEntitlements**](AuthorizationAPI.md#CheckEntitlements) | **Post** /entitlements/check | Check permissions on a resource.
[**GetResourceInstanceEntitlements**](AuthorizationAPI.md#GetResourceInstanceEntitlements) | **Get** /users/{userName}/entitlements/{resourceType}/instances/{resourceInstanceName} | Return details of the user&#39;s permissions for a resource instance.
[**GetResourceTypeEntitlements**](AuthorizationAPI.md#GetResourceTypeEntitlements) | **Get** /users/{userName}/entitlements/{resourceType} | Return details of the user&#39;s permissions for a resource type.
[**ListResourceInstanceEntitlements**](AuthorizationAPI.md#ListResourceInstanceEntitlements) | **Get** /users/{userName}/entitlements/{resourceType}/instances | List the user&#39;s permissions for all instances of a resource type.



## CheckEntitlements

> PermissionItem CheckEntitlements(ctx).PermissionItem(permissionItem).AcceptVersion(acceptVersion).Execute()

Check permissions on a resource.



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
	permissionItem := *openapiclient.NewPermissionItem("TODO", "worker_1234-job-5678", "GenericWorkJobItem", "b7f6c5d1-2a44-4f9a-9d77-6e5a8c9d4e1b") // PermissionItem | Data required to check permissions on a resource.
	acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthorizationAPI.CheckEntitlements(context.Background()).PermissionItem(permissionItem).AcceptVersion(acceptVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationAPI.CheckEntitlements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckEntitlements`: PermissionItem
	fmt.Fprintf(os.Stdout, "Response from `AuthorizationAPI.CheckEntitlements`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckEntitlementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **permissionItem** | [**PermissionItem**](PermissionItem.md) | Data required to check permissions on a resource. | 
 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 

### Return type

[**PermissionItem**](PermissionItem.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResourceInstanceEntitlements

> InstancePermissionItem GetResourceInstanceEntitlements(ctx, resourceInstanceName, resourceType, userName).AcceptVersion(acceptVersion).Execute()

Return details of the user's permissions for a resource instance.



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
	resourceInstanceName := "resourceInstanceName_example" // string | The identifier of the resource.
	resourceType := "resourceType_example" // string | The type of resource.
	userName := "userName_example" // string | The identifier of the user.
	acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthorizationAPI.GetResourceInstanceEntitlements(context.Background(), resourceInstanceName, resourceType, userName).AcceptVersion(acceptVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationAPI.GetResourceInstanceEntitlements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResourceInstanceEntitlements`: InstancePermissionItem
	fmt.Fprintf(os.Stdout, "Response from `AuthorizationAPI.GetResourceInstanceEntitlements`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceInstanceName** | **string** | The identifier of the resource. | 
**resourceType** | **string** | The type of resource. | 
**userName** | **string** | The identifier of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResourceInstanceEntitlementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 

### Return type

[**InstancePermissionItem**](InstancePermissionItem.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResourceTypeEntitlements

> ResourceTypePermissionItem GetResourceTypeEntitlements(ctx, resourceType, userName).AcceptVersion(acceptVersion).Execute()

Return details of the user's permissions for a resource type.



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
	resourceType := "resourceType_example" // string | The type of resource.
	userName := "userName_example" // string | The identifier of the user.
	acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthorizationAPI.GetResourceTypeEntitlements(context.Background(), resourceType, userName).AcceptVersion(acceptVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationAPI.GetResourceTypeEntitlements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResourceTypeEntitlements`: ResourceTypePermissionItem
	fmt.Fprintf(os.Stdout, "Response from `AuthorizationAPI.GetResourceTypeEntitlements`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceType** | **string** | The type of resource. | 
**userName** | **string** | The identifier of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResourceTypeEntitlementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 

### Return type

[**ResourceTypePermissionItem**](ResourceTypePermissionItem.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListResourceInstanceEntitlements

> InstancePermissionCollection ListResourceInstanceEntitlements(ctx, resourceType, userName).AcceptVersion(acceptVersion).Embed(embed).IfNoneMatch(ifNoneMatch).Limit(limit).Offset(offset).Execute()

List the user's permissions for all instances of a resource type.



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
	resourceType := "resourceType_example" // string | The type of resource.
	userName := "userName_example" // string | The identifier of the user.
	acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)
	embed := false // bool | Embedding: The whether or not to embed resources into the collection (rather than return links). (optional) (default to false)
	ifNoneMatch := "ifNoneMatch_example" // string | Caching: Optional header to improve performance. The value of this header should be the `ETag` of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content. (optional)
	limit := int32(20) // int32 | Paging: The maximum number of items to return in a resource. (optional) (default to 20)
	offset := int32(0) // int32 | Paging:  The index of the first item to return in the resource. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthorizationAPI.ListResourceInstanceEntitlements(context.Background(), resourceType, userName).AcceptVersion(acceptVersion).Embed(embed).IfNoneMatch(ifNoneMatch).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationAPI.ListResourceInstanceEntitlements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListResourceInstanceEntitlements`: InstancePermissionCollection
	fmt.Fprintf(os.Stdout, "Response from `AuthorizationAPI.ListResourceInstanceEntitlements`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceType** | **string** | The type of resource. | 
**userName** | **string** | The identifier of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListResourceInstanceEntitlementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 
 **embed** | **bool** | Embedding: The whether or not to embed resources into the collection (rather than return links). | [default to false]
 **ifNoneMatch** | **string** | Caching: Optional header to improve performance. The value of this header should be the &#x60;ETag&#x60; of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content. | 
 **limit** | **int32** | Paging: The maximum number of items to return in a resource. | [default to 20]
 **offset** | **int32** | Paging:  The index of the first item to return in the resource. | [default to 0]

### Return type

[**InstancePermissionCollection**](InstancePermissionCollection.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

