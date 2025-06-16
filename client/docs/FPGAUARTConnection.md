<!--
Copyright (C) 2020-2025 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# FPGAUARTConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaudRate** | **int32** | Baud rate for the UART connection | [default to 38400]
**TtyPort** | **string** | TTY port for the UART connection | 

## Methods

### NewFPGAUARTConnection

`func NewFPGAUARTConnection(baudRate int32, ttyPort string, ) *FPGAUARTConnection`

NewFPGAUARTConnection instantiates a new FPGAUARTConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFPGAUARTConnectionWithDefaults

`func NewFPGAUARTConnectionWithDefaults() *FPGAUARTConnection`

NewFPGAUARTConnectionWithDefaults instantiates a new FPGAUARTConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaudRate

`func (o *FPGAUARTConnection) GetBaudRate() int32`

GetBaudRate returns the BaudRate field if non-nil, zero value otherwise.

### GetBaudRateOk

`func (o *FPGAUARTConnection) GetBaudRateOk() (*int32, bool)`

GetBaudRateOk returns a tuple with the BaudRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaudRate

`func (o *FPGAUARTConnection) SetBaudRate(v int32)`

SetBaudRate sets BaudRate field to given value.


### GetTtyPort

`func (o *FPGAUARTConnection) GetTtyPort() string`

GetTtyPort returns the TtyPort field if non-nil, zero value otherwise.

### GetTtyPortOk

`func (o *FPGAUARTConnection) GetTtyPortOk() (*string, bool)`

GetTtyPortOk returns a tuple with the TtyPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtyPort

`func (o *FPGAUARTConnection) SetTtyPort(v string)`

SetTtyPort sets TtyPort field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


