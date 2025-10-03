<!--
Copyright (C) 2020-2025 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# \ServiceAccountAPI

All URIs are relative to *https://all.api.keil.arm.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServiceAccount**](ServiceAccountAPI.md#CreateServiceAccount) | **Post** /service-accounts/ | Create a service account.
[**DeleteServiceAccount**](ServiceAccountAPI.md#DeleteServiceAccount) | **Delete** /service-accounts/{serviceAccountName} | Delete a service account
[**EditServiceAccount**](ServiceAccountAPI.md#EditServiceAccount) | **Put** /service-accounts/{serviceAccountName} | Update a service account
[**GetServiceAccount**](ServiceAccountAPI.md#GetServiceAccount) | **Get** /service-accounts/{serviceAccountName} | Get a service account
[**ListServiceAccounts**](ServiceAccountAPI.md#ListServiceAccounts) | **Get** /service-accounts/ | List all service accounts



## CreateServiceAccount

> ServiceAccountItem CreateServiceAccount(ctx).ServiceAccountItem(serviceAccountItem).AcceptVersion(acceptVersion).Execute()

Create a service account.



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
	serviceAccountItem := *openapiclient.NewServiceAccountItem(*openapiclient.NewServiceAccountItemLinks(*openapiclient.NewHalLinkData("/endpoint/?limit=20&offset=0")), "TODO", "CreatedBy_example", "88a6137e-1d99-4cde-8db8-015312f7d5e6") // ServiceAccountItem | Data required to create a new service account.
	acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceAccountAPI.CreateServiceAccount(context.Background()).ServiceAccountItem(serviceAccountItem).AcceptVersion(acceptVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountAPI.CreateServiceAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateServiceAccount`: ServiceAccountItem
	fmt.Fprintf(os.Stdout, "Response from `ServiceAccountAPI.CreateServiceAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateServiceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceAccountItem** | [**ServiceAccountItem**](ServiceAccountItem.md) | Data required to create a new service account. | 
 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 

### Return type

[**ServiceAccountItem**](ServiceAccountItem.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteServiceAccount

> DeleteServiceAccount(ctx, serviceAccountName).AcceptVersion(acceptVersion).Execute()

Delete a service account



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
	serviceAccountName := "serviceAccountName_example" // string | Unique ID of the service account.
	acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServiceAccountAPI.DeleteServiceAccount(context.Background(), serviceAccountName).AcceptVersion(acceptVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountAPI.DeleteServiceAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceAccountName** | **string** | Unique ID of the service account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServiceAccountRequest struct via the builder pattern


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


## EditServiceAccount

> ServiceAccountItem EditServiceAccount(ctx, serviceAccountName).IfMatch(ifMatch).ServiceAccountItem(serviceAccountItem).AcceptVersion(acceptVersion).Execute()

Update a service account



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
	serviceAccountName := "serviceAccountName_example" // string | Unique ID of the service account.
	ifMatch := "ifMatch_example" // string | Conditional Requests: This is required in order to perform an update of a resource. The value of this header should be the `ETag` of the resource when read (before being subsequently modified by the client).
	serviceAccountItem := *openapiclient.NewServiceAccountItem(*openapiclient.NewServiceAccountItemLinks(*openapiclient.NewHalLinkData("/endpoint/?limit=20&offset=0")), "TODO", "CreatedBy_example", "88a6137e-1d99-4cde-8db8-015312f7d5e6") // ServiceAccountItem | The new definition of the service account.
	acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceAccountAPI.EditServiceAccount(context.Background(), serviceAccountName).IfMatch(ifMatch).ServiceAccountItem(serviceAccountItem).AcceptVersion(acceptVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountAPI.EditServiceAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditServiceAccount`: ServiceAccountItem
	fmt.Fprintf(os.Stdout, "Response from `ServiceAccountAPI.EditServiceAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceAccountName** | **string** | Unique ID of the service account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditServiceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **string** | Conditional Requests: This is required in order to perform an update of a resource. The value of this header should be the &#x60;ETag&#x60; of the resource when read (before being subsequently modified by the client). | 
 **serviceAccountItem** | [**ServiceAccountItem**](ServiceAccountItem.md) | The new definition of the service account. | 
 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 

### Return type

[**ServiceAccountItem**](ServiceAccountItem.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceAccount

> ServiceAccountItem GetServiceAccount(ctx, serviceAccountName).AcceptVersion(acceptVersion).Execute()

Get a service account



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
	serviceAccountName := "serviceAccountName_example" // string | Unique ID of the service account.
	acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceAccountAPI.GetServiceAccount(context.Background(), serviceAccountName).AcceptVersion(acceptVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountAPI.GetServiceAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceAccount`: ServiceAccountItem
	fmt.Fprintf(os.Stdout, "Response from `ServiceAccountAPI.GetServiceAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceAccountName** | **string** | Unique ID of the service account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 

### Return type

[**ServiceAccountItem**](ServiceAccountItem.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServiceAccounts

> ServiceAccountCollection ListServiceAccounts(ctx).AcceptVersion(acceptVersion).Embed(embed).IfNoneMatch(ifNoneMatch).Limit(limit).Offset(offset).Execute()

List all service accounts



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
	resp, r, err := apiClient.ServiceAccountAPI.ListServiceAccounts(context.Background()).AcceptVersion(acceptVersion).Embed(embed).IfNoneMatch(ifNoneMatch).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountAPI.ListServiceAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListServiceAccounts`: ServiceAccountCollection
	fmt.Fprintf(os.Stdout, "Response from `ServiceAccountAPI.ListServiceAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListServiceAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 
 **embed** | **bool** | Embedding: The whether or not to embed resources into the collection (rather than return links). | [default to false]
 **ifNoneMatch** | **string** | Caching: Optional header to improve performance. The value of this header should be the &#x60;ETag&#x60; of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content. | 
 **limit** | **int32** | Paging: The maximum number of items to return in a resource. | [default to 20]
 **offset** | **int32** | Paging:  The index of the first item to return in the resource. | [default to 0]

### Return type

[**ServiceAccountCollection**](ServiceAccountCollection.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

