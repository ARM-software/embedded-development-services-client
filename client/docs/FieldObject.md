<!--
Copyright (C) 2020-2023 Arm Limited or its affiliates and Contributors. All rights reserved.
SPDX-License-Identifier: Apache-2.0
-->
# FieldObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldName** | **string** | Name of the field name in the request which has failed validation. | [readonly] 
**FieldPath** | Pointer to **string** | Field name, possibly including the path of the field which caused the error. | [optional] [readonly] 
**Message** | **string** | A human readable message, which should provide an explanation and possible corrective actions. | [readonly] 

## Methods

### NewFieldObject

`func NewFieldObject(fieldName string, message string, ) *FieldObject`

NewFieldObject instantiates a new FieldObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldObjectWithDefaults

`func NewFieldObjectWithDefaults() *FieldObject`

NewFieldObjectWithDefaults instantiates a new FieldObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldName

`func (o *FieldObject) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *FieldObject) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *FieldObject) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.


### GetFieldPath

`func (o *FieldObject) GetFieldPath() string`

GetFieldPath returns the FieldPath field if non-nil, zero value otherwise.

### GetFieldPathOk

`func (o *FieldObject) GetFieldPathOk() (*string, bool)`

GetFieldPathOk returns a tuple with the FieldPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldPath

`func (o *FieldObject) SetFieldPath(v string)`

SetFieldPath sets FieldPath field to given value.

### HasFieldPath

`func (o *FieldObject) HasFieldPath() bool`

HasFieldPath returns a boolean if a field has been set.

### GetMessage

`func (o *FieldObject) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *FieldObject) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *FieldObject) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


