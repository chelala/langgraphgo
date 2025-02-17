# Cron

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CronId** | **string** | The ID of the cron. | 
**ThreadId** | **string** | The ID of the thread. | 
**EndTime** | **time.Time** | The end date to stop running the cron. | 
**Schedule** | **string** | The schedule to run, cron format. | 
**CreatedAt** | **time.Time** | The time the cron was created. | 
**UpdatedAt** | **time.Time** | The last time the cron was updated. | 
**Payload** | **map[string]interface{}** | The run payload to use for creating new run. | 

## Methods

### NewCron

`func NewCron(cronId string, threadId string, endTime time.Time, schedule string, createdAt time.Time, updatedAt time.Time, payload map[string]interface{}, ) *Cron`

NewCron instantiates a new Cron object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCronWithDefaults

`func NewCronWithDefaults() *Cron`

NewCronWithDefaults instantiates a new Cron object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCronId

`func (o *Cron) GetCronId() string`

GetCronId returns the CronId field if non-nil, zero value otherwise.

### GetCronIdOk

`func (o *Cron) GetCronIdOk() (*string, bool)`

GetCronIdOk returns a tuple with the CronId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronId

`func (o *Cron) SetCronId(v string)`

SetCronId sets CronId field to given value.


### GetThreadId

`func (o *Cron) GetThreadId() string`

GetThreadId returns the ThreadId field if non-nil, zero value otherwise.

### GetThreadIdOk

`func (o *Cron) GetThreadIdOk() (*string, bool)`

GetThreadIdOk returns a tuple with the ThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadId

`func (o *Cron) SetThreadId(v string)`

SetThreadId sets ThreadId field to given value.


### GetEndTime

`func (o *Cron) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *Cron) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *Cron) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.


### GetSchedule

`func (o *Cron) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *Cron) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *Cron) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.


### GetCreatedAt

`func (o *Cron) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Cron) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Cron) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Cron) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Cron) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Cron) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetPayload

`func (o *Cron) GetPayload() map[string]interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *Cron) GetPayloadOk() (*map[string]interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *Cron) SetPayload(v map[string]interface{})`

SetPayload sets Payload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


