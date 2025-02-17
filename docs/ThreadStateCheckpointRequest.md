# ThreadStateCheckpointRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Checkpoint** | [**CheckpointConfig**](CheckpointConfig.md) | The checkpoint to get the state for. | 
**Subgraphs** | Pointer to **bool** | Include subgraph states. | [optional] 

## Methods

### NewThreadStateCheckpointRequest

`func NewThreadStateCheckpointRequest(checkpoint CheckpointConfig, ) *ThreadStateCheckpointRequest`

NewThreadStateCheckpointRequest instantiates a new ThreadStateCheckpointRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreadStateCheckpointRequestWithDefaults

`func NewThreadStateCheckpointRequestWithDefaults() *ThreadStateCheckpointRequest`

NewThreadStateCheckpointRequestWithDefaults instantiates a new ThreadStateCheckpointRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckpoint

`func (o *ThreadStateCheckpointRequest) GetCheckpoint() CheckpointConfig`

GetCheckpoint returns the Checkpoint field if non-nil, zero value otherwise.

### GetCheckpointOk

`func (o *ThreadStateCheckpointRequest) GetCheckpointOk() (*CheckpointConfig, bool)`

GetCheckpointOk returns a tuple with the Checkpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckpoint

`func (o *ThreadStateCheckpointRequest) SetCheckpoint(v CheckpointConfig)`

SetCheckpoint sets Checkpoint field to given value.


### GetSubgraphs

`func (o *ThreadStateCheckpointRequest) GetSubgraphs() bool`

GetSubgraphs returns the Subgraphs field if non-nil, zero value otherwise.

### GetSubgraphsOk

`func (o *ThreadStateCheckpointRequest) GetSubgraphsOk() (*bool, bool)`

GetSubgraphsOk returns a tuple with the Subgraphs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubgraphs

`func (o *ThreadStateCheckpointRequest) SetSubgraphs(v bool)`

SetSubgraphs sets Subgraphs field to given value.

### HasSubgraphs

`func (o *ThreadStateCheckpointRequest) HasSubgraphs() bool`

HasSubgraphs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


