<!--
Copyright (C) 2020-2025 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# \FPGAEntitlementsAPI

All URIs are relative to *https://all.api.keil.arm.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateFPGAEntitlement**](FPGAEntitlementsAPI.md#UpdateFPGAEntitlement) | **Put** /fpga-entitlements | Update an FPGA entitlement that defines which users are permitted to interact with an FPGA.



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
	fPGAEntitlementItem := *openapiclient.NewFPGAEntitlementItem("TODO", "Fpga_example", "samsung", []string{"Users_example"}) // FPGAEntitlementItem | 
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

