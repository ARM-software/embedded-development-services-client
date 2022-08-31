# BoardItemLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collection** | [**HalLinkData**](HalLinkData.md) |  | 
**Curies** | Pointer to [**[]HalLinkData**](HalLinkData.md) | Links to external services. | [optional] [readonly] 
**Device** | Pointer to [**[]HalLinkData**](HalLinkData.md) | Links to Devices mounted on the Board. | [optional] 
**Self** | [**HalLinkData**](HalLinkData.md) |  | 

## Methods

### NewBoardItemLinks

`func NewBoardItemLinks(collection HalLinkData, self HalLinkData, ) *BoardItemLinks`

NewBoardItemLinks instantiates a new BoardItemLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoardItemLinksWithDefaults

`func NewBoardItemLinksWithDefaults() *BoardItemLinks`

NewBoardItemLinksWithDefaults instantiates a new BoardItemLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollection

`func (o *BoardItemLinks) GetCollection() HalLinkData`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *BoardItemLinks) GetCollectionOk() (*HalLinkData, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *BoardItemLinks) SetCollection(v HalLinkData)`

SetCollection sets Collection field to given value.


### GetCuries

`func (o *BoardItemLinks) GetCuries() []HalLinkData`

GetCuries returns the Curies field if non-nil, zero value otherwise.

### GetCuriesOk

`func (o *BoardItemLinks) GetCuriesOk() (*[]HalLinkData, bool)`

GetCuriesOk returns a tuple with the Curies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCuries

`func (o *BoardItemLinks) SetCuries(v []HalLinkData)`

SetCuries sets Curies field to given value.

### HasCuries

`func (o *BoardItemLinks) HasCuries() bool`

HasCuries returns a boolean if a field has been set.

### SetCuriesNil

`func (o *BoardItemLinks) SetCuriesNil(b bool)`

 SetCuriesNil sets the value for Curies to be an explicit nil

### UnsetCuries
`func (o *BoardItemLinks) UnsetCuries()`

UnsetCuries ensures that no value is present for Curies, not even an explicit nil
### GetDevice

`func (o *BoardItemLinks) GetDevice() []HalLinkData`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *BoardItemLinks) GetDeviceOk() (*[]HalLinkData, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *BoardItemLinks) SetDevice(v []HalLinkData)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *BoardItemLinks) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetSelf

`func (o *BoardItemLinks) GetSelf() HalLinkData`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *BoardItemLinks) GetSelfOk() (*HalLinkData, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *BoardItemLinks) SetSelf(v HalLinkData)`

SetSelf sets Self field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


