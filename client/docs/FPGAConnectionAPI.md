<!--
Copyright (C) 2020-2025 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# \FPGAConnectionAPI

All URIs are relative to *https://all.api.keil.arm.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFpgaConnection**](FPGAConnectionAPI.md#GetFpgaConnection) | **Get** /fpga-connections/{connectionName} | Get connection information
[**ListFpgaConnections**](FPGAConnectionAPI.md#ListFpgaConnections) | **Get** /fpga-connections/ | List available FPGA connections.
[**StartFpgaConnection**](FPGAConnectionAPI.md#StartFpgaConnection) | **Get** /fpga-connections/{connectionName}/connect | starts a websocket connection
[**TerminateFpgaConnection**](FPGAConnectionAPI.md#TerminateFpgaConnection) | **Post** /fpga-connections/{connectionName}/terminate | Terminates all websocket connections to the application running on the FPGA



## GetFpgaConnection

> FPGAConnectionItem GetFpgaConnection(ctx, connectionName).AcceptVersion(acceptVersion).Execute()

Get connection information



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
	connectionName := "connectionName_example" // string | The identifier of the connection
	acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FPGAConnectionAPI.GetFpgaConnection(context.Background(), connectionName).AcceptVersion(acceptVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FPGAConnectionAPI.GetFpgaConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFpgaConnection`: FPGAConnectionItem
	fmt.Fprintf(os.Stdout, "Response from `FPGAConnectionAPI.GetFpgaConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionName** | **string** | The identifier of the connection | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFpgaConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 

### Return type

[**FPGAConnectionItem**](FPGAConnectionItem.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFpgaConnections

> FPGAConnectionCollection ListFpgaConnections(ctx).AcceptVersion(acceptVersion).Embed(embed).IfNoneMatch(ifNoneMatch).Limit(limit).Offset(offset).Execute()

List available FPGA connections.



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
	resp, r, err := apiClient.FPGAConnectionAPI.ListFpgaConnections(context.Background()).AcceptVersion(acceptVersion).Embed(embed).IfNoneMatch(ifNoneMatch).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FPGAConnectionAPI.ListFpgaConnections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFpgaConnections`: FPGAConnectionCollection
	fmt.Fprintf(os.Stdout, "Response from `FPGAConnectionAPI.ListFpgaConnections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFpgaConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 
 **embed** | **bool** | Embedding: The whether or not to embed resources into the collection (rather than return links). | [default to false]
 **ifNoneMatch** | **string** | Caching: Optional header to improve performance. The value of this header should be the &#x60;ETag&#x60; of the resource when last read. If this is provided and there have been no changes to the resource then a 304 will be returned without content. | 
 **limit** | **int32** | Paging: The maximum number of items to return in a resource. | [default to 20]
 **offset** | **int32** | Paging:  The index of the first item to return in the resource. | [default to 0]

### Return type

[**FPGAConnectionCollection**](FPGAConnectionCollection.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartFpgaConnection

> *os.File StartFpgaConnection(ctx, connectionName).Upgrade(upgrade).Connection(connection).SecWebSocketProtocol(secWebSocketProtocol).AcceptVersion(acceptVersion).Execute()

starts a websocket connection



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
	connectionName := "connectionName_example" // string | The identifier of the connection
	upgrade := "upgrade_example" // string | Header used to upgrade an already-established client/server connection to a different protocol  (https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Upgrade). It is for instance used when establishing a Websocket connection (https://en.wikipedia.org/wiki/WebSocket#Opening_handshake) (default to "WebSocket")
	connection := "connection_example" // string | Header controlling whether the network connection (https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Connection). It is for instance used when establishing a Websocket connection (https://en.wikipedia.org/wiki/WebSocket#Opening_handshake) (default to "Upgrade")
	secWebSocketProtocol := "secWebSocketProtocol_example" // string | Header used to define a sub protocol to use in the communication  (https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Sec-WebSocket-Protocol). It is for instance used when establishing a Websocket connection (https://en.wikipedia.org/wiki/WebSocket#Opening_handshake) (optional)
	acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FPGAConnectionAPI.StartFpgaConnection(context.Background(), connectionName).Upgrade(upgrade).Connection(connection).SecWebSocketProtocol(secWebSocketProtocol).AcceptVersion(acceptVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FPGAConnectionAPI.StartFpgaConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartFpgaConnection`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `FPGAConnectionAPI.StartFpgaConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionName** | **string** | The identifier of the connection | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartFpgaConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upgrade** | **string** | Header used to upgrade an already-established client/server connection to a different protocol  (https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Upgrade). It is for instance used when establishing a Websocket connection (https://en.wikipedia.org/wiki/WebSocket#Opening_handshake) | [default to &quot;WebSocket&quot;]
 **connection** | **string** | Header controlling whether the network connection (https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Connection). It is for instance used when establishing a Websocket connection (https://en.wikipedia.org/wiki/WebSocket#Opening_handshake) | [default to &quot;Upgrade&quot;]
 **secWebSocketProtocol** | **string** | Header used to define a sub protocol to use in the communication  (https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Sec-WebSocket-Protocol). It is for instance used when establishing a Websocket connection (https://en.wikipedia.org/wiki/WebSocket#Opening_handshake) | 
 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TerminateFpgaConnection

> FPGAConnectionItem TerminateFpgaConnection(ctx, connectionName).AcceptVersion(acceptVersion).Execute()

Terminates all websocket connections to the application running on the FPGA



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
	connectionName := "connectionName_example" // string | The identifier of the connection
	acceptVersion := "1.0.0" // string | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FPGAConnectionAPI.TerminateFpgaConnection(context.Background(), connectionName).AcceptVersion(acceptVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FPGAConnectionAPI.TerminateFpgaConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TerminateFpgaConnection`: FPGAConnectionItem
	fmt.Fprintf(os.Stdout, "Response from `FPGAConnectionAPI.TerminateFpgaConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionName** | **string** | The identifier of the connection | 

### Other Parameters

Other parameters are passed through a pointer to a apiTerminateFpgaConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptVersion** | **string** | Versioning: Optional header to request a specific version of the API. While it is possible to specify a particular major, minor or patch version it is not recommended for production use cases. Only the major version number should be specified as minor and patch versions can be updated without warning. | 

### Return type

[**FPGAConnectionItem**](FPGAConnectionItem.md)

### Authorization

[TokenAuth](../README.md#TokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

