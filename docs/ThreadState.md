# ThreadState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Values** | [**Values**](Values.md) |  | 
**Next** | **[]string** |  | 
**Tasks** | Pointer to [**[]ThreadStateTasksInner**](ThreadStateTasksInner.md) |  | [optional] 
**Checkpoint** | [**CheckpointConfig**](CheckpointConfig.md) |  | 
**Metadata** | **map[string]interface{}** |  | 
**CreatedAt** | **string** |  | 
**ParentCheckpoint** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewThreadState

`func NewThreadState(values Values, next []string, checkpoint CheckpointConfig, metadata map[string]interface{}, createdAt string, ) *ThreadState`

NewThreadState instantiates a new ThreadState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreadStateWithDefaults

`func NewThreadStateWithDefaults() *ThreadState`

NewThreadStateWithDefaults instantiates a new ThreadState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValues

`func (o *ThreadState) GetValues() Values`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ThreadState) GetValuesOk() (*Values, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ThreadState) SetValues(v Values)`

SetValues sets Values field to given value.


### GetNext

`func (o *ThreadState) GetNext() []string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *ThreadState) GetNextOk() (*[]string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *ThreadState) SetNext(v []string)`

SetNext sets Next field to given value.


### GetTasks

`func (o *ThreadState) GetTasks() []ThreadStateTasksInner`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *ThreadState) GetTasksOk() (*[]ThreadStateTasksInner, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *ThreadState) SetTasks(v []ThreadStateTasksInner)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *ThreadState) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetCheckpoint

`func (o *ThreadState) GetCheckpoint() CheckpointConfig`

GetCheckpoint returns the Checkpoint field if non-nil, zero value otherwise.

### GetCheckpointOk

`func (o *ThreadState) GetCheckpointOk() (*CheckpointConfig, bool)`

GetCheckpointOk returns a tuple with the Checkpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckpoint

`func (o *ThreadState) SetCheckpoint(v CheckpointConfig)`

SetCheckpoint sets Checkpoint field to given value.


### GetMetadata

`func (o *ThreadState) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ThreadState) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ThreadState) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetCreatedAt

`func (o *ThreadState) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ThreadState) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ThreadState) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetParentCheckpoint

`func (o *ThreadState) GetParentCheckpoint() map[string]interface{}`

GetParentCheckpoint returns the ParentCheckpoint field if non-nil, zero value otherwise.

### GetParentCheckpointOk

`func (o *ThreadState) GetParentCheckpointOk() (*map[string]interface{}, bool)`

GetParentCheckpointOk returns a tuple with the ParentCheckpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCheckpoint

`func (o *ThreadState) SetParentCheckpoint(v map[string]interface{})`

SetParentCheckpoint sets ParentCheckpoint field to given value.

### HasParentCheckpoint

`func (o *ThreadState) HasParentCheckpoint() bool`

HasParentCheckpoint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


