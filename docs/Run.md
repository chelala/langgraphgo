# Run

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RunId** | **string** | The ID of the run. | 
**ThreadId** | **string** | The ID of the thread. | 
**AssistantId** | **string** | The assistant that was used for this run. | 
**CreatedAt** | **time.Time** | The time the run was created. | 
**UpdatedAt** | **time.Time** | The last time the run was updated. | 
**Status** | **string** | The status of the run. One of &#39;pending&#39;, &#39;error&#39;, &#39;success&#39;, &#39;timeout&#39;, &#39;interrupted&#39;. | 
**Metadata** | **map[string]interface{}** | The run metadata. | 
**Kwargs** | **map[string]interface{}** |  | 
**MultitaskStrategy** | **string** | Strategy to handle concurrent runs on the same thread. | 

## Methods

### NewRun

`func NewRun(runId string, threadId string, assistantId string, createdAt time.Time, updatedAt time.Time, status string, metadata map[string]interface{}, kwargs map[string]interface{}, multitaskStrategy string, ) *Run`

NewRun instantiates a new Run object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunWithDefaults

`func NewRunWithDefaults() *Run`

NewRunWithDefaults instantiates a new Run object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRunId

`func (o *Run) GetRunId() string`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *Run) GetRunIdOk() (*string, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *Run) SetRunId(v string)`

SetRunId sets RunId field to given value.


### GetThreadId

`func (o *Run) GetThreadId() string`

GetThreadId returns the ThreadId field if non-nil, zero value otherwise.

### GetThreadIdOk

`func (o *Run) GetThreadIdOk() (*string, bool)`

GetThreadIdOk returns a tuple with the ThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadId

`func (o *Run) SetThreadId(v string)`

SetThreadId sets ThreadId field to given value.


### GetAssistantId

`func (o *Run) GetAssistantId() string`

GetAssistantId returns the AssistantId field if non-nil, zero value otherwise.

### GetAssistantIdOk

`func (o *Run) GetAssistantIdOk() (*string, bool)`

GetAssistantIdOk returns a tuple with the AssistantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssistantId

`func (o *Run) SetAssistantId(v string)`

SetAssistantId sets AssistantId field to given value.


### GetCreatedAt

`func (o *Run) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Run) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Run) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Run) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Run) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Run) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetStatus

`func (o *Run) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Run) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Run) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetMetadata

`func (o *Run) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Run) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Run) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetKwargs

`func (o *Run) GetKwargs() map[string]interface{}`

GetKwargs returns the Kwargs field if non-nil, zero value otherwise.

### GetKwargsOk

`func (o *Run) GetKwargsOk() (*map[string]interface{}, bool)`

GetKwargsOk returns a tuple with the Kwargs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKwargs

`func (o *Run) SetKwargs(v map[string]interface{})`

SetKwargs sets Kwargs field to given value.


### GetMultitaskStrategy

`func (o *Run) GetMultitaskStrategy() string`

GetMultitaskStrategy returns the MultitaskStrategy field if non-nil, zero value otherwise.

### GetMultitaskStrategyOk

`func (o *Run) GetMultitaskStrategyOk() (*string, bool)`

GetMultitaskStrategyOk returns a tuple with the MultitaskStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultitaskStrategy

`func (o *Run) SetMultitaskStrategy(v string)`

SetMultitaskStrategy sets MultitaskStrategy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


