# CheckpointConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ThreadId** | Pointer to **string** | Unique identifier for the thread associated with this checkpoint. | [optional] 
**CheckpointNs** | Pointer to **string** | Namespace for the checkpoint, used for organization and retrieval. | [optional] 
**CheckpointId** | Pointer to **string** | Optional unique identifier for the checkpoint itself. | [optional] 
**CheckpointMap** | Pointer to **map[string]interface{}** | Optional dictionary containing checkpoint-specific data. | [optional] 

## Methods

### NewCheckpointConfig

`func NewCheckpointConfig() *CheckpointConfig`

NewCheckpointConfig instantiates a new CheckpointConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckpointConfigWithDefaults

`func NewCheckpointConfigWithDefaults() *CheckpointConfig`

NewCheckpointConfigWithDefaults instantiates a new CheckpointConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThreadId

`func (o *CheckpointConfig) GetThreadId() string`

GetThreadId returns the ThreadId field if non-nil, zero value otherwise.

### GetThreadIdOk

`func (o *CheckpointConfig) GetThreadIdOk() (*string, bool)`

GetThreadIdOk returns a tuple with the ThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadId

`func (o *CheckpointConfig) SetThreadId(v string)`

SetThreadId sets ThreadId field to given value.

### HasThreadId

`func (o *CheckpointConfig) HasThreadId() bool`

HasThreadId returns a boolean if a field has been set.

### GetCheckpointNs

`func (o *CheckpointConfig) GetCheckpointNs() string`

GetCheckpointNs returns the CheckpointNs field if non-nil, zero value otherwise.

### GetCheckpointNsOk

`func (o *CheckpointConfig) GetCheckpointNsOk() (*string, bool)`

GetCheckpointNsOk returns a tuple with the CheckpointNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckpointNs

`func (o *CheckpointConfig) SetCheckpointNs(v string)`

SetCheckpointNs sets CheckpointNs field to given value.

### HasCheckpointNs

`func (o *CheckpointConfig) HasCheckpointNs() bool`

HasCheckpointNs returns a boolean if a field has been set.

### GetCheckpointId

`func (o *CheckpointConfig) GetCheckpointId() string`

GetCheckpointId returns the CheckpointId field if non-nil, zero value otherwise.

### GetCheckpointIdOk

`func (o *CheckpointConfig) GetCheckpointIdOk() (*string, bool)`

GetCheckpointIdOk returns a tuple with the CheckpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckpointId

`func (o *CheckpointConfig) SetCheckpointId(v string)`

SetCheckpointId sets CheckpointId field to given value.

### HasCheckpointId

`func (o *CheckpointConfig) HasCheckpointId() bool`

HasCheckpointId returns a boolean if a field has been set.

### GetCheckpointMap

`func (o *CheckpointConfig) GetCheckpointMap() map[string]interface{}`

GetCheckpointMap returns the CheckpointMap field if non-nil, zero value otherwise.

### GetCheckpointMapOk

`func (o *CheckpointConfig) GetCheckpointMapOk() (*map[string]interface{}, bool)`

GetCheckpointMapOk returns a tuple with the CheckpointMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckpointMap

`func (o *CheckpointConfig) SetCheckpointMap(v map[string]interface{})`

SetCheckpointMap sets CheckpointMap field to given value.

### HasCheckpointMap

`func (o *CheckpointConfig) HasCheckpointMap() bool`

HasCheckpointMap returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


