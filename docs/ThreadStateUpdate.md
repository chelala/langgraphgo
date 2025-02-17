# ThreadStateUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Values** | Pointer to [**NullableValues1**](Values1.md) |  | [optional] 
**Checkpoint** | Pointer to [**CheckpointConfig**](CheckpointConfig.md) | The checkpoint to update the state of. | [optional] 
**AsNode** | Pointer to **string** | Update the state as if this node had just executed. | [optional] 

## Methods

### NewThreadStateUpdate

`func NewThreadStateUpdate() *ThreadStateUpdate`

NewThreadStateUpdate instantiates a new ThreadStateUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreadStateUpdateWithDefaults

`func NewThreadStateUpdateWithDefaults() *ThreadStateUpdate`

NewThreadStateUpdateWithDefaults instantiates a new ThreadStateUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValues

`func (o *ThreadStateUpdate) GetValues() Values1`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ThreadStateUpdate) GetValuesOk() (*Values1, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ThreadStateUpdate) SetValues(v Values1)`

SetValues sets Values field to given value.

### HasValues

`func (o *ThreadStateUpdate) HasValues() bool`

HasValues returns a boolean if a field has been set.

### SetValuesNil

`func (o *ThreadStateUpdate) SetValuesNil(b bool)`

 SetValuesNil sets the value for Values to be an explicit nil

### UnsetValues
`func (o *ThreadStateUpdate) UnsetValues()`

UnsetValues ensures that no value is present for Values, not even an explicit nil
### GetCheckpoint

`func (o *ThreadStateUpdate) GetCheckpoint() CheckpointConfig`

GetCheckpoint returns the Checkpoint field if non-nil, zero value otherwise.

### GetCheckpointOk

`func (o *ThreadStateUpdate) GetCheckpointOk() (*CheckpointConfig, bool)`

GetCheckpointOk returns a tuple with the Checkpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckpoint

`func (o *ThreadStateUpdate) SetCheckpoint(v CheckpointConfig)`

SetCheckpoint sets Checkpoint field to given value.

### HasCheckpoint

`func (o *ThreadStateUpdate) HasCheckpoint() bool`

HasCheckpoint returns a boolean if a field has been set.

### GetAsNode

`func (o *ThreadStateUpdate) GetAsNode() string`

GetAsNode returns the AsNode field if non-nil, zero value otherwise.

### GetAsNodeOk

`func (o *ThreadStateUpdate) GetAsNodeOk() (*string, bool)`

GetAsNodeOk returns a tuple with the AsNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsNode

`func (o *ThreadStateUpdate) SetAsNode(v string)`

SetAsNode sets AsNode field to given value.

### HasAsNode

`func (o *ThreadStateUpdate) HasAsNode() bool`

HasAsNode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


