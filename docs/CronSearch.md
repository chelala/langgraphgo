# CronSearch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssistantId** | Pointer to **string** | The assistant ID or graph name to search for. | [optional] 
**ThreadId** | Pointer to **string** | The thread ID to search for. | [optional] 
**Limit** | Pointer to **int32** | The maximum number of results to return. | [optional] [default to 10]
**Offset** | Pointer to **int32** | The number of results to skip. | [optional] [default to 0]

## Methods

### NewCronSearch

`func NewCronSearch() *CronSearch`

NewCronSearch instantiates a new CronSearch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCronSearchWithDefaults

`func NewCronSearchWithDefaults() *CronSearch`

NewCronSearchWithDefaults instantiates a new CronSearch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssistantId

`func (o *CronSearch) GetAssistantId() string`

GetAssistantId returns the AssistantId field if non-nil, zero value otherwise.

### GetAssistantIdOk

`func (o *CronSearch) GetAssistantIdOk() (*string, bool)`

GetAssistantIdOk returns a tuple with the AssistantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssistantId

`func (o *CronSearch) SetAssistantId(v string)`

SetAssistantId sets AssistantId field to given value.

### HasAssistantId

`func (o *CronSearch) HasAssistantId() bool`

HasAssistantId returns a boolean if a field has been set.

### GetThreadId

`func (o *CronSearch) GetThreadId() string`

GetThreadId returns the ThreadId field if non-nil, zero value otherwise.

### GetThreadIdOk

`func (o *CronSearch) GetThreadIdOk() (*string, bool)`

GetThreadIdOk returns a tuple with the ThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadId

`func (o *CronSearch) SetThreadId(v string)`

SetThreadId sets ThreadId field to given value.

### HasThreadId

`func (o *CronSearch) HasThreadId() bool`

HasThreadId returns a boolean if a field has been set.

### GetLimit

`func (o *CronSearch) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *CronSearch) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *CronSearch) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *CronSearch) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *CronSearch) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *CronSearch) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *CronSearch) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *CronSearch) HasOffset() bool`

HasOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


