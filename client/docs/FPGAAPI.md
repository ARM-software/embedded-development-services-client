<!--
Copyright (C) 2020-2025 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# \FPGAAPI

All URIs are relative to *https://all.api.keil.arm.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SunsetFpga**](FPGAAPI.md#SunsetFpga) | **Post** /fpgas/{fpgaName}/sunset | Initiate the sunsetting of the FPGA as part of a decommission



## SunsetFpga

> FPGAJobItem SunsetFpga(ctx, fpgaName).AcceptVersion(acceptVersion).Execute()

Initiate the sunsetting of the FPGA as part of a decommission



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
	fpgaName := "fpgaName_example" // string | Unique ID of an FPGA.
	acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FPGAAPI.SunsetFpga(context.Background(), fpgaName).AcceptVersion(acceptVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FPGAAPI.SunsetFpga``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SunsetFpga`: FPGAJobItem
	fmt.Fprintf(os.Stdout, "Response from `FPGAAPI.SunsetFpga`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fpgaName** | **string** | Unique ID of an FPGA. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSunsetFpgaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 

### Return type

[**FPGAJobItem**](FPGAJobItem.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

