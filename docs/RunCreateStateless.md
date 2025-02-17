# RunCreateStateless

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssistantId** | [**RunCreateStatefulAssistantId**](RunCreateStatefulAssistantId.md) |  | 
**Input** | Pointer to [**NullableInput1**](Input1.md) |  | [optional] 
**Command** | Pointer to [**NullableCommand**](Command.md) |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Metadata to assign to the run. | [optional] 
**Config** | Pointer to [**Config2**](Config2.md) |  | [optional] 
**Webhook** | Pointer to **string** | Webhook to call after LangGraph API call is done. | [optional] 
**InterruptBefore** | Pointer to [**InterruptBefore**](InterruptBefore.md) |  | [optional] 
**InterruptAfter** | Pointer to [**InterruptAfter**](InterruptAfter.md) |  | [optional] 
**StreamMode** | Pointer to [**StreamMode**](StreamMode.md) |  | [optional] [default to [values]]
**FeedbackKeys** | Pointer to **[]string** | Feedback keys to assign to run. | [optional] 
**StreamSubgraphs** | Pointer to **bool** | Whether to stream output from subgraphs. | [optional] [default to false]
**OnCompletion** | Pointer to **string** | Whether to delete or keep the thread created for a stateless run. Must be one of &#39;delete&#39; or &#39;keep&#39;. | [optional] [default to "delete"]
**OnDisconnect** | Pointer to **string** | The disconnect mode to use. Must be one of &#39;cancel&#39; or &#39;continue&#39;. | [optional] [default to "cancel"]
**AfterSeconds** | Pointer to **int32** | The number of seconds to wait before starting the run. Use to schedule future runs. | [optional] 

## Methods

### NewRunCreateStateless

`func NewRunCreateStateless(assistantId RunCreateStatefulAssistantId, ) *RunCreateStateless`

NewRunCreateStateless instantiates a new RunCreateStateless object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunCreateStatelessWithDefaults

`func NewRunCreateStatelessWithDefaults() *RunCreateStateless`

NewRunCreateStatelessWithDefaults instantiates a new RunCreateStateless object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssistantId

`func (o *RunCreateStateless) GetAssistantId() RunCreateStatefulAssistantId`

GetAssistantId returns the AssistantId field if non-nil, zero value otherwise.

### GetAssistantIdOk

`func (o *RunCreateStateless) GetAssistantIdOk() (*RunCreateStatefulAssistantId, bool)`

GetAssistantIdOk returns a tuple with the AssistantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssistantId

`func (o *RunCreateStateless) SetAssistantId(v RunCreateStatefulAssistantId)`

SetAssistantId sets AssistantId field to given value.


### GetInput

`func (o *RunCreateStateless) GetInput() Input1`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *RunCreateStateless) GetInputOk() (*Input1, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *RunCreateStateless) SetInput(v Input1)`

SetInput sets Input field to given value.

### HasInput

`func (o *RunCreateStateless) HasInput() bool`

HasInput returns a boolean if a field has been set.

### SetInputNil

`func (o *RunCreateStateless) SetInputNil(b bool)`

 SetInputNil sets the value for Input to be an explicit nil

### UnsetInput
`func (o *RunCreateStateless) UnsetInput()`

UnsetInput ensures that no value is present for Input, not even an explicit nil
### GetCommand

`func (o *RunCreateStateless) GetCommand() Command`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *RunCreateStateless) GetCommandOk() (*Command, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *RunCreateStateless) SetCommand(v Command)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *RunCreateStateless) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### SetCommandNil

`func (o *RunCreateStateless) SetCommandNil(b bool)`

 SetCommandNil sets the value for Command to be an explicit nil

### UnsetCommand
`func (o *RunCreateStateless) UnsetCommand()`

UnsetCommand ensures that no value is present for Command, not even an explicit nil
### GetMetadata

`func (o *RunCreateStateless) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RunCreateStateless) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RunCreateStateless) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RunCreateStateless) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetConfig

`func (o *RunCreateStateless) GetConfig() Config2`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *RunCreateStateless) GetConfigOk() (*Config2, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *RunCreateStateless) SetConfig(v Config2)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *RunCreateStateless) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetWebhook

`func (o *RunCreateStateless) GetWebhook() string`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *RunCreateStateless) GetWebhookOk() (*string, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *RunCreateStateless) SetWebhook(v string)`

SetWebhook sets Webhook field to given value.

### HasWebhook

`func (o *RunCreateStateless) HasWebhook() bool`

HasWebhook returns a boolean if a field has been set.

### GetInterruptBefore

`func (o *RunCreateStateless) GetInterruptBefore() InterruptBefore`

GetInterruptBefore returns the InterruptBefore field if non-nil, zero value otherwise.

### GetInterruptBeforeOk

`func (o *RunCreateStateless) GetInterruptBeforeOk() (*InterruptBefore, bool)`

GetInterruptBeforeOk returns a tuple with the InterruptBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterruptBefore

`func (o *RunCreateStateless) SetInterruptBefore(v InterruptBefore)`

SetInterruptBefore sets InterruptBefore field to given value.

### HasInterruptBefore

`func (o *RunCreateStateless) HasInterruptBefore() bool`

HasInterruptBefore returns a boolean if a field has been set.

### GetInterruptAfter

`func (o *RunCreateStateless) GetInterruptAfter() InterruptAfter`

GetInterruptAfter returns the InterruptAfter field if non-nil, zero value otherwise.

### GetInterruptAfterOk

`func (o *RunCreateStateless) GetInterruptAfterOk() (*InterruptAfter, bool)`

GetInterruptAfterOk returns a tuple with the InterruptAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterruptAfter

`func (o *RunCreateStateless) SetInterruptAfter(v InterruptAfter)`

SetInterruptAfter sets InterruptAfter field to given value.

### HasInterruptAfter

`func (o *RunCreateStateless) HasInterruptAfter() bool`

HasInterruptAfter returns a boolean if a field has been set.

### GetStreamMode

`func (o *RunCreateStateless) GetStreamMode() StreamMode`

GetStreamMode returns the StreamMode field if non-nil, zero value otherwise.

### GetStreamModeOk

`func (o *RunCreateStateless) GetStreamModeOk() (*StreamMode, bool)`

GetStreamModeOk returns a tuple with the StreamMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamMode

`func (o *RunCreateStateless) SetStreamMode(v StreamMode)`

SetStreamMode sets StreamMode field to given value.

### HasStreamMode

`func (o *RunCreateStateless) HasStreamMode() bool`

HasStreamMode returns a boolean if a field has been set.

### GetFeedbackKeys

`func (o *RunCreateStateless) GetFeedbackKeys() []string`

GetFeedbackKeys returns the FeedbackKeys field if non-nil, zero value otherwise.

### GetFeedbackKeysOk

`func (o *RunCreateStateless) GetFeedbackKeysOk() (*[]string, bool)`

GetFeedbackKeysOk returns a tuple with the FeedbackKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedbackKeys

`func (o *RunCreateStateless) SetFeedbackKeys(v []string)`

SetFeedbackKeys sets FeedbackKeys field to given value.

### HasFeedbackKeys

`func (o *RunCreateStateless) HasFeedbackKeys() bool`

HasFeedbackKeys returns a boolean if a field has been set.

### GetStreamSubgraphs

`func (o *RunCreateStateless) GetStreamSubgraphs() bool`

GetStreamSubgraphs returns the StreamSubgraphs field if non-nil, zero value otherwise.

### GetStreamSubgraphsOk

`func (o *RunCreateStateless) GetStreamSubgraphsOk() (*bool, bool)`

GetStreamSubgraphsOk returns a tuple with the StreamSubgraphs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamSubgraphs

`func (o *RunCreateStateless) SetStreamSubgraphs(v bool)`

SetStreamSubgraphs sets StreamSubgraphs field to given value.

### HasStreamSubgraphs

`func (o *RunCreateStateless) HasStreamSubgraphs() bool`

HasStreamSubgraphs returns a boolean if a field has been set.

### GetOnCompletion

`func (o *RunCreateStateless) GetOnCompletion() string`

GetOnCompletion returns the OnCompletion field if non-nil, zero value otherwise.

### GetOnCompletionOk

`func (o *RunCreateStateless) GetOnCompletionOk() (*string, bool)`

GetOnCompletionOk returns a tuple with the OnCompletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnCompletion

`func (o *RunCreateStateless) SetOnCompletion(v string)`

SetOnCompletion sets OnCompletion field to given value.

### HasOnCompletion

`func (o *RunCreateStateless) HasOnCompletion() bool`

HasOnCompletion returns a boolean if a field has been set.

### GetOnDisconnect

`func (o *RunCreateStateless) GetOnDisconnect() string`

GetOnDisconnect returns the OnDisconnect field if non-nil, zero value otherwise.

### GetOnDisconnectOk

`func (o *RunCreateStateless) GetOnDisconnectOk() (*string, bool)`

GetOnDisconnectOk returns a tuple with the OnDisconnect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnDisconnect

`func (o *RunCreateStateless) SetOnDisconnect(v string)`

SetOnDisconnect sets OnDisconnect field to given value.

### HasOnDisconnect

`func (o *RunCreateStateless) HasOnDisconnect() bool`

HasOnDisconnect returns a boolean if a field has been set.

### GetAfterSeconds

`func (o *RunCreateStateless) GetAfterSeconds() int32`

GetAfterSeconds returns the AfterSeconds field if non-nil, zero value otherwise.

### GetAfterSecondsOk

`func (o *RunCreateStateless) GetAfterSecondsOk() (*int32, bool)`

GetAfterSecondsOk returns a tuple with the AfterSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfterSeconds

`func (o *RunCreateStateless) SetAfterSeconds(v int32)`

SetAfterSeconds sets AfterSeconds field to given value.

### HasAfterSeconds

`func (o *RunCreateStateless) HasAfterSeconds() bool`

HasAfterSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


