# ThreadStateSearch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **int32** | The maximum number of states to return. | [optional] [default to 10]
**Before** | Pointer to [**CheckpointConfig**](CheckpointConfig.md) | Return states before this checkpoint. | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Filter states by metadata key-value pairs. | [optional] 
**Checkpoint** | Pointer to [**CheckpointConfig**](CheckpointConfig.md) | Return states for this subgraph. | [optional] 

## Methods

### NewThreadStateSearch

`func NewThreadStateSearch() *ThreadStateSearch`

NewThreadStateSearch instantiates a new ThreadStateSearch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreadStateSearchWithDefaults

`func NewThreadStateSearchWithDefaults() *ThreadStateSearch`

NewThreadStateSearchWithDefaults instantiates a new ThreadStateSearch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *ThreadStateSearch) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ThreadStateSearch) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ThreadStateSearch) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ThreadStateSearch) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetBefore

`func (o *ThreadStateSearch) GetBefore() CheckpointConfig`

GetBefore returns the Before field if non-nil, zero value otherwise.

### GetBeforeOk

`func (o *ThreadStateSearch) GetBeforeOk() (*CheckpointConfig, bool)`

GetBeforeOk returns a tuple with the Before field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBefore

`func (o *ThreadStateSearch) SetBefore(v CheckpointConfig)`

SetBefore sets Before field to given value.

### HasBefore

`func (o *ThreadStateSearch) HasBefore() bool`

HasBefore returns a boolean if a field has been set.

### GetMetadata

`func (o *ThreadStateSearch) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ThreadStateSearch) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ThreadStateSearch) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ThreadStateSearch) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCheckpoint

`func (o *ThreadStateSearch) GetCheckpoint() CheckpointConfig`

GetCheckpoint returns the Checkpoint field if non-nil, zero value otherwise.

### GetCheckpointOk

`func (o *ThreadStateSearch) GetCheckpointOk() (*CheckpointConfig, bool)`

GetCheckpointOk returns a tuple with the Checkpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckpoint

`func (o *ThreadStateSearch) SetCheckpoint(v CheckpointConfig)`

SetCheckpoint sets Checkpoint field to given value.

### HasCheckpoint

`func (o *ThreadStateSearch) HasCheckpoint() bool`

HasCheckpoint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


