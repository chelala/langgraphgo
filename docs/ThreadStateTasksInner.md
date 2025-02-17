# ThreadStateTasksInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Error** | Pointer to **string** |  | [optional] 
**Interrupts** | Pointer to **[]interface{}** |  | [optional] 
**Checkpoint** | Pointer to [**CheckpointConfig**](CheckpointConfig.md) |  | [optional] 
**State** | Pointer to [**ThreadState**](ThreadState.md) |  | [optional] 

## Methods

### NewThreadStateTasksInner

`func NewThreadStateTasksInner(id string, name string, ) *ThreadStateTasksInner`

NewThreadStateTasksInner instantiates a new ThreadStateTasksInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreadStateTasksInnerWithDefaults

`func NewThreadStateTasksInnerWithDefaults() *ThreadStateTasksInner`

NewThreadStateTasksInnerWithDefaults instantiates a new ThreadStateTasksInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ThreadStateTasksInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ThreadStateTasksInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ThreadStateTasksInner) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ThreadStateTasksInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ThreadStateTasksInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ThreadStateTasksInner) SetName(v string)`

SetName sets Name field to given value.


### GetError

`func (o *ThreadStateTasksInner) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ThreadStateTasksInner) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ThreadStateTasksInner) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ThreadStateTasksInner) HasError() bool`

HasError returns a boolean if a field has been set.

### GetInterrupts

`func (o *ThreadStateTasksInner) GetInterrupts() []interface{}`

GetInterrupts returns the Interrupts field if non-nil, zero value otherwise.

### GetInterruptsOk

`func (o *ThreadStateTasksInner) GetInterruptsOk() (*[]interface{}, bool)`

GetInterruptsOk returns a tuple with the Interrupts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterrupts

`func (o *ThreadStateTasksInner) SetInterrupts(v []interface{})`

SetInterrupts sets Interrupts field to given value.

### HasInterrupts

`func (o *ThreadStateTasksInner) HasInterrupts() bool`

HasInterrupts returns a boolean if a field has been set.

### GetCheckpoint

`func (o *ThreadStateTasksInner) GetCheckpoint() CheckpointConfig`

GetCheckpoint returns the Checkpoint field if non-nil, zero value otherwise.

### GetCheckpointOk

`func (o *ThreadStateTasksInner) GetCheckpointOk() (*CheckpointConfig, bool)`

GetCheckpointOk returns a tuple with the Checkpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckpoint

`func (o *ThreadStateTasksInner) SetCheckpoint(v CheckpointConfig)`

SetCheckpoint sets Checkpoint field to given value.

### HasCheckpoint

`func (o *ThreadStateTasksInner) HasCheckpoint() bool`

HasCheckpoint returns a boolean if a field has been set.

### GetState

`func (o *ThreadStateTasksInner) GetState() ThreadState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ThreadStateTasksInner) GetStateOk() (*ThreadState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ThreadStateTasksInner) SetState(v ThreadState)`

SetState sets State field to given value.

### HasState

`func (o *ThreadStateTasksInner) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


