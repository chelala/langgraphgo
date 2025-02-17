# AssistantVersionsSearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to **map[string]interface{}** | Metadata to filter versions by. Exact match filter for each KV pair. | [optional] 
**Limit** | Pointer to **int32** | The maximum number of versions to return. | [optional] [default to 10]
**Offset** | Pointer to **int32** | The number of versions to skip. | [optional] [default to 0]

## Methods

### NewAssistantVersionsSearchRequest

`func NewAssistantVersionsSearchRequest() *AssistantVersionsSearchRequest`

NewAssistantVersionsSearchRequest instantiates a new AssistantVersionsSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssistantVersionsSearchRequestWithDefaults

`func NewAssistantVersionsSearchRequestWithDefaults() *AssistantVersionsSearchRequest`

NewAssistantVersionsSearchRequestWithDefaults instantiates a new AssistantVersionsSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *AssistantVersionsSearchRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AssistantVersionsSearchRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AssistantVersionsSearchRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AssistantVersionsSearchRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetLimit

`func (o *AssistantVersionsSearchRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *AssistantVersionsSearchRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *AssistantVersionsSearchRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *AssistantVersionsSearchRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *AssistantVersionsSearchRequest) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *AssistantVersionsSearchRequest) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *AssistantVersionsSearchRequest) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *AssistantVersionsSearchRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


