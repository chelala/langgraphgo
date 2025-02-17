# AssistantSearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to **map[string]interface{}** | Metadata to filter by. Exact match filter for each KV pair. | [optional] 
**GraphId** | Pointer to **string** | The ID of the graph to filter by. The graph ID is normally set in your langgraph.json configuration. | [optional] 
**Limit** | Pointer to **int32** | The maximum number of results to return. | [optional] [default to 10]
**Offset** | Pointer to **int32** | The number of results to skip. | [optional] [default to 0]

## Methods

### NewAssistantSearchRequest

`func NewAssistantSearchRequest() *AssistantSearchRequest`

NewAssistantSearchRequest instantiates a new AssistantSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssistantSearchRequestWithDefaults

`func NewAssistantSearchRequestWithDefaults() *AssistantSearchRequest`

NewAssistantSearchRequestWithDefaults instantiates a new AssistantSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *AssistantSearchRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AssistantSearchRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AssistantSearchRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AssistantSearchRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetGraphId

`func (o *AssistantSearchRequest) GetGraphId() string`

GetGraphId returns the GraphId field if non-nil, zero value otherwise.

### GetGraphIdOk

`func (o *AssistantSearchRequest) GetGraphIdOk() (*string, bool)`

GetGraphIdOk returns a tuple with the GraphId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphId

`func (o *AssistantSearchRequest) SetGraphId(v string)`

SetGraphId sets GraphId field to given value.

### HasGraphId

`func (o *AssistantSearchRequest) HasGraphId() bool`

HasGraphId returns a boolean if a field has been set.

### GetLimit

`func (o *AssistantSearchRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *AssistantSearchRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *AssistantSearchRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *AssistantSearchRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *AssistantSearchRequest) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *AssistantSearchRequest) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *AssistantSearchRequest) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *AssistantSearchRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


