# AssistantCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssistantId** | Pointer to **string** | The ID of the assistant. If not provided, a random UUID will be generated. | [optional] 
**GraphId** | **string** | The ID of the graph the assistant should use. The graph ID is normally set in your langgraph.json configuration. | 
**Config** | Pointer to **map[string]interface{}** | Configuration to use for the graph. Useful when graph is configurable and you want to create different assistants based on different configurations. | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Metadata to add to assistant. | [optional] 
**IfExists** | Pointer to **string** | How to handle duplicate creation. Must be either &#39;raise&#39; (raise error if duplicate), or &#39;do_nothing&#39; (return existing assistant). | [optional] [default to "raise"]
**Name** | Pointer to **string** | The name of the assistant. Defaults to &#39;Untitled&#39;. | [optional] 

## Methods

### NewAssistantCreate

`func NewAssistantCreate(graphId string, ) *AssistantCreate`

NewAssistantCreate instantiates a new AssistantCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssistantCreateWithDefaults

`func NewAssistantCreateWithDefaults() *AssistantCreate`

NewAssistantCreateWithDefaults instantiates a new AssistantCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssistantId

`func (o *AssistantCreate) GetAssistantId() string`

GetAssistantId returns the AssistantId field if non-nil, zero value otherwise.

### GetAssistantIdOk

`func (o *AssistantCreate) GetAssistantIdOk() (*string, bool)`

GetAssistantIdOk returns a tuple with the AssistantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssistantId

`func (o *AssistantCreate) SetAssistantId(v string)`

SetAssistantId sets AssistantId field to given value.

### HasAssistantId

`func (o *AssistantCreate) HasAssistantId() bool`

HasAssistantId returns a boolean if a field has been set.

### GetGraphId

`func (o *AssistantCreate) GetGraphId() string`

GetGraphId returns the GraphId field if non-nil, zero value otherwise.

### GetGraphIdOk

`func (o *AssistantCreate) GetGraphIdOk() (*string, bool)`

GetGraphIdOk returns a tuple with the GraphId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphId

`func (o *AssistantCreate) SetGraphId(v string)`

SetGraphId sets GraphId field to given value.


### GetConfig

`func (o *AssistantCreate) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AssistantCreate) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AssistantCreate) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AssistantCreate) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetMetadata

`func (o *AssistantCreate) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AssistantCreate) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AssistantCreate) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AssistantCreate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetIfExists

`func (o *AssistantCreate) GetIfExists() string`

GetIfExists returns the IfExists field if non-nil, zero value otherwise.

### GetIfExistsOk

`func (o *AssistantCreate) GetIfExistsOk() (*string, bool)`

GetIfExistsOk returns a tuple with the IfExists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIfExists

`func (o *AssistantCreate) SetIfExists(v string)`

SetIfExists sets IfExists field to given value.

### HasIfExists

`func (o *AssistantCreate) HasIfExists() bool`

HasIfExists returns a boolean if a field has been set.

### GetName

`func (o *AssistantCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AssistantCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AssistantCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AssistantCreate) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


