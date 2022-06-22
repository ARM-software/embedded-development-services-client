# PagingMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | Paging metadata: The number of items returned in this message. | [readonly] 
**Ctime** | **time.Time** | Creation Time: UTC date and time (in RFC3339 format) when the resource was created. If this is a system created resource, this will be a fixed time unaffected by user actions. | [readonly] 
**Etime** | Pointer to **NullableTime** | Expiry Time: UTC date and time (in RFC3339 format) when the resource will be removed automatically by the system and become unavailable. | [optional] [readonly] 
**Limit** | **int32** | Paging metadata: The limit on the number of items to return. | [readonly] 
**Mtime** | **time.Time** | Last Modification Time: UTC date and time (in RFC3339 format) when the resource was last updated. For a resource that cannot be modified this will be the same as &#x60;ctime&#x60;. | [readonly] 
**Offset** | **int32** | Paging metadata: The index of the first item returned. | [readonly] 
**Total** | **int32** | Paging metadata: Total number of items that can be paged through. | [readonly] 

## Methods

### NewPagingMetadata

`func NewPagingMetadata(count int32, ctime time.Time, limit int32, mtime time.Time, offset int32, total int32, ) *PagingMetadata`

NewPagingMetadata instantiates a new PagingMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPagingMetadataWithDefaults

`func NewPagingMetadataWithDefaults() *PagingMetadata`

NewPagingMetadataWithDefaults instantiates a new PagingMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PagingMetadata) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PagingMetadata) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PagingMetadata) SetCount(v int32)`

SetCount sets Count field to given value.


### GetCtime

`func (o *PagingMetadata) GetCtime() time.Time`

GetCtime returns the Ctime field if non-nil, zero value otherwise.

### GetCtimeOk

`func (o *PagingMetadata) GetCtimeOk() (*time.Time, bool)`

GetCtimeOk returns a tuple with the Ctime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCtime

`func (o *PagingMetadata) SetCtime(v time.Time)`

SetCtime sets Ctime field to given value.


### GetEtime

`func (o *PagingMetadata) GetEtime() time.Time`

GetEtime returns the Etime field if non-nil, zero value otherwise.

### GetEtimeOk

`func (o *PagingMetadata) GetEtimeOk() (*time.Time, bool)`

GetEtimeOk returns a tuple with the Etime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtime

`func (o *PagingMetadata) SetEtime(v time.Time)`

SetEtime sets Etime field to given value.

### HasEtime

`func (o *PagingMetadata) HasEtime() bool`

HasEtime returns a boolean if a field has been set.

### SetEtimeNil

`func (o *PagingMetadata) SetEtimeNil(b bool)`

 SetEtimeNil sets the value for Etime to be an explicit nil

### UnsetEtime
`func (o *PagingMetadata) UnsetEtime()`

UnsetEtime ensures that no value is present for Etime, not even an explicit nil
### GetLimit

`func (o *PagingMetadata) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *PagingMetadata) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *PagingMetadata) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetMtime

`func (o *PagingMetadata) GetMtime() time.Time`

GetMtime returns the Mtime field if non-nil, zero value otherwise.

### GetMtimeOk

`func (o *PagingMetadata) GetMtimeOk() (*time.Time, bool)`

GetMtimeOk returns a tuple with the Mtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtime

`func (o *PagingMetadata) SetMtime(v time.Time)`

SetMtime sets Mtime field to given value.


### GetOffset

`func (o *PagingMetadata) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *PagingMetadata) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *PagingMetadata) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetTotal

`func (o *PagingMetadata) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PagingMetadata) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PagingMetadata) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


