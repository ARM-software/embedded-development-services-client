# DeviceItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**DeviceItemLinks**](DeviceItemLinks.md) |  | 
**Metadata** | [**NullableCommonMetadata**](CommonMetadata.md) |  | 
**Family** | Pointer to **NullableString** | Device Family the Device belongs to. | [optional] 
**Processors** | [**[]ProcessorItem**](ProcessorItem.md) | Array of processors within the Device. | 
**Slug** | **string** | Unique human-readable identifier for the Device | [readonly] 
**SourcePackId** | [**DeviceItemSourcePackId**](DeviceItemSourcePackId.md) |  | 
**SubFamily** | Pointer to **NullableString** | Sub-Device Family the Device belongs to. | [optional] 
**Title** | **string** | Human-readable name of the Device. | 
**Vendor** | **string** | Vendor of the Device. | 

## Methods

### NewDeviceItem

`func NewDeviceItem(links DeviceItemLinks, metadata NullableCommonMetadata, processors []ProcessorItem, slug string, sourcePackId DeviceItemSourcePackId, title string, vendor string, ) *DeviceItem`

NewDeviceItem instantiates a new DeviceItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceItemWithDefaults

`func NewDeviceItemWithDefaults() *DeviceItem`

NewDeviceItemWithDefaults instantiates a new DeviceItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *DeviceItem) GetLinks() DeviceItemLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DeviceItem) GetLinksOk() (*DeviceItemLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DeviceItem) SetLinks(v DeviceItemLinks)`

SetLinks sets Links field to given value.


### GetMetadata

`func (o *DeviceItem) GetMetadata() CommonMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DeviceItem) GetMetadataOk() (*CommonMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DeviceItem) SetMetadata(v CommonMetadata)`

SetMetadata sets Metadata field to given value.


### SetMetadataNil

`func (o *DeviceItem) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *DeviceItem) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
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

### SetFamilyNil

`func (o *DeviceItem) SetFamilyNil(b bool)`

 SetFamilyNil sets the value for Family to be an explicit nil

### UnsetFamily
`func (o *DeviceItem) UnsetFamily()`

UnsetFamily ensures that no value is present for Family, not even an explicit nil
### GetProcessors

`func (o *DeviceItem) GetProcessors() []ProcessorItem`

GetProcessors returns the Processors field if non-nil, zero value otherwise.

### GetProcessorsOk

`func (o *DeviceItem) GetProcessorsOk() (*[]ProcessorItem, bool)`

GetProcessorsOk returns a tuple with the Processors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessors

`func (o *DeviceItem) SetProcessors(v []ProcessorItem)`

SetProcessors sets Processors field to given value.


### GetSlug

`func (o *DeviceItem) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *DeviceItem) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *DeviceItem) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetSourcePackId

`func (o *DeviceItem) GetSourcePackId() DeviceItemSourcePackId`

GetSourcePackId returns the SourcePackId field if non-nil, zero value otherwise.

### GetSourcePackIdOk

`func (o *DeviceItem) GetSourcePackIdOk() (*DeviceItemSourcePackId, bool)`

GetSourcePackIdOk returns a tuple with the SourcePackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePackId

`func (o *DeviceItem) SetSourcePackId(v DeviceItemSourcePackId)`

SetSourcePackId sets SourcePackId field to given value.


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

### SetSubFamilyNil

`func (o *DeviceItem) SetSubFamilyNil(b bool)`

 SetSubFamilyNil sets the value for SubFamily to be an explicit nil

### UnsetSubFamily
`func (o *DeviceItem) UnsetSubFamily()`

UnsetSubFamily ensures that no value is present for SubFamily, not even an explicit nil
### GetTitle

`func (o *DeviceItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DeviceItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DeviceItem) SetTitle(v string)`

SetTitle sets Title field to given value.


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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


