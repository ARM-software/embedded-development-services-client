<!--
Copyright (C) 2020-2025 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# \AuthorizationAPI

All URIs are relative to *https://all.api.keil.arm.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckPermission**](AuthorizationAPI.md#CheckPermission) | **Post** /entitlements/check | Check permissions on a resource.



## CheckPermission

> PermissionItem CheckPermission(ctx).PermissionItem(permissionItem).AcceptVersion(acceptVersion).Execute()

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
	permissionItem := *openapiclient.NewPermissionItem([]string{"Operations_example"}, "fpga-connection", "worker_1234-job-5678", "9c2d1b2f-3d4e-456a-a7fc-12b5d9c6e8a4", "GenericWorkJobItem", "b7f6c5d1-2a44-4f9a-9d77-6e5a8c9d4e1b", "arm_pat_HGfmzxFblhr6AXiTnyaUqd1pgA9waIiO") // PermissionItem | Data required to check permissions on a resource.
	acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthorizationAPI.CheckPermission(context.Background()).PermissionItem(permissionItem).AcceptVersion(acceptVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationAPI.CheckPermission``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckPermission`: PermissionItem
	fmt.Fprintf(os.Stdout, "Response from `AuthorizationAPI.CheckPermission`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckPermissionRequest struct via the builder pattern


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

