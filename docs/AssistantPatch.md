# AssistantPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GraphId** | Pointer to **string** | The ID of the graph the assistant should use. The graph ID is normally set in your langgraph.json configuration. If not provided, assistant will keep pointing to same graph. | [optional] 
**Config** | Pointer to **map[string]interface{}** | Configuration to use for the graph. Useful when graph is configurable and you want to update the assistant&#39;s configuration. | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Metadata to merge with existing assistant metadata. | [optional] 
**Name** | Pointer to **string** | The new name for the assistant. If not provided, assistant will keep its current name. | [optional] 

## Methods

### NewAssistantPatch

`func NewAssistantPatch() *AssistantPatch`

NewAssistantPatch instantiates a new AssistantPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssistantPatchWithDefaults

`func NewAssistantPatchWithDefaults() *AssistantPatch`

NewAssistantPatchWithDefaults instantiates a new AssistantPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGraphId

`func (o *AssistantPatch) GetGraphId() string`

GetGraphId returns the GraphId field if non-nil, zero value otherwise.

### GetGraphIdOk

`func (o *AssistantPatch) GetGraphIdOk() (*string, bool)`

GetGraphIdOk returns a tuple with the GraphId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphId

`func (o *AssistantPatch) SetGraphId(v string)`

SetGraphId sets GraphId field to given value.

### HasGraphId

`func (o *AssistantPatch) HasGraphId() bool`

HasGraphId returns a boolean if a field has been set.

### GetConfig

`func (o *AssistantPatch) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AssistantPatch) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AssistantPatch) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AssistantPatch) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetMetadata

`func (o *AssistantPatch) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AssistantPatch) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AssistantPatch) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AssistantPatch) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *AssistantPatch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AssistantPatch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AssistantPatch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AssistantPatch) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


