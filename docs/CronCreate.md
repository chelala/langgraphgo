# CronCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schedule** | **string** | The cron schedule to execute this job on. | 
**AssistantId** | [**CronCreateAssistantId**](CronCreateAssistantId.md) |  | 
**Input** | Pointer to [**Input**](Input.md) |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Metadata to assign to the cron job runs. | [optional] 
**Config** | Pointer to [**Config2**](Config2.md) |  | [optional] 
**Webhook** | Pointer to **string** | Webhook to call after LangGraph API call is done. | [optional] 
**InterruptBefore** | Pointer to [**InterruptBefore**](InterruptBefore.md) |  | [optional] 
**InterruptAfter** | Pointer to [**InterruptAfter**](InterruptAfter.md) |  | [optional] 
**MultitaskStrategy** | Pointer to **string** | Multitask strategy to use. Must be one of &#39;reject&#39;, &#39;interrupt&#39;, &#39;rollback&#39;, or &#39;enqueue&#39;. | [optional] [default to "reject"]

## Methods

### NewCronCreate

`func NewCronCreate(schedule string, assistantId CronCreateAssistantId, ) *CronCreate`

NewCronCreate instantiates a new CronCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCronCreateWithDefaults

`func NewCronCreateWithDefaults() *CronCreate`

NewCronCreateWithDefaults instantiates a new CronCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchedule

`func (o *CronCreate) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *CronCreate) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *CronCreate) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.


### GetAssistantId

`func (o *CronCreate) GetAssistantId() CronCreateAssistantId`

GetAssistantId returns the AssistantId field if non-nil, zero value otherwise.

### GetAssistantIdOk

`func (o *CronCreate) GetAssistantIdOk() (*CronCreateAssistantId, bool)`

GetAssistantIdOk returns a tuple with the AssistantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssistantId

`func (o *CronCreate) SetAssistantId(v CronCreateAssistantId)`

SetAssistantId sets AssistantId field to given value.


### GetInput

`func (o *CronCreate) GetInput() Input`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *CronCreate) GetInputOk() (*Input, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *CronCreate) SetInput(v Input)`

SetInput sets Input field to given value.

### HasInput

`func (o *CronCreate) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetMetadata

`func (o *CronCreate) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CronCreate) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CronCreate) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CronCreate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetConfig

`func (o *CronCreate) GetConfig() Config2`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CronCreate) GetConfigOk() (*Config2, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CronCreate) SetConfig(v Config2)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *CronCreate) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetWebhook

`func (o *CronCreate) GetWebhook() string`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *CronCreate) GetWebhookOk() (*string, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *CronCreate) SetWebhook(v string)`

SetWebhook sets Webhook field to given value.

### HasWebhook

`func (o *CronCreate) HasWebhook() bool`

HasWebhook returns a boolean if a field has been set.

### GetInterruptBefore

`func (o *CronCreate) GetInterruptBefore() InterruptBefore`

GetInterruptBefore returns the InterruptBefore field if non-nil, zero value otherwise.

### GetInterruptBeforeOk

`func (o *CronCreate) GetInterruptBeforeOk() (*InterruptBefore, bool)`

GetInterruptBeforeOk returns a tuple with the InterruptBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterruptBefore

`func (o *CronCreate) SetInterruptBefore(v InterruptBefore)`

SetInterruptBefore sets InterruptBefore field to given value.

### HasInterruptBefore

`func (o *CronCreate) HasInterruptBefore() bool`

HasInterruptBefore returns a boolean if a field has been set.

### GetInterruptAfter

`func (o *CronCreate) GetInterruptAfter() InterruptAfter`

GetInterruptAfter returns the InterruptAfter field if non-nil, zero value otherwise.

### GetInterruptAfterOk

`func (o *CronCreate) GetInterruptAfterOk() (*InterruptAfter, bool)`

GetInterruptAfterOk returns a tuple with the InterruptAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterruptAfter

`func (o *CronCreate) SetInterruptAfter(v InterruptAfter)`

SetInterruptAfter sets InterruptAfter field to given value.

### HasInterruptAfter

`func (o *CronCreate) HasInterruptAfter() bool`

HasInterruptAfter returns a boolean if a field has been set.

### GetMultitaskStrategy

`func (o *CronCreate) GetMultitaskStrategy() string`

GetMultitaskStrategy returns the MultitaskStrategy field if non-nil, zero value otherwise.

### GetMultitaskStrategyOk

`func (o *CronCreate) GetMultitaskStrategyOk() (*string, bool)`

GetMultitaskStrategyOk returns a tuple with the MultitaskStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultitaskStrategy

`func (o *CronCreate) SetMultitaskStrategy(v string)`

SetMultitaskStrategy sets MultitaskStrategy field to given value.

### HasMultitaskStrategy

`func (o *CronCreate) HasMultitaskStrategy() bool`

HasMultitaskStrategy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


