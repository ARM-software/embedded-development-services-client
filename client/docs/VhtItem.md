# VhtItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**NullableVhtItemLinks**](VhtItemLinks.md) |  | 
**Metadata** | [**NullableCommonMetadata**](CommonMetadata.md) |  | 
**AdditionalTools** | Pointer to [**[]AdditionalTool**](AdditionalTool.md) |  | [optional] 
**Board** | **string** | The board Target. | [readonly] 
**BoardRevision** | Pointer to **string** | The revision of the board Target. | [optional] [readonly] 
**Deprecated** | **bool** | True if this VHT is scheduled to be removed from the service. | [readonly] 
**DeprecationInfo** | Pointer to [**DeprecationInfo**](DeprecationInfo.md) |  | [optional] 
**Model** | Pointer to **NullableString** | The model used | [optional] [readonly] 
**Name** | **string** | Unique ID of the VHT. | [readonly] 
**Title** | **string** | Human readable name of the VHT. | [readonly] 
**Vendor** | **string** | The vendor who supplies the VHT. | [readonly] 
**VirtualInterfaces** | Pointer to [**[]VirtualInterface**](VirtualInterface.md) |  | [optional] 

## Methods

### NewVhtItem

`func NewVhtItem(links NullableVhtItemLinks, metadata NullableCommonMetadata, board string, deprecated bool, name string, title string, vendor string, ) *VhtItem`

NewVhtItem instantiates a new VhtItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVhtItemWithDefaults

`func NewVhtItemWithDefaults() *VhtItem`

NewVhtItemWithDefaults instantiates a new VhtItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *VhtItem) GetLinks() VhtItemLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *VhtItem) GetLinksOk() (*VhtItemLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *VhtItem) SetLinks(v VhtItemLinks)`

SetLinks sets Links field to given value.


### SetLinksNil

`func (o *VhtItem) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *VhtItem) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetMetadata

`func (o *VhtItem) GetMetadata() CommonMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *VhtItem) GetMetadataOk() (*CommonMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *VhtItem) SetMetadata(v CommonMetadata)`

SetMetadata sets Metadata field to given value.


### SetMetadataNil

`func (o *VhtItem) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *VhtItem) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetAdditionalTools

`func (o *VhtItem) GetAdditionalTools() []AdditionalTool`

GetAdditionalTools returns the AdditionalTools field if non-nil, zero value otherwise.

### GetAdditionalToolsOk

`func (o *VhtItem) GetAdditionalToolsOk() (*[]AdditionalTool, bool)`

GetAdditionalToolsOk returns a tuple with the AdditionalTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalTools

`func (o *VhtItem) SetAdditionalTools(v []AdditionalTool)`

SetAdditionalTools sets AdditionalTools field to given value.

### HasAdditionalTools

`func (o *VhtItem) HasAdditionalTools() bool`

HasAdditionalTools returns a boolean if a field has been set.

### GetBoard

`func (o *VhtItem) GetBoard() string`

GetBoard returns the Board field if non-nil, zero value otherwise.

### GetBoardOk

`func (o *VhtItem) GetBoardOk() (*string, bool)`

GetBoardOk returns a tuple with the Board field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoard

`func (o *VhtItem) SetBoard(v string)`

SetBoard sets Board field to given value.


### GetBoardRevision

`func (o *VhtItem) GetBoardRevision() string`

GetBoardRevision returns the BoardRevision field if non-nil, zero value otherwise.

### GetBoardRevisionOk

`func (o *VhtItem) GetBoardRevisionOk() (*string, bool)`

GetBoardRevisionOk returns a tuple with the BoardRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoardRevision

`func (o *VhtItem) SetBoardRevision(v string)`

SetBoardRevision sets BoardRevision field to given value.

### HasBoardRevision

`func (o *VhtItem) HasBoardRevision() bool`

HasBoardRevision returns a boolean if a field has been set.

### GetDeprecated

`func (o *VhtItem) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *VhtItem) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *VhtItem) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.


### GetDeprecationInfo

`func (o *VhtItem) GetDeprecationInfo() DeprecationInfo`

GetDeprecationInfo returns the DeprecationInfo field if non-nil, zero value otherwise.

### GetDeprecationInfoOk

`func (o *VhtItem) GetDeprecationInfoOk() (*DeprecationInfo, bool)`

GetDeprecationInfoOk returns a tuple with the DeprecationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecationInfo

`func (o *VhtItem) SetDeprecationInfo(v DeprecationInfo)`

SetDeprecationInfo sets DeprecationInfo field to given value.

### HasDeprecationInfo

`func (o *VhtItem) HasDeprecationInfo() bool`

HasDeprecationInfo returns a boolean if a field has been set.

### GetModel

`func (o *VhtItem) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *VhtItem) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *VhtItem) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *VhtItem) HasModel() bool`

HasModel returns a boolean if a field has been set.

### SetModelNil

`func (o *VhtItem) SetModelNil(b bool)`

 SetModelNil sets the value for Model to be an explicit nil

### UnsetModel
`func (o *VhtItem) UnsetModel()`

UnsetModel ensures that no value is present for Model, not even an explicit nil
### GetName

`func (o *VhtItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VhtItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VhtItem) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *VhtItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *VhtItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *VhtItem) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetVendor

`func (o *VhtItem) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *VhtItem) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *VhtItem) SetVendor(v string)`

SetVendor sets Vendor field to given value.


### GetVirtualInterfaces

`func (o *VhtItem) GetVirtualInterfaces() []VirtualInterface`

GetVirtualInterfaces returns the VirtualInterfaces field if non-nil, zero value otherwise.

### GetVirtualInterfacesOk

`func (o *VhtItem) GetVirtualInterfacesOk() (*[]VirtualInterface, bool)`

GetVirtualInterfacesOk returns a tuple with the VirtualInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualInterfaces

`func (o *VhtItem) SetVirtualInterfaces(v []VirtualInterface)`

SetVirtualInterfaces sets VirtualInterfaces field to given value.

### HasVirtualInterfaces

`func (o *VhtItem) HasVirtualInterfaces() bool`

HasVirtualInterfaces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


