# VendorItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**VendorItemLinks**](VendorItemLinks.md) |  | 
**Metadata** | [**NullableCommonMetadata**](CommonMetadata.md) |  | 
**Description** | Pointer to **string** | Description of the Vendor | [optional] 
**ExternalId** | Pointer to **string** | Value used to identify this Vendor, in the data source used to create this Vendor. | [optional] 
**Id** | **string** | Unique ID of the Vendor. | [readonly] 
**Logo** | Pointer to **string** | URL to the Vendors logo image | [optional] 
**Slug** | **string** | Web-addressable slug of the Vendor name | [readonly] 
**Title** | **string** | Human-readable name of the Vendor. | 
**Type** | Pointer to **string** | Type of data used to populate the &#x60;external_id&#x60; field | [optional] 
**Website** | Pointer to **string** | URL to the Vendors website | [optional] 

## Methods

### NewVendorItem

`func NewVendorItem(links VendorItemLinks, metadata NullableCommonMetadata, id string, slug string, title string, ) *VendorItem`

NewVendorItem instantiates a new VendorItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVendorItemWithDefaults

`func NewVendorItemWithDefaults() *VendorItem`

NewVendorItemWithDefaults instantiates a new VendorItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *VendorItem) GetLinks() VendorItemLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *VendorItem) GetLinksOk() (*VendorItemLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *VendorItem) SetLinks(v VendorItemLinks)`

SetLinks sets Links field to given value.


### GetMetadata

`func (o *VendorItem) GetMetadata() CommonMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *VendorItem) GetMetadataOk() (*CommonMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *VendorItem) SetMetadata(v CommonMetadata)`

SetMetadata sets Metadata field to given value.


### SetMetadataNil

`func (o *VendorItem) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *VendorItem) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetDescription

`func (o *VendorItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VendorItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VendorItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VendorItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalId

`func (o *VendorItem) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *VendorItem) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *VendorItem) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *VendorItem) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetId

`func (o *VendorItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VendorItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VendorItem) SetId(v string)`

SetId sets Id field to given value.


### GetLogo

`func (o *VendorItem) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *VendorItem) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *VendorItem) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *VendorItem) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetSlug

`func (o *VendorItem) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *VendorItem) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *VendorItem) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetTitle

`func (o *VendorItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *VendorItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *VendorItem) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetType

`func (o *VendorItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VendorItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VendorItem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VendorItem) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWebsite

`func (o *VendorItem) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *VendorItem) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *VendorItem) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *VendorItem) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


