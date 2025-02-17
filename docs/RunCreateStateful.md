# RunCreateStateful

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssistantId** | [**RunCreateStatefulAssistantId**](RunCreateStatefulAssistantId.md) |  | 
**Checkpoint** | Pointer to [**CheckpointConfig**](CheckpointConfig.md) | The checkpoint to resume from. | [optional] 
**Input** | Pointer to [**NullableInput1**](Input1.md) |  | [optional] 
**Command** | Pointer to [**NullableCommand**](Command.md) |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Metadata to assign to the run. | [optional] 
**Config** | Pointer to [**Config2**](Config2.md) |  | [optional] 
**Webhook** | Pointer to **string** | Webhook to call after LangGraph API call is done. | [optional] 
**InterruptBefore** | Pointer to [**InterruptBefore**](InterruptBefore.md) |  | [optional] 
**InterruptAfter** | Pointer to [**InterruptAfter**](InterruptAfter.md) |  | [optional] 
**StreamMode** | Pointer to [**StreamMode**](StreamMode.md) |  | [optional] [default to [values]]
**StreamSubgraphs** | Pointer to **bool** | Whether to stream output from subgraphs. | [optional] [default to false]
**OnDisconnect** | Pointer to **string** | The disconnect mode to use. Must be one of &#39;cancel&#39; or &#39;continue&#39;. | [optional] [default to "cancel"]
**FeedbackKeys** | Pointer to **[]string** | Feedback keys to assign to run. | [optional] 
**MultitaskStrategy** | Pointer to **string** | Multitask strategy to use. Must be one of &#39;reject&#39;, &#39;interrupt&#39;, &#39;rollback&#39;, or &#39;enqueue&#39;. | [optional] [default to "reject"]
**IfNotExists** | Pointer to **string** | How to handle missing thread. Must be either &#39;reject&#39; (raise error if missing), or &#39;create&#39; (create new thread). | [optional] [default to "reject"]
**AfterSeconds** | Pointer to **int32** | The number of seconds to wait before starting the run. Use to schedule future runs. | [optional] 

## Methods

### NewRunCreateStateful

`func NewRunCreateStateful(assistantId RunCreateStatefulAssistantId, ) *RunCreateStateful`

NewRunCreateStateful instantiates a new RunCreateStateful object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunCreateStatefulWithDefaults

`func NewRunCreateStatefulWithDefaults() *RunCreateStateful`

NewRunCreateStatefulWithDefaults instantiates a new RunCreateStateful object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssistantId

`func (o *RunCreateStateful) GetAssistantId() RunCreateStatefulAssistantId`

GetAssistantId returns the AssistantId field if non-nil, zero value otherwise.

### GetAssistantIdOk

`func (o *RunCreateStateful) GetAssistantIdOk() (*RunCreateStatefulAssistantId, bool)`

GetAssistantIdOk returns a tuple with the AssistantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssistantId

`func (o *RunCreateStateful) SetAssistantId(v RunCreateStatefulAssistantId)`

SetAssistantId sets AssistantId field to given value.


### GetCheckpoint

`func (o *RunCreateStateful) GetCheckpoint() CheckpointConfig`

GetCheckpoint returns the Checkpoint field if non-nil, zero value otherwise.

### GetCheckpointOk

`func (o *RunCreateStateful) GetCheckpointOk() (*CheckpointConfig, bool)`

GetCheckpointOk returns a tuple with the Checkpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckpoint

`func (o *RunCreateStateful) SetCheckpoint(v CheckpointConfig)`

SetCheckpoint sets Checkpoint field to given value.

### HasCheckpoint

`func (o *RunCreateStateful) HasCheckpoint() bool`

HasCheckpoint returns a boolean if a field has been set.

### GetInput

`func (o *RunCreateStateful) GetInput() Input1`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *RunCreateStateful) GetInputOk() (*Input1, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *RunCreateStateful) SetInput(v Input1)`

SetInput sets Input field to given value.

### HasInput

`func (o *RunCreateStateful) HasInput() bool`

HasInput returns a boolean if a field has been set.

### SetInputNil

`func (o *RunCreateStateful) SetInputNil(b bool)`

 SetInputNil sets the value for Input to be an explicit nil

### UnsetInput
`func (o *RunCreateStateful) UnsetInput()`

UnsetInput ensures that no value is present for Input, not even an explicit nil
### GetCommand

`func (o *RunCreateStateful) GetCommand() Command`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *RunCreateStateful) GetCommandOk() (*Command, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *RunCreateStateful) SetCommand(v Command)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *RunCreateStateful) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### SetCommandNil

`func (o *RunCreateStateful) SetCommandNil(b bool)`

 SetCommandNil sets the value for Command to be an explicit nil

### UnsetCommand
`func (o *RunCreateStateful) UnsetCommand()`

UnsetCommand ensures that no value is present for Command, not even an explicit nil
### GetMetadata

`func (o *RunCreateStateful) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RunCreateStateful) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RunCreateStateful) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RunCreateStateful) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetConfig

`func (o *RunCreateStateful) GetConfig() Config2`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *RunCreateStateful) GetConfigOk() (*Config2, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *RunCreateStateful) SetConfig(v Config2)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *RunCreateStateful) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetWebhook

`func (o *RunCreateStateful) GetWebhook() string`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *RunCreateStateful) GetWebhookOk() (*string, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *RunCreateStateful) SetWebhook(v string)`

SetWebhook sets Webhook field to given value.

### HasWebhook

`func (o *RunCreateStateful) HasWebhook() bool`

HasWebhook returns a boolean if a field has been set.

### GetInterruptBefore

`func (o *RunCreateStateful) GetInterruptBefore() InterruptBefore`

GetInterruptBefore returns the InterruptBefore field if non-nil, zero value otherwise.

### GetInterruptBeforeOk

`func (o *RunCreateStateful) GetInterruptBeforeOk() (*InterruptBefore, bool)`

GetInterruptBeforeOk returns a tuple with the InterruptBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterruptBefore

`func (o *RunCreateStateful) SetInterruptBefore(v InterruptBefore)`

SetInterruptBefore sets InterruptBefore field to given value.

### HasInterruptBefore

`func (o *RunCreateStateful) HasInterruptBefore() bool`

HasInterruptBefore returns a boolean if a field has been set.

### GetInterruptAfter

`func (o *RunCreateStateful) GetInterruptAfter() InterruptAfter`

GetInterruptAfter returns the InterruptAfter field if non-nil, zero value otherwise.

### GetInterruptAfterOk

`func (o *RunCreateStateful) GetInterruptAfterOk() (*InterruptAfter, bool)`

GetInterruptAfterOk returns a tuple with the InterruptAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterruptAfter

`func (o *RunCreateStateful) SetInterruptAfter(v InterruptAfter)`

SetInterruptAfter sets InterruptAfter field to given value.

### HasInterruptAfter

`func (o *RunCreateStateful) HasInterruptAfter() bool`

HasInterruptAfter returns a boolean if a field has been set.

### GetStreamMode

`func (o *RunCreateStateful) GetStreamMode() StreamMode`

GetStreamMode returns the StreamMode field if non-nil, zero value otherwise.

### GetStreamModeOk

`func (o *RunCreateStateful) GetStreamModeOk() (*StreamMode, bool)`

GetStreamModeOk returns a tuple with the StreamMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamMode

`func (o *RunCreateStateful) SetStreamMode(v StreamMode)`

SetStreamMode sets StreamMode field to given value.

### HasStreamMode

`func (o *RunCreateStateful) HasStreamMode() bool`

HasStreamMode returns a boolean if a field has been set.

### GetStreamSubgraphs

`func (o *RunCreateStateful) GetStreamSubgraphs() bool`

GetStreamSubgraphs returns the StreamSubgraphs field if non-nil, zero value otherwise.

### GetStreamSubgraphsOk

`func (o *RunCreateStateful) GetStreamSubgraphsOk() (*bool, bool)`

GetStreamSubgraphsOk returns a tuple with the StreamSubgraphs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamSubgraphs

`func (o *RunCreateStateful) SetStreamSubgraphs(v bool)`

SetStreamSubgraphs sets StreamSubgraphs field to given value.

### HasStreamSubgraphs

`func (o *RunCreateStateful) HasStreamSubgraphs() bool`

HasStreamSubgraphs returns a boolean if a field has been set.

### GetOnDisconnect

`func (o *RunCreateStateful) GetOnDisconnect() string`

GetOnDisconnect returns the OnDisconnect field if non-nil, zero value otherwise.

### GetOnDisconnectOk

`func (o *RunCreateStateful) GetOnDisconnectOk() (*string, bool)`

GetOnDisconnectOk returns a tuple with the OnDisconnect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnDisconnect

`func (o *RunCreateStateful) SetOnDisconnect(v string)`

SetOnDisconnect sets OnDisconnect field to given value.

### HasOnDisconnect

`func (o *RunCreateStateful) HasOnDisconnect() bool`

HasOnDisconnect returns a boolean if a field has been set.

### GetFeedbackKeys

`func (o *RunCreateStateful) GetFeedbackKeys() []string`

GetFeedbackKeys returns the FeedbackKeys field if non-nil, zero value otherwise.

### GetFeedbackKeysOk

`func (o *RunCreateStateful) GetFeedbackKeysOk() (*[]string, bool)`

GetFeedbackKeysOk returns a tuple with the FeedbackKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedbackKeys

`func (o *RunCreateStateful) SetFeedbackKeys(v []string)`

SetFeedbackKeys sets FeedbackKeys field to given value.

### HasFeedbackKeys

`func (o *RunCreateStateful) HasFeedbackKeys() bool`

HasFeedbackKeys returns a boolean if a field has been set.

### GetMultitaskStrategy

`func (o *RunCreateStateful) GetMultitaskStrategy() string`

GetMultitaskStrategy returns the MultitaskStrategy field if non-nil, zero value otherwise.

### GetMultitaskStrategyOk

`func (o *RunCreateStateful) GetMultitaskStrategyOk() (*string, bool)`

GetMultitaskStrategyOk returns a tuple with the MultitaskStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultitaskStrategy

`func (o *RunCreateStateful) SetMultitaskStrategy(v string)`

SetMultitaskStrategy sets MultitaskStrategy field to given value.

### HasMultitaskStrategy

`func (o *RunCreateStateful) HasMultitaskStrategy() bool`

HasMultitaskStrategy returns a boolean if a field has been set.

### GetIfNotExists

`func (o *RunCreateStateful) GetIfNotExists() string`

GetIfNotExists returns the IfNotExists field if non-nil, zero value otherwise.

### GetIfNotExistsOk

`func (o *RunCreateStateful) GetIfNotExistsOk() (*string, bool)`

GetIfNotExistsOk returns a tuple with the IfNotExists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIfNotExists

`func (o *RunCreateStateful) SetIfNotExists(v string)`

SetIfNotExists sets IfNotExists field to given value.

### HasIfNotExists

`func (o *RunCreateStateful) HasIfNotExists() bool`

HasIfNotExists returns a boolean if a field has been set.

### GetAfterSeconds

`func (o *RunCreateStateful) GetAfterSeconds() int32`

GetAfterSeconds returns the AfterSeconds field if non-nil, zero value otherwise.

### GetAfterSecondsOk

`func (o *RunCreateStateful) GetAfterSecondsOk() (*int32, bool)`

GetAfterSecondsOk returns a tuple with the AfterSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfterSeconds

`func (o *RunCreateStateful) SetAfterSeconds(v int32)`

SetAfterSeconds sets AfterSeconds field to given value.

### HasAfterSeconds

`func (o *RunCreateStateful) HasAfterSeconds() bool`

HasAfterSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


