<!--
Copyright (C) 2020-2025 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# \FPGAEntitlementAPI

All URIs are relative to *https://all.api.keil.arm.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RemoveFPGAEntitlement**](FPGAEntitlementAPI.md#RemoveFPGAEntitlement) | **Delete** /fpga-entitlements/{fpgaEntitlementName} | Remove all the permissions associated with an FPGA entitlement



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
	r, err := apiClient.FPGAEntitlementAPI.RemoveFPGAEntitlement(context.Background(), fpgaEntitlementName).AcceptVersion(acceptVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FPGAEntitlementAPI.RemoveFPGAEntitlement``: %v\n", err)
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

