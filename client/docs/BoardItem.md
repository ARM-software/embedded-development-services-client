# BoardItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**BoardItemLinks**](BoardItemLinks.md) |  | 
**Metadata** | [**NullableCommonMetadata**](CommonMetadata.md) |  | 
**DebugInterfaces** | Pointer to [**[]ADebugInterface**](ADebugInterface.md) | Array of debug interfaces for the Board. | [optional] 
**Description** | **string** | Description of the Board. | 
**DetectCode** | Pointer to **NullableString** | Mbed detection code for debugging. | [optional] 
**Features** | Pointer to [**[]BoardItemFeaturesInner**](BoardItemFeaturesInner.md) | Array of features for the Board. | [optional] 
**Id** | **string** | Unique ID of the Board. | [readonly] 
**Revision** | **string** | Revision of the Board. | 
**Summary** | **string** | Brief summary of the Board. | 
**Title** | **string** | Human-readable name of the Board. | 

## Methods

### NewBoardItem

`func NewBoardItem(links BoardItemLinks, metadata NullableCommonMetadata, description string, id string, revision string, summary string, title string, ) *BoardItem`

NewBoardItem instantiates a new BoardItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoardItemWithDefaults

`func NewBoardItemWithDefaults() *BoardItem`

NewBoardItemWithDefaults instantiates a new BoardItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *BoardItem) GetLinks() BoardItemLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BoardItem) GetLinksOk() (*BoardItemLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BoardItem) SetLinks(v BoardItemLinks)`

SetLinks sets Links field to given value.


### GetMetadata

`func (o *BoardItem) GetMetadata() CommonMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BoardItem) GetMetadataOk() (*CommonMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BoardItem) SetMetadata(v CommonMetadata)`

SetMetadata sets Metadata field to given value.


### SetMetadataNil

`func (o *BoardItem) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *BoardItem) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetDebugInterfaces

`func (o *BoardItem) GetDebugInterfaces() []ADebugInterface`

GetDebugInterfaces returns the DebugInterfaces field if non-nil, zero value otherwise.

### GetDebugInterfacesOk

`func (o *BoardItem) GetDebugInterfacesOk() (*[]ADebugInterface, bool)`

GetDebugInterfacesOk returns a tuple with the DebugInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugInterfaces

`func (o *BoardItem) SetDebugInterfaces(v []ADebugInterface)`

SetDebugInterfaces sets DebugInterfaces field to given value.

### HasDebugInterfaces

`func (o *BoardItem) HasDebugInterfaces() bool`

HasDebugInterfaces returns a boolean if a field has been set.

### SetDebugInterfacesNil

`func (o *BoardItem) SetDebugInterfacesNil(b bool)`

 SetDebugInterfacesNil sets the value for DebugInterfaces to be an explicit nil

### UnsetDebugInterfaces
`func (o *BoardItem) UnsetDebugInterfaces()`

UnsetDebugInterfaces ensures that no value is present for DebugInterfaces, not even an explicit nil
### GetDescription

`func (o *BoardItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BoardItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BoardItem) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDetectCode

`func (o *BoardItem) GetDetectCode() string`

GetDetectCode returns the DetectCode field if non-nil, zero value otherwise.

### GetDetectCodeOk

`func (o *BoardItem) GetDetectCodeOk() (*string, bool)`

GetDetectCodeOk returns a tuple with the DetectCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectCode

`func (o *BoardItem) SetDetectCode(v string)`

SetDetectCode sets DetectCode field to given value.

### HasDetectCode

`func (o *BoardItem) HasDetectCode() bool`

HasDetectCode returns a boolean if a field has been set.

### SetDetectCodeNil

`func (o *BoardItem) SetDetectCodeNil(b bool)`

 SetDetectCodeNil sets the value for DetectCode to be an explicit nil

### UnsetDetectCode
`func (o *BoardItem) UnsetDetectCode()`

UnsetDetectCode ensures that no value is present for DetectCode, not even an explicit nil
### GetFeatures

`func (o *BoardItem) GetFeatures() []BoardItemFeaturesInner`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *BoardItem) GetFeaturesOk() (*[]BoardItemFeaturesInner, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *BoardItem) SetFeatures(v []BoardItemFeaturesInner)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *BoardItem) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### SetFeaturesNil

`func (o *BoardItem) SetFeaturesNil(b bool)`

 SetFeaturesNil sets the value for Features to be an explicit nil

### UnsetFeatures
`func (o *BoardItem) UnsetFeatures()`

UnsetFeatures ensures that no value is present for Features, not even an explicit nil
### GetId

`func (o *BoardItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BoardItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BoardItem) SetId(v string)`

SetId sets Id field to given value.


### GetRevision

`func (o *BoardItem) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *BoardItem) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *BoardItem) SetRevision(v string)`

SetRevision sets Revision field to given value.


### GetSummary

`func (o *BoardItem) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *BoardItem) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *BoardItem) SetSummary(v string)`

SetSummary sets Summary field to given value.


### GetTitle

`func (o *BoardItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BoardItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BoardItem) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


