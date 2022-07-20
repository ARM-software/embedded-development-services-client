# DeviceItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Family** | Pointer to **string** | Device Family the Device belonds to. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the Device. | [optional] [readonly] 
**Processors** | Pointer to [**[]DeviceItemProcessorsInner**](DeviceItemProcessorsInner.md) | Array of processors within the Device. | [optional] [readonly] 
**SubFamily** | Pointer to **string** | Sub-Device Family the Device belonds to. | [optional] [readonly] 
**Vendor** | Pointer to **string** | Vendor of the Device. | [optional] [readonly] 

## Methods

### NewDeviceItem

`func NewDeviceItem() *DeviceItem`

NewDeviceItem instantiates a new DeviceItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceItemWithDefaults

`func NewDeviceItemWithDefaults() *DeviceItem`

NewDeviceItemWithDefaults instantiates a new DeviceItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFamily

`func (o *DeviceItem) GetFamily() string`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *DeviceItem) GetFamilyOk() (*string, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *DeviceItem) SetFamily(v string)`

SetFamily sets Family field to given value.

### HasFamily

`func (o *DeviceItem) HasFamily() bool`

HasFamily returns a boolean if a field has been set.

### GetName

`func (o *DeviceItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeviceItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProcessors

`func (o *DeviceItem) GetProcessors() []DeviceItemProcessorsInner`

GetProcessors returns the Processors field if non-nil, zero value otherwise.

### GetProcessorsOk

`func (o *DeviceItem) GetProcessorsOk() (*[]DeviceItemProcessorsInner, bool)`

GetProcessorsOk returns a tuple with the Processors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessors

`func (o *DeviceItem) SetProcessors(v []DeviceItemProcessorsInner)`

SetProcessors sets Processors field to given value.

### HasProcessors

`func (o *DeviceItem) HasProcessors() bool`

HasProcessors returns a boolean if a field has been set.

### GetSubFamily

`func (o *DeviceItem) GetSubFamily() string`

GetSubFamily returns the SubFamily field if non-nil, zero value otherwise.

### GetSubFamilyOk

`func (o *DeviceItem) GetSubFamilyOk() (*string, bool)`

GetSubFamilyOk returns a tuple with the SubFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubFamily

`func (o *DeviceItem) SetSubFamily(v string)`

SetSubFamily sets SubFamily field to given value.

### HasSubFamily

`func (o *DeviceItem) HasSubFamily() bool`

HasSubFamily returns a boolean if a field has been set.

### GetVendor

`func (o *DeviceItem) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *DeviceItem) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *DeviceItem) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *DeviceItem) HasVendor() bool`

HasVendor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


