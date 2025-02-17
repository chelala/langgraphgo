# Assistant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssistantId** | **string** | The ID of the assistant. | 
**GraphId** | **string** | The ID of the graph. | 
**Config** | [**Config1**](Config1.md) |  | 
**CreatedAt** | **time.Time** | The time the assistant was created. | 
**UpdatedAt** | **time.Time** | The last time the assistant was updated. | 
**Metadata** | **map[string]interface{}** | The assistant metadata. | 
**Version** | Pointer to **int32** | The version of the assistant | [optional] 
**Name** | Pointer to **string** | The name of the assistant | [optional] 

## Methods

### NewAssistant

`func NewAssistant(assistantId string, graphId string, config Config1, createdAt time.Time, updatedAt time.Time, metadata map[string]interface{}, ) *Assistant`

NewAssistant instantiates a new Assistant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssistantWithDefaults

`func NewAssistantWithDefaults() *Assistant`

NewAssistantWithDefaults instantiates a new Assistant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssistantId

`func (o *Assistant) GetAssistantId() string`

GetAssistantId returns the AssistantId field if non-nil, zero value otherwise.

### GetAssistantIdOk

`func (o *Assistant) GetAssistantIdOk() (*string, bool)`

GetAssistantIdOk returns a tuple with the AssistantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssistantId

`func (o *Assistant) SetAssistantId(v string)`

SetAssistantId sets AssistantId field to given value.


### GetGraphId

`func (o *Assistant) GetGraphId() string`

GetGraphId returns the GraphId field if non-nil, zero value otherwise.

### GetGraphIdOk

`func (o *Assistant) GetGraphIdOk() (*string, bool)`

GetGraphIdOk returns a tuple with the GraphId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphId

`func (o *Assistant) SetGraphId(v string)`

SetGraphId sets GraphId field to given value.


### GetConfig

`func (o *Assistant) GetConfig() Config1`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *Assistant) GetConfigOk() (*Config1, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *Assistant) SetConfig(v Config1)`

SetConfig sets Config field to given value.


### GetCreatedAt

`func (o *Assistant) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Assistant) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Assistant) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Assistant) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Assistant) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Assistant) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetMetadata

`func (o *Assistant) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Assistant) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Assistant) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetVersion

`func (o *Assistant) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Assistant) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Assistant) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Assistant) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetName

`func (o *Assistant) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Assistant) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Assistant) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Assistant) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


