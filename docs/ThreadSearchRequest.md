# ThreadSearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to **map[string]interface{}** | Thread metadata to filter on. | [optional] 
**Values** | Pointer to **map[string]interface{}** | State values to filter on. | [optional] 
**Status** | Pointer to **string** | Thread status to filter on. | [optional] 
**Limit** | Pointer to **int32** | Maximum number to return. | [optional] [default to 10]
**Offset** | Pointer to **int32** | Offset to start from. | [optional] [default to 0]

## Methods

### NewThreadSearchRequest

`func NewThreadSearchRequest() *ThreadSearchRequest`

NewThreadSearchRequest instantiates a new ThreadSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreadSearchRequestWithDefaults

`func NewThreadSearchRequestWithDefaults() *ThreadSearchRequest`

NewThreadSearchRequestWithDefaults instantiates a new ThreadSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *ThreadSearchRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ThreadSearchRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ThreadSearchRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ThreadSearchRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetValues

`func (o *ThreadSearchRequest) GetValues() map[string]interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ThreadSearchRequest) GetValuesOk() (*map[string]interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ThreadSearchRequest) SetValues(v map[string]interface{})`

SetValues sets Values field to given value.

### HasValues

`func (o *ThreadSearchRequest) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetStatus

`func (o *ThreadSearchRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ThreadSearchRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ThreadSearchRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ThreadSearchRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLimit

`func (o *ThreadSearchRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ThreadSearchRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ThreadSearchRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ThreadSearchRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *ThreadSearchRequest) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ThreadSearchRequest) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ThreadSearchRequest) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ThreadSearchRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


